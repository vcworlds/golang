package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type BlockController struct {
	beego.Controller
}

type Block struct {
	Index      int      // 区块索引号
	PrevHash   string   // 父块哈希值
	Data       []string // 区块中需保存的记录
	Timestamp  string   // 区块生成的时间戳
	Difficulty int      // 区块难度
	Nonce      int
	MerkleRoot string
	BlockHash  string
}

func (c *BlockController) BlockMerkleRoot(data []string) string {
	// 计算默克尔根
	calcTxs := make([]string, len(data))
	copy(calcTxs, data)

	if len(calcTxs) == 1 {
		return calcTxs[0]
	}

	for len(calcTxs) > 1 {
		if len(calcTxs)%2 == 1 {
			calcTxs = append(calcTxs, calcTxs[len(calcTxs)-1])
		}

		subHashRoots := make([]string, 0)

		for i := 0; i < len(calcTxs); i += 2 {
			joinStr := calcTxs[i] + calcTxs[i+1]
			hash := sha256.Sum256([]byte(joinStr))
			subHashRoots = append(subHashRoots, hex.EncodeToString(hash[:]))
		}

		calcTxs = subHashRoots
	}

	return calcTxs[0]
}

func (c *BlockController) BlockHash(block Block) string {
	// 区块哈希计算
	blockHeader := c.getFormattedBlockHeader(&block)
	hash := sha256.Sum256([]byte(blockHeader))
	return hex.EncodeToString(hash[:])
}

func (c *BlockController) getFormattedBlockHeader(block *Block) string {
	return fmt.Sprintf("%d%s%v%s%x%d", block.Index, block.PrevHash, block.Data, block.Timestamp, block.Difficulty, block.Nonce)
}

func (c *BlockController) CreateBlock() {
	index := 1
	prevHash := "0"                     // 初始区块，所以设置为默认值
	blockData := []string{"交易1", "交易2"} // 示例数据
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	difficulty := 0 // 设置难度值
	nonce := 0

	merkleRoot := c.BlockMerkleRoot(blockData)

	block := Block{
		Index:      index,
		PrevHash:   prevHash,
		Data:       blockData,
		Timestamp:  timestamp,
		Difficulty: difficulty,
		Nonce:      nonce,
		MerkleRoot: merkleRoot,
	}

	blockHash := c.BlockHash(block)
	block.BlockHash = blockHash

	// 以JSON格式返回区块
	c.Data["json"] = &block
	c.ServeJSON()
}
