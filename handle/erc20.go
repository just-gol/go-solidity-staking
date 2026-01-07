package handle

import (
	"go-solidity-staking/models"
	"go-solidity-staking/service"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ERC20TokenHandle struct {
	svc service.ERC20TokenService
}

func NewERC20Handler(svc service.ERC20TokenService) *ERC20TokenHandle {
	return &ERC20TokenHandle{svc: svc}
}

func (e *ERC20TokenHandle) Approve(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	spenderAddress := common.HexToAddress(ctx.PostForm("spenderAddress"))
	parseInt, err := strconv.ParseInt(ctx.PostForm("value"), 10, 64)
	if err != nil {
		models.Error(ctx, "Error parsing amount")
		return
	}
	value := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	approve, err := e.svc.Approve(ctx.Request.Context(), contractAddress, spenderAddress, privateKey, value)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, approve.Hash().Hex())
}

// Transfer
// contractAddress = ERC20 合约地址
// to = 用户地址
// privateKeyStr = 持币者（部署者）私钥
// value = 乘 1e18 后的数量
// /*
func (e *ERC20TokenHandle) Transfer(ctx *gin.Context) {
	privateKey, err := parsePrivateKey(ctx.PostForm("privateKeyStr"))
	if err != nil {
		models.Error(ctx, "Error parsing private key")
		return
	}
	contractAddress := common.HexToAddress(ctx.PostForm("contractAddress"))
	to := common.HexToAddress(ctx.PostForm("to"))
	parseInt, err := strconv.ParseInt(ctx.PostForm("value"), 10, 64)
	if err != nil {
		models.Error(ctx, "Error parsing amount")
		return
	}
	value := new(big.Int).Mul(
		big.NewInt(parseInt),
		big.NewInt(1e18),
	)
	approve, err := e.svc.Transfer(ctx.Request.Context(), contractAddress, to, privateKey, value)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, approve.Hash().Hex())
}

func (e *ERC20TokenHandle) BalanceOf(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	to := common.HexToAddress(ctx.Query("to"))
	balanceOf, err := e.svc.BalanceOf(ctx.Request.Context(), contractAddress, to)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, balanceOf)
}

func (e *ERC20TokenHandle) Allowance(ctx *gin.Context) {
	contractAddress := common.HexToAddress(ctx.Query("contractAddress"))
	ownerAddress := common.HexToAddress(ctx.Query("ownerAddress"))
	spenderAddress := common.HexToAddress(ctx.Query("spenderAddress"))
	allowance, err := e.svc.Allowance(ctx.Request.Context(), contractAddress, ownerAddress, spenderAddress)
	if err != nil {
		models.Error(ctx, err.Error())
		return
	}
	models.Success(ctx, allowance)
}
