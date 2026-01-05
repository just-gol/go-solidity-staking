package bootstrap

import (
	"context"
	"go-solidity-staking/handle"
	"go-solidity-staking/routers"
	"go-solidity-staking/service"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func NewApp() (*gin.Engine, error) {
	config, err := ini.Load("../config.ini")
	if err != nil {
		return nil, err
	}
	wsClient, err := ethclient.Dial(config.Section("url").Key("ws_url").String())
	if err != nil {
		return nil, err
	}
	listenerService := service.NewListenerService(wsClient)
	contractAddress := common.HexToAddress(config.Section("eth").Key("contract_address").String())
	rpcClient, err := ethclient.Dial(config.Section("url").Key("rpc_url").String())
	if err != nil {
		return nil, err
	}
	stakingService := service.NewStakingService(rpcClient)
	stakingHandle := handle.NewStakingHandle(stakingService)
	go func() {
		// 调用区块链回访
		if err := listenerService.ReplayFromLast(
			context.Background(),
			contractAddress,
			config.Section("eth").Key("start_block").MustUint64(0),
			config.Section("eth").Key("confirmations").MustUint64(1),
		); err != nil {
			log.Printf("Error replaying from last: %v", err)
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
	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r, stakingHandle)
	return r, nil
}
