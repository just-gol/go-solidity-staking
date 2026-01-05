package routers

import (
	"go-solidity-staking/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handle *handle.StakingHandle) {
	group := r.Group("/api")
	{
		group.POST("/stake", handle.Stake)
	}
}
