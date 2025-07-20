package main

import "github.com/ethereum/go-ethereum/ethclient"

func getEthClient() *ethclient.Client {
	dial, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/" + AlchemyKey)
	if err != nil {
		panic(err)
	}
	return dial
}

func getEthWssClient() *ethclient.Client {
	dial, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/" + AlchemyKey)
	if err != nil {
		panic(err)
	}
	return dial
}
