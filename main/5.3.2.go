package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// 定义任务函数
func task1(a, b int) {
	result := a + b
	fmt.Printf("%s execute simple task a+b is %d+%d=%d\n", time.Now(), a, b, result)
}

func main() {
	// 创建定时任务调度器
	c := cron.New()

	// 添加定时任务，每10秒执行一次
	c.AddFunc("@every 10s", func() {
		task1(1, 2)
	})

	// 启动定时任务
	c.Start()

	// 阻塞主程序，以便定时任务持续执行
	select {}
}
