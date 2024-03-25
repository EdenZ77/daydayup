package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node 双向链表的节点
type Node struct {
	value int
	last  *Node
	next  *Node
}

// HeadTail 排序后的头尾节点
type HeadTail struct {
	h *Node
	t *Node
}

// Info 分区信息
type Info struct {
	lh *Node
	lt *Node
	ls int
	rh *Node
	rt *Node
	rs int
	eh *Node
	et *Node
}

// QuickSort 双向链表的快速排序
func QuickSort(head *Node) *Node {
	if head == nil {
		return nil
	}
	N := 0
	c := head
	var e *Node
	for c != nil {
		N++
		e = c
		c = c.next
	}
	ht := process(head, e, N)
	return ht.h
}

// Process 处理排序
func process(L, R *Node, N int) HeadTail {
	if L == nil {
		return HeadTail{}
	}
	if L == R {
		return HeadTail{h: L, t: R}
	}

	// L..R是不止一个节点的链表
	// 随机得到一个随机节点
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(N)
	randomNode := L
	for randomIndex != 0 {
		randomNode = randomNode.next
		randomIndex--
	}

	// 把随机节点从原来的链表中分离出来
	if randomNode == L {
		L = randomNode.next
		L.last = nil
	} else if randomNode == R {
		R = randomNode.last
		R.next = nil
	} else {
		randomNode.last.next = randomNode.next
		randomNode.next.last = randomNode.last
	}
	randomNode.last = nil
	randomNode.next = nil

	info := partition(L, randomNode)

	// < randomNode的部分去排序
	lht := process(info.lh, info.lt, info.ls)
	// > randomNode的部分去排序
	rht := process(info.rh, info.rt, info.rs)

	// 将左部分、等于部分、右部分连接起来
	if lht.h != nil {
		lht.t.next = info.eh
		info.eh.last = lht.t
	}
	if rht.h != nil {
		info.et.next = rht.h
		rht.h.last = info.et
	}

	h := lht.h
	if h == nil {
		h = info.eh
	}
	t := rht.t
	if t == nil {
		t = info.et
	}
	return HeadTail{h: h, t: t}
}

// Partition 分区
func partition(L, pivot *Node) Info {
	var lh, lt *Node
	ls := 0
	var rh, rt *Node
	rs := 0
	eh := pivot
	et := pivot
	var tmp *Node
	for L != nil {
		tmp = L.next
		L.next = nil
		L.last = nil
		if L.value < pivot.value {
			ls++
			if lh == nil {
				lh = L
				lt = L
			} else {
				lt.next = L
				L.last = lt
				lt = L
			}
		} else if L.value > pivot.value {
			rs++
			if rh == nil {
				rh = L
				rt = L
			} else {
				rt.next = L
				L.last = rt
				rt = L
			}
		} else {
			et.next = L
			L.last = et
			et = L
		}
		L = tmp
	}
	return Info{lh: lh, lt: lt, ls: ls, rh: rh, rt: rt, rs: rs, eh: eh, et: et}
}

// 测试用例
func main() {
	// 创建测试链表
	n1 := &Node{value: 4}
	n2 := &Node{value: 2}
	n3 := &Node{value: 6}
	n4 := &Node{value: 5}
	n5 := &Node{value: 1}
	n6 := &Node{value: 8}
	n1.next = n2
	n2.last = n1
	n2.next = n3
	n3.last = n2
	n3.next = n4
	n4.last = n3
	n4.next = n5
	n5.last = n4
	n5.next = n6
	n6.last = n5

	// 执行排序
	sortedHead := QuickSort(n1)

	// 打印排序结果
	printDoubleLinkedList(sortedHead)
}

// 打印双向链表
func printDoubleLinkedList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.value)
		head = head.next
	}
	fmt.Println()
}
