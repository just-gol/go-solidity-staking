package handle

import (
	"crypto/ecdsa"
	"go-solidity-staking/logger"
	"go-solidity-staking/models"
	"go-solidity-staking/service"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logger.WithModule("api").WithError(err).Error("stake parse amount failed")
		models.Error(ctx, "Error parsing amount")
		return
	}
	amount := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "stake",
		"contract": contractAddress.Hex(),
		"amount":   amount.String(),
	}).Info("stake request")
	stake, err := s.svc.Stake(ctx.Request.Context(), contractAddress, privateKey, amount)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("stake failed")
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
		logger.WithModule("api").WithError(err).Error("withdraw parse amount failed")
		models.Error(ctx, "Error parsing amount")
		return
	}
	amount := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "withdraw",
		"contract": contractAddress.Hex(),
		"amount":   amount.String(),
	}).Info("withdraw request")
	withdraw, err := s.svc.WithdrawStakedTokens(ctx.Request.Context(), contractAddress, privateKey, amount)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("withdraw failed")
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
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "reward",
		"contract": contractAddress.Hex(),
	}).Info("reward request")
	getReward, err := s.svc.GetReward(ctx.Request.Context(), contractAddress, privateKey)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("get reward failed")
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
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "update reward rate",
		"contract": contractAddress.Hex(),
	}).Info("update reward rate request")
	updateRewardRate, err := s.svc.UpdateRewardRate(ctx.Request.Context(), contractAddress, privateKey, new(big.Int).SetInt64(parseInt))
	if err != nil {
		logger.WithModule("api").WithError(err).Error("update reward rate failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, updateRewardRate.Hash().Hex())
}

func (s *StakingHandle) Earned(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "earned",
		"contract": contractAddress.Hex(),
		"account":  account.Hex(),
	}).Info("earned request")
	earned, err := s.svc.Earned(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("earned failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, earned)
}

func (s *StakingHandle) StakedBalance(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "staked_balance",
		"contract": contractAddress.Hex(),
		"account":  account.Hex(),
	}).Info("staked balance request")
	balance, err := s.svc.StakedBalance(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("staked balance failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, balance)
}

func (s *StakingHandle) RewardPerToken(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "reward_per_token",
		"contract": contractAddress.Hex(),
	}).Info("reward per token request")
	value, err := s.svc.RewardPerToken(ctx.Request.Context(), contractAddress)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("reward per token failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) RewardPerTokenStored(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "reward_per_token_stored",
		"contract": contractAddress.Hex(),
	}).Info("reward per token stored request")
	value, err := s.svc.RewardPerTokenStored(ctx.Request.Context(), contractAddress)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("reward_per_token_stored call failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) RewardRate(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "reward_rate",
		"contract": contractAddress.Hex(),
	}).Info("reward rate request")
	value, err := s.svc.RewardRate(ctx.Request.Context(), contractAddress)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("reward rate failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) LastUpdateTime(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "last_update_time",
		"contract": contractAddress.Hex(),
	}).Info("last update time request")
	value, err := s.svc.LastUpdateTime(ctx.Request.Context(), contractAddress)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("last update time failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) UserRewardPerTokenPaid(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "user_reward_per_token_paid",
		"contract": contractAddress.Hex(),
		"account":  account.Hex(),
	}).Info("user reward per token paid request")
	value, err := s.svc.UserRewardPerTokenPaid(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("user reward per token paid failed")
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, value)
}

func (s *StakingHandle) Rewards(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	account := common.HexToAddress(ctx.Query("account"))
	logger.WithModule("api").WithFields(logrus.Fields{
		"action":   "rewards",
		"contract": contractAddress.Hex(),
		"account":  account.Hex(),
	}).Info("rewards request")
	value, err := s.svc.Rewards(ctx.Request.Context(), contractAddress, account)
	if err != nil {
		logger.WithModule("api").WithError(err).Error("rewards failed")
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
