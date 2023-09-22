package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating a new person")
			return new(Person)
		},
	}
}

/*
首先，需要初始化 Pool，唯一需要的就是设置好 New 函数。
当调用 Get 方法时，如果池子里缓存了对象，就直接返回缓存的对象。如果没有存货，则调用 New 函数创建一个新的对象。
*/
func main() {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("首次从pool中获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	// 这个时候上面已经get了，pool里面没有对象了，所以会调用New函数创建一个新的对象
	person2 := pool.Get().(*Person)
	person2.Name = "second"
	fmt.Println("Pool 没有对象了，调用 Get: ", person2)
	// 向pool中放入两个对象，里面是队列，先进先出
	pool.Put(person2)
	pool.Put(&Person{
		Name: "three",
	})
	fmt.Println("现在pool里面有两个对象了")
	fmt.Println("试试", pool.Get().(*Person)) // second
	fmt.Println("试试", pool.Get().(*Person)) // three

}
