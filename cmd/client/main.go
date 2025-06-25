package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"

	vrfBinding "github.com/Layr-Labs/hourglass-avs-template/pkg/bindings/VRF"
)

func main() {
	app := &cli.App{
		Name:  "vrf-client",
		Usage: "VRF AVS client for testing contract interactions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "rpc-url",
				Aliases:  []string{"r"},
				Usage:    "Ethereum RPC URL",
				Required: true,
				EnvVars:  []string{"RPC_URL"},
			},
			&cli.StringFlag{
				Name:     "private-key",
				Aliases:  []string{"k"},
				Usage:    "Private key (hex string with or without 0x prefix)",
				Required: true,
				EnvVars:  []string{"PRIVATE_KEY"},
			},
			&cli.StringFlag{
				Name:     "contract-address",
				Aliases:  []string{"a"},
				Usage:    "VRF contract address",
				Required: true,
				EnvVars:  []string{"VRF_CONTRACT_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "ws-url",
				Aliases:  []string{"w"},
				Usage:    "WebSocket URL for event subscription (e.g., ws://localhost:8546)",
				Required: false,
				EnvVars:  []string{"WS_URL"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "request",
				Aliases: []string{"req"},
				Usage:   "Request randomness from VRF contract",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "seed",
						Aliases:  []string{"s"},
						Usage:    "Seed value for randomness generation",
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "wait",
						Aliases: []string{"wait-for-result"},
						Usage:   "Wait for randomness result via WebSocket subscription",
						Value:   false,
					},
				},
				Action: requestRandomness,
			},
			{
				Name:    "status",
				Aliases: []string{"st"},
				Usage:   "Check status of a randomness request",
				Flags: []cli.Flag{
					&cli.Uint64Flag{
						Name:     "request-id",
						Aliases:  []string{"id"},
						Usage:    "Request ID to check",
						Required: true,
					},
				},
				Action: checkStatus,
			},
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "Get contract information",
				Action:  getInfo,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func setupClient(c *cli.Context) (*ethclient.Client, *bind.TransactOpts, *vrfBinding.VRF, common.Address, error) {
	// Connect to Ethereum client
	client, err := ethclient.Dial(c.String("rpc-url"))
	if err != nil {
		return nil, nil, nil, common.Address{}, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Parse private key (handle both with and without 0x prefix)
	privateKeyHex := c.String("private-key")
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, nil, nil, common.Address{}, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Get chain ID from the RPC endpoint
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, nil, common.Address{}, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, nil, common.Address{}, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Parse contract address and create VRF binding
	contractAddr := common.HexToAddress(c.String("contract-address"))
	vrfContract, err := vrfBinding.NewVRF(contractAddr, client)
	if err != nil {
		return nil, nil, nil, common.Address{}, fmt.Errorf("failed to create VRF contract binding: %w", err)
	}

	return client, auth, vrfContract, contractAddr, nil
}

func requestRandomness(c *cli.Context) error {
	client, auth, vrfContract, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()

	seed := []byte(c.String("seed"))
	waitForResult := c.Bool("wait")
	
	fmt.Printf("Requesting randomness...\n")
	fmt.Printf("  Contract: %s\n", contractAddr.Hex())
	fmt.Printf("  Seed: %s\n", string(seed))
	fmt.Printf("  From: %s\n", auth.From.Hex())

	// Set up WebSocket subscription if waiting for result
	var wsClient *ethclient.Client
	var vrfWSContract *vrfBinding.VRF
	if waitForResult {
		wsURL := c.String("ws-url")
		if wsURL == "" {
			return fmt.Errorf("WebSocket URL is required when using --wait flag")
		}
		
		fmt.Printf("Setting up WebSocket connection to: %s\n", wsURL)
		wsClient, err = ethclient.Dial(wsURL)
		if err != nil {
			return fmt.Errorf("failed to connect to WebSocket: %w", err)
		}
		defer wsClient.Close()
		
		vrfWSContract, err = vrfBinding.NewVRF(contractAddr, wsClient)
		if err != nil {
			return fmt.Errorf("failed to create WebSocket contract binding: %w", err)
		}
	}

	// Call the requestRandomness function on the VRF contract
	tx, err := vrfContract.RequestRandomness(auth, seed)
	if err != nil {
		return fmt.Errorf("failed to call requestRandomness: %w", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Printf("Waiting for transaction confirmation...\n")

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	fmt.Printf("Transaction confirmed in block: %d\n", receipt.BlockNumber.Uint64())
	
	// Parse logs to get the request ID
	var requestID [32]byte
	var requestIDFound bool
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0].Hex() == "0xadbe210fbb3f1373d938d421ee71b9db071c3d86ccaf8851d6bda23e70c6b532" { // Updated event signature hash
			copy(requestID[:], log.Topics[1][:])
			fmt.Printf("Request ID: 0x%x\n", requestID)
			requestIDFound = true
			break
		}
	}

	if !requestIDFound {
		return fmt.Errorf("could not find request ID in transaction logs")
	}

	if !waitForResult {
		fmt.Printf("Randomness request submitted successfully. Use 'status' command to check completion.\n")
		return nil
	}

	// Subscribe to RandomnessFulfilled events and wait for our request
	fmt.Printf("Waiting for randomness result...\n")
	return waitForRandomnessResult(wsClient, vrfWSContract, requestID)
}

func waitForRandomnessResult(wsClient *ethclient.Client, vrfContract *vrfBinding.VRF, requestID [32]byte) error {
	// Create a context with timeout (5 minutes)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Subscribe to RandomnessFulfilled events
	eventsChan := make(chan *vrfBinding.VRFRandomnessFulfilled)
	sub, err := vrfContract.WatchRandomnessFulfilled(
		&bind.WatchOpts{Context: ctx},
		eventsChan,
		[][32]byte{requestID}, // Filter by our request ID
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to RandomnessFulfilled events: %w", err)
	}
	defer sub.Unsubscribe()

	fmt.Printf("Subscribed to RandomnessFulfilled events for request ID: 0x%x\n", requestID)

	// Wait for the event or timeout
	select {
	case event := <-eventsChan:
		fmt.Printf("\n🎉 Randomness Result Received!\n")
		fmt.Printf("  Request ID: 0x%x\n", event.RequestId)
		fmt.Printf("  Result: %s\n", event.Result.String())
		fmt.Printf("  Block Number: %d\n", event.Raw.BlockNumber)
		fmt.Printf("  Transaction Hash: %s\n", event.Raw.TxHash.Hex())
		return nil

	case err := <-sub.Err():
		return fmt.Errorf("subscription error: %w", err)

	case <-ctx.Done():
		return fmt.Errorf("timeout waiting for randomness result (5 minutes)")
	}
}

func checkStatus(c *cli.Context) error {
	client, _, vrfContract, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()

	// Convert request ID from uint64 to [32]byte
	requestIDUint := c.Uint64("request-id")
	var requestID [32]byte
	binary.BigEndian.PutUint64(requestID[24:], requestIDUint) // Put uint64 in the last 8 bytes
	
	fmt.Printf("Checking status for request ID: 0x%x\n", requestID)
	fmt.Printf("Contract: %s\n", contractAddr.Hex())
	
	// Get request details
	request, err := vrfContract.GetRequest(&bind.CallOpts{}, requestID)
	if err != nil {
		return fmt.Errorf("failed to get request details: %w", err)
	}

	fmt.Printf("\nRequest Details:\n")
	fmt.Printf("  Requester: %s\n", request.Requester.Hex())
	// TaskHash field no longer exists in the struct
	fmt.Printf("  Block Number: %s\n", request.BlockNumber.String())
	fmt.Printf("  Fulfilled: %t\n", request.Fulfilled)
	
	if request.Fulfilled {
		fmt.Printf("  Result: %s\n", request.Result.String())
	} else {
		fmt.Printf("  Result: Pending\n")
	}

	return nil
}

func getInfo(c *cli.Context) error {
	client, _, vrfContract, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()
	
	fmt.Printf("VRF Contract Information\n")
	fmt.Printf("========================\n")
	fmt.Printf("Contract Address: %s\n", contractAddr.Hex())
	
	// Get basic contract info
	code, err := client.CodeAt(context.Background(), contractAddr, nil)
	if err != nil {
		return fmt.Errorf("failed to get contract code: %w", err)
	}
	
	if len(code) == 0 {
		fmt.Printf("Status: No contract deployed at this address\n")
		return nil
	} else {
		fmt.Printf("Status: Contract deployed (code size: %d bytes)\n", len(code))
	}

	// Get VRF-specific contract info  
	// Note: GetRequestCounter method no longer exists in the updated contract
	fmt.Printf("VRF Contract is operational\n")

	operatorSet, err := vrfContract.ExecutorOperatorSet(&bind.CallOpts{})
	if err == nil {
		fmt.Printf("Executor Operator Set:\n")
		fmt.Printf("  AVS: %s\n", operatorSet.Avs.Hex())
		fmt.Printf("  ID: %d\n", operatorSet.Id)
	}

	taskMailbox, err := vrfContract.TaskMailbox(&bind.CallOpts{})
	if err == nil {
		fmt.Printf("Task Mailbox: %s\n", taskMailbox.Hex())
	}

	// Get network info
	chainID, err := client.ChainID(context.Background())
	if err == nil {
		fmt.Printf("Chain ID: %s\n", chainID.String())
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err == nil {
		fmt.Printf("Latest Block: %d\n", blockNumber)
	}

	return nil
}

// Helper function to convert private key to address
func privateKeyToAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}