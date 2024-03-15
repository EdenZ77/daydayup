package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

// 移除链表中等于给定值num的所有节点
func removeValue(head *Node, num int) *Node {
	// 找到第一个不需要删除的节点作为新的头节点
	for head != nil && head.value == num {
		head = head.next
	}
	// 初始化当前节点为新头节点
	cur := head
	var pre *Node // 初始化前一个节点
	// 遍历链表
	for cur != nil {
		if cur.value == num {
			// 如果当前节点需要删除，则将前一个节点的next指向当前节点的下一个节点
			pre.next = cur.next
		} else {
			// 如果当前节点不需要删除，则更新前一个节点为当前节点
			pre = cur
		}
		cur = cur.next // 移动当前节点
	}
	return head
}

// 辅助打印链表函数
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.value)
		head = head.next
	}
	fmt.Println()
}

func main() {
	// 构造测试链表
	values := []int{1, 2, 3, 4, 2, 5, 2, 2}
	//values := []int{1}
	var head *Node
	var cur *Node
	for _, val := range values {
		if head == nil {
			head = &Node{value: val}
			cur = head
		} else {
			cur.next = &Node{value: val}
			cur = cur.next
		}
	}
	fmt.Println("Original linked list:")
	printLinkedList(head)

	// 删除特定值
	num := 1
	head = removeValue(head, num)
	fmt.Printf("Linked list after removing %d:\n", num)
	printLinkedList(head)
}
