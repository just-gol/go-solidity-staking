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
		models.Error(ctx, "Error staking")
		return
	}
	models.Success(ctx, stake.Hash().Hex())
}

func parsePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	if len(hexKey) >= 2 && hexKey[:2] == "0x" {
		hexKey = hexKey[2:]
	}
	return crypto.HexToECDSA(hexKey)
}
