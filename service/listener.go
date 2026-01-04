package service

import (
	"context"
	"errors"
	"go-solidity-staking/models"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type ListenerService interface {
	ReplayFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error
	StartReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration)
}

type listenerService struct {
	client *ethclient.Client
}

func NewListenerService(client *ethclient.Client) ListenerService {
	return &listenerService{client: client}
}

func (l *listenerService) ReplayFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error {
	// 读取上次同步的区块
	key := syncKey(contractAddress)
	lastBlock, err := l.getSyncBlock(key)
	if err != nil {
		return err
	}
	if lastBlock == 0 && starkBlock > 0 {
		lastBlock = starkBlock - 1
	}
	// 获取最新区块
	latestHeader, err := l.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	latest := latestHeader.Number.Uint64()
	// 区块已同步
	if lastBlock > latest {
		return nil
	}
	// 回放
	return nil
}

func (l *listenerService) StartReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration) {

}

func (l *listenerService) getSyncBlock(key string) (uint64, error) {
	var state models.SyncState
	err := models.DB.Where("name = ?", key).First(&state).Error
	if err == nil {
		return state.BlockNumber, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	return 0, err
}

// 区块地址
func syncKey(contractAddress common.Address) string {
	return "staking" + strings.ToLower(contractAddress.Hex())
}
