package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/md5"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/btcsuite/btcutil/bech32"
)

// SignData signs the provided data using the private key
func SignData(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// VerifySignature verifies the provided signature against the provided data and public key
func VerifySignature(data []byte, signature []byte, publicKey *rsa.PublicKey) (bool, error) {
	hash := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		if err == rsa.ErrVerification {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GenKeys() (*rsa.PrivateKey, *rsa.PublicKey, []byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Encode the public key for sharing
	publicKey := &privateKey.PublicKey
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return privateKey, publicKey, publicKeyBytes, privateKeyBytes
}

// encodeToBinary encodes the provided value to a binary format using gob encoding
func encodeToBinary(value interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func getAddressBech32(pref string,pubB []byte) string {
	hash := md5.Sum(pubB)
	conv, err := bech32.ConvertBits(hash[:], 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	encoded, err := bech32.Encode(pref, conv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return encoded
}

func main() {
	priv, pub, pubB, _ := GenKeys()

	// Demonstrate signing and verifying a transaction
	transaction := "hiiii"
	transactionBytes, err := encodeToBinary(transaction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signature, err := SignData(transactionBytes, priv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	verifyResult, err := VerifySignature(transactionBytes, signature, pub)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Transaction verified:", verifyResult)

	fmt.Printf("bech32 addr: %v\n", getAddressBech32("Z", pubB)) // Z17zeaxfjctc9j6kms7kck5nurx5026guh
	fmt.Printf("len %v\n", len(getAddressBech32("Z", pubB))) // 34
	// fmt.Printf("pubSize %v\n", len(pubB))
	// fmt.Printf("addrSize %v\n", len([]byte(getAddressBech32("Z", pubB))))
}
