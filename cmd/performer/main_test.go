package main

import (
	"encoding/binary"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
	
	"github.com/Layr-Labs/hourglass-avs-template/pkg/bindings/VRF"
)

func TestVRFTaskValidation(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	taskWorker := NewTaskWorker(logger)

	t.Run("ValidVDFTask", func(t *testing.T) {
		payload := createValidVDFPayload(t, []byte("test-seed-123"))
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err != nil {
			t.Errorf("ValidateTask failed for valid payload: %v", err)
		}
	})

	t.Run("InvalidPayloadTooShort", func(t *testing.T) {
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  []byte("short"),
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err == nil {
			t.Error("Expected validation to fail for short payload")
		}
	})

	t.Run("UnsupportedRandomnessType", func(t *testing.T) {
		payload := createInvalidRandomnessTypePayload(t)
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err == nil {
			t.Error("Expected validation to fail for unsupported randomness type")
		}
	})

	t.Run("EmptySeed", func(t *testing.T) {
		payload := createValidVDFPayload(t, []byte{})
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err == nil {
			t.Error("Expected validation to fail for empty seed")
		}
	})

	t.Run("SeedTooLong", func(t *testing.T) {
		largeSeed := make([]byte, 1025) // Exceeds 1024 byte limit
		for i := range largeSeed {
			largeSeed[i] = byte(i % 256)
		}
		payload := createValidVDFPayload(t, largeSeed)
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err == nil {
			t.Error("Expected validation to fail for oversized seed")
		}
	})

	t.Run("BackwardCompatibilityOldVDFType", func(t *testing.T) {
		// Test that old VDF type (0) still works for backward compatibility
		payload := createVDFPayloadWithType(t, []byte("test-seed"), RandomnessType(0))
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err != nil {
			t.Errorf("ValidateTask failed for old VDF type (0): %v", err)
		}
	})

	t.Run("ProperABIEncodingFromContract", func(t *testing.T) {
		// Test with proper ABI encoding from the contract's view functions
		payload := createContractEncodedVDFPayload(t, []byte("test"))
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		err := taskWorker.ValidateTask(taskRequest)
		if err != nil {
			t.Errorf("ValidateTask failed for contract ABI encoding: %v", err)
		}
	})
}

func TestVRFTaskHandling(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	taskWorker := NewTaskWorker(logger)

	t.Run("SuccessfulVDFComputation", func(t *testing.T) {
		seed := []byte("deterministic-test-seed")
		payload := createValidVDFPayload(t, seed)
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		resp, err := taskWorker.HandleTask(taskRequest)
		if err != nil {
			t.Fatalf("HandleTask failed: %v", err)
		}

		if resp == nil {
			t.Fatal("Response is nil")
		}

		if string(resp.TaskId) != "test-task-id" {
			t.Errorf("Expected task ID 'test-task-id', got '%s'", string(resp.TaskId))
		}

		if len(resp.Result) == 0 {
			t.Error("Expected non-empty result")
		}

		t.Logf("VDF computation result length: %d bytes", len(resp.Result))
	})

	t.Run("DeterministicResults", func(t *testing.T) {
		seed := []byte("deterministic-seed-for-testing")
		payload := createValidVDFPayload(t, seed)
		
		taskRequest := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-id"),
			Payload:  payload,
			Metadata: []byte("test-metadata"),
		}

		// Run computation twice with same seed
		resp1, err1 := taskWorker.HandleTask(taskRequest)
		resp2, err2 := taskWorker.HandleTask(taskRequest)

		if err1 != nil || err2 != nil {
			t.Fatalf("HandleTask failed: err1=%v, err2=%v", err1, err2)
		}

		// Results should be deterministic (same seed = same result)
		if len(resp1.Result) != len(resp2.Result) {
			t.Error("Result lengths differ between computations")
		}

		// Compare result bytes
		for i := range resp1.Result {
			if resp1.Result[i] != resp2.Result[i] {
				t.Error("Results differ between computations with same seed")
				break
			}
		}
	})

	t.Run("DifferentSeedsDifferentResults", func(t *testing.T) {
		seed1 := []byte("seed-one")
		seed2 := []byte("seed-two")
		
		payload1 := createValidVDFPayload(t, seed1)
		payload2 := createValidVDFPayload(t, seed2)
		
		taskRequest1 := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-1"),
			Payload:  payload1,
			Metadata: []byte("test-metadata"),
		}
		
		taskRequest2 := &performerV1.TaskRequest{
			TaskId:   []byte("test-task-2"),
			Payload:  payload2,
			Metadata: []byte("test-metadata"),
		}

		resp1, err1 := taskWorker.HandleTask(taskRequest1)
		resp2, err2 := taskWorker.HandleTask(taskRequest2)

		if err1 != nil || err2 != nil {
			t.Fatalf("HandleTask failed: err1=%v, err2=%v", err1, err2)
		}

		// Results should be different for different seeds
		if len(resp1.Result) == len(resp2.Result) {
			allSame := true
			for i := range resp1.Result {
				if resp1.Result[i] != resp2.Result[i] {
					allSame = false
					break
				}
			}
			if allSame {
				t.Error("Different seeds produced identical results")
			}
		}
	})
}

