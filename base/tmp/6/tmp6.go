package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

func main() {
	//s := "https://1rx-pharmacy.com/c1146/bestsellers.html"
	//index := strings.Index(s, "://")
	//if index != -1 {
	//	fmt.Println(s[index+3:])
	//}

	s := "0-0.com/Ebay/eBay%20NY%2038ac%20Hilltop%20wonderland%20overlooking%20the%20Southern%20Tier.htm"
	index := strings.Index(s, "/")
	if index != -1 {
		fmt.Println(s[:index])
	}
	//s := "0-n-0.ru"
	//strArr := strings.Fields(strings.TrimSpace(s))
	//fmt.Println(strArr)
	//fmt.Println(len(strArr))
	//fmt.Println(strArr[0])
	//ch := make(chan error)
	//go func() {
	//	<-ch
	//}()
	//ch <- nil
	//fmt.Println("main")

	//file, err := os.Create("test.txt")
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//file.WriteString("baidu.com" + "\t" + "T" + "\n")
	//file.WriteString("1.baidu.com" + "\t" + "F" + "\n")
	//countWc()
	//testConcurrentWrite()
}

func testConcurrentWrite() {
	txtFile, err := os.OpenFile("55555.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777) // O_TRUNC 清空重写
	if err != nil {
		fmt.Println("WriteDataToTxt os.OpenFile() err:", err)
		return
	}
	defer txtFile.Close()
	// txtFile.WriteString() // os操作文件-效率低

	bufWriter := bufio.NewWriter(txtFile)
	var wg sync.WaitGroup
	limitChan := make(chan struct{}, runtime.GOMAXPROCS(runtime.NumCPU())) // 最大并发协程数
	var mutex sync.Mutex

	for i := 0; i < 10000000; i++ { // 写1w行测试
		limitChan <- struct{}{}
		wg.Add(1)

		go func(j int) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("WriteDataToTxt panic: %v,stack: %s\n", e, debug.Stack())
				}

				wg.Done()
				<-limitChan
			}()

			// 模拟业务逻辑：先整合所有数据，然后再统一写WriteString()
			strId := fmt.Sprintf("%v", j)
			strName := fmt.Sprintf(" user_%v", j)
			strScore := fmt.Sprintf(" %d", j*10)

			// 要加锁/解锁，否则 bufWriter.WriteString 写入数据有问题
			mutex.Lock()
			_, err := bufWriter.WriteString(strId + strName + strScore + "\n")
			if err != nil {
				fmt.Printf("WriteDataToTxt WriteString err: %v\n", err)
				return
			}
			mutex.Unlock()
			// bufWriter.Flush() // 刷入磁盘（错误示例：WriteDataToTxt err: short write，short write；因为循环太快，有时写入的数据量太小了）
		}(i)
	}
	wg.Wait()
	// 刷入磁盘（正确示例，bufio 通过 flush 操作将缓冲写入真实的文件的，所以一定要在关闭文件之前先flush，否则会造成数据丢失的情况）
	bufWriter.Flush()
}

func countWc() {
	f, err := os.Open("55555.txt")
	if err != nil {
		log.Fatalf("os.open alexa_top.txt err: %s", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 5*1024*1024)
	scanner.Buffer(buf, 5*1024*1024)
	var count int64
	for scanner.Scan() {
		count++
	}
	if scanner.Err() != nil {
		log.Fatalf("scanner err: %s", scanner.Err())
		return
	}
	fmt.Printf("count: %d\n", count)
}
