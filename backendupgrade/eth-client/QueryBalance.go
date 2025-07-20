package main

import (
	"context"
	"eth-client/token"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

//func main() {
//	//queryEthBalance()
//	//queryPendingEth()
//	queryERC20Balance()
//}

func queryEthBalance() {
	client := getEthClient()
	address := getAccountAddress()
	// if pass block number as nil, it will query the latest block
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Balance of address", address.Hex(), "in wei is", balance)                   // in wei
	fmt.Println("Balance of address", address.Hex(), "in ether is", convertWei2Eth(balance)) // in ether
}

func queryPendingEth() {
	client := getEthClient()
	address := getAccountAddress()
	at, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Pending Balance of address", address.Hex(), "in wei is", at)                   // in wei
	fmt.Println("Pending Balance of address", address.Hex(), "in ether is", convertWei2Eth(at)) // in ether
}

func convertWei2Eth(wei *big.Int) string {
	fBalance := new(big.Float)
	fBalance.SetString(wei.String())
	quo := fBalance.Quo(fBalance, big.NewFloat(1e18))
	return quo.String()
}

func queryERC20Balance() {
	client := getEthClient()
	address := getContractAddress()
	instance := getInstance(address, client)
	accountAddress := getAccountAddress()
	balance, err := instance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance of address", address.Hex(), "in ERC20 is", balance)

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("symbol: %s\n", symbol)
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		fmt.Println(decimals)
		log.Fatal(err)

	}
	fmt.Printf("decimals: %v\n", decimals)

	fBal := new(big.Float)
	fBal.SetString(balance.String())
	value := fBal.Quo(fBal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value)
}

func getInstance(address common.Address, client *ethclient.Client) *token.Token {
	instance, err := token.NewToken(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func getContractAddress() common.Address {
	var contractAddress = "0x89A479a8A2b438627DaDc5864D5b750942D7a6D8"
	tokenAddress := common.HexToAddress(contractAddress)
	return tokenAddress
}
func getAccountAddress() common.Address {
	address := common.HexToAddress("0x01e823A90a39D6A9A50F35Bcb69f8e85E9f8d361")
	return address
}
