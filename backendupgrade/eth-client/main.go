package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

const AlchemyKey string = "xxxxx"

//func main() {
//	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/" + AlchemyKey)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(client.ChainID(context.Background()))
//	//QueryBlockInfo(client)
//	//QueryTransactionInfo(client)
//	//QueryTransactionReceiptInfo(client)
//
//}

// Latest Block Number: 8735188
// header hash string :  0x3c47c380f978e1d3ffdb3fd8199841e28c451ff66370938acf9c59ae388717e5
// header hash hex :  0x3c47c380f978e1d3ffdb3fd8199841e28c451ff66370938acf9c59ae388717e5
func QueryBlockInfo(client *ethclient.Client) {
	// Query the latest block number
	// If we don't pass a block number, it defaults to the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	headerNumberStr := header.Number.String()
	fmt.Println("Latest Block Number:", headerNumberStr)
	fmt.Println("header hash string : ", header.Hash().String())
	fmt.Println("header hash hex : ", header.Hash().Hex())
	fmt.Println("header difficulty ", header.Difficulty)
	fmt.Println("header gas limit ", header.GasLimit)
	fmt.Println("header nonce ", header.Nonce)
	fmt.Println("timestamp ", header.Time)

	//use block hash to query block info
	hash, _ := client.BlockByHash(context.Background(), header.Hash())
	fmt.Println("Block Number:", hash.Number().String())

	// Query the latest block
	block, _ := client.BlockByNumber(context.Background(), nil)
	blockNumber := block.Number().String()
	fmt.Println("Latest Block:", blockNumber)
	fmt.Println("Latest Block Hash:", block.Hash().String())

	//上一个节点的hash
	fmt.Println("Latest Block Parent Hash:", block.ParentHash().String())
	fmt.Println("Latest Block Nonce:", block.Nonce())

	// Query the latest block hash
	count, _ := client.TransactionCount(context.Background(), block.Hash())
	fmt.Println("Latest Block Transaction Count:", count)
}

/*
*

	Transaction Count in Block: 195
	Transaction Hash: 0x21e80697ff1bd0f4669db73f5e19529cf06289f600182f55d7ba5a50bd10f528
	Transaction Value: 200000000000000000
	Transaction Gas Price: 100000000000
	Transaction Nonce: 812851
	Transaction Sender: 0x2CdA41645F2dBffB852a605E92B185501801FC28
	Transaction Receipt Status: 1
	Transaction Receipt Block Number: 8735188
	Transaction Receipt Block Hash: 0x3c47c380f978e1d3ffdb3fd8199841e28c451ff66370938acf9c59ae388717e5
	Transaction Receipt Gas Used: 21000
	Transaction Receipt Cumulative Gas Used: 21000
	Transaction Receipt Contract Address: 0x0000000000000000000000000000000000000000
	Transaction by Hash: 0x21e80697ff1bd0f4669db73f5e19529cf06289f600182f55d7ba5a50bd10f528
	Transaction Pending: false
	Transaction by Index Hash: 0x21e80697ff1bd0f4669db73f5e19529cf06289f600182f55d7ba5a50bd10f528
	Transaction by Index Value: 200000000000000000
*/
func QueryTransactionInfo(client *ethclient.Client) {
	chainId, _ := client.ChainID(context.Background())
	block, _ := client.BlockByNumber(context.Background(), big.NewInt(8735188))
	transactions := block.Transactions()
	fmt.Println("Transaction Count in Block:", len(transactions))
	var firstTxHash common.Hash
	for _, tx := range transactions {
		hex := tx.Hash().Hex()

		fmt.Println("Transaction Hash:", hex)
		fmt.Println("Transaction Value:", tx.Value().String())
		fmt.Println("Transaction Gas Price:", tx.GasPrice().String())
		fmt.Println("Transaction Nonce:", tx.Nonce())

		//signer := types.NewEIP155Signer(chainId)
		//sender, _ := types.Sender(signer, tx)
		//fmt.Println("Transaction Sender:", sender.Hex())
		//receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
		firstTxHash = tx.Hash()
		if sender, err := types.Sender(types.NewEIP155Signer(chainId), tx); err == nil {
			fmt.Println("Transaction Sender:", sender.Hex())
			receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
			fmt.Println("Transaction Receipt Status:", receipt.Status)
			fmt.Println("Transaction Receipt Block Number:", receipt.BlockNumber.String())
			fmt.Println("Transaction Receipt Block Hash:", receipt.BlockHash.Hex())
			fmt.Println("Transaction Receipt Gas Used:", receipt.GasUsed)
			fmt.Println("Transaction Receipt Cumulative Gas Used:", receipt.CumulativeGasUsed)
			fmt.Println("Transaction Receipt Contract Address:", receipt.ContractAddress.Hex())

		}

		break
	}

	if firstTxHash == (common.Hash{}) {
		fmt.Println("No transactions found in the block.")
		return
	}

	// 某个交易的hash
	tx, pending, err := client.TransactionByHash(context.Background(), firstTxHash)
	if err != nil {
		panic("Transaction not found")
	}
	fmt.Println("Transaction by Hash:", tx.Hash().Hex())
	fmt.Println("Transaction Pending:", pending)

	txByIndex, err1 := client.TransactionInBlock(context.Background(), block.Hash(), 0)
	if err1 != nil {
		panic("Transaction not found by index")

	}
	fmt.Println("Transaction by Index Hash:", txByIndex.Hash().Hex())
	fmt.Println("Transaction by Index Value:", txByIndex.Value().String())
}

/*
*Transaction Count in Block: 195
1
[]
0x21e80697ff1bd0f4669db73f5e19529cf06289f600182f55d7ba5a50bd10f528
0
0x0000000000000000000000000000000000000000
*/
func QueryTransactionReceiptInfo(client *ethclient.Client) {
	receipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(8735188))
	if err != nil {
		panic("Transaction receipt not found")
	}
	fmt.Println("Transaction Count in Block:", len(receipts))
	for _, receipt := range receipts {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}
}
