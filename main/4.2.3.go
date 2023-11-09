package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

type Transaction struct {
	From      string
	To        string
	Message   string
	Timestamp time.Time
	Signature []byte
}

func main() {
	// 1. 创建两个测试账户，一个模拟发送方，一个模拟接收方
	fmt.Println("********** 生成第一个账户 **********")
	sendAccount := createAccount()
	fmt.Println("********** 生成第二个账户 **********")
	recAccount := createAccount()

	// 2. 生成一个测试交易内容
	tx := Transaction{
		From:      sendAccount.Address,
		To:        recAccount.Address,
		Message:   "测试交易",
		Timestamp: time.Now(),
		Signature: nil,
	}

	// 3. 对交易进行签名
	tx.Signature = signTransaction(tx, sendAccount.PrivateKey)

	// 4. 验证交易正确性
	verificationRes := verifyTransaction(tx, sendAccount.PublicKey)
	fmt.Printf("验证交易，结果为: %v\n", verificationRes)
}

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

func createAccount() *Account {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey
	address := "0x" + fmt.Sprintf("%x", publicKey.X)
	return &Account{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
}

func signTransaction(tx Transaction, privateKey *ecdsa.PrivateKey) []byte {
	message := tx.ID()
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)
	signature := append(r.Bytes(), s.Bytes()...)
	return signature
}

func verifyTransaction(tx Transaction, publicKey *ecdsa.PublicKey) bool {
	message := tx.ID()
	r := new(big.Int).SetBytes(tx.Signature[:32])
	s := new(big.Int).SetBytes(tx.Signature[32:])
	return ecdsa.Verify(publicKey, message, r, s)
}

func (tx *Transaction) ID() []byte {
	message := fmt.Sprintf("%s:%s:%s:%s", tx.From, tx.To, tx.Message, tx.Timestamp)
	hash := sha256.Sum256([]byte(message))
	return hash[:]
}
