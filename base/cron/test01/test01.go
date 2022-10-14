package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	// 创建一个默认的cron对象
	//c := cron.New()

	// 自定义解析器
	c := cron.New(cron.WithSeconds())

	// Seconds field, optional
	//c := cron.New(cron.WithParser(cron.NewParser(
	//	cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	//)))

	// 添加一个任务，每30s执行一次
	_, err := c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("every 30 seconds execute")
	})
	if err != nil {
		log.Fatal(err)
		//return
	}
	c.Start()
	// 允许往正在执行的cron中添加任务
	_, err = c.AddFunc("@every 30s", func() {
		fmt.Println("every day")
	})
	if err != nil {
		log.Fatal(err)
		//return
	}
	select {}
}
