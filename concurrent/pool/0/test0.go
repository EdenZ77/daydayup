package main

import (
	"fmt"
	"sync"
)

/*
参考资料：https://www.cnblogs.com/qcrao-2018/p/12736031.html

sync.Pool 是 sync 包下的一个组件，可以作为保存临时取还对象的一个“池子”。个人觉得它的名字有一定的误导性，因为 Pool 里装的对象可以被无通知地被回收，可能 sync.Cache 是一个更合适的名字。
*/
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

// 首先，需要初始化 Pool，唯一需要的就是设置好 New 函数。当调用 Get 方法时，如果池子里缓存了对象，就直接返回缓存的对象。如果没有存货，则调用 New 函数创建一个新的对象。
// 另外，我们发现 Get 方法取出来的对象和上次 Put 进去的对象实际上是同一个，Pool 没有做任何“清空”的处理。但我们不应当对此有任何假设，因为在实际的并发使用场景中，无法保证这种顺序，最好的做法是在 Put 前，将对象清空。
func main() {
	initPool()

	p := pool.Get().(*Person) // 第一次从池中拿，由于池中什么都没有，所以创建对象
	fmt.Println("首次从pool中获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p) // 将带有属性值的对象put到池中，那么下一次取出来的就是这个带有属性值的对象，所以有时候需要我们根据业务判断是否需要清理对象的属性值

	// Pool 里已有一个对象：&{first}，调用 Get:  &{first}
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
