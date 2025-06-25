package main

import (
	"encoding/binary"
	"testing"

	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
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
	// Create VDF params
	vdfParams := make([]byte, 32+len(seed))
	// Seed length as uint256 (32 bytes)
	binary.BigEndian.PutUint64(vdfParams[24:32], uint64(len(seed)))
	// Seed data
	copy(vdfParams[32:], seed)

	// Create task payload
	payload := make([]byte, 96+len(vdfParams))
	
	// Randomness type as uint256 (32 bytes) - VDF = 0
	payload[31] = byte(VDF)
	
	// VDF params length as uint256 (32 bytes)
	binary.BigEndian.PutUint64(payload[88:96], uint64(len(vdfParams)))
	
	// VDF params
	copy(payload[96:], vdfParams)

	return payload
}

// Helper function to create invalid randomness type payload
func createInvalidRandomnessTypePayload(t *testing.T) []byte {
	seed := []byte("test-seed")
	
	// Create VDF params
	vdfParams := make([]byte, 32+len(seed))
	binary.BigEndian.PutUint64(vdfParams[24:32], uint64(len(seed)))
	copy(vdfParams[32:], seed)

	// Create task payload with invalid randomness type
	payload := make([]byte, 96+len(vdfParams))
	
	// Invalid randomness type (255)
	payload[31] = 255
	
	// VDF params length
	binary.BigEndian.PutUint64(payload[88:96], uint64(len(vdfParams)))
	
	// VDF params
	copy(payload[96:], vdfParams)

	return payload
}
