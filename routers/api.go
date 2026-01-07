package routers

import (
	"go-solidity-staking/handle"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine, handle *handle.StakingHandle, tokenHandle *handle.ERC20TokenHandle) {
	group := r.Group("/api")
	{
		group.POST("/stake", handle.Stake)
		group.POST("/withdrawStakedTokens", handle.WithdrawStakedTokens)
		group.GET("/getReward", handle.GetReward)
		group.POST("/updateRewardRate", handle.UpdateRewardRate)
		group.GET("/earned", handle.Earned)
		group.GET("/stakedBalance", handle.StakedBalance)
		group.GET("/rewardPerToken", handle.RewardPerToken)
		group.GET("/rewardPerTokenStored", handle.RewardPerTokenStored)
		group.GET("/rewardRate", handle.RewardRate)
		group.GET("/lastUpdateTime", handle.LastUpdateTime)
		group.GET("/userRewardPerTokenPaid", handle.UserRewardPerTokenPaid)
		group.GET("/rewards", handle.Rewards)
		group.POST("/approve", tokenHandle.Approve)
		group.POST("/transfer", tokenHandle.Transfer)
		group.GET("/balanceOf", tokenHandle.BalanceOf)
		group.GET("/allowance", tokenHandle.Allowance)
	}
}
