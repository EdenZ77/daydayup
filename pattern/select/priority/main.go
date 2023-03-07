package main

import "fmt"

/*
现在我们有一个需求：我们有一个函数会持续不间断地从ch1和ch2中分别接收任务1和任务2，

如何确保当ch1和ch2同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2呢？

上面的需求虽然是我编的，但是关于在select中实现优先级在实际生产中是有实际应用场景的，例如K8s的controller中就有关于上面这个技巧的实际使用示例
// kubernetes/pkg/controller/nodelifecycle/scheduler/taint_manager.go

*/

func worker(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			// 当任务2到的时候，再次检查任务1，如果任务1也到了则先处理，任务2后处理
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}

func main() {

}
