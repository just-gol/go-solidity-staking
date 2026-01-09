package service

import (
	"context"
	"encoding/json"
	"errors"
	"go-solidity-staking/gen/erc20"
	"go-solidity-staking/gen/staking"
	"go-solidity-staking/logger"
	"go-solidity-staking/models"
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
	ReplayERC20FromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error
	StartERC20ReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration)
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
		latest = latest - (confirmations - 1)
	}
	// 区块已同步
	if lastBlock > latest {
		return nil
	}
	// 回放
	return l.replayRange(ctx, contractAddress, lastBlock+1, latest)
}

func (l *listenerService) StartReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := l.ReplayFromLast(ctx, contractAddress, starkBlock, confirmations); err != nil {
				logger.WithModule("listener").WithError(err).Error("start replay loop failed")
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
	if err := l.consumeStakingStaked(s, ctx, start, &endCopy); err != nil {
		return err
	}
	if err := l.consumeStakingWithdrawn(s, ctx, start, &endCopy); err != nil {
		return err
	}
	if err := l.consumeStakingRewardsClaimed(s, ctx, start, &endCopy); err != nil {
		return err
	}
	if err := l.consumeStakingRewardRateUpdated(s, ctx, start, &endCopy); err != nil {
		return err
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

func (l *listenerService) handleWithdrawn(ev *staking.StakingWithdrawn) {
	if ev == nil {
		return
	}
	ok, err := l.recordEvent(ev.Raw, "withdrawn")
	if err != nil || !ok {
		return
	}
	_ = l.setSyncBlock(syncKey("staking", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) handleRewardsClaimed(ev *staking.StakingRewardsClaimed) {
	if ev == nil {
		return
	}
	ok, err := l.recordEvent(ev.Raw, "rewards_claimed")
	if err != nil || !ok {
		return
	}
	_ = l.setSyncBlock(syncKey("staking", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) handleRewardRateUpdated(ev *staking.StakingRewardRateUpdated) {
	if ev == nil {
		return
	}
	signature := ""
	if len(ev.Raw.Topics) > 0 {
		signature = ev.Raw.Topics[0].Hex()
	}
	indexedMap := map[string]string{
		"signature":       signature,
		"new_reward_rate": ev.NewRewardRate.String(),
	}
	ok, err := l.recordEventMap(ev.Raw, "reward_rate_updated", indexedMap)
	if err != nil || !ok {
		return
	}
	_ = l.setSyncBlock(syncKey("staking", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) ReplayERC20FromLast(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64) error {
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
		latest = latest - (confirmations - 1)
	}
	if lastBlock > latest {
		return nil
	}
	return l.replayERC20Range(ctx, contractAddress, lastBlock+1, latest)
}

func (l *listenerService) StartERC20ReplayLoop(ctx context.Context, contractAddress common.Address, starkBlock uint64, confirmations uint64, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := l.ReplayERC20FromLast(ctx, contractAddress, starkBlock, confirmations); err != nil {
				logger.WithModule("listener").WithError(err).Error("start erc20 replay loop failed")
			}
		}
	}
}

func (l *listenerService) replayERC20Range(ctx context.Context, contractAddress common.Address, start uint64, end uint64) error {
	token, err := erc20.NewErc20(contractAddress, l.client)
	if err != nil {
		return err
	}
	endCopy := end
	if err := l.consumeErc20Transfer(token, ctx, start, &endCopy); err != nil {
		return err
	}
	if err := l.consumeErc20Approval(token, ctx, start, &endCopy); err != nil {
		return err
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
	indexedMap := map[string]string{
		"from":  ev.From.Hex(),
		"to":    ev.To.Hex(),
		"value": ev.Value.String(),
	}
	return l.recordEventMap(ev.Raw, "erc20_transfer", indexedMap)
}

func (l *listenerService) handleErc20Approval(ev *erc20.Erc20Approval) {
	if ev == nil {
		return
	}
	ok, err := l.recordErc20Approval(ev)
	if err != nil || !ok {
		return
	}
	_ = l.setSyncBlock(syncKey("erc20_transfer", ev.Raw.Address), ev.Raw.BlockNumber)
}

func (l *listenerService) recordErc20Approval(ev *erc20.Erc20Approval) (bool, error) {
	indexedMap := map[string]string{
		"owner":   ev.Owner.Hex(),
		"spender": ev.Spender.Hex(),
		"value":   ev.Value.String(),
	}
	return l.recordEventMap(ev.Raw, "erc20_approval", indexedMap)
}

func (l *listenerService) recordEvent(logEntry types.Log, eventName string) (bool, error) {
	if len(logEntry.Topics) < 3 {
		return false, nil
	}
	indexedMap := map[string]string{
		"user":   common.BytesToAddress(logEntry.Topics[1].Bytes()[12:]).Hex(),
		"amount": new(big.Int).SetBytes(logEntry.Topics[2].Bytes()).String(),
	}
	return l.recordEventMap(logEntry, eventName, indexedMap)
}

func (l *listenerService) recordEventMap(logEntry types.Log, eventName string, indexedMap map[string]string) (bool, error) {
	signature := ""
	if len(logEntry.Topics) > 0 {
		signature = logEntry.Topics[0].Hex()
	}
	indexedMap["signature"] = signature
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

func (l *listenerService) consumeStakingStaked(s *staking.Staking, ctx context.Context, start uint64, end *uint64) error {
	iter, err := s.FilterStaked(&bind.FilterOpts{Start: start, End: end, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleStaked(iter.Event)
	}
	return nil
}

func (l *listenerService) consumeStakingWithdrawn(s *staking.Staking, ctx context.Context, start uint64, end *uint64) error {
	iter, err := s.FilterWithdrawn(&bind.FilterOpts{Start: start, End: end, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleWithdrawn(iter.Event)
	}
	return nil
}

func (l *listenerService) consumeStakingRewardsClaimed(s *staking.Staking, ctx context.Context, start uint64, end *uint64) error {
	iter, err := s.FilterRewardsClaimed(&bind.FilterOpts{Start: start, End: end, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleRewardsClaimed(iter.Event)
	}
	return nil
}

func (l *listenerService) consumeStakingRewardRateUpdated(s *staking.Staking, ctx context.Context, start uint64, end *uint64) error {
	iter, err := s.FilterRewardRateUpdated(&bind.FilterOpts{Start: start, End: end, Context: ctx})
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleRewardRateUpdated(iter.Event)
	}
	return nil
}

func (l *listenerService) consumeErc20Transfer(token *erc20.Erc20, ctx context.Context, start uint64, end *uint64) error {
	iter, err := token.FilterTransfer(&bind.FilterOpts{Start: start, End: end, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleErc20Transfer(iter.Event)
	}
	return nil
}

func (l *listenerService) consumeErc20Approval(token *erc20.Erc20, ctx context.Context, start uint64, end *uint64) error {
	iter, err := token.FilterApproval(&bind.FilterOpts{Start: start, End: end, Context: ctx}, nil, nil)
	if err != nil {
		return err
	}
	defer iter.Close()
	for iter.Next() {
		l.handleErc20Approval(iter.Event)
	}
	return nil
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
