package main

import (
	"fmt"
	"sync"
)

/*
第二种死锁的场景有点隐蔽。我们知道，有活跃 reader 的时候，writer 会等待，如果我们在 reader 的读操作时调用 writer 的写操作（它会调用 Lock 方法），
那么，这个 reader 和 writer 就会形成互相依赖的死锁状态。Reader 想等待 writer 完成后再释放锁，而 writer 需要这个 reader 释放锁之后，
才能不阻塞地继续执行。这是一个读写锁常见的死锁场景。反过来也一样，writer写操作过程中调用reader读操作，也会导致死锁


*/

func foo(l *sync.RWMutex) {
	fmt.Println("in foo")
	//l.RLock()
	l.Lock()
	bar(l)
	l.Unlock()
	//l.RUnlock()

}

func bar(l *sync.RWMutex) {
	l.RLock()
	fmt.Println("in bar")
	l.RUnlock()
}

func main() {
	l := &sync.RWMutex{}
	foo(l)
}
