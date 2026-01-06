package main

import (
	"context"
	"fmt"
	staking "go-solidity-staking/gen"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/ini.v1"
)

func main() {
	config, err := ini.Load("./config/staking.ini")
	if err != nil {
		log.Fatalf("ini load error:%v", err)
	}
	rpcUrl := config.Section("url").Key("rpc_url").String()
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf(" ethclient.Dial error:%v", err)
	}
	privateKeyStr := config.Section("eth").Key("private_key").String()
	privateKey, err := crypto.HexToECDSA(privateKeyStr[2:])
	if err != nil {
		log.Fatalf("parses a secp256k1 private key error:%v", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainID error:%v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("NewKeyedTransactorWithChainID error:%v", err)
	}
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)
	deployStaking, transaction, _, err := staking.DeployStaking(auth, client, publicKey, publicKey, publicKey)
	if err != nil {
		log.Fatalf("deploy error:%v", err)
	}
	fmt.Printf("Deploying staking contract successfully:%s\n", deployStaking.Hex())
	fmt.Printf("Transaction Hash: %s", transaction.Hash().Hex())
}
