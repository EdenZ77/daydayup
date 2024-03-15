package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
}

type DoubleNode struct {
	value int
	prev  *DoubleNode
	next  *DoubleNode
}

// 反转单链表
/*
在遍历列表时，将当前节点的 next 指针改为指向前一个元素。必须事先存储当前节点后面一个元素。不要忘记在最后返回新的头引用！
*/
func reverseLinkedList(head *Node) *Node {
	var prev *Node
	var next *Node
	for head != nil {
		// 保存下一个节点
		next = head.next
		// 当前节点的 next 指针指向前一个元素
		head.next = prev
		// 更新 prev 为当前节点
		prev = head
		// 更新当前节点为下一个节点
		head = next
	}
	return prev
}

// 反转双向链表
func reverseDoubleList(head *DoubleNode) *DoubleNode {
	var prev *DoubleNode
	var next *DoubleNode
	for head != nil {
		// 保存下一个节点
		next = head.next
		// 当前节点的 next 指针指向前一个元素
		head.next = prev
		// 当前节点的 prev 指针指向下一个元素
		head.prev = next
		// 更新 prev 为当前节点
		prev = head
		// 更新当前节点为下一个节点
		head = next
	}
	return prev
}

// 生成随机单链表
func generateRandomLinkedList(length int, maxValue int) *Node {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(length + 1)
	if size == 0 {
		return nil
	}
	head := &Node{value: rand.Intn(maxValue + 1)}
	current := head
	for size > 1 {
		size--
		current.next = &Node{value: rand.Intn(maxValue + 1)}
		current = current.next
	}
	return head
}

// 生成随机双向链表
func generateRandomDoubleList(length int, maxValue int) *DoubleNode {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(length + 1)
	if size == 0 {
		return nil
	}
	head := &DoubleNode{value: rand.Intn(maxValue + 1)}
	current := head
	for size > 1 {
		size--
		newNode := &DoubleNode{value: rand.Intn(maxValue + 1)}
		current.next = newNode
		newNode.prev = current
		current = newNode
	}
	return head
}

// 打印单链表
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Print(head.value, " ")
		head = head.next
	}
	fmt.Println()
}

// 打印双向链表
func printDoubleList(head *DoubleNode) {
	for head != nil {
		fmt.Print(head.value, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	// 测试单链表反转
	head := generateRandomLinkedList(10, 100)
	fmt.Println("Original linked list:")
	printLinkedList(head)
	reversedHead := reverseLinkedList(head)
	fmt.Println("Reversed linked list:")
	printLinkedList(reversedHead)

	// 测试双向链表反转
	doubleHead := generateRandomDoubleList(10, 100)
	fmt.Println("Original double linked list:")
	printDoubleList(doubleHead)
	reversedDoubleHead := reverseDoubleList(doubleHead)
	fmt.Println("Reversed double linked list:")
	printDoubleList(reversedDoubleHead)
}
