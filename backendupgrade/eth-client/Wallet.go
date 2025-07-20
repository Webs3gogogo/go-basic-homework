package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

//func main() {
//	createWallet()
//}

func createWallet() {

	key := createPrivateKey()

	recoveryPrivateKeyWithString("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")

	keyString := crypto.FromECDSA(key)
	fmt.Println("keyString: ", keyString)
	fmt.Println(hexutil.Encode(keyString))
	fmt.Println(hexutil.Encode(keyString)[2:])

	//Get public key from private key
	publicKey := key.Public()

	//convert public key to ECDSA type first due to the fact that crypto package uses ECDSA for Ethereum
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	pubKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("pubKeyBytes: ", pubKeyBytes)
	fmt.Println("publicKeyECDSA: ", publicKeyECDSA)
	// 0x04 is the prefix for uncompressed public keys in Ethereum
	fmt.Println("public key :", hexutil.Encode(pubKeyBytes)[4:])

	//Get wallet address from public key

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("address: ", address)
	fmt.Println("address: ", address.Hex())
	fmt.Println("address: ", address.String())

	// Following is the same as above but using crypto.PubkeyToAddress directly
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:]) // skip the first byte (0x04) . Because the first byte is not part of the address - 04
	fmt.Println("hash: ", hash)
	fmt.Println("hash: ", hexutil.Encode(hash.Sum(nil)))
	fmt.Println("address from hash: ", hexutil.Encode(hash.Sum(nil)[12:])) // last 20 bytes are the address

}

func createPrivateKey() *ecdsa.PrivateKey {
	// create a new private key ( From ethereum/go-ethereum/crypto package )
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	fmt.Println("privateKey: ", privateKey)

	return privateKey
}

func recoveryPrivateKeyWithString(key string) *ecdsa.PrivateKey {
	toECDSA, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("toECDSA: ", toECDSA)
	return toECDSA
}
