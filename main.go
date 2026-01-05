package main

import "go-solidity-staking/bootstrap"

func main() {
	app, err := bootstrap.NewApp()
	if err != nil {
		panic(err)
	}
	_ = app.Run(":8080")
}
