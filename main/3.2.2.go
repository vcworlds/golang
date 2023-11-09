package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// 定义初始比特值
const INITIAL_BITS = 0x1e777777

// 区块结构体
type Block struct {
	Index       int
	PrevHash    string
	Data        string
	Timestamp   string
	Bits        int
	Nonce       int
	ElapsedTime string
	BlockHash   string
}

// 转换区块为JSON格式
func (b *Block) ToJSON() string {
	return fmt.Sprintf(`{
		"bits": "%08x",
		"block_hash": "%s",
		"data": "%s",
		"index": %d,
		"nonce": "%08x",
		"prev_hash": "%s",
		"timestamp": "%s"
	}`, b.Bits, b.BlockHash, b.Data, b.Index, b.Nonce, b.PrevHash, b.Timestamp)
}

// 计算区块哈希值
func (b *Block) CalculateBlockHash() string {
	blockHeader := fmt.Sprintf("%d%s%s%s%x%d",
		b.Index, b.PrevHash, b.Data, b.Timestamp, b.Bits, b.Nonce)
	h := sha256.New()
	h.Write([]byte(blockHeader))
	blockHash := fmt.Sprintf("%x", h.Sum(nil))
	b.BlockHash = blockHash
	return blockHash
}

// 区块链结构体
type Blockchain struct {
	Chain []Block
}

// 创建区块链
func NewBlockchain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{[]Block{genesisBlock}}
}

// 创建创世区块
func CreateGenesisBlock() Block {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	return Block{
		Index:       0,
		PrevHash:    "0000000000000000000000000000000000000000000000000000000000000000",
		Data:        "这是创世区块",
		Timestamp:   timestamp,
		Bits:        INITIAL_BITS,
		Nonce:       0,
		ElapsedTime: "",
		BlockHash:   "",
	}
}

// 添加新区块
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := CreateNewBlock(prevBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

// 创建新区块
func CreateNewBlock(prevBlock Block, data string) Block {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	newBlock := Block{
		Index:       prevBlock.Index + 1,
		PrevHash:    prevBlock.BlockHash,
		Data:        data,
		Timestamp:   timestamp,
		Bits:        prevBlock.Bits,
		Nonce:       0,
		ElapsedTime: "",
		BlockHash:   "",
	}
	return newBlock
}

func main() {
	// 创建区块链
	blockchain := NewBlockchain()

	// 输出创世区块信息
	genesisBlock := blockchain.Chain[0]
	fmt.Println("创世区块信息:")
	fmt.Println(genesisBlock.ToJSON())

	// 间隔1秒后创建一个新区块并输出
	time.Sleep(1 * time.Second)
	blockchain.AddBlock("新增的区块")
	newBlock := blockchain.Chain[len(blockchain.Chain)-1]
	fmt.Println("新区块信息:")
	fmt.Println(newBlock.ToJSON())
}
