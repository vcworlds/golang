package main

import (
	"crypto/sha256"
	"fmt"
)

func calcMerkleRoot(data []string) string {
	// 复制交易数据到临时数组
	calcTxs := make([]string, len(data))
	copy(calcTxs, data)

	// 如果只有一个交易，返回该交易的哈希值
	if len(calcTxs) == 1 {
		return calcTxs[0]
	}

	// 循环计算哈希值
	for len(calcTxs) > 1 {
		// 如果交易个数为奇数，复制最后一个交易信息进行补全
		if len(calcTxs)%2 == 1 {
			calcTxs = append(calcTxs, calcTxs[len(calcTxs)-1])
		}
		subHashRoots := []string{}
		// 每两个交易进行组合，生成新的哈希值
		for i := 0; i < len(calcTxs); i += 2 {
			joinStr := calcTxs[i] + calcTxs[i+1]
			// 生成的哈希值进行累加
			subHashRoots = append(subHashRoots, fmt.Sprintf("%x", sha256.Sum256([]byte(joinStr))))
		}
		// 将累加得出的新的哈希列表赋值为下一次循环的内容
		calcTxs = subHashRoots
	}
	return calcTxs[0]
}

func main() {
	transactions := []string{
		"Transaction 1 Data",
		"Transaction 2 Data",
		"Transaction 3 Data",
		"Transaction 4 Data",
	}
	merkleRoot := calcMerkleRoot(transactions)
	fmt.Println("计算的默克尔根为:", merkleRoot)
}
