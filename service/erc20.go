package service

import (
	"context"
	"crypto/ecdsa"
	"go-solidity-staking/gen/erc20"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ERC20TokenService interface {
	Approve(ctx context.Context, contractAddress common.Address, spenderAddress common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error)
}

type erc20TokenService struct {
	client *ethclient.Client
}

func NewERC20Token(client *ethclient.Client) ERC20TokenService {
	return &erc20TokenService{client: client}
}

func (e *erc20TokenService) Approve(ctx context.Context, contractAddress common.Address, spenderAddress common.Address, privateKey *ecdsa.PrivateKey, value *big.Int) (*types.Transaction, error) {
	client := e.client
	newErc20, err := erc20.NewErc20(contractAddress, client)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return newErc20.Approve(auth, spenderAddress, value)
}
