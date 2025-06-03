package main

import (
	"context"
	"fmt"
	"time"
)

/*
通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，
所以一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，
当一次请求到达我们设置的超时时间，就会及时取消，不再往下执行。
withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已，他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，
通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。
*/
func main() {
	HttpHandler()
}

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func HttpHandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	deal(ctx)
}

func deal(ctx context.Context) {
	ctx1, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		// 超时时间由这个ctx到根节点路径上超时时间最短的ctx决定，子节点cancel不影响父节点
		case <-ctx1.Done():
			fmt.Println("case ======")
			fmt.Println(ctx1.Err())      // context deadline exceeded
			fmt.Println(ctx1.Deadline()) // 2023-05-13 10:38:09.2290093 +0800 CST m=+3.003439401 true 这语句循环输出都是一样的
			//return
		default:
			fmt.Printf("deal time is %d\n", i)
		}
	}
}
