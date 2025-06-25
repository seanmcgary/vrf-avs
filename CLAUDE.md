# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a VRF (Verifiable Random Function) AVS built using the Hourglass framework for EigenLayer. The project implements a task-based Actively Validated Service (AVS) with both onchain and offchain components.

### Key Architecture Components

- **Performer** (`cmd/main.go`): The core AVS logic that handles task validation and execution
- **Smart Contracts** (`contracts/`): L1 and L2 contracts including TaskAVSRegistrar and AVSTaskHook
- **Configuration** (`config/`): Project and environment-specific settings
- **Hourglass Infrastructure** (`.hourglass/`): Docker compose and configuration for Aggregator/Executor
- **DevKit Scripts** (`.devkit/`): Build and deployment automation scripts

## Common Development Commands

### Building
```bash
# Build the performer binary
make build

# Install dependencies
make deps

# Build container
make build/container
```

### Testing
```bash
# Run all tests
make test

# Run Go tests with verbose output
go test ./... -v -p 1

# Run a specific test
go test ./cmd -run Test_TaskRequestPayload -v
```

### DevKit Commands (for full AVS development)
```bash
# Build the entire AVS (contracts + binaries)
devkit avs build

# Start local development network
devkit avs devnet start

# Stop development network
devkit avs devnet stop

# List running containers
devkit avs devnet list

# Simulate task execution
devkit avs call -- signature="(uint256,string)" args='(5,"hello")'
```

### Contract Development
```bash
# Navigate to contracts directory
cd contracts

# Build contracts
forge build

# Run contract tests
forge test

# Deploy contracts locally (handled by devkit)
forge script script/DeployMyContracts.s.sol
```

## Code Structure

### TaskWorker Implementation (`cmd/main.go`)
The main AVS logic is implemented in the TaskWorker struct with two key methods:
- `ValidateTask()`: Validates incoming task requests before execution
- `HandleTask()`: Contains the core AVS business logic for processing tasks

### Smart Contracts (`contracts/src/`)
- `l1-contracts/TaskAVSRegistrar.sol`: Handles operator registration on L1
- `l2-contracts/AVSTaskHook.sol`: Validates task lifecycle on L2
- `HelloWorld.sol`: Example contract for deployment testing
- `script/DeployMyContracts.s.sol`: Custom contract deployment script

### Configuration Files
- `config/config.yaml`: Project-level configuration (name, version, context)
- `config/contexts/devnet.yaml`: Development environment settings
- `.env`: Environment variables for L1/L2 fork URLs (copy from `.env.example`)

### Infrastructure Components
- `.hourglass/docker-compose.yml`: Aggregator and Executor services
- `.hourglass/config/`: Template configurations for Hourglass components
- `.devkit/scripts/`: Build, deployment, and utility scripts
- `keystores/`: Operator keystore files for local development

## Development Guidelines

### Go Code Conventions
- Use structured logging with zap logger (`tw.logger.Sugar().Infow(...)`)
- Follow dependency injection patterns for testability
- Implement proper error handling in both task validation and execution
- Keep TaskWorker methods focused on single responsibilities

### Framework Dependencies
- `github.com/Layr-Labs/hourglass-monorepo/ponos`: Core Hourglass framework
- `github.com/Layr-Labs/protocol-apis`: EigenLayer protocol APIs and protobuf types
- `go.uber.org/zap`: Structured logging

### Testing
- Run tests with `go test ./... -v -p 1` (sequential execution)
- Focus on testing both ValidateTask and HandleTask methods
- Use table-driven tests for multiple scenarios

## Development Workflow

### Initial Setup
1. Copy `.env.example` to `.env` and configure L1/L2 fork URLs
2. Ensure Docker is running for DevKit operations
3. Run `devkit avs build` to build contracts and binaries
4. Start development network with `devkit avs devnet start`

### Task Development Cycle
1. Implement AVS logic in `ValidateTask()` and `HandleTask()` methods in `cmd/main.go`
2. Add corresponding tests in `cmd/main_test.go`
3. Test locally with `make test` or `go test ./cmd -v`
4. Simulate task execution with `devkit avs call`
5. Iterate based on results

### Contract Development
1. Add custom contracts to `contracts/src/`
2. Update `contracts/script/DeployMyContracts.s.sol` for deployment
3. Test contracts with `forge test` in contracts directory
4. Rebuild with `devkit avs build` to include changes

## Important Notes

- This is an alpha-stage project and not production-ready
- The project uses the Hourglass task-based AVS framework
- All DevKit commands must be run from the project root (directory containing `config/` folder)
- Environment variables should be configured in `.env` file (see `.env.example`)
- The Hourglass framework handles Aggregator/Executor orchestration via Docker containers
- Keystores in `keystores/` directory are for local development only