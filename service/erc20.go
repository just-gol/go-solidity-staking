package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-solidity-staking/gen/erc20"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ERC20TokenService interface {
	Approve(ctx context.Context, contractAddress common.Address, spenderAddress common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error)
	Transfer(ctx context.Context, contractAddress common.Address, to common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error)
	BalanceOf(ctx context.Context, contractAddress common.Address, to common.Address) (*big.Int, error)
	Allowance(ctx context.Context, contractAddress common.Address, ownerAddress common.Address, spenderAddress common.Address) (*big.Int, error)
}

type erc20TokenService struct {
	client *ethclient.Client
}

func NewERC20TokenService(client *ethclient.Client) ERC20TokenService {
	return &erc20TokenService{client: client}
}

func (e *erc20TokenService) Approve(ctx context.Context, contractAddress common.Address, spenderAddress common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	client := e.client
	newErc20, err := erc20.NewErc20(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new erc20 contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newErc20.Approve(auth, spenderAddress, value)
	if err != nil {
		return nil, fmt.Errorf("approve tx: %w", err)
	}
	return tx, nil
}
func (e *erc20TokenService) Transfer(ctx context.Context, contractAddress common.Address, to common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	client := e.client
	newErc20, err := erc20.NewErc20(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new erc20 contract: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	tx, err := newErc20.Transfer(auth, to, value)
	if err != nil {
		return nil, fmt.Errorf("transfer tx: %w", err)
	}
	return tx, nil
}
func (e *erc20TokenService) BalanceOf(ctx context.Context, contractAddress common.Address, to common.Address) (*big.Int, error) {
	client := e.client
	newErc20, err := erc20.NewErc20(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new erc20 contract: %w", err)
	}
	value, err := newErc20.BalanceOf(&bind.CallOpts{Context: ctx}, to)
	if err != nil {
		return nil, fmt.Errorf("balanceOf call: %w", err)
	}
	return value, nil
}

func (e *erc20TokenService) Allowance(ctx context.Context, contractAddress common.Address, ownerAddress common.Address, spenderAddress common.Address) (*big.Int, error) {
	client := e.client
	newErc20, err := erc20.NewErc20(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("new erc20 contract: %w", err)
	}
	value, err := newErc20.Allowance(&bind.CallOpts{Context: ctx}, ownerAddress, spenderAddress)
	if err != nil {
		return nil, fmt.Errorf("allowance call: %w", err)
	}
	return value, nil
}
