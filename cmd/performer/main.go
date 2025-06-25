package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"go.uber.org/zap"
)

// This offchain binary is run by Operators running the Hourglass Executor. It contains
// the business logic of the AVS and performs worked based on the tasked sent to it.
// The Hourglass Aggregator ingests tasks from the TaskMailbox and distributes work
// to Executors configured to run the AVS Performer. Performers execute the work and
// return the result to the Executor where the result is signed and return to the
// Aggregator to place in the outbox once the signing threshold is met.

// VRF-specific data structures matching our Solidity contracts

// RandomnessType represents the type of randomness generation
type RandomnessType uint8

const (
	UNKNOWN RandomnessType = 0
	VDF     RandomnessType = 1
)

// VDFParams contains parameters for VDF randomness generation
type VDFParams struct {
	Seed []byte
}

// VDFResult contains the result of VDF computation
type VDFResult struct {
	Result [32]byte // uint256 as 32-byte array
}

// TaskPayload represents the task payload structure
type TaskPayload struct {
	RandomnessType   RandomnessType
	RandomnessParams []byte // ABI-encoded params based on randomnessType
}

// TaskResponsePayload represents the task response structure
type TaskResponsePayload struct {
	RandomnessType RandomnessType
	RandomValue    []byte // ABI-encoded result based on randomnessType
}

// VRF-specific errors
var (
	ErrInvalidPayload  = errors.New("invalid task payload")
	ErrUnsupportedType = errors.New("unsupported randomness type")
	ErrInvalidSeed     = errors.New("invalid seed parameters")
	ErrVDFComputation  = errors.New("VDF computation failed")
	ErrEncodingFailed  = errors.New("encoding failed")
)

type TaskWorker struct {
	logger *zap.Logger
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	return &TaskWorker{
		logger: logger,
	}
}

func (tw *TaskWorker) ValidateTask(t *performerV1.TaskRequest) error {
	tw.logger.Sugar().Infow("Validating VRF task",
		zap.String("taskId", string(t.TaskId)),
		zap.Int("payloadLength", len(t.Payload)),
	)

	// Decode the task payload
	taskPayload, err := tw.decodeTaskPayload(t.Payload)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to decode task payload", "error", err)
		return fmt.Errorf("%w: %v", ErrInvalidPayload, err)
	}

	// Validate randomness type - support both old (0) and new (1) VDF values for backward compatibility
	if taskPayload.RandomnessType != VDF && taskPayload.RandomnessType != RandomnessType(0) {
		tw.logger.Sugar().Errorw("Unsupported randomness type",
			zap.Uint8("type", uint8(taskPayload.RandomnessType)),
		)
		return ErrUnsupportedType
	}

	// Decode and validate VDF parameters
	vdfParams, err := tw.decodeVDFParams(taskPayload.RandomnessParams)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to decode VDF params", "error", err)
		return fmt.Errorf("%w: %v", ErrInvalidSeed, err)
	}

	// Validate seed constraints
	if len(vdfParams.Seed) == 0 {
		tw.logger.Sugar().Errorw("Empty seed provided")
		return fmt.Errorf("%w: seed cannot be empty", ErrInvalidSeed)
	}

	if len(vdfParams.Seed) > 1024 {
		tw.logger.Sugar().Errorw("Seed too long", "length", len(vdfParams.Seed))
		return fmt.Errorf("%w: seed length exceeds maximum (1024 bytes)", ErrInvalidSeed)
	}

	tw.logger.Sugar().Infow("Task validation successful",
		"randomnessType", taskPayload.RandomnessType,
		"seedLength", len(vdfParams.Seed),
	)

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Handling VRF task",
		zap.String("taskId", string(t.TaskId)),
		zap.Int("payloadLength", len(t.Payload)),
	)

	// Decode the task payload
	taskPayload, err := tw.decodeTaskPayload(t.Payload)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to decode task payload", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
	}

	// Decode VDF parameters
	vdfParams, err := tw.decodeVDFParams(taskPayload.RandomnessParams)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to decode VDF params", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrInvalidSeed, err)
	}

	tw.logger.Sugar().Infow("Processing VDF computation",
		"randomnessType", taskPayload.RandomnessType,
		"seedLength", len(vdfParams.Seed),
	)

	// Perform VDF computation (stubbed for now)
	vdfResult, err := tw.computeVDF(vdfParams)
	if err != nil {
		tw.logger.Sugar().Errorw("VDF computation failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrVDFComputation, err)
	}

	// Encode the VDF result
	encodedResult, err := tw.encodeVDFResult(vdfResult)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to encode VDF result", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrEncodingFailed, err)
	}

	// Create task response payload
	responsePayload := TaskResponsePayload{
		RandomnessType: VDF,
		RandomValue:    encodedResult,
	}

	// Encode the response payload
	resultBytes, err := tw.encodeTaskResponsePayload(responsePayload)
	if err != nil {
		tw.logger.Sugar().Errorw("Failed to encode response payload", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrEncodingFailed, err)
	}

	tw.logger.Sugar().Infow("VRF task completed successfully",
		"taskId", string(t.TaskId),
		"resultLength", len(resultBytes),
		"vdfResult", fmt.Sprintf("0x%x", vdfResult.Result[:8]), // Log first 8 bytes for debugging
	)

	return &performerV1.TaskResponse{
		TaskId: t.TaskId,
		Result: resultBytes,
	}, nil
}

// Helper functions for encoding/decoding and VDF computation

