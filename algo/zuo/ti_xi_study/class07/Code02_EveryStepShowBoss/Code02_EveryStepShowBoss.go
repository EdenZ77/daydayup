package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Customer struct {
	id        int
	buy       int
	enterTime int
	index     int // 堆中的索引
}

// CandidateHeap maintains the candidates outside of the daddy list.
type CandidateHeap []*Customer

func (h CandidateHeap) Len() int { return len(h) }

// Less 谁购买数量多，谁大，如果购买数量一样，谁先进来，谁大
func (h CandidateHeap) Less(i, j int) bool {
	if h[i].buy == h[j].buy {
		return h[i].enterTime < h[j].enterTime
	}
	return h[i].buy > h[j].buy
}
func (h CandidateHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *CandidateHeap) Push(x interface{}) { *h = append(*h, x.(*Customer)) }
func (h *CandidateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// DaddyHeap maintains the top k customers.
type DaddyHeap []*Customer

func (h DaddyHeap) Len() int { return len(h) }
func (h DaddyHeap) Less(i, j int) bool {
	if h[i].buy == h[j].buy {
		return h[i].enterTime < h[j].enterTime
	}
	return h[i].buy < h[j].buy
}
func (h DaddyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *DaddyHeap) Push(x interface{}) { *h = append(*h, x.(*Customer)) }
func (h *DaddyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// WhosYourDaddy maintains the entire logic for the customers' buy and refund operations.
type WhosYourDaddy struct {
	customers  map[int]*Customer
	candHeap   CandidateHeap
	daddyHeap  DaddyHeap
	daddyLimit int
}

// NewWhosYourDaddy creates a new WhosYourDaddy instance.
func NewWhosYourDaddy(limit int) *WhosYourDaddy {
	return &WhosYourDaddy{
		customers:  make(map[int]*Customer),
		candHeap:   make(CandidateHeap, 0),
		daddyHeap:  make(DaddyHeap, 0),
		daddyLimit: limit,
	}
}

// operate performs the buy or refund operation and adjusts the customers' positions.
func (w *WhosYourDaddy) operate(time int, id int, buyOrRefund bool) {
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
	}
	// Check if in candidate or daddy list
	inCand := c.index < len(w.candHeap) && w.candHeap[c.index] == c
	inDaddy := c.index < len(w.daddyHeap) && w.daddyHeap[c.index] == c

	if !inCand && !inDaddy {
		if w.daddyHeap.Len() < w.daddyLimit {
			c.enterTime = time
			heap.Push(&w.daddyHeap, c)
		} else {
			c.enterTime = time
			heap.Push(&w.candHeap, c)
		}
	} else if inCand {
		if c.buy == 0 {
			heap.Remove(&w.candHeap, c.index)
		} else {
			heap.Fix(&w.candHeap, c.index)
		}
	} else {
		if c.buy == 0 {
			heap.Remove(&w.daddyHeap, c.index)
		} else {
			heap.Fix(&w.daddyHeap, c.index)
		}
	}
	w.daddyMove(time)
}

// getDaddies returns the list of daddy customers.
func (w *WhosYourDaddy) getDaddies() []int {
	var ans []int
	for _, c := range w.daddyHeap {
		ans = append(ans, c.id)
	}
	return ans
}

// daddyMove adjusts the daddy list after each operation.
func (w *WhosYourDaddy) daddyMove(time int) {
	if w.candHeap.Len() == 0 {
		return
	}
	if w.daddyHeap.Len() < w.daddyLimit {
		p := heap.Pop(&w.candHeap).(*Customer)
		p.enterTime = time
		heap.Push(&w.daddyHeap, p)
	} else {
		if w.candHeap[0].buy > w.daddyHeap[0].buy {
			oldDaddy := heap.Pop(&w.daddyHeap).(*Customer)
			newDaddy := heap.Pop(&w.candHeap).(*Customer)
			oldDaddy.enterTime = time
			newDaddy.enterTime = time
			heap.Push(&w.daddyHeap, newDaddy)
			heap.Push(&w.candHeap, oldDaddy)
		}
	}
}

// topK simulates the buy and refund operations and returns the list of top k customers at each step.
func topK(arr []int, op []bool, k int) [][]int {
	ans := make([][]int, 0)
	whosYourDaddy := NewWhosYourDaddy(k)
	for i, id := range arr {
		whosYourDaddy.operate(i, id, op[i])
		ans = append(ans, whosYourDaddy.getDaddies())
	}
	return ans
}

// compare 比较器 干完所有的事，模拟，不优化
func compare(arr []int, op []bool, k int) [][]int {
	// key: 用户id value: 用户信息
	mapCustomers := make(map[int]*Customer)
	// 候选区
	var cands []*Customer
	// 父亲区
	var daddies []*Customer
	// 返回的答案
	var ans [][]int

	for i, id := range arr {
		buyOrRefund := op[i]
		// 没有发生：用户购买数为0并且又退货了
		if !buyOrRefund && mapCustomers[id] == nil {
			ans = append(ans, getCurAns(daddies))
			continue
		}
		// 没有发生：用户购买数为0并且又退货了
		// 用户之前购买数是0，此时买货事件
		// 用户之前购买数>0， 此时买货
		// 用户之前购买数>0, 此时退货
		if mapCustomers[id] == nil {
			// 用户之前不存在，初始化然后放到map里，后面调整buy和enterTime
			mapCustomers[id] = &Customer{id: id, buy: 0, enterTime: 0}
		}

		c := mapCustomers[id]
		if buyOrRefund {
			c.buy++
		} else {
			c.buy--
		}
		// 如果退货之后购买数为0，删除用户信息，同时也要从候选区和父亲区删除(逻辑在下面)
		if c.buy == 0 {
			delete(mapCustomers, id)
		}
		// 之前不在候选区也不在父亲区，说明是新用户，或者是之前的用户，但是之前的用户已经被淘汰了(购买数清空了)
		if !contains(cands, c) && !contains(daddies, c) {
			// 父亲区没满，直接放进去
			if len(daddies) < k {
				c.enterTime = i
				daddies = append(daddies, c)
				// 父亲区满了，放到候选区，等待后续的调整
			} else {
				c.enterTime = i
				cands = append(cands, c)
			}
		}
		// 复杂度就主要发生在里面，对于每个事件都需要遍历清空购买数为0的用户
		// 以及对候选区和父亲区进行排序和调整
		cleanZeroBuy(&cands)
		cleanZeroBuy(&daddies)
		// 将候选区和父亲区按照规则排序
		sort.Sort(CandidateHeap(cands))
		sort.Sort(DaddyHeap(daddies))
		move(&cands, &daddies, k, i)
		ans = append(ans, getCurAns(daddies))
	}
	return ans
}

// move 调整候选区和父亲区的逻辑
func move(cands, daddies *[]*Customer, k, time int) {
	// 候选区为空，不用调整。比如父亲区还没有满，候选区就是空的
	if len(*cands) == 0 {
		return
	}
	// 父亲区没满，直接把候选区的第一个放到父亲区。比如父亲区都是1票，候选区也是1票，此刻的事件父亲区某个用户退货，所以就出现了父亲区没满的情况但是候选区有用户
	if len(*daddies) < k {
		c := (*cands)[0]
		c.enterTime = time
		// 直接将候选区的第一个(最大的一个)放到父亲区
		*daddies = append(*daddies, c)
		// 候选区的第一个已经放到父亲区了，所以候选区的第一个要删除
		*cands = (*cands)[1:]
		// 父亲区满了，而且候选区也不为空，此时要进行调整
	} else {
		// 候选区的第一个大于父亲区的第一个(最小值)，交换
		if (*cands)[0].buy > (*daddies)[0].buy {
			oldDaddy := (*daddies)[0]
			newDaddy := (*cands)[0]
			*daddies = (*daddies)[1:]
			*cands = (*cands)[1:]
			newDaddy.enterTime = time
			oldDaddy.enterTime = time
			*daddies = append(*daddies, newDaddy)
			*cands = append(*cands, oldDaddy)
		}
	}
}

// cleanZeroBuy removes customers with a buy count of zero from the list.
func cleanZeroBuy(arr *[]*Customer) {
	var noZero []*Customer
	for _, c := range *arr {
		if c.buy != 0 {
			noZero = append(noZero, c)
		}
	}
	*arr = noZero
}

// getCurAns extracts the current list of daddy customers.
func getCurAns(daddies []*Customer) []int {
	ans := []int{}
	for _, c := range daddies {
		ans = append(ans, c.id)
	}
	return ans
}

// contains checks if a customer is present in the given list.
func contains(list []*Customer, customer *Customer) bool {
	for _, c := range list {
		if c == customer {
			return true
		}
	}
	return false
}

// sameAnswer compares two lists of lists of integers for equality.
func sameAnswer(ans1, ans2 [][]int) bool {
	if len(ans1) != len(ans2) {
		return false
	}
	for i := range ans1 {
		cur1 := append([]int(nil), ans1[i]...)
		cur2 := append([]int(nil), ans2[i]...)
		sort.Ints(cur1)
		sort.Ints(cur2)
		if len(cur1) != len(cur2) {
			return false
		}
		for j := range cur1 {
			if cur1[j] != cur2[j] {
				return false
			}
		}
	}
	return true
}

// randomData generates random test data for the simulation.
func randomData(maxValue, maxLen int) ([]int, []bool) {
	len := rand.Intn(maxLen) + 1
	arr := make([]int, len)
	op := make([]bool, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(maxValue)
		op[i] = rand.Intn(2) == 1
	}
	return arr, op
}

func main() {
	rand.Seed(time.Now().UnixNano())

	maxValue := 10
	maxLen := 100
	maxK := 6
	testTimes := 100000
	fmt.Println("Testing starts")

	for i := 0; i < testTimes; i++ {
		arr, op := randomData(maxValue, maxLen)
		k := rand.Intn(maxK) + 1
		ans1 := topK(arr, op, k)
		ans2 := compare(arr, op, k)
		if !sameAnswer(ans1, ans2) {
			fmt.Println("Error found!")
			fmt.Println("arr:", arr)
			fmt.Println("op:", op)
			fmt.Println("k:", k)
			fmt.Println("ans1:", ans1)
			fmt.Println("ans2:", ans2)
			break
		}
	}

	fmt.Println("Testing ends")
}