func TestEncodingDecoding(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	taskWorker := NewTaskWorker(logger)

	t.Run("TaskPayloadRoundTrip", func(t *testing.T) {
		originalSeed := []byte("test-seed-for-roundtrip")
		payload := createValidVDFPayload(t, originalSeed)

		// Decode the payload we just created
		decoded, err := taskWorker.decodeTaskPayload(payload)
		if err != nil {
			t.Fatalf("Failed to decode payload: %v", err)
		}

		if decoded.RandomnessType != VDF {
			t.Errorf("Expected randomness type %d, got %d", VDF, decoded.RandomnessType)
		}

		// Decode VDF params
		vdfParams, err := taskWorker.decodeVDFParams(decoded.RandomnessParams)
		if err != nil {
			t.Fatalf("Failed to decode VDF params: %v", err)
		}

		if len(vdfParams.Seed) != len(originalSeed) {
			t.Errorf("Seed length mismatch: expected %d, got %d", len(originalSeed), len(vdfParams.Seed))
		}

		for i := range originalSeed {
			if vdfParams.Seed[i] != originalSeed[i] {
				t.Error("Decoded seed doesn't match original")
				break
			}
		}
	})
}

// Helper function to create a valid VDF payload for testing
func createValidVDFPayload(t *testing.T, seed []byte) []byte {
	return createVDFPayloadWithType(t, seed, VDF)
}

// Helper function to create VDF payload with specific randomness type using proper ABI encoding
func createVDFPayloadWithType(t *testing.T, seed []byte, randomnessType RandomnessType) []byte {
	// Use proper ABI encoding like the contract does
	
	// Create VDFParams struct
	vdfParams := struct {
		Seed []byte
	}{
		Seed: seed,
	}
	
	// Encode VDFParams
	vdfParamsType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "seed", Type: "bytes"},
	})
	if err != nil {
		t.Fatalf("Failed to create VDFParams type: %v", err)
	}
	
	vdfParamsArgs := abi.Arguments{{Type: vdfParamsType}}
	encodedVDFParams, err := vdfParamsArgs.Pack(vdfParams)
	if err != nil {
		t.Fatalf("Failed to encode VDFParams: %v", err)
	}
	
	// Create TaskPayload struct
	taskPayload := struct {
		RandomnessType uint8
		RandomnessParams []byte
	}{
		RandomnessType: uint8(randomnessType),
		RandomnessParams: encodedVDFParams,
	}
	
	// Encode TaskPayload
	taskPayloadType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "randomnessType", Type: "uint8"},
		{Name: "randomnessParams", Type: "bytes"},
	})
	if err != nil {
		t.Fatalf("Failed to create TaskPayload type: %v", err)
	}
	
	taskPayloadArgs := abi.Arguments{{Type: taskPayloadType}}
	encodedPayload, err := taskPayloadArgs.Pack(taskPayload)
	if err != nil {
		t.Fatalf("Failed to encode TaskPayload: %v", err)
	}
	
	return encodedPayload
}

// Helper function to create invalid randomness type payload
func createInvalidRandomnessTypePayload(t *testing.T) []byte {
	seed := []byte("test-seed")
	// Use invalid randomness type (255)
	return createVDFPayloadWithType(t, seed, RandomnessType(255))
}

// Helper function to create proper ABI-encoded payload using the contract
func createContractEncodedVDFPayload(t *testing.T, seed []byte) []byte {
	// Set up a simulated backend for testing
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)}
	backend := backends.NewSimulatedBackend(alloc, 10000000)
	
	// Deploy the VRF contract
	// Note: This is a simplified deployment - in real tests you'd need proper constructor args
	mockOperatorSet := VRF.OperatorSet{
		Avs: auth.From,
		Id:  uint32(1),
	}
	
	contractAddr, _, contract, err := VRF.DeployVRF(auth, backend, auth.From, mockOperatorSet)
	if err != nil {
		t.Fatalf("Failed to deploy VRF contract: %v", err)
	}
	backend.Commit()
	
	// Call the contract's encoding function
	callOpts := &bind.CallOpts{}
	encodedPayload, err := contract.EncodeVDFTaskPayload(callOpts, seed)
	if err != nil {
		t.Fatalf("Failed to encode VDF task payload: %v", err)
	}
	
	t.Logf("Contract deployed at %s, encoded payload length: %d", contractAddr.Hex(), len(encodedPayload))
	return encodedPayload
}

// Helper function to create proper ABI-encoded payload (simulating Solidity encoding)
func createProperABIPayload(t *testing.T, seed []byte, randomnessType RandomnessType) []byte {
	// Encode VDF params (seed)
	vdfParamsLength := 32 + len(seed) // 32 bytes for length + seed data
	paramsData := make([]byte, vdfParamsLength)
	binary.BigEndian.PutUint64(paramsData[24:32], uint64(len(seed))) // seed length as uint256
	copy(paramsData[32:], seed) // seed data
	
	// Calculate total payload size
	// 32 bytes (randomnessType) + 32 bytes (offset) + 32 bytes (params length) + params data
	totalSize := 64 + vdfParamsLength
	payload := make([]byte, totalSize)
	
	// First 32 bytes: randomnessType as uint256
	binary.BigEndian.PutUint64(payload[24:32], uint64(randomnessType))
	
	// Second 32 bytes: offset to params data (0x40 = 64)
	binary.BigEndian.PutUint64(payload[56:64], 64)
	
	// Third 32 bytes: length of params data
	binary.BigEndian.PutUint64(payload[88:96], uint64(vdfParamsLength))
	
	// Remaining bytes: actual params data
	copy(payload[96:], paramsData)
	
	return payload
}
