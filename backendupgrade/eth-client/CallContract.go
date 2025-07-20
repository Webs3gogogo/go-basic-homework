package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
)

func main() {
	callContractByABI()
}

func callContractMethod() {
	client := getEthClient()
	store := getStore()
	privateKey := getPrivateKey()
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// 设置item
	var key [32]byte
	var value [32]byte

	copy(key[:], "demoKey")
	copy(value[:], "demoValue")

	tx, err := store.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SetItem success!")
	fmt.Println("Transaction hash:", tx.Hash().Hex())
	//查询item
	items, err := store.Items(&bind.CallOpts{
		Context: context.Background(),
	}, key)

	if err != nil {
		panic(err)
	}
	fmt.Println("Item value:", items)
	itemString := string(items[:])
	fmt.Println("Item value as string:", itemString)
	fmt.Println("is value saving in contract equals to origin value:", items == value)
}
func callContractByABI() {

	client := getEthClient()
	privateKey := getPrivateKey()
	public := privateKey.Public()
	publicKeyECDSA, _ := public.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(0)

	// 替换成需要调用的data
	abiData := getABIData()
	methodName := "setItem"
	var key [32]byte
	var value [32]byte
	copy(key[:], "demoKey2")
	copy(value[:], "demoValue2")
	//`Pack` 是 go-ethereum 的 `abi.ABI` 类型的方法，用于将合约方法名和参数编码为以太坊合约调用所需的字节数据（即 ABI 编码）。
	input, _ := abiData.Pack(methodName, key, value)

	//----------------Start: 另外一种直接call的abi的方式 --------
	//methodSignature := []byte("setItem(bytes32,bytes32)")
	//methodSelector := crypto.Keccak256(methodSignature)[:4]
	//
	//var key [32]byte
	//var value [32]byte
	//copy(key[:], []byte("demo_save_key_no_use_abi"))
	//copy(value[:], []byte("demo_save_value_no_use_abi_11111"))
	//
	//// 组合调用数据
	//var input []byte
	//input = append(input, methodSelector...)
	//input = append(input, key[:]...)
	//input = append(input, value[:]...)
	//----------------End: 另外一种直接call的abi的方式 --------

	toAddress := common.HexToAddress("0x2Ee0a2b1F564941985Cb9594B12a57846b076DFB")
	tx := types.NewTx(&types.LegacyTx{
		To:       &toAddress,
		Nonce:    nonce,
		Value:    amount,
		Gas:      300000,
		GasPrice: price,
		Data:     input,
	})
	fmt.Println(tx.Hash().Hex())
	//私钥签名
	chainId, _ := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signTx)
	fmt.Printf("Transaction sent: %s\n", signTx.Hash().Hex())
	_, err = waitForReceipt(client, signTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	callInput, err := abiData.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := toAddress
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: callInput,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unppackedData [32]byte
	// UnpackIntoInterface 用于将 ABI 编码的结果解码为 Go 语言中的数据结构。
	err = abiData.UnpackIntoInterface(&unppackedData, "items", result)
	if err != nil {
		log.Fatal(err)
	}

	// --------start : 直接call的方式 ----------------
	//result, err := client.CallContract(context.Background(), callMsg, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var unpacked [32]byte
	//copy(unpacked[:], result)
	//fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
	// --------end : 直接call的方式 ----------------

	fmt.Println("item value:", unppackedData)
	fmt.Println("is value saving in contract equals to origin value:", unppackedData == value)

	//--------------start: 直接call的方式 ----------------
	//itemsSignature := []byte("items(bytes32)")
	//itemsSelector := crypto.Keccak256(itemsSignature)[:4]
	//
	//var callInput []byte
	//callInput = append(callInput, itemsSelector...)
	//callInput = append(callInput, key[:]...)
	//
	//to := common.HexToAddress(contractAddr)
	//callMsg := ethereum.CallMsg{
	//	To:   &to,
	//	Data: callInput,
	//}
	// --------------------end: 直接call的方式 ----------------

}

func getABIData() abi.ABI {
	json, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}
	return json

}
