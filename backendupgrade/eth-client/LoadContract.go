package main

import (
	"eth-client/store"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

//func main() {
//	loadContract()
//}

func loadContract() {
	client := getEthClient()
	address := getDeployContractAddress()
	newStore, err := store.NewStore(common.HexToAddress(address), client)
	if err != nil {
		panic(err)
	}
	version, err := newStore.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract loaded successfully!")
	fmt.Println("Contract version:", version)
}

func getDeployContractAddress() string {
	return "0x2Ee0a2b1F564941985Cb9594B12a57846b076DFB"
}

func getStore() *store.Store {
	client := getEthClient()
	address := getDeployContractAddress()
	newStore, err := store.NewStore(common.HexToAddress(address), client)
	if err != nil {
		panic(err)
	}
	return newStore
}
