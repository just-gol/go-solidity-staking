package service

import (
	"context"
	"crypto/ecdsa"
	staking "go-solidity-staking/gen"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type StakingService interface {
	Stake(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error)
}

type stakingService struct {
	client *ethclient.Client
}

func NewStakingService(client *ethclient.Client) StakingService {
	return &stakingService{client: client}
}
func (s *stakingService) Stake(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	client := s.client
	newStaking, err := staking.NewStaking(contractAddress, client)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return nil, err
	}
	return newStaking.Stake(auth, amount)
}
