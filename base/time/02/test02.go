package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t) // 2022-07-17 22:41:06.001567 +0800 CST m=+0.000057466

	//时间增加 1小时
	fmt.Println(t.Add(time.Hour * 1)) // 2022-07-17 23:41:06.001567 +0800 CST m=+3600.000057466
	//时间增加 15 分钟
	fmt.Println(t.Add(time.Minute * 15)) // 2022-07-17 22:56:06.001567 +0800 CST m=+900.000057466
	//时间增加 10 秒钟
	fmt.Println(t.Add(time.Second * 10)) // 2022-07-17 22:41:16.001567 +0800 CST m=+10.000057466

	//时间减少 1 小时
	fmt.Println(t.Add(-time.Hour * 1)) // 2022-07-17 21:41:06.001567 +0800 CST m=-3599.999942534
	//时间减少 15 分钟
	fmt.Println(t.Add(-time.Minute * 15)) // 2022-07-17 22:26:06.001567 +0800 CST m=-899.999942534
	//时间减少 10 秒钟
	fmt.Println(t.Add(-time.Second * 10)) // 2022-07-17 22:40:56.001567 +0800 CST m=-9.999942534

	time.Sleep(time.Second * 5)
	t2 := time.Now()
	// 计算 t 到 t2 的持续时间
	fmt.Println(t2.Sub(t)) // 5.004318874s
	// 1 年之后的时间
	//t3 := t2.AddDate(1, 0, 0)
	// 计算从 t 到当前的持续时间   这一般计算当前时间-过去某个时间，它们之间的duration，
	fmt.Println(time.Since(t)) // 5.004442316s 如果机器足够快的话，这个时间可能和上面的输出一样
	// 计算现在到明年的持续时间    这个一般计算未来某个时间-当前时间，它们之间的duration，如果是当前时间-未来时间，输出就带有“-”负号
	fmt.Println(time.Until(t)) // 8759h59m59.999864s
}
