package main

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_API_KEY")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	log.Println("Connected to Sepolia")

	
	contractAddress := common.HexToAddress("YOUR_CONTRACT_ADDRESS")    //-----> Multisignature wallet address

	
	balance, err := client.BalanceAt(nil, contractAddress, nil)   // ----->  For Fetching balance
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	log.Printf("Balance: %s", balance.String())
}
