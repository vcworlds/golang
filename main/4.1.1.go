package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func generateKeys() (*ecdsa.PrivateKey, []byte, []byte, error) {
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, nil, err
	}

	// 生成公钥
	publicKey := &privateKey.PublicKey

	// 将私钥转换为PEM格式
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, nil, err
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})

	// 将公钥转换为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, nil, err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})

	return privateKey, privateKeyPEM, publicKeyPEM, nil
}

func main() {
	_, privateKeyPEM, publicKeyPEM, err := generateKeys()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Private Key:\n%s\n", privateKeyPEM)
	fmt.Printf("Public Key:\n%s\n", publicKeyPEM)
}
