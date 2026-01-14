package models

import "time"

type ERC20EventApproval struct {
	ID          uint
	TxHash      string
	LogIndex    uint
	BlockNumber uint64
	Contract    string
	Owner       string
	Spender     string
	Value       string
	CreatedAt   time.Time
}

func (ERC20EventApproval) TableName() string {
	return "erc20_event_approval"
}
