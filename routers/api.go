package routers

import (
	"go-solidity-staking/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handle *handle.StakingHandle, tokenHandle *handle.ERC20TokenHandle) {
	group := r.Group("/api")
	{
		group.POST("/stake", handle.Stake)
		group.POST("/approve", tokenHandle.Approve)
		group.POST("/transfer", tokenHandle.Transfer)
		group.POST("/balanceOf", tokenHandle.BalanceOf)
		group.POST("/allowance", tokenHandle.Allowance)
	}
}
