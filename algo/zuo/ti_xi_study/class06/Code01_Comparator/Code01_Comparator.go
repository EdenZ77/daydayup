package main

import (
	"fmt"
	"sort"
)

// Student 结构体，表示学生
type Student struct {
	name string
	id   int
	age  int
}

// IdAgeComparator 结构体，实现了 sort.Interface，用于比较学生
type IdAgeComparator []Student

func (comp IdAgeComparator) Len() int {
	return len(comp)
}

func (comp IdAgeComparator) Swap(i, j int) {
	comp[i], comp[j] = comp[j], comp[i]
}

// Less 方法，根据 id 升序排列，如果 id 相同，则根据 age 降序排列
func (comp IdAgeComparator) Less(i, j int) bool {
	if comp[i].id != comp[j].id {
		return comp[i].id < comp[j].id
	}
	return comp[i].age > comp[j].age
}

// printStudents 打印学生信息
func printStudents(students []Student) {
	for _, student := range students {
		fmt.Printf("Name: %s, Id: %d, Age: %d\n", student.name, student.id, student.age)
	}
}

// MyComp 实现了自定义的整数比较器，用于降序排列
type MyComp []int

func (comp MyComp) Len() int {
	return len(comp)
}

func (comp MyComp) Swap(i, j int) {
	comp[i], comp[j] = comp[j], comp[i]
}

func (comp MyComp) Less(i, j int) bool {
	return comp[i] > comp[j] // 降序
}

func main() {
	// 整数数组排序
	arr := []int{5, 4, 3, 2, 7, 9, 1, 0}
	sort.Sort(MyComp(arr))
	fmt.Println(arr)

	fmt.Println("===========================")

	// 学生数组排序
	students := []Student{
		{"A", 4, 40},
		{"B", 4, 21},
		{"C", 3, 12},
		{"D", 3, 62},
		{"E", 3, 42},
	}

	// 使用 IdAgeComparator 进行排序
	sort.Sort(IdAgeComparator(students))
	printStudents(students)
}
