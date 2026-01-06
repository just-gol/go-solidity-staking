package service

import (
	"context"
	"encoding/json"
	"errors"
	"go-solidity-staking/gen/erc20"
	"go-solidity-staking/gen/staking"
	"go-solidity-staking/models"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type ListenerService interface {
	ReplayFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error
	StartReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration)
	ReplayERC20TransfersFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error
	StartERC20TransferReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration)
}

type listenerService struct {
	client *ethclient.Client
}

func NewListenerService(client *ethclient.Client) ListenerService {
	return &listenerService{client: client}
}

func (l *listenerService) ReplayFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error {
	// 读取上次同步的区块
	key := syncKey("staking", contractAddress)
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
	// 只回放到最新已确认的区块
	// 避免刚出块就被回滚导致数据错
	if confirmations > 1 && latest >= confirmations-1 {
		latest = confirmations - 1
	}
	// 区块已同步
	if lastBlock > latest {
		return nil
	}
	// 回放
	return l.replayRange(ctx, contractAddress, lastBlock+1, latest)
}

func (l *listenerService) StartReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration) {
	timer := time.NewTimer(interval)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := l.ReplayFromLast(ctx, contractAddress, starkBlock, confirmations); err != nil {
				log.Printf("start replay loop:%v", err)
			}
		}
	}
}

func (l *listenerService) replayRange(ctx context.Context, contractAddress common.Address, start uint64, end uint64) error {
	s, err := staking.NewStaking(contractAddress, l.client)
	if err != nil {
		return err
	}
	endCopy := end
	stakedIter, err := s.FilterStaked(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer stakedIter.Close()
	for stakedIter.Next() {
		l.handleStaked(stakedIter.Event)
	}
	return l.setSyncBlock(syncKey("staking", contractAddress), end)
}
func (l *listenerService) handleStaked(ev *staking.StakingStaked) {
	if ev == nil {
		return
	}
	ok, err := l.recordEvent(ev.Raw, "staked")
	if err != nil || !ok {
		return
	}
	// 更新区块高度
	_ = l.setSyncBlock(syncKey("staking", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) ReplayERC20TransfersFromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error {
	key := syncKey("erc20_transfer", contractAddress)
	lastBlock, err := l.getSyncBlock(key)
	if err != nil {
		return err
	}
	if lastBlock == 0 && starkBlock > 0 {
		lastBlock = starkBlock - 1
	}
	latestHeader, err := l.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	latest := latestHeader.Number.Uint64()
	if confirmations > 1 && latest >= confirmations-1 {
		latest = confirmations - 1
	}
	if lastBlock > latest {
		return nil
	}
	return l.replayERC20TransferRange(ctx, contractAddress, lastBlock+1, latest)
}

func (l *listenerService) StartERC20TransferReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration) {
	timer := time.NewTimer(interval)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := l.ReplayERC20TransfersFromLast(ctx, contractAddress, starkBlock, confirmations); err != nil {
				log.Printf("start erc20 transfer replay loop:%v", err)
			}
		}
	}
}

func (l *listenerService) replayERC20TransferRange(ctx context.Context, contractAddress common.Address, start uint64, end uint64) error {
	token, err := erc20.NewErc20(contractAddress, l.client)
	if err != nil {
		return err
	}
	endCopy := end
	transferIter, err := token.FilterTransfer(&bind.FilterOpts{Start: start, End: &endCopy, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer transferIter.Close()
	for transferIter.Next() {
		l.handleErc20Transfer(transferIter.Event)
	}
	return l.setSyncBlock(syncKey("erc20_transfer", contractAddress), end)
}

func (l *listenerService) handleErc20Transfer(ev *erc20.Erc20Transfer) {
	if ev == nil {
		return
	}
	ok, err := l.recordErc20Transfer(ev)
	if err != nil || !ok {
		return
	}
	_ = l.setSyncBlock(syncKey("erc20_transfer", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) recordErc20Transfer(ev *erc20.Erc20Transfer) (bool, error) {
	signature := ""
	if len(ev.Raw.Topics) > 0 {
		signature = ev.Raw.Topics[0].Hex()
	}
	indexedMap := map[string]string{
		"signature": signature,
		"from":      ev.From.Hex(),
		"to":        ev.To.Hex(),
		"value":     ev.Value.String(),
	}
	marshal, err := json.Marshal(indexedMap)
	if err != nil {
		return false, err
	}
	entry := models.EventLog{
		TxHash:      ev.Raw.TxHash.Hex(),
		LogIndex:    ev.Raw.Index,
		BlockNumber: ev.Raw.BlockNumber,
		Event:       "erc20_transfer",
		EventArgs:   string(marshal),
		Contract:    ev.Raw.Address.Hex(),
	}
	result := models.DB.Where("tx_hash=? and log_index=?", entry.TxHash, entry.LogIndex).FirstOrCreate(&entry)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
func (l *listenerService) recordEvent(logEntry types.Log, eventName string) (bool, error) {
	if len(logEntry.Topics) < 3 {
		return false, nil
	}
	signature := logEntry.Topics[0].Hex()
	address := common.BytesToAddress(logEntry.Topics[1].Bytes()[12:])
	amount := new(big.Int).SetBytes(logEntry.Topics[2].Bytes())
	indexedMap := map[string]string{
		"signature": signature,
		"user":      address.Hex(),
		"amount":    amount.String(),
	}
	marshal, err := json.Marshal(indexedMap)
	if err != nil {
		return false, err
	}
	entry := models.EventLog{
		TxHash:      logEntry.TxHash.Hex(),
		LogIndex:    logEntry.Index,
		BlockNumber: logEntry.BlockNumber,
		Event:       eventName,
		EventArgs:   string(marshal),
		Contract:    logEntry.Address.Hex(),
	}
	result := models.DB.Where("tx_hash=? and log_index=?", entry.TxHash, entry.LogIndex).FirstOrCreate(&entry)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (l *listenerService) setSyncBlock(key string, block uint64) error {
	state := models.SyncState{Name: key, BlockNumber: block}
	return models.DB.Where("name=?", key).Assign(models.SyncState{BlockNumber: block}).FirstOrCreate(&state).Error
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
func syncKey(prefix string, contractAddress common.Address) string {
	return prefix + strings.ToLower(contractAddress.Hex())
}
