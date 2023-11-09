package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func job1() {
	fmt.Println(time.Now(), "executed job1!")
}

func main() {
	c := cron.New()

	// 每5秒执行一次任务
	c.AddFunc("*/5 * * * * *", job1)

	c.Start()

	select {}
}
