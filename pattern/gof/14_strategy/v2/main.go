package main

/*
参考资料：GPT

策略模式（Strategy Pattern）是一种行为设计模式，它定义了一系列算法，并将每个算法封装起来，使它们可以互换使用。
策略模式让算法的变化独立于使用算法的客户端。
*/
import (
	"fmt"
	"sort"
)

// SortStrategy 排序策略接口
type SortStrategy interface {
	Sort([]int)
}

// BubbleSortStrategy 冒泡排序策略
type BubbleSortStrategy struct{}

// Sort 实现冒泡排序算法
func (b BubbleSortStrategy) Sort(data []int) {
	fmt.Println("Sorting using bubble sort")
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// QuickSortStrategy 快速排序策略
type QuickSortStrategy struct{}

// Sort 实现快速排序算法
func (q QuickSortStrategy) Sort(data []int) {
	fmt.Println("Sorting using quick sort")
	sort.Ints(data) // 这里为了简化，直接使用了sort包的Ints方法
}

// Context 上下文
type Context struct {
	strategy SortStrategy
}

// SetStrategy 设置排序策略
func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

// Sort 调用策略方法进行排序
func (c *Context) Sort(data []int) {
	c.strategy.Sort(data)
}

func main() {
	data := []int{9, 5, 3, 7, 1}
	context := Context{}

	// 策略模式允许客户端在运行时选择算法或者行为，从而使算法的变更和使用算法的客户端解耦。

	// 使用冒泡排序策略
	context.SetStrategy(BubbleSortStrategy{})
	context.Sort(data)
	fmt.Println(data)

	// 使用快速排序策略
	context.SetStrategy(QuickSortStrategy{})
	context.Sort(data)
	fmt.Println(data)
}
