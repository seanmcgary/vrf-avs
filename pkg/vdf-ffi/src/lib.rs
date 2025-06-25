// FFI bindings for VDF library for use with Go CGO
use std::ffi::{CStr, CString};
use std::os::raw::{c_char, c_int, c_uchar};
use std::slice;
use vdf::{PietrzakVDFParams, VDFParams, VDF};

/// Result codes for VDF operations
#[repr(C)]
pub enum VDFResultCode {
    Success = 0,
    InvalidInput = 1,
    ComputationError = 2,
    VerificationError = 3,
    InternalError = 4,
}

/// VDF computation result structure
#[repr(C)]
pub struct VDFResult {
    pub code: VDFResultCode,
    pub output_len: usize,
    pub output_ptr: *mut c_uchar,
}

/// Opaque handle for VDF parameters
pub struct VDFHandle;

/// Initialize VDF with given difficulty (key size in bits)
#[no_mangle]
pub extern "C" fn vdf_new(difficulty: c_int) -> *mut VDFHandle {
    if difficulty <= 0 || difficulty > 4096 {
        return std::ptr::null_mut();
    }
    
    let params = Box::new(PietrzakVDFParams(difficulty as u16));
    Box::into_raw(params) as *mut VDFHandle
}

/// Free VDF parameters
#[no_mangle]
pub extern "C" fn vdf_free(params: *mut VDFHandle) {
    if !params.is_null() {
        unsafe {
            let _ = Box::from_raw(params as *mut PietrzakVDFParams);
        }
    }
}

/// Solve VDF with given challenge and iterations
#[no_mangle]
pub extern "C" fn vdf_solve(
    params: *const VDFHandle,
    challenge: *const c_uchar,
    challenge_len: usize,
    iterations: u64,
) -> VDFResult {
    if params.is_null() || challenge.is_null() || challenge_len == 0 {
        return VDFResult {
            code: VDFResultCode::InvalidInput,
            output_len: 0,
            output_ptr: std::ptr::null_mut(),
        };
    }

    let params = unsafe { &*(params as *const PietrzakVDFParams) };
    let challenge_slice = unsafe { slice::from_raw_parts(challenge, challenge_len) };

    let vdf = params.new();
    
    match vdf.solve(challenge_slice, iterations) {
        Ok(solution) => {
            let output_len = solution.len();
            let output_vec = solution.into_boxed_slice();
            let output_ptr = Box::into_raw(output_vec) as *mut c_uchar;
            
            VDFResult {
                code: VDFResultCode::Success,
                output_len,
                output_ptr,
            }
        }
        Err(_) => VDFResult {
            code: VDFResultCode::ComputationError,
            output_len: 0,
            output_ptr: std::ptr::null_mut(),
        },
    }
}

/// Verify VDF solution
#[no_mangle]
pub extern "C" fn vdf_verify(
    params: *const VDFHandle,
    challenge: *const c_uchar,
    challenge_len: usize,
    iterations: u64,
    solution: *const c_uchar,
    solution_len: usize,
) -> VDFResultCode {
    if params.is_null() || challenge.is_null() || solution.is_null() || 
       challenge_len == 0 || solution_len == 0 {
        return VDFResultCode::InvalidInput;
    }

    let params = unsafe { &*(params as *const PietrzakVDFParams) };
    let challenge_slice = unsafe { slice::from_raw_parts(challenge, challenge_len) };
    let solution_slice = unsafe { slice::from_raw_parts(solution, solution_len) };

    let vdf = params.new();
    
    match vdf.verify(challenge_slice, iterations, solution_slice) {
        Ok(_) => VDFResultCode::Success,
        Err(_) => VDFResultCode::VerificationError,
    }
}

/// Free memory allocated by vdf_solve
#[no_mangle]
pub extern "C" fn vdf_free_result(result: *mut c_uchar, len: usize) {
    if !result.is_null() && len > 0 {
        unsafe {
            let _ = Box::from_raw(slice::from_raw_parts_mut(result, len));
        }
    }
}

/// Get error message for a result code
#[no_mangle]
pub extern "C" fn vdf_error_message(code: VDFResultCode) -> *const c_char {
    let message = match code {
        VDFResultCode::Success => "Success",
        VDFResultCode::InvalidInput => "Invalid input parameters",
        VDFResultCode::ComputationError => "VDF computation failed",
        VDFResultCode::VerificationError => "VDF verification failed",
        VDFResultCode::InternalError => "Internal error",
    };
    
    match CString::new(message) {
        Ok(c_str) => c_str.into_raw(),
        Err(_) => std::ptr::null(),
    }
}

/// Free error message string
#[no_mangle]
pub extern "C" fn vdf_free_string(s: *mut c_char) {
    if !s.is_null() {
        unsafe {
            let _ = CString::from_raw(s);
        }
    }
}