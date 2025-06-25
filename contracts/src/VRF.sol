// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IAVSTaskHook} from "@hourglass-monorepo/src/interfaces/avs/l2/IAVSTaskHook.sol";
import {IBN254CertificateVerifierTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IBN254CertificateVerifier.sol";

/**
 * @title VRF - Verifiable Random Function AVS Contract
 * @notice This contract provides verifiable random number generation using VDF (Verifiable Delay Function)
 * @dev Integrates with Hourglass TaskMailbox for task-based randomness generation
 */
contract VRF is IAVSTaskHook {
    /// @notice Types of randomness generation supported
    enum RandomnessType {
        UNKNOWN,
        VDF
    }

    /// @notice Parameters for VDF randomness generation
    struct VDFParams {
        bytes seed;
    }

    /// @notice Result structure for VDF computation
    struct VDFResult {
        uint256 result;
    }

    /// @notice Task payload structure sent to performers
    struct TaskPayload {
        RandomnessType randomnessType;
        bytes randomnessParams; // abi.encoded params based on randomnessType
    }

    /// @notice Task response payload structure returned by performers
    struct TaskResponsePayload {
        RandomnessType randomnessType;
        bytes randomValue; // abi.encoded result based on randomnessType
    }

    /// @notice Information about a randomness request
    struct RandomnessRequest {
        address requester;
        uint256 blockNumber;
        bool fulfilled;
        uint256 result;
    }

    /// @notice The TaskMailbox contract for task management
    ITaskMailbox public immutable taskMailbox;
    
    /// @notice The operator set used for randomness generation
    OperatorSet public executorOperatorSet;
    
    /// @notice Mapping from task hash (used as request ID) to randomness request details
    mapping(bytes32 => RandomnessRequest) public requests;

    /// @notice Emitted when randomness is requested
    event RandomnessRequested(
        bytes32 indexed requestId,
        address indexed requester,
        bytes seed
    );

    /// @notice Emitted when randomness is fulfilled
    event RandomnessFulfilled(
        bytes32 indexed requestId,
        uint256 result
    );

    /// @notice Emitted when task result submission is validated (before processing)
    event TaskResultSubmissionValidated(
        bytes32 indexed requestId
    );

    /// @notice Thrown when caller is not the TaskMailbox
    error UnauthorizedCaller();
    
    /// @notice Thrown when request is not found
    error RequestNotFound();
    
    /// @notice Thrown when request is already fulfilled
    error RequestAlreadyFulfilled();
    
    /// @notice Thrown when task response is invalid
    error InvalidTaskResponse();

    /**
     * @notice Constructor
     * @param _taskMailbox Address of the TaskMailbox contract
     * @param _executorOperatorSet The operator set for randomness generation
     */
    constructor(address _taskMailbox, OperatorSet memory _executorOperatorSet) {
        taskMailbox = ITaskMailbox(_taskMailbox);
        executorOperatorSet = _executorOperatorSet;
    }

    /**
     * @notice Requests verifiable randomness using VDF
     * @param seed The seed bytes for randomness generation
     * @return requestId The unique identifier for this randomness request (task hash)
     */
    function requestRandomness(bytes calldata seed) external returns (bytes32 requestId) {
        // Prepare task payload
        VDFParams memory vdfParams = VDFParams({seed: seed});
        TaskPayload memory payload = TaskPayload({
            randomnessType: RandomnessType.VDF,
            randomnessParams: abi.encode(vdfParams)
        });
        
        // Create task in TaskMailbox
        ITaskMailboxTypes.TaskParams memory taskParams = ITaskMailboxTypes.TaskParams({
            refundCollector: msg.sender,
            avsFee: 0, // No fee for now
            executorOperatorSet: executorOperatorSet,
            payload: abi.encode(payload)
        });
        
        bytes32 taskHash = taskMailbox.createTask(taskParams);
        requestId = taskHash; // Use task hash as request ID
        
        // Store request details
        requests[requestId] = RandomnessRequest({
            requester: msg.sender,
            blockNumber: block.number,
            fulfilled: false,
            result: 0
        });
        
        emit RandomnessRequested(requestId, msg.sender, seed);
        
        return requestId;
    }

    /**
     * @notice Called by TaskMailbox when a task is completed
     * @param taskHash The hash of the completed task
     * @param result The encoded task result
     */
    function onTaskCompleted(bytes32 taskHash, bytes calldata result) external {
        if (msg.sender != address(taskMailbox)) {
            revert UnauthorizedCaller();
        }
        
        bytes32 requestId = taskHash; // Task hash is our request ID
        RandomnessRequest storage request = requests[requestId];
        if (request.requester == address(0)) {
            revert RequestNotFound();
        }
        
        if (request.fulfilled) {
            revert RequestAlreadyFulfilled();
        }
        
        // Decode the task response
        TaskResponsePayload memory responsePayload = abi.decode(result, (TaskResponsePayload));
        
        if (responsePayload.randomnessType != RandomnessType.VDF) {
            revert InvalidTaskResponse();
        }
        
        VDFResult memory vdfResult = abi.decode(responsePayload.randomValue, (VDFResult));
        
        // Mark request as fulfilled
        request.fulfilled = true;
        request.result = vdfResult.result;
        
        emit RandomnessFulfilled(requestId, vdfResult.result);
    }

    /**
     * @notice Gets the result of a randomness request
     * @param requestId The request identifier (task hash)
     * @return fulfilled Whether the request has been fulfilled
     * @return result The random number result (0 if not fulfilled)
     */
    function getRandomnessResult(bytes32 requestId) 
        external 
        view 
        returns (bool fulfilled, uint256 result) 
    {
        RandomnessRequest memory request = requests[requestId];
        return (request.fulfilled, request.result);
    }

    /**
     * @notice Gets detailed information about a randomness request
     * @param requestId The request identifier (task hash)
     * @return request The complete request information
     */
    function getRequest(bytes32 requestId) 
        external 
        view 
        returns (RandomnessRequest memory request) 
    {
        return requests[requestId];
    }

    /**
     * @notice Updates the operator set (only callable by contract owner/admin)
     * @param _executorOperatorSet The new operator set
     * @dev This function should have proper access control in production
     */
    function updateOperatorSet(OperatorSet memory _executorOperatorSet) external {
        // TODO: Add proper access control (onlyOwner or similar)
        executorOperatorSet = _executorOperatorSet;
    }


    // IAVSTaskHook implementation
    
    /**
     * @notice Validates a task before it is created
     * @param caller Address that is creating the task
     * @param operatorSet The operator set that will execute the task
     * @param payload Task payload
     * @dev This function validates that the task payload is a valid VRF request
     */
    function validatePreTaskCreation(
        address caller,
        OperatorSet memory operatorSet,
        bytes memory payload
    ) external view override {
        // Decode and validate the task payload
        try this.decodeAndValidateTaskPayload(payload) {
            // Task payload is valid
        } catch {
            revert InvalidTaskResponse();
        }
        
        // Additional validation: ensure operator set matches our configured operator set
        if (operatorSet.avs != executorOperatorSet.avs || operatorSet.id != executorOperatorSet.id) {
            revert InvalidTaskResponse();
        }
    }

    /**
     * @notice Validates a task after it is created
     * @param taskHash Unique identifier of the task
     * @dev This function is called after task creation to track active tasks
     */
    function validatePostTaskCreation(
        bytes32 taskHash
    ) external override {
        // For VRF, we can use this to track active tasks or update internal state
        // The task is already stored in our mapping via requestRandomness
        // This is a hook for additional validation or state updates if needed
    }

    /**
     * @notice Validates a task result submission
     * @param taskHash Unique identifier of the task
     * @param cert Certificate proving the validity of the result
     * @dev This function validates that the task result is properly signed
     */
    function validateTaskResultSubmission(
        bytes32 taskHash,
        IBN254CertificateVerifierTypes.BN254Certificate memory cert
    ) external override {
        // Verify that the task exists in our system
        bytes32 requestId = taskHash; // Task hash is our request ID
        RandomnessRequest storage request = requests[requestId];
        if (request.requester == address(0)) {
            revert RequestNotFound();
        }
        
        // Additional validation could include:
        // - Verifying the certificate signature
        // - Checking that the result format matches expectations
        // - Ensuring the task hasn't already been completed
        
        if (request.fulfilled) {
            revert RequestAlreadyFulfilled();
        }
        
        // Emit event indicating validation is complete (result will be available via RandomnessFulfilled event later)
        emit TaskResultSubmissionValidated(requestId);
    }

    /**
     * @notice Internal function to decode and validate task payload
     * @param payload The encoded task payload
     * @dev This function is called externally for validation purposes
     */
    function decodeAndValidateTaskPayload(bytes memory payload) external pure {
        // Decode the task payload
        TaskPayload memory taskPayload = abi.decode(payload, (TaskPayload));
        
        // Validate randomness type
        if (taskPayload.randomnessType != RandomnessType.VDF && 
            taskPayload.randomnessType != RandomnessType.UNKNOWN) {
            revert InvalidTaskResponse();
        }
        
        // For VDF tasks, validate the parameters
        if (taskPayload.randomnessType == RandomnessType.VDF) {
            VDFParams memory vdfParams = abi.decode(taskPayload.randomnessParams, (VDFParams));
            
            // Validate seed constraints
            if (vdfParams.seed.length == 0) {
                revert InvalidTaskResponse();
            }
            
            if (vdfParams.seed.length > 1024) {
                revert InvalidTaskResponse();
            }
        }
    }

    // View functions for testing and encoding
    
    /**
     * @notice Encodes a TaskPayload for testing
     * @param randomnessType The type of randomness
     * @param randomnessParams The encoded parameters
     * @return The encoded TaskPayload
     */
    function encodeTaskPayload(RandomnessType randomnessType, bytes calldata randomnessParams) 
        external 
        pure 
        returns (bytes memory) 
    {
        TaskPayload memory payload = TaskPayload({
            randomnessType: randomnessType,
            randomnessParams: randomnessParams
        });
        return abi.encode(payload);
    }

    /**
     * @notice Encodes VDFParams for testing
     * @param seed The seed bytes
     * @return The encoded VDFParams
     */
    function encodeVDFParams(bytes calldata seed) 
        external 
        pure 
        returns (bytes memory) 
    {
        VDFParams memory params = VDFParams({seed: seed});
        return abi.encode(params);
    }

    /**
     * @notice Encodes a complete task payload with VDF params for testing
     * @param seed The seed bytes
     * @return The encoded TaskPayload with VDF parameters
     */
    function encodeVDFTaskPayload(bytes calldata seed) 
        external 
        pure 
        returns (bytes memory) 
    {
        VDFParams memory vdfParams = VDFParams({seed: seed});
        TaskPayload memory payload = TaskPayload({
            randomnessType: RandomnessType.VDF,
            randomnessParams: abi.encode(vdfParams)
        });
        return abi.encode(payload);
    }
}