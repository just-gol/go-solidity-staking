package models

import "time"

type ERC20EventTransfer struct {
	ID          uint
	TxHash      string
	LogIndex    uint
	BlockNumber uint64
	Contract    string
	From        string
	To          string
	Value       string
	CreatedAt   time.Time
}

func (ERC20EventTransfer) TableName() string {
	return "erc20_event_transfer"
}
