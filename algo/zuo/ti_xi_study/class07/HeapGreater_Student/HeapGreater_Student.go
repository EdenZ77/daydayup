package main

import (
	"fmt"
)

type Student struct {
	age  int
	name string
}

type HeapGreater struct {
	heap []*Student
	// indexMap 用于记录每个学生在堆中的位置
	indexMap map[*Student]int
	comp     func(a, b *Student) bool
}

func NewHeapGreater(comp func(a, b *Student) bool) *HeapGreater {
	return &HeapGreater{
		heap:     make([]*Student, 0),
		indexMap: make(map[*Student]int),
		comp:     comp,
	}
}

func (h *HeapGreater) isEmpty() bool {
	return len(h.heap) == 0
}

func (h *HeapGreater) size() int {
	return len(h.heap)
}

// contains 判断学生是否在堆中
func (h *HeapGreater) contains(s *Student) bool {
	_, exists := h.indexMap[s]
	return exists
}

// peek 获取堆顶元素
func (h *HeapGreater) peek() *Student {
	return h.heap[0]
}

// push 将学生放入堆中
func (h *HeapGreater) push(s *Student) {
	h.heap = append(h.heap, s)
	// 记录学生在堆中的位置
	h.indexMap[s] = len(h.heap) - 1
	// 调整堆
	h.heapInsert(len(h.heap) - 1)
}

// pop 弹出堆顶元素
func (h *HeapGreater) pop() *Student {
	ans := h.heap[0]
	// 交换堆顶元素和最后一个元素
	h.swap(0, len(h.heap)-1)
	// 删除学生在反向索引中的位置
	delete(h.indexMap, ans)
	// 删除堆顶元素
	h.heap = h.heap[:len(h.heap)-1]
	// 向下调整堆
	h.heapify(0)
	return ans
}

// remove 删除堆中的学生
func (h *HeapGreater) remove(s *Student) {
	lastIndex := len(h.heap) - 1
	// 获取并移除堆中的最后一个元素
	replace := h.heap[lastIndex]
	h.heap = h.heap[:lastIndex]
	// 获取要删除的元素的索引
	if index, exists := h.indexMap[s]; exists {
		delete(h.indexMap, s)
		// 如果要删除的元素不是最后一个元素
		if index != lastIndex {
			// 将最后一个元素放在被删除元素的位置
			h.heap[index] = replace
			h.indexMap[replace] = index
			// 重新平衡堆
			h.resign(replace)
		}
	}
}

func (h *HeapGreater) resign(s *Student) {
	h.heapInsert(h.indexMap[s])
	h.heapify(h.indexMap[s])
}

func (h *HeapGreater) getAllElements() []*Student {
	return h.heap
}

func (h *HeapGreater) heapInsert(index int) {
	// 使用比较器来判断是否需要交换，比较器在初始化时传入
	for index > 0 && h.comp(h.heap[index], h.heap[(index-1)/2]) {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater) heapify(index int) {
	left := 2*index + 1
	for left < len(h.heap) {
		smallest := left
		if right := left + 1; right < len(h.heap) && h.comp(h.heap[right], h.heap[left]) {
			smallest = right
		}
		// 如果左右孩子中较小的孩子都比当前节点大，那么不需要继续下沉
		if h.comp(h.heap[index], h.heap[smallest]) {
			break
		}
		h.swap(index, smallest)
		// 继续下沉
		index = smallest
		left = 2*index + 1
	}
}

func (h *HeapGreater) swap(i, j int) {
	// 交换两个学生在反向索引中的位置
	h.indexMap[h.heap[i]], h.indexMap[h.heap[j]] = j, i
	// 交换两个学生在堆中的位置
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func main() {
	// 初始化一些学生
	students := []*Student{
		{age: 17, name: "A同学"},
		{age: 10, name: "B同学"},
		{age: 29, name: "C同学"},
		{age: 33, name: "D同学"},
		{age: 54, name: "E同学"},
		{age: 93, name: "F同学"},
		{age: 12, name: "12同学"},
	}

	// 创建一个最小堆，比较学生的年龄
	heap := NewHeapGreater(func(a, b *Student) bool {
		return a.age < b.age
	})

	// 将所有学生放入堆中
	for _, student := range students {
		heap.push(student)
	}

	// 修改一个学生的年龄并重新调整堆
	students[4].age = 4
	heap.resign(students[4])

	// 弹出所有学生，它们应该按照年龄顺序弹出
	//for !heap.isEmpty() {
	//	student := heap.pop()
	//	fmt.Printf("年龄: %d, 名字: %s\n", student.age, student.name)
	//}
	fmt.Println("=====================================")
	heap.remove(students[len(students)-1])
	for !heap.isEmpty() {
		student := heap.pop()
		fmt.Printf("年龄: %d, 名字: %s\n", student.age, student.name)
	}
}
