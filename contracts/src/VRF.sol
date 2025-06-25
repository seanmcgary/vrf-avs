// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

/**
 * @title VRF - Verifiable Random Function AVS Contract
 * @notice This contract provides verifiable random number generation using VDF (Verifiable Delay Function)
 * @dev Integrates with Hourglass TaskMailbox for task-based randomness generation
 */
contract VRF {
    /// @notice Types of randomness generation supported
    enum RandomnessType {
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
        bytes32 taskHash;
        uint256 blockNumber;
        bool fulfilled;
        uint256 result;
    }

    /// @notice The TaskMailbox contract for task management
    ITaskMailbox public immutable taskMailbox;
    
    /// @notice The operator set used for randomness generation
    OperatorSet public executorOperatorSet;
    
    /// @notice Counter for generating unique request IDs
    uint256 private requestCounter;
    
    /// @notice Mapping from request ID to randomness request details
    mapping(uint256 => RandomnessRequest) public requests;
    
    /// @notice Mapping from task hash to request ID
    mapping(bytes32 => uint256) public taskHashToRequestId;

    /// @notice Emitted when randomness is requested
    event RandomnessRequested(
        uint256 indexed requestId,
        address indexed requester,
        bytes32 indexed taskHash,
        bytes seed
    );

    /// @notice Emitted when randomness is fulfilled
    event RandomnessFulfilled(
        uint256 indexed requestId,
        bytes32 indexed taskHash,
        uint256 result
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
     * @return requestId The unique identifier for this randomness request
     */
    function requestRandomness(bytes calldata seed) external returns (uint256 requestId) {
        requestId = ++requestCounter;
        
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
        
        // Store request details
        requests[requestId] = RandomnessRequest({
            requester: msg.sender,
            taskHash: taskHash,
            blockNumber: block.number,
            fulfilled: false,
            result: 0
        });
        
        taskHashToRequestId[taskHash] = requestId;
        
        emit RandomnessRequested(requestId, msg.sender, taskHash, seed);
        
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
        
        uint256 requestId = taskHashToRequestId[taskHash];
        if (requestId == 0) {
            revert RequestNotFound();
        }
        
        RandomnessRequest storage request = requests[requestId];
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
        
        emit RandomnessFulfilled(requestId, taskHash, vdfResult.result);
    }

    /**
     * @notice Gets the result of a randomness request
     * @param requestId The request identifier
     * @return fulfilled Whether the request has been fulfilled
     * @return result The random number result (0 if not fulfilled)
     */
    function getRandomnessResult(uint256 requestId) 
        external 
        view 
        returns (bool fulfilled, uint256 result) 
    {
        RandomnessRequest memory request = requests[requestId];
        return (request.fulfilled, request.result);
    }

    /**
     * @notice Gets detailed information about a randomness request
     * @param requestId The request identifier
     * @return request The complete request information
     */
    function getRequest(uint256 requestId) 
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

    /**
     * @notice Gets the current request counter
     * @return The number of requests made so far
     */
    function getRequestCounter() external view returns (uint256) {
        return requestCounter;
    }
}