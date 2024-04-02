package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id        int
	buy       int
	enterTime int
}

// CandidateHeap 是一个最大堆，用于存储候选人
type CandidateHeap []Customer

func (h CandidateHeap) Len() int { return len(h) }
func (h CandidateHeap) Less(i, j int) bool {
	if h[i].buy == h[j].buy {
		return h[i].enterTime < h[j].enterTime
	}
	return h[i].buy > h[j].buy
}
func (h CandidateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *CandidateHeap) Push(x interface{}) {
	*h = append(*h, x.(Customer))
}
func (h *CandidateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// DaddyHeap 是一个最小堆，用于存储爸爸级别的顾客
type DaddyHeap []Customer

func (h DaddyHeap) Len() int { return len(h) }
func (h DaddyHeap) Less(i, j int) bool {
	if h[i].buy == h[j].buy {
		return h[i].enterTime < h[j].enterTime
	}
	return h[i].buy < h[j].buy
}
func (h DaddyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *DaddyHeap) Push(x interface{}) {
	*h = append(*h, x.(Customer))
}
func (h *DaddyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// WhosYourDaddy 维护了两个堆结构，分别是候选人堆（CandidateHeap）和爸爸堆（DaddyHeap）
type WhosYourDaddy struct {
	customers   map[int]*Customer
	candHeap    CandidateHeap
	daddyHeap   DaddyHeap
	daddyLimit  int
	currentTime int
}

// NewWhosYourDaddy 创建一个新的 WhosYourDaddy 实例
func NewWhosYourDaddy(limit int) *WhosYourDaddy {
	return &WhosYourDaddy{
		customers:   make(map[int]*Customer),
		candHeap:    make(CandidateHeap, 0, limit),
		daddyHeap:   make(DaddyHeap, 0, limit),
		daddyLimit:  limit,
		currentTime: 0,
	}
}

// Operate 处理买入或退货事件
func (w *WhosYourDaddy) Operate(id int, buyOrRefund bool) {
	if !buyOrRefund && w.customers[id] == nil {
		return
	}

	if w.customers[id] == nil {
		w.customers[id] = &Customer{id: id, buy: 0, enterTime: 0}
	}

	c := w.customers[id]
	if buyOrRefund {
		c.buy++
	} else {
		c.buy--
	}

	if c.buy == 0 {
		delete(w.customers, id)
	} else {
		c.enterTime = w.currentTime
		heap.Push(&w.candHeap, *c)
	}

	w.currentTime++
	w.daddyMove()
}

// daddyMove 确保爸爸堆（DaddyHeap）拥有购买次数最多的顾客
func (w *WhosYourDaddy) daddyMove() {
	if w.candHeap.Len() == 0 {
		return
	}

	if w.daddyHeap.Len() < w.daddyLimit {
		c := heap.Pop(&w.candHeap).(Customer)
		c.enterTime = w.currentTime
		heap.Push(&w.daddyHeap, c)
	} else {
		if w.candHeap[0].buy > w.daddyHeap[0].buy {
			oldDaddy := heap.Pop(&w.daddyHeap).(Customer)
			newDaddy := heap.Pop(&w.candHeap).(Customer)
			oldDaddy.enterTime = w.currentTime
			newDaddy.enterTime = w.currentTime
			heap.Push(&w.daddyHeap, newDaddy)
			heap.Push(&w.candHeap, oldDaddy)
		}
	}
}

// GetDaddies 返回当前爸爸堆中顾客的 ID 列表
func (w *WhosYourDaddy) GetDaddies() []int {
	daddies := make([]int, w.daddyHeap.Len())
	for i, customer := range w.daddyHeap {
		daddies[i] = customer.id
	}
	return daddies
}

func main() {
	// 使用示例
	rand.Seed(time.Now().UnixNano())
	whosYourDaddy := NewWhosYourDaddy(5)

	for i := 0; i < 10; i++ {
		id := rand.Intn(10)
		buyOrRefund := rand.Intn(2) == 1
		whosYourDaddy.Operate(id, buyOrRefund)
		daddies := whosYourDaddy.GetDaddies()
		fmt.Printf("Event #%d: ID=%d, Buy/Refund: %v\n", i, id, buyOrRefund)
		fmt.Printf("Current Daddies: %v\n", daddies)
	}
}
