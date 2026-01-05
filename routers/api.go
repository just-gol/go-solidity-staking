package routers

import (
	"go-solidity-staking/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handle *handle.StakingHandle) {
	group := r.Group("/staking")
	{
		group.POST("/stake", handle.Stake)
	}
}
