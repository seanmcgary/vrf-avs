package vdf

import (
	"bytes"
	"testing"
)

func TestVDFBasicOperation(t *testing.T) {
	// Create VDF with moderate difficulty for testing
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}
	defer vdf.Close()

	challenge := []byte{0xaa} // Use simple byte challenge like VDF examples
	iterations := uint64(100) // Must be even for Pietrzak VDF

	// Solve VDF
	solution, err := vdf.Solve(challenge, iterations)
	if err != nil {
		t.Fatalf("Failed to solve VDF: %v", err)
	}

	if len(solution) == 0 {
		t.Fatal("Solution is empty")
	}

	t.Logf("VDF solution length: %d bytes", len(solution))

	// Verify solution
	err = vdf.Verify(challenge, iterations, solution)
	if err != nil {
		t.Fatalf("Failed to verify VDF solution: %v", err)
	}

	t.Log("VDF solution verified successfully")
}

func TestVDFDeterministic(t *testing.T) {
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}
	defer vdf.Close()

	challenge := []byte{0xbb} // Use simple byte challenge
	iterations := uint64(100) // Must be even and >= 66 for Pietrzak VDF

	// Solve twice
	solution1, err := vdf.Solve(challenge, iterations)
	if err != nil {
		t.Fatalf("Failed to solve VDF (first): %v", err)
	}

	solution2, err := vdf.Solve(challenge, iterations)
	if err != nil {
		t.Fatalf("Failed to solve VDF (second): %v", err)
	}

	// Solutions should be identical
	if !bytes.Equal(solution1, solution2) {
		t.Fatal("VDF solutions are not deterministic")
	}

	t.Log("VDF produces deterministic results")
}

func TestVDFDifferentChallenges(t *testing.T) {
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}
	defer vdf.Close()

	challenge1 := []byte{0xcc}
	challenge2 := []byte{0xdd}
	iterations := uint64(100) // Must be even and >= 66 for Pietrzak VDF

	solution1, err := vdf.Solve(challenge1, iterations)
	if err != nil {
		t.Fatalf("Failed to solve VDF for challenge1: %v", err)
	}

	solution2, err := vdf.Solve(challenge2, iterations)
	if err != nil {
		t.Fatalf("Failed to solve VDF for challenge2: %v", err)
	}

	// Different challenges should produce different solutions
	if bytes.Equal(solution1, solution2) {
		t.Fatal("Different challenges produced identical solutions")
	}

	// Verify both solutions
	if err := vdf.Verify(challenge1, iterations, solution1); err != nil {
		t.Fatalf("Failed to verify solution1: %v", err)
	}

	if err := vdf.Verify(challenge2, iterations, solution2); err != nil {
		t.Fatalf("Failed to verify solution2: %v", err)
	}

	t.Log("Different challenges produce different valid solutions")
}

func TestVDFInvalidInput(t *testing.T) {
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}
	defer vdf.Close()

	// Test empty challenge
	_, err = vdf.Solve([]byte{}, 100)
	if err != ErrInvalidInput {
		t.Fatalf("Expected ErrInvalidInput for empty challenge, got: %v", err)
	}

	// Test verification with empty challenge
	solution := []byte("fake-solution")
	err = vdf.Verify([]byte{}, 100, solution)
	if err != ErrInvalidInput {
		t.Fatalf("Expected ErrInvalidInput for empty challenge in verify, got: %v", err)
	}

	// Test verification with empty solution
	challenge := []byte("test-challenge")
	err = vdf.Verify(challenge, 100, []byte{})
	if err != ErrInvalidInput {
		t.Fatalf("Expected ErrInvalidInput for empty solution, got: %v", err)
	}

	t.Log("Invalid input handling works correctly")
}

func TestVDFInvalidSolution(t *testing.T) {
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}
	defer vdf.Close()

	challenge := []byte{0xee}
	iterations := uint64(100) // Must be even and >= 66 for Pietrzak VDF
	fakeSolution := []byte("this-is-not-a-valid-solution")

	// Try to verify invalid solution
	err = vdf.Verify(challenge, iterations, fakeSolution)
	if err != ErrVerificationError {
		t.Fatalf("Expected ErrVerificationError for fake solution, got: %v", err)
	}

	t.Log("Invalid solution correctly rejected")
}

func TestVDFInvalidDifficulty(t *testing.T) {
	// Test invalid difficulties
	testCases := []int{0, -1, 5000}

	for _, difficulty := range testCases {
		_, err := NewVDF(difficulty)
		if err == nil {
			t.Fatalf("Expected error for invalid difficulty %d", difficulty)
		}
		t.Logf("Correctly rejected difficulty %d: %v", difficulty, err)
	}
}

func TestVDFClose(t *testing.T) {
	vdf, err := NewVDF(1024)
	if err != nil {
		t.Fatalf("Failed to create VDF: %v", err)
	}

	// Close the VDF
	vdf.Close()

	// Try to use closed VDF
	_, err = vdf.Solve([]byte("test"), 10)
	if err == nil {
		t.Fatal("Expected error when using closed VDF")
	}

	// Close again should be safe
	vdf.Close()

	t.Log("VDF close handling works correctly")
}