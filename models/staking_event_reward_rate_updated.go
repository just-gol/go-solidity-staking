package models

import "time"

type StakingEventRewardRateUpdated struct {
	ID            uint
	TxHash        string
	LogIndex      uint
	BlockNumber   uint64
	Contract      string
	NewRewardRate string
	CreatedAt     time.Time
}

func (StakingEventRewardRateUpdated) TableName() string {
	return "staking_event_reward_rate_updated"
}
