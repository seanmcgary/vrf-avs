# vrf-avs

As the name suggests, this project is an EigenLayer AVS built with Hourglass that provides random numbers.

To start, the way it will provide randomness is using a VDF, more specifically the implementation found at https://github.com/poanetwork/vdf.

## Contract structure

The smart contract (./contracts/src/VRF.sol) follows a similar pattern to chainlink/gelato:

A `requestRandomness` function which takes a caller-provided seed as bytes. This will queue a message into the Hourglass TaskMailbox.

The AVS will then ingest the task, generate the randomness across multiple Executors, aggregate the results with signatures and return the result in the TaskMailbox.

The VRF.sol contract needs to implement the functions found in `./contracts/src/l2-contracts/AVSTaskHook.sol` and update the TaskMailbox taskHook address to point at our contract

### Task structure

```solidity
enum RandomnessType {
    VDF
}

struct VDFParams {
    bytes seed;
}
struct VDFResult {
    uint256 result;
}

struct TaskPayload {
    RandomnessType randomnessType;
    bytes randomnessParams; // e.g. abi.encode(VDFParams({ seed: seed }))
}

struct TaskResponsePayload {
    RandomnessType randomnessType;
    bytes randomValue; // e.g. abi.encode(VDFResult({ result: vdfResult }))
}
```

## Performer logic

Built in Go, located at `./cmd/main.go`

The performer is pretty straightforward, it will:

- Decode the task payload
- Based on what the `RandomnessType` is, it will decode the `randomnessParams` and execute the appropriate logic.
- It will then encode the result in the `TaskResponsePayload` and send it back to the TaskMailbox.

Rather than shelling out to the `vdf-cli` from poanetwork/vdf, we should find a way to compile the VDF library (its in rust) to a shared object and call it directly from Go. This will allow us to avoid the overhead of spawning a new process and make the Performer more efficient.
