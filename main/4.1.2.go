package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

var curve = elliptic.P256()

// 创建密钥对应的种子

func createSeed() []byte {
	seed := make([]byte, 32)
	for {
		_, err := rand.Read(seed)
		if err != nil {
			fmt.Println("随机数生成失败:", err)
		} else {
			break
		}
	}
	return seed
}

// 使用种子创建私钥

func createPrivateKey(seed []byte) *ecdsa.PrivateKey {
	priv, err := ecdsa.GenerateKey(curve, bytes.NewReader(seed))
	if err != nil {
		fmt.Println("私钥生成失败:", err)
		return nil
	}
	return priv
}

// 使用私钥生成公钥

func createPublicKey(privateKey *ecdsa.PrivateKey) []byte {
	return elliptic.Marshal(curve, privateKey.X, privateKey.Y)
}

// 生成地址

func createAddress(publicKey []byte) string {
	// 1. 使用SHA-256计算公钥的哈希值
	sha256Hash := sha256.New()
	sha256Hash.Write(publicKey)
	sha256Sum := sha256Hash.Sum(nil)

	// 2. 使用RIPEMD-160计算新的哈希值
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(sha256Sum)
	ripemd160Sum := ripemd160Hasher.Sum(nil)

	// 3. 添加版本前缀，这里默认使用0x00
	versionPrefix := []byte{0x00}
	payload := append(versionPrefix, ripemd160Sum...)

	// 4. 计算双SHA-256哈希值
	sha256Hash.Reset()
	sha256Hash.Write(payload)
	sha256Sum = sha256Hash.Sum(nil)
	sha256Hash.Reset()
	sha256Hash.Write(sha256Sum)
	doubleSHA256Sum := sha256Hash.Sum(nil)

	// 5. 取前4个字节作为校验和
	checksum := doubleSHA256Sum[:4]

	// 6. 将校验和附加到有效负载后面
	payloadWithChecksum := append(payload, checksum...)

	// 7. 使用Base58编码生成最终的地址
	address := base58.Encode(payloadWithChecksum)

	return address
}

func main() {
	seed := createSeed()
	if seed == nil {
		return
	}

	privateKey := createPrivateKey(seed)
	if privateKey == nil {
		return
	}

	publicKey := createPublicKey(privateKey)
	address := createAddress(publicKey)

	fmt.Println("生成的地址是:", address)
	fmt.Println("私钥:", hex.EncodeToString(privateKey.D.Bytes()))
	fmt.Println("公钥:", hex.EncodeToString(publicKey))
}
