package bootstrap

import (
	"context"
	"go-solidity-staking/handle"
	"go-solidity-staking/logger"
	"go-solidity-staking/routers"
	"go-solidity-staking/service"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func NewApp() (*gin.Engine, error) {
	logger.Init()
	config, err := ini.Load("./config/staking.ini")
	if err != nil {
		logger.WithModule("bootstrap").WithError(err).Error("load config failed")
		return nil, err
	}
	wsClient, err := ethclient.Dial(config.Section("url").Key("ws_url").String())
	if err != nil {
		logger.WithModule("bootstrap").WithError(err).Error("dial ws failed")
		return nil, err
	}
	listenerService := service.NewListenerService(wsClient)
	contractAddress := common.HexToAddress(config.Section("eth").Key("contract_address").String())
	stakingTokenAddressStr := config.Section("eth").Key("staking_token").String()
	stakingTokenAddress := common.HexToAddress(stakingTokenAddressStr)
	rewardTokenAddressStr := config.Section("eth").Key("reward_token").String()
	rewardTokenAddress := common.HexToAddress(rewardTokenAddressStr)
	rpcClient, err := ethclient.Dial(config.Section("url").Key("rpc_url").String())
	if err != nil {
		logger.WithModule("bootstrap").WithError(err).Error("dial rpc failed")
		return nil, err
	}
	// 质押
	stakingService := service.NewStakingService(rpcClient)
	stakingHandle := handle.NewStakingHandle(stakingService)

	//ERC20
	tokenService := service.NewERC20TokenService(rpcClient)
	tokenHandle := handle.NewERC20Handler(tokenService)

	go func() {
		// 调用区块链回放
		if err := listenerService.ReplayFromLast(
			context.Background(),
			contractAddress,
			config.Section("eth").Key("start_block").MustUint64(0),
			config.Section("eth").Key("confirmations").MustUint64(1),
		); err != nil {
			logger.WithModule("listener").WithError(err).Error("replay from last failed")
			return
		}
		listenerService.StartReplayLoop(
			context.Background(),
			contractAddress,
			config.Section("eth").Key("start_block").MustUint64(0),
			config.Section("eth").Key("confirmations").MustUint64(1),
			time.Duration(config.Section("eth").Key("interval").MustUint64(1))*time.Second,
		)
	}()
	funcERC20(stakingTokenAddressStr, stakingTokenAddress, listenerService, config)
	funcERC20(rewardTokenAddressStr, rewardTokenAddress, listenerService, config)
	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r, stakingHandle, tokenHandle)
	return r, nil
}

func funcERC20(addressStr string, tokenAddress common.Address, listenerService service.ListenerService, config *ini.File) {
	if addressStr != "" && tokenAddress != (common.Address{}) {
		go func() {
			if err := listenerService.ReplayERC20FromLast(
				context.Background(),
				tokenAddress,
				config.Section("eth").Key("start_block").MustUint64(0),
				config.Section("eth").Key("confirmations").MustUint64(1),
			); err != nil {
				logger.WithModule("listener").WithError(err).Error("replay erc20 from last failed")
				return
			}
			listenerService.StartERC20ReplayLoop(
				context.Background(),
				tokenAddress,
				config.Section("eth").Key("start_block").MustUint64(0),
				config.Section("eth").Key("confirmations").MustUint64(1),
				time.Duration(config.Section("eth").Key("interval").MustUint64(1))*time.Second,
			)
		}()
	}
}
