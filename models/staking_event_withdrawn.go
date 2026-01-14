package models

import "time"

type StakingEventWithdrawn struct {
	ID          uint
	TxHash      string
	LogIndex    uint
	BlockNumber uint64
	Contract    string
	User        string
	Amount      string
	CreatedAt   time.Time
}

func (StakingEventWithdrawn) TableName() string {
	return "staking_event_withdrawn"
}
