package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//创建一个周期性的定时器，每3秒执行一次的定时器
	ticker := time.NewTicker(3 * time.Second)

	//定义计数器
	count := 1
	fmt.Println("first 当前时间为:", time.Now(), "count = ", count)

	go func() {
		for {
			//从定时器中获取数据
			t := <-ticker.C
			count++
			fmt.Println("当前时间为:", t, "count = ", count)
			if count == 10 {
				// 如果周期性定时被消费10次后就停止该定时器
				ticker.Stop()
				// 调用runtime.goExit()将立即终止当前goroutine执行
				// runtime.Goexit函数在终止调用它的Goroutine之前会先执行该Groution中还没有执行的defer语句
				runtime.Goexit()
			}
		}
	}()

	time.Sleep(time.Second * 35)
	fmt.Println("main over")

}

//func worker() {
//	heartbeat := time.NewTicker(30 * time.Second)
//	defer heartbeat.Stop()
//	for {
//		select {
//		case <-c:
//			// ... do some stuff
//		case <- heartbeat.C:
//			//... do heartbeat stuff
//		}
//	}
//}
