package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
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
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newStaking.Stake(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("stake tx: %w", err)
	}
	return tx, nil
}
func (s *stakingService) WithdrawStakedTokens(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	client := s.client
	newStaking, err := staking.NewStaking(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newStaking.WithdrawStakedTokens(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("withdraw tx: %w", err)
	}
	return tx, nil
}
func (s *stakingService) GetReward(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client := s.client
	newStaking, err := staking.NewStaking(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newStaking.GetReward(auth)
	if err != nil {
		return nil, fmt.Errorf("getReward tx: %w", err)
	}
	return tx, nil
}
func (s *stakingService) UpdateRewardRate(ctx context.Context, contractAddress common.Address, privateKey *ecdsa.PrivateKey, newRewardRate *big.Int) (*types.Transaction, error) {
	client := s.client
	newStaking, err := staking.NewStaking(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newStaking.UpdateRewardRate(auth, newRewardRate)
	if err != nil {
		return nil, fmt.Errorf("updateRewardRate tx: %w", err)
	}
	return tx, nil
}

func (s *stakingService) Earned(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.Earned(&bind.CallOpts{Context: ctx}, account)
	if err != nil {
		return nil, fmt.Errorf("earned call: %w", err)
	}
	return value, nil
}

func (s *stakingService) StakedBalance(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.StakedBalance(&bind.CallOpts{Context: ctx}, account)
	if err != nil {
		return nil, fmt.Errorf("stakedBalance call: %w", err)
	}
	return value, nil
}

func (s *stakingService) RewardPerToken(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.RewardPerToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("rewardPerToken call: %w", err)
	}
	return value, nil
}

func (s *stakingService) RewardPerTokenStored(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.RewardPerTokenStored(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("rewardPerTokenStored call: %w", err)
	}
	return value, nil
}

func (s *stakingService) RewardRate(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.RewardRate(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("rewardRate call: %w", err)
	}
	return value, nil
}

func (s *stakingService) LastUpdateTime(ctx context.Context, contractAddress common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.LastUpdateTime(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("lastUpdateTime call: %w", err)
	}
	return value, nil
}

func (s *stakingService) UserRewardPerTokenPaid(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.UserRewardPerTokenPaid(&bind.CallOpts{Context: ctx}, account)
	if err != nil {
		return nil, fmt.Errorf("userRewardPerTokenPaid call: %w", err)
	}
	return value, nil
}

func (s *stakingService) Rewards(ctx context.Context, contractAddress common.Address, account common.Address) (*big.Int, error) {
	newStaking, err := staking.NewStaking(contractAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("new staking contract: %w", err)
	}
	value, err := newStaking.Rewards(&bind.CallOpts{Context: ctx}, account)
	if err != nil {
		return nil, fmt.Errorf("rewards call: %w", err)
	}
	return value, nil
}
