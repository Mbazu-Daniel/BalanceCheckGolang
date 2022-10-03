package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
}