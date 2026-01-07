package handle

import (
	"crypto/ecdsa"
	"go-solidity-staking/models"
	"go-solidity-staking/service"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type StakingHandle struct {
	svc service.StakingService
}

func NewStakingHandle(svc service.StakingService) *StakingHandle {
	return &StakingHandle{svc: svc}
}

func (s *StakingHandle) Stake(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	parseInt, err := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
	if err != nil {
		models.Error(ctx, "Error parsing amount")
		return
	}
	amount := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	stake, err := s.svc.Stake(ctx.Request.Context(), contractAddress, privateKey, amount)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, stake.Hash().Hex())
}
func (s *StakingHandle) WithdrawStakedTokens(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	parseInt, err := strconv.ParseInt(ctx.PostForm("amount"), 10, 64)
	if err != nil {
		models.Error(ctx, "Error parsing amount")
		return
	}
	amount := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	withdraw, err := s.svc.WithdrawStakedTokens(ctx.Request.Context(), contractAddress, privateKey, amount)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, withdraw.Hash().Hex())
}

func (s *StakingHandle) GetReward(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.Query("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	getReward, err := s.svc.GetReward(ctx.Request.Context(), contractAddress, privateKey)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, getReward.Hash().Hex())
}
func (s *StakingHandle) UpdateRewardRate(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	parseInt, err := strconv.ParseInt(ctx.PostForm("newRewardRate"), 10, 64)
	if err != nil {
		models.Error(ctx, "Error parsing amount")
		return
	}
	updateRewardRate, err := s.svc.UpdateRewardRate(ctx.Request.Context(), contractAddress, privateKey, new(big.Int).SetInt64(parseInt))
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, updateRewardRate.Hash().Hex())
}

func (s *StakingHandle) Earned(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	earned, err := s.svc.Earned(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, earned)
}

func (s *StakingHandle) StakedBalance(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	balance, err := s.svc.StakedBalance(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, balance)
}

func (s *StakingHandle) RewardPerToken(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	value, err := s.svc.RewardPerToken(ctx.Request.Context(), contractAddress)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) RewardPerTokenStored(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	value, err := s.svc.RewardPerTokenStored(ctx.Request.Context(), contractAddress)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) RewardRate(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	value, err := s.svc.RewardRate(ctx.Request.Context(), contractAddress)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) LastUpdateTime(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	value, err := s.svc.LastUpdateTime(ctx.Request.Context(), contractAddress)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) UserRewardPerTokenPaid(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	value, err := s.svc.UserRewardPerTokenPaid(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) Rewards(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	value, err := s.svc.Rewards(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func parsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}
	return crypto.HexToECDSA(hexKey)
}
