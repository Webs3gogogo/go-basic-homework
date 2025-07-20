package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

//func main() {
//	subscribe()
//}

func subscribe() {
	client := getEthWssClient()
	//创建一个只接收区块头的通道
	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case errors := <-sub.Err():
			log.Fatal(errors)
		case header := <-headers:
			fmt.Println("New block header received:", header.Number.String(), " Hash:", header.Hash().String())
		}
	}

}