// decodeTaskPayload decodes the task payload from bytes using proper ABI decoding
func (tw *TaskWorker) decodeTaskPayload(payload []byte) (*TaskPayload, error) {
	// Define ABI type for TaskPayload struct: (uint8, bytes)
	tupleType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "randomnessType", Type: "uint8"},
		{Name: "randomnessParams", Type: "bytes"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type: %v", err)
	}

	taskPayloadArgs := abi.Arguments{
		{Type: tupleType},
	}

	// Decode the payload
	decoded, err := taskPayloadArgs.Unpack(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TaskPayload: %v", err)
	}

	if len(decoded) != 1 {
		return nil, errors.New("invalid TaskPayload structure")
	}

	// ABI unpacking returns a struct with named fields
	// The type is: struct { RandomnessType uint8; RandomnessParams []uint8 }
	tupleStruct := decoded[0]

	// Use reflection to extract the struct fields
	structValue := reflect.ValueOf(tupleStruct)
	if structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %T", decoded[0])
	}

	// Extract RandomnessType field
	randomnessTypeField := structValue.FieldByName("RandomnessType")
	if !randomnessTypeField.IsValid() {
		return nil, errors.New("missing RandomnessType field")
	}
	randomnessType := uint8(randomnessTypeField.Uint())

	// Extract RandomnessParams field  
	randomnessParamsField := structValue.FieldByName("RandomnessParams")
	if !randomnessParamsField.IsValid() {
		return nil, errors.New("missing RandomnessParams field")
	}

	// Convert []uint8 to []byte
	var randomnessParams []byte
	if randomnessParamsField.Kind() == reflect.Slice {
		slice := randomnessParamsField.Interface().([]uint8)
		randomnessParams = make([]byte, len(slice))
		for i, v := range slice {
			randomnessParams[i] = byte(v)
		}
	} else {
		return nil, errors.New("invalid RandomnessParams field type")
	}

	return &TaskPayload{
		RandomnessType:   RandomnessType(randomnessType),
		RandomnessParams: randomnessParams,
	}, nil
}

// decodeVDFParams decodes VDF parameters from the encoded bytes using ABI
func (tw *TaskWorker) decodeVDFParams(paramsBytes []byte) (*VDFParams, error) {
	// Define ABI type for VDFParams struct: (bytes)
	tupleType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "seed", Type: "bytes"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create VDFParams tuple type: %v", err)
	}

	vdfParamsArgs := abi.Arguments{
		{Type: tupleType},
	}

	// Decode the params
	decoded, err := vdfParamsArgs.Unpack(paramsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode VDFParams: %v", err)
	}

	if len(decoded) != 1 {
		return nil, errors.New("invalid VDFParams structure")
	}

	// ABI unpacking returns a struct with named fields
	// The type is: struct { Seed []uint8 }
	tupleStruct := decoded[0]

	// Use reflection to extract the struct fields
	structValue := reflect.ValueOf(tupleStruct)
	if structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %T", decoded[0])
	}

	// Extract Seed field
	seedField := structValue.FieldByName("Seed")
	if !seedField.IsValid() {
		return nil, errors.New("missing Seed field")
	}

	// Convert []uint8 to []byte
	var seed []byte
	if seedField.Kind() == reflect.Slice {
		slice := seedField.Interface().([]uint8)
		seed = make([]byte, len(slice))
		for i, v := range slice {
			seed[i] = byte(v)
		}
	} else {
		return nil, errors.New("invalid Seed field type")
	}

	return &VDFParams{
		Seed: seed,
	}, nil
}

// computeVDF performs VDF computation (stubbed implementation)
func (tw *TaskWorker) computeVDF(params *VDFParams) (*VDFResult, error) {
	tw.logger.Sugar().Infow("Computing VDF (STUBBED)",
		"seedLength", len(params.Seed),
		"seedPrefix", fmt.Sprintf("0x%x", params.Seed[:min(8, len(params.Seed))]),
	)

	// STUBBED VDF COMPUTATION
	// TODO: Replace with actual VDF computation from poanetwork/vdf
	// For now, we'll use a deterministic hash-based approach

	// Create a deterministic "random" number based on the seed
	hasher := sha256.New()
	hasher.Write(params.Seed)
	hasher.Write([]byte("vdf-stub-computation")) // Add some uniqueness
	hash := hasher.Sum(nil)

	var result [32]byte
	copy(result[:], hash)

	tw.logger.Sugar().Infow("VDF computation completed (STUBBED)",
		"result", fmt.Sprintf("0x%x", result[:8]),
	)

	return &VDFResult{
		Result: result,
	}, nil
}

// encodeVDFResult encodes a VDF result for the response
func (tw *TaskWorker) encodeVDFResult(result *VDFResult) ([]byte, error) {
	// Simple ABI-like encoding: result as uint256 (32 bytes)
	return result.Result[:], nil
}

// encodeTaskResponsePayload encodes the task response payload
func (tw *TaskWorker) encodeTaskResponsePayload(payload TaskResponsePayload) ([]byte, error) {
	// Simple binary encoding matching our decoding logic
	// In production, use proper ABI encoding

	result := make([]byte, 0, 96+len(payload.RandomValue))

	// Encode randomness type as uint256 (32 bytes)
	typeBytes := make([]byte, 32)
	typeBytes[31] = byte(payload.RandomnessType)
	result = append(result, typeBytes...)

	// Encode random value length as uint256 (32 bytes)
	lengthBytes := make([]byte, 32)
	binary.BigEndian.PutUint64(lengthBytes[24:32], uint64(len(payload.RandomValue)))
	result = append(result, lengthBytes...)

	// Encode random value
	result = append(result, payload.RandomValue...)

	return result, nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ctx := context.Background()
	l, _ := zap.NewProduction()

	w := NewTaskWorker(l)

	pp, err := server.NewPonosPerformerWithRpcServer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, w, l)
	if err != nil {
		panic(fmt.Errorf("failed to create performer: %w", err))
	}

	if err := pp.Start(ctx); err != nil {
		panic(err)
	}
}
