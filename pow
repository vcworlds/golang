package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index      int
	PrevHash   string
	Data       []string
	Timestamp  time.Time
	Difficulty uint32
	Nonce      uint32
	MerkleRoot string
	BlockHash  string
	Target     uint64
	HashPrefix string
}

func (b *Block) SetTarget(difficultyBits uint32) {
	exponent := difficultyBits >> 24
	coefficient := difficultyBits & 0xffffff
	b.Target = uint64(coefficient) << (8 * (uint64(exponent) - 3))
}
func (b *Block) CalculateMerkleRoot() string {
	//默克尔跟
	transactionHashes := make([]string, len(b.Data))
	for i, tx := range b.Data {
		hash := sha256.Sum256([]byte(tx))
		transactionHashes[i] = hex.EncodeToString(hash[:])
	}

	for len(transactionHashes) > 1 {
		if len(transactionHashes)%2 == 1 {
			transactionHashes = append(transactionHashes, transactionHashes[len(transactionHashes)-1])
		}
		subHashRoots := make([]string, 0)
		for i := 0; i < len(transactionHashes); i += 2 {
			combined := transactionHashes[i] + transactionHashes[i+1]
			hash := sha256.Sum256([]byte(combined))
			subHashRoots = append(subHashRoots, hex.EncodeToString(hash[:]))
		}
		transactionHashes = subHashRoots
	}

	return transactionHashes[0]
}

func (b *Block) CalculateHash() string {
	//哈希值
	blockHeader := fmt.Sprintf("%d%s%s%v%x%d", b.Index, b.PrevHash, b.MerkleRoot, b.Timestamp, b.Difficulty, b.Nonce)
	hash := sha256.Sum256([]byte(blockHeader))
	return hex.EncodeToString(hash[:])
}

func (b *Block) PoW() {
	for {
		b.Nonce++
		b.BlockHash = b.CalculateHash()
		if b.IsValid() {
			break
		}
	}
}

func (b *Block) IsValid() bool {
	hashPrefix := b.BlockHash[:len(b.HashPrefix)]
	return hashPrefix == b.HashPrefix
}

func main() {
	// 区块难度
	difficultyBits := uint32(0x1e11ffff)

	// 示例区块
	block := Block{
		Index:      1,
		PrevHash:   "0000000000000000000000000000000000000000000000000000000000000000",
		Data:       []string{"123456"},
		Timestamp:  time.Now(),
		Difficulty: difficultyBits,
		Nonce:      0,
		MerkleRoot: "",
		BlockHash:  "",
		Target:     0,
		HashPrefix: "000000",
	}

	// 设置目标值
	block.SetTarget(difficultyBits)

	// 计算默克尔根
	block.MerkleRoot = block.CalculateMerkleRoot()

	// 执行PoW算法
	block.PoW()

	fmt.Println("********************完成计算！********************")
	fmt.Printf("总共计算了: %d 次\n", block.Nonce)
	fmt.Printf("目标值为: %s\n", hex.EncodeToString([]byte{0x1e, 0x11, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}))
	fmt.Printf("区块哈希值为: %s\n", block.BlockHash)
}
