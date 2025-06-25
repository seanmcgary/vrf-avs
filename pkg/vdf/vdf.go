package vdf

/*
#cgo CFLAGS: -I../vdf-ffi/target/include/vdf-ffi
#cgo LDFLAGS: -L../vdf-ffi/target/release -lvdf_ffi -lgmp
#include "vdf_ffi.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// VDF errors
var (
	ErrInvalidInput      = errors.New("invalid input parameters")
	ErrComputationError  = errors.New("VDF computation failed")
	ErrVerificationError = errors.New("VDF verification failed")
	ErrInternalError     = errors.New("internal error")
)

// VDF represents a Verifiable Delay Function instance
type VDF struct {
	params *C.struct_VDFHandle
}

// NewVDF creates a new VDF instance with the specified difficulty
// Difficulty should be a power of 2, typically 1024, 2048, or 4096
func NewVDF(difficulty int) (*VDF, error) {
	if difficulty <= 0 || difficulty > 4096 {
		return nil, fmt.Errorf("difficulty must be between 1 and 4096, got %d", difficulty)
	}

	params := C.vdf_new(C.int(difficulty))
	if params == nil {
		return nil, fmt.Errorf("failed to create VDF with difficulty %d", difficulty)
	}

	return &VDF{params: params}, nil
}

// Close frees the VDF resources
func (v *VDF) Close() {
	if v.params != nil {
		C.vdf_free(v.params)
		v.params = nil
	}
}

// Solve computes a VDF proof for the given challenge and iterations
func (v *VDF) Solve(challenge []byte, iterations uint64) ([]byte, error) {
	if v.params == nil {
		return nil, errors.New("VDF instance is closed")
	}
	if len(challenge) == 0 {
		return nil, ErrInvalidInput
	}

	challengePtr := (*C.uchar)(unsafe.Pointer(&challenge[0]))
	result := C.vdf_solve(v.params, challengePtr, C.size_t(len(challenge)), C.ulonglong(iterations))

	switch result.code {
	case C.Success:
		if result.output_len == 0 || result.output_ptr == nil {
			return nil, ErrInternalError
		}
		
		// Convert C array to Go slice
		output := C.GoBytes(unsafe.Pointer(result.output_ptr), C.int(result.output_len))
		
		// Free the C memory
		C.vdf_free_result(result.output_ptr, result.output_len)
		
		return output, nil
	case C.InvalidInput:
		return nil, ErrInvalidInput
	case C.ComputationError:
		return nil, ErrComputationError
	case C.InternalError:
		return nil, ErrInternalError
	default:
		return nil, fmt.Errorf("unknown VDF error code: %d", int(result.code))
	}
}

// Verify checks if a VDF solution is valid
func (v *VDF) Verify(challenge []byte, iterations uint64, solution []byte) error {
	if v.params == nil {
		return errors.New("VDF instance is closed")
	}
	if len(challenge) == 0 || len(solution) == 0 {
		return ErrInvalidInput
	}

	challengePtr := (*C.uchar)(unsafe.Pointer(&challenge[0]))
	solutionPtr := (*C.uchar)(unsafe.Pointer(&solution[0]))
	
	code := C.vdf_verify(
		v.params,
		challengePtr, C.size_t(len(challenge)),
		C.ulonglong(iterations),
		solutionPtr, C.size_t(len(solution)),
	)

	switch code {
	case C.Success:
		return nil
	case C.InvalidInput:
		return ErrInvalidInput
	case C.VerificationError:
		return ErrVerificationError
	case C.InternalError:
		return ErrInternalError
	default:
		return fmt.Errorf("unknown VDF error code: %d", int(code))
	}
}

// GetErrorMessage returns a human-readable error message for a VDF error
func GetErrorMessage(err error) string {
	switch err {
	case ErrInvalidInput:
		return "Invalid input parameters"
	case ErrComputationError:
		return "VDF computation failed"
	case ErrVerificationError:
		return "VDF verification failed"
	case ErrInternalError:
		return "Internal error"
	default:
		return err.Error()
	}
}