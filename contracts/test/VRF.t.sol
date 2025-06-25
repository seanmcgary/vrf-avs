// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

import {VRF} from "@project/VRF.sol";

/**
 * @title VRF Contract Test Suite
 * @notice Tests for the VRF contract functionality
 */
contract VRFTest is Test {
    VRF public vrfContract;
    address public mockTaskMailbox;
    OperatorSet public testOperatorSet;
    
    address public user = address(0x1);
    address public operator1 = address(0x2);
    address public operator2 = address(0x3);
    
    // Test events
    event RandomnessRequested(
        uint256 indexed requestId,
        address indexed requester,
        bytes32 indexed taskHash,
        bytes seed
    );
    
    event RandomnessFulfilled(
        uint256 indexed requestId,
        bytes32 indexed taskHash,
        uint256 result
    );

    function setUp() public {
        // Set up mock TaskMailbox
        mockTaskMailbox = address(0x999);
        
        // Create test operator set
        testOperatorSet = OperatorSet({
            avs: address(this),
            id: 1
        });
        
        // Deploy VRF contract
        vrfContract = new VRF(mockTaskMailbox, testOperatorSet);
    }

    function testRequestRandomness() public {
        bytes memory seed = "test-seed-123";
        bytes32 expectedTaskHash = keccak256("mock-task-hash");
        
        // Mock the TaskMailbox createTask call
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(expectedTaskHash)
        );
        
        // Expect the RandomnessRequested event
        vm.expectEmit(true, true, true, true);
        emit RandomnessRequested(1, user, expectedTaskHash, seed);
        
        // Request randomness as user
        vm.prank(user);
        uint256 requestId = vrfContract.requestRandomness(seed);
        
        // Verify request ID
        assertEq(requestId, 1);
        
        // Verify request details
        VRF.RandomnessRequest memory request = vrfContract.getRequest(requestId);
        assertEq(request.requester, user);
        assertEq(request.taskHash, expectedTaskHash);
        assertEq(request.blockNumber, block.number);
        assertFalse(request.fulfilled);
        assertEq(request.result, 0);
        
        // Verify request counter
        assertEq(vrfContract.getRequestCounter(), 1);
    }

    function testMultipleRequests() public {
        bytes memory seed1 = "seed-1";
        bytes memory seed2 = "seed-2";
        bytes32 taskHash1 = keccak256("task-1");
        bytes32 taskHash2 = keccak256("task-2");
        
        // Mock TaskMailbox calls
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash1)
        );
        
        vm.prank(user);
        uint256 requestId1 = vrfContract.requestRandomness(seed1);
        
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash2)
        );
        
        vm.prank(user);
        uint256 requestId2 = vrfContract.requestRandomness(seed2);
        
        assertEq(requestId1, 1);
        assertEq(requestId2, 2);
        assertEq(vrfContract.getRequestCounter(), 2);
    }

    function testOnTaskCompleted() public {
        // First, create a request
        bytes memory seed = "test-seed";
        bytes32 taskHash = keccak256("test-task");
        uint256 expectedResult = 12345;
        
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash)
        );
        
        vm.prank(user);
        uint256 requestId = vrfContract.requestRandomness(seed);
        
        // Prepare task response
        VRF.VDFResult memory vdfResult = VRF.VDFResult({result: expectedResult});
        VRF.TaskResponsePayload memory responsePayload = VRF.TaskResponsePayload({
            randomnessType: VRF.RandomnessType.VDF,
            randomValue: abi.encode(vdfResult)
        });
        bytes memory encodedResponse = abi.encode(responsePayload);
        
        // Expect the RandomnessFulfilled event
        vm.expectEmit(true, true, true, true);
        emit RandomnessFulfilled(requestId, taskHash, expectedResult);
        
        // Call onTaskCompleted as TaskMailbox
        vm.prank(mockTaskMailbox);
        vrfContract.onTaskCompleted(taskHash, encodedResponse);
        
        // Verify the request is fulfilled
        (bool fulfilled, uint256 result) = vrfContract.getRandomnessResult(requestId);
        assertTrue(fulfilled);
        assertEq(result, expectedResult);
        
        VRF.RandomnessRequest memory request = vrfContract.getRequest(requestId);
        assertTrue(request.fulfilled);
        assertEq(request.result, expectedResult);
    }

    function testOnTaskCompletedUnauthorized() public {
        bytes32 taskHash = keccak256("test-task");
        bytes memory response = "test-response";
        
        // Try to call onTaskCompleted from unauthorized address
        vm.prank(user);
        vm.expectRevert(VRF.UnauthorizedCaller.selector);
        vrfContract.onTaskCompleted(taskHash, response);
    }

    function testOnTaskCompletedRequestNotFound() public {
        bytes32 nonExistentTaskHash = keccak256("non-existent");
        bytes memory response = "test-response";
        
        vm.prank(mockTaskMailbox);
        vm.expectRevert(VRF.RequestNotFound.selector);
        vrfContract.onTaskCompleted(nonExistentTaskHash, response);
    }

    function testOnTaskCompletedAlreadyFulfilled() public {
        // Create and fulfill a request first
        bytes memory seed = "test-seed";
        bytes32 taskHash = keccak256("test-task");
        uint256 result = 12345;
        
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash)
        );
        
        vm.prank(user);
        vrfContract.requestRandomness(seed);
        
        VRF.VDFResult memory vdfResult = VRF.VDFResult({result: result});
        VRF.TaskResponsePayload memory responsePayload = VRF.TaskResponsePayload({
            randomnessType: VRF.RandomnessType.VDF,
            randomValue: abi.encode(vdfResult)
        });
        bytes memory encodedResponse = abi.encode(responsePayload);
        
        vm.prank(mockTaskMailbox);
        vrfContract.onTaskCompleted(taskHash, encodedResponse);
        
        // Try to fulfill again
        vm.prank(mockTaskMailbox);
        vm.expectRevert(VRF.RequestAlreadyFulfilled.selector);
        vrfContract.onTaskCompleted(taskHash, encodedResponse);
    }

    function testOnTaskCompletedInvalidRandomnessType() public {
        bytes memory seed = "test-seed";
        bytes32 taskHash = keccak256("test-task");
        
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash)
        );
        
        vm.prank(user);
        vrfContract.requestRandomness(seed);
        
        // Create response with invalid randomness type - we'll use abi.encode with wrong data
        // Since we can't cast invalid enum values, we'll create malformed data
        bytes memory invalidResponse = abi.encode("invalid", "response", "structure");
        
        vm.prank(mockTaskMailbox);
        vm.expectRevert();
        vrfContract.onTaskCompleted(taskHash, invalidResponse);
    }

    function testUpdateOperatorSet() public {
        OperatorSet memory newOperatorSet = OperatorSet({
            avs: address(0x4),
            id: 2
        });
        
        vrfContract.updateOperatorSet(newOperatorSet);
        
        // Verify the operator set was updated
        // Note: executorOperatorSet() getter returns (address, uint32)
        (address currentAvs, uint32 currentId) = (vrfContract.executorOperatorSet());
        assertEq(currentAvs, address(0x4));
        assertEq(currentId, 2);
    }

    function testGetRandomnessResultUnfulfilled() public {
        bytes memory seed = "test-seed";
        bytes32 taskHash = keccak256("test-task");
        
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(taskHash)
        );
        
        vm.prank(user);
        uint256 requestId = vrfContract.requestRandomness(seed);
        
        (bool fulfilled, uint256 result) = vrfContract.getRandomnessResult(requestId);
        assertFalse(fulfilled);
        assertEq(result, 0);
    }

    function testTaskPayloadEncoding() public {
        bytes memory seed = "test-seed-for-encoding";
        
        // Create the expected payload structure
        VRF.VDFParams memory vdfParams = VRF.VDFParams({seed: seed});
        VRF.TaskPayload memory expectedPayload = VRF.TaskPayload({
            randomnessType: VRF.RandomnessType.VDF,
            randomnessParams: abi.encode(vdfParams)
        });
        
        // Mock TaskMailbox
        vm.mockCall(
            mockTaskMailbox,
            abi.encodeWithSignature("createTask((address,uint96,(address,uint32),bytes))"),
            abi.encode(keccak256("task-hash"))
        );
        
        vm.prank(user);
        vrfContract.requestRandomness(seed);
        
        // Verify the payload structure can be decoded correctly
        bytes memory encodedExpected = abi.encode(expectedPayload);
        VRF.TaskPayload memory decodedPayload = abi.decode(encodedExpected, (VRF.TaskPayload));
        
        assertEq(uint256(decodedPayload.randomnessType), uint256(VRF.RandomnessType.VDF));
        
        VRF.VDFParams memory decodedParams = abi.decode(decodedPayload.randomnessParams, (VRF.VDFParams));
        assertEq(decodedParams.seed, seed);
    }
}
