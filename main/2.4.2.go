package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Sign函数使用提供的私钥对数据进行签名。
func Sign(data []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, err
	}

	// 将r和s转换为字节数组，并将它们连接在一起以生成签名。
	rBytes := r.Bytes()
	sBytes := s.Bytes()
	signature := append(rBytes, sBytes...)

	return signature, nil
}

// Verify函数使用提供的公钥和签名来验证数据。
func Verify(data, signature []byte, publicKey *ecdsa.PublicKey) bool {
	hash := sha256.Sum256(data)

	// 将签名分成r和s两部分。
	rBytes := signature[:len(signature)/2]
	sBytes := signature[len(signature)/2:]

	// 将r和s恢复为大整数。
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	// 使用公钥验证签名。
	return ecdsa.Verify(publicKey, hash[:], r, s)
}

func main() {
	// 1. 创建要签名的原始数据
	data := []byte("hello world")

	// 2. 生成私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("生成私钥时出错:", err)
		return
	}

	// 3. 从私钥获取公钥
	publicKey := &privateKey.PublicKey

	// 4. 签名
	signature, err := Sign(data, privateKey)
	if err != nil {
		fmt.Println("签名时出错:", err)
		return
	}

	// 5. 验证签名
	valid := Verify(data, signature, publicKey)
	if valid {
		fmt.Println("签名验证成功")
	} else {
		fmt.Println("签名验证失败")
	}

	fmt.Printf("数据: %s\n", string(data))
	fmt.Printf("私钥是:%x\n", privateKey) // 打印私钥的十六进制表示形式
	fmt.Printf("公钥: %x\n", publicKey)
	fmt.Printf("签名: %s\n", hex.EncodeToString(signature))
}
