package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

/*
*
Account1 - private：XX

Account2- private :
XX
*/
//func main() {
//	//transferToken()
//	getTx("XX")
//}

func transferEth() {
	//create eth client
	client := getEthClient()
	fmt.Println(client.ChainID(context.Background()))

	//load private key
	privateKey, err := crypto.HexToECDSA("xxxx")
	if err != nil {
		fmt.Println("Error loading private key:", err)
		return
	}

	// Get Nonce by method PendingNonceAt
	// But this method need pass public address, so we need to get public key first
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ECDSA:", ok)
		return
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nounce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		fmt.Println("Error getting nonce:", err)
		return
	}
	value := big.NewInt(100000000000)
	gasLimit := uint64(21000)           // in units
	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	//get gas price by method SuggestGasPrice

	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	recipientAddressStr := "0x37F70b4A295DEae161De2751020A3596d8A85c3a"
	// figure out the recipient address
	recipientAddress := common.HexToAddress(recipientAddressStr)
	// Create a new transaction

	tx := types.NewTransaction(nounce, recipientAddress, value, gasLimit, gasPrice, nil)
	//newTx := types.NewTx(&types.LegacyTx{
	//	Nonce:    nounce,
	//	To:       recipientAddress,
	//	Value:    value,
	//	Gas:      gasLimit,
	//	GasPrice: gasPrice,
	//	Data:     nil, // no data for a simple transfer
	//})

	// Sign the transaction with the private key、
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Send the signed transaction to the network
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	//  https://sepolia.etherscan.io/tx/0x294ec599b87d5efc716e5c845da628089074cde89f438a15702fbf23cf3e05e3
	fmt.Printf("Transaction sent: %s\n", signTx.Hash().Hex())
}

func transferToken() {
	// 代币转账不需要传递value
	value := big.NewInt(0)

	// 设置 to 地址
	toAddress := common.HexToAddress("0xc7FdDB54FD4FE0689Ead3Ce0a340415AF17a10d4")

	/// 获取合约地址
	var contractAddress = "0xEEF2f28c77BA1Fc196688C51499e5e74f92684a8"
	tokenAddress := common.HexToAddress(contractAddress)
	fmt.Println("tokenAddress: ", tokenAddress.Hex())
	transferFnSignature := []byte("transfer(address,uint256)")

	// 计算函数签名的Keccak256哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodId := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodId))

	// 左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	// 设置要转账的金额
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	// 构建交易数据
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//评估 gas limit
	client := getEthClient()
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(gasLimit)

	privateKey := getPrivateKey()

	public := privateKey.Public()
	publicECDSA := public.(*ecdsa.PublicKey)
	publicAddress := crypto.PubkeyToAddress(*publicECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	fmt.Println("nonce: ", nonce)
	price, err := client.SuggestGasPrice(context.Background())
	fmt.Println("price: ", price)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, price, data)

	id, _ := client.ChainID(context.Background())
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(id), privateKey)
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signTx.Hash().Hex())

}

func getPrivateKey() *ecdsa.PrivateKey {
	//load private key
	privateKey, err := crypto.HexToECDSA("xxxxx")
	if err != nil {
		fmt.Println("Error loading private key:", err)
		panic(err)
	}
	return privateKey
}

func getTx(tx string) {
	client := getEthClient()
	txHash := common.HexToHash(tx)
	transaction, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction: ", transaction)
	fmt.Println("Is Pending: ", isPending)

	// 获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receipt: ", receipt)

	// 获取交易的区块号
	blockNumber := receipt.BlockNumber
	fmt.Println("Block Number: ", blockNumber)
}
