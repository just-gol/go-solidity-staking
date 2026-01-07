package service

import (
	"context"
	"crypto/ecdsa"
	"go-solidity-staking/gen/staking"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type StakingService interface {
	Stake(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error)
	WithdrawStakedTokens(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error)
	GetReward(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error)
	UpdateRewardRate(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, newRewardRate *big.Int) (*types.Transaction, error)
	Earned(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error)
	StakedBalance(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error)
	RewardPerToken(ctx context.Context, contractAddress common.Address) (*big.Int, error)
	RewardPerTokenStored(ctx context.Context, contractAddress common.Address) (*big.Int, error)
	RewardRate(ctx context.Context, contractAddress common.Address) (*big.Int, error)
	LastUpdateTime(ctx context.Context, contractAddress common.Address) (*big.Int, error)
	UserRewardPerTokenPaid(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error)
	Rewards(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error)
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
func (s *stakingService) WithdrawStakedTokens(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
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
	return newStaking.WithdrawStakedTokens(auth, amount)
}
func (s *stakingService) GetReward(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
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
	return newStaking.GetReward(auth)
}
func (s *stakingService) UpdateRewardRate(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, newRewardRate *big.Int) (*types.Transaction, error) {
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
	return newStaking.UpdateRewardRate(auth, newRewardRate)
}

func (s *stakingService) Earned(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.Earned(&bind.CallOpts{Context: ctx}, account)
}

func (s *stakingService) StakedBalance(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.StakedBalance(&bind.CallOpts{Context: ctx}, account)
}

func (s *stakingService) RewardPerToken(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.RewardPerToken(&bind.CallOpts{Context: ctx})
}

func (s *stakingService) RewardPerTokenStored(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.RewardPerTokenStored(&bind.CallOpts{Context: ctx})
}

func (s *stakingService) RewardRate(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.RewardRate(&bind.CallOpts{Context: ctx})
}

func (s *stakingService) LastUpdateTime(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.LastUpdateTime(&bind.CallOpts{Context: ctx})
}

func (s *stakingService) UserRewardPerTokenPaid(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.UserRewardPerTokenPaid(&bind.CallOpts{Context: ctx}, account)
}

func (s *stakingService) Rewards(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, err
	}
	return newStaking.Rewards(&bind.CallOpts{Context: ctx}, account)
}
