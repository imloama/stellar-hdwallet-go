package hdwallet

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/stellar/go/keypair"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"io"
)

type HDWallet struct {
	Mnemonic string `json:"mnemonic "` // mnemonic
}

func Random()*HDWallet{
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)


	// https://github.com/stellar/stellar-protocol/blob/master/ecosystem/sep-0005.md
	// m/44'/148'/0'
	child1, _ := masterKey.NewChildKey(0x8000002c) //44
	child2, _ := child1.NewChildKey(0x80000094)//148

	//e0eec84fe165cd427cb7bc9b6cfdef0555aa1cb6f9043ff1fe986c3c8ddd22e3
	//ss, _ := child2.Serialize()
	fmt.Printf("m/44'/148' key: %s\n" ,hex.EncodeToString(child2.Key))
	child3, _ := child2.NewChildKey(0)

	fmt.Println("stellar private key: ", child3)


	var rawSeed [32]byte

	reader := bytes.NewReader(child3.Key)
	_, err := io.ReadFull(reader, rawSeed[:])
	if err!=nil{
		fmt.Printf("error : %s\n", err)
	}

	kp, _ := keypair.FromRawSeed(rawSeed)

	fmt.Printf("stellar accountid:%s\n", kp.Address())
	fmt.Printf("stellar secret:%s\n", kp.Seed())


	return &HDWallet{}
}

func FromMnemonic(mnemonic string){
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)


	// https://github.com/stellar/stellar-protocol/blob/master/ecosystem/sep-0005.md
	// m/44'/148'/0'
	child1, _ := masterKey.NewChildKey(44)//0x8000002c) //44
	child2, _ := child1.NewChildKey(148)//0x80000094)//148

	//e0eec84fe165cd427cb7bc9b6cfdef0555aa1cb6f9043ff1fe986c3c8ddd22e3
	//ss, _ := child2.Serialize()
	fmt.Printf("m/44'/148' key: %s\n" ,hex.EncodeToString(child2.Key))
	child3, _ := child2.NewChildKey(0)

	fmt.Println("stellar seed key: ", hex.EncodeToString(child3.Key))


	var rawSeed [32]byte

	reader := bytes.NewReader(child3.Key)
	_, err := io.ReadFull(reader, rawSeed[:])
	if err!=nil{
		fmt.Printf("error : %s\n", err)
	}

	kp, _ := keypair.FromRawSeed(rawSeed)

	fmt.Printf("stellar accountid:%s\n", kp.Address())
	fmt.Printf("stellar secret:%s\n", kp.Seed())

}