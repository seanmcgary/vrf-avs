package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

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
	
	fmt.Printf("Requesting randomness...\n")
	fmt.Printf("  Contract: %s\n", contractAddr.Hex())
	fmt.Printf("  Seed: %s\n", string(seed))
	fmt.Printf("  From: %s\n", auth.From.Hex())

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
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0].Hex() == "0x152b260fdc6b51380aea47b7a68bd79722f0459b2603c789e9e9f476007e07be" {
			requestID := new(big.Int).SetBytes(log.Topics[1][:])
			fmt.Printf("Request ID: %s\n", requestID.String())
			break
		}
	}

	return nil
}

func checkStatus(c *cli.Context) error {
	client, _, vrfContract, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()

	requestID := new(big.Int).SetUint64(c.Uint64("request-id"))
	
	fmt.Printf("Checking status for request ID: %s\n", requestID.String())
	fmt.Printf("Contract: %s\n", contractAddr.Hex())
	
	// Get request details
	request, err := vrfContract.GetRequest(&bind.CallOpts{}, requestID)
	if err != nil {
		return fmt.Errorf("failed to get request details: %w", err)
	}

	fmt.Printf("\nRequest Details:\n")
	fmt.Printf("  Requester: %s\n", request.Requester.Hex())
	fmt.Printf("  Task Hash: 0x%x\n", request.TaskHash)
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
	requestCounter, err := vrfContract.GetRequestCounter(&bind.CallOpts{})
	if err == nil {
		fmt.Printf("Request Counter: %s\n", requestCounter.String())
	}

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