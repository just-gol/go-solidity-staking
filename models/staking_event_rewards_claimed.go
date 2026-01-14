package models

import "time"

type StakingEventRewardsClaimed struct {
	ID          uint
	TxHash      string
	LogIndex    uint
	BlockNumber uint64
	Contract    string
	User        string
	Amount      string
	CreatedAt   time.Time
}

func (StakingEventRewardsClaimed) TableName() string {
	return "staking_event_rewards_claimed"
}
