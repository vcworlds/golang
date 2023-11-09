package main

import (
	"crypto/sha256" //用于处理SHA-256哈希函数的Go语言包
	"encoding/hex"  //用于处理十六进制编码和解码
	"fmt"
)

func main() {
	words := "Hello World"
	hash := sha256.Sum256([]byte(words))    //将字符串words转换为字节数组，并使用sha256.Sum256函数计算其SHA-256哈希值。结果存储在变量hash中。
	hashCode := hex.EncodeToString(hash[:]) //使用hex.EncodeToString函数将哈希值hash转换为十六进制字符串表示，并将结果存储在变量hashCode中。
	fmt.Println(hashCode)                   //打印变量hashCode的值，即SHA-256哈希值的十六进制表示。
}

//#   1. 引入hashlib依赖包
//import hashlib
//#   2. 定义需要加密的变量
//words = 'Hello World'
//#   3. 使用SHA-256算法加密
//hash_code = hashlib.sha256(words.encode()).hexdigest()
//#   4. 输出加密后的结果
//print(hash_code)
