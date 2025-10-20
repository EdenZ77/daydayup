package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"k8s.io/client-go/util/workqueue"
)

func main() {
	fmt.Println("Starting workqueue example...")

	// 1. 创建一个基本工作队列
	queue := workqueue.New()

	// 2. 创建一个等待组，确保所有goroutine完成
	var wg sync.WaitGroup

	// 3. 启动生产者goroutine
	wg.Add(1)
	go producer(&wg, queue)

	// 4. 启动多个消费者goroutine
	consumerCount := 3
	wg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go consumer(i, &wg, queue)
	}

	// 5. 等待一段时间让生产者和消费者工作
	time.Sleep(5 * time.Second)

	// 6. 关闭队列
	fmt.Println("Shutting down queue...")
	queue.ShutDown()

	// 7. 等待所有goroutine完成
	wg.Wait()
	fmt.Println("All goroutines completed. Exiting.")
}

// 生产者函数，向队列添加项目
func producer(wg *sync.WaitGroup, queue workqueue.Interface) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// 随机延迟，模拟实际生产环境
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		item := "item-" + strconv.Itoa(i)
		fmt.Printf("Producer: Adding %s to queue\n", item)

		// 使用Add方法添加项目到队列
		queue.Add(item)

		// 每添加几个项目后，打印队列长度
		if i%3 == 0 {
			fmt.Printf("Queue length: %d\n", queue.Len())
		}
	}

	fmt.Println("Producer finished adding items")
}

// 消费者函数，从队列获取并处理项目
func consumer(id int, wg *sync.WaitGroup, queue workqueue.Interface) {
	defer wg.Done()

	for {
		// 检查队列是否正在关闭
		if queue.ShuttingDown() {
			fmt.Printf("Consumer %d: Queue is shutting down, stopping...\n", id)
			return
		}

		// 从队列获取项目
		item, shutdown := queue.Get()

		// 如果队列已关闭，退出循环
		if shutdown {
			fmt.Printf("Consumer %d: Queue is shut down\n", id)
			return
		}

		// 处理项目
		fmt.Printf("Consumer %d: Processing %s\n", id, item)

		// 模拟处理时间
		processingTime := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(processingTime)

		// 随机决定是否处理成功
		success := rand.Intn(2) == 0

		if success {
			fmt.Printf("Consumer %d: Successfully processed %s\n", id, item)
			// 标记项目处理完成
			queue.Done(item)
		} else {
			fmt.Printf("Consumer %d: Failed to process %s, re-adding to queue\n", id, item)
			// 处理失败，将项目重新加入队列
			queue.Add(item)
			// 标记项目处理完成（即使失败也要调用Done）
			queue.Done(item)
		}

		// 打印当前队列长度
		fmt.Printf("Consumer %d: Queue length is now %d\n", id, queue.Len())
	}
}
