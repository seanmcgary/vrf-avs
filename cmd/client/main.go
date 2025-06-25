package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
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
				Usage:    "Private key (hex string without 0x prefix)",
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

func setupClient(c *cli.Context) (*ethclient.Client, *bind.TransactOpts, common.Address, error) {
	// Connect to Ethereum client
	client, err := ethclient.Dial(c.String("rpc-url"))
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(c.String("private-key"))
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Parse contract address
	contractAddr := common.HexToAddress(c.String("contract-address"))

	return client, auth, contractAddr, nil
}

func requestRandomness(c *cli.Context) error {
	client, auth, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()

	seed := []byte(c.String("seed"))
	
	fmt.Printf("Requesting randomness...\n")
	fmt.Printf("  Contract: %s\n", contractAddr.Hex())
	fmt.Printf("  Seed: %s\n", string(seed))
	fmt.Printf("  From: %s\n", auth.From.Hex())

	// For now, we'll make a simple contract call
	// TODO: Generate Go bindings for the VRF contract
	fmt.Printf("Contract call would be made here with seed: %s\n", string(seed))
	fmt.Printf("Note: Full contract interaction requires generated Go bindings\n")

	return nil
}

func checkStatus(c *cli.Context) error {
	client, _, contractAddr, err := setupClient(c)
	if err != nil {
		return err
	}
	defer client.Close()

	requestID := c.Uint64("request-id")
	
	fmt.Printf("Checking status for request ID: %d\n", requestID)
	fmt.Printf("Contract: %s\n", contractAddr.Hex())
	
	// TODO: Add contract call to check request status
	fmt.Printf("Status check would be performed here\n")
	fmt.Printf("Note: Full contract interaction requires generated Go bindings\n")

	return nil
}

func getInfo(c *cli.Context) error {
	client, _, contractAddr, err := setupClient(c)
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
	} else {
		fmt.Printf("Status: Contract deployed (code size: %d bytes)\n", len(code))
	}

	// Get network info
	chainID, err := client.NetworkID(context.Background())
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