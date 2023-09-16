package main

import (
	"context"
	"fmt"
	"time"
)

/*
https://blog.csdn.net/a568283992/article/details/123438301
泛型+管道 实现一个阻塞队列
*/

type RejectHandler func(ctx context.Context) bool

type Queue[T any] interface {
	Push(value T)
	TryPush(value T, timeout time.Duration) bool
	Poll() T
	TryPoll(timeout time.Duration) (T, bool)
}

type BlockingQueue[T any] struct {
	q       chan T // 阻塞队列使用通道表示
	limit   int    // 阻塞队列的大小
	ctx     context.Context
	handler RejectHandler
}

func NewBlockingQueue[T any](ctx context.Context, queueSize int) *BlockingQueue[T] {
	return &BlockingQueue[T]{
		q:     make(chan T, queueSize),
		limit: queueSize,
		ctx:   ctx,
	}
}

func (b *BlockingQueue[T]) SetRejectHandler(handler RejectHandler) {
	b.handler = handler
}

func (b *BlockingQueue[T]) Push(value T) {
	ok := true
	// 是否有拒绝处理程序
	if b.handler != nil {
		select {
		case b.q <- value:
			return
		// 当chan了满了，不会阻塞，而是调用处理程序，根据处理结果看是否再次入栈
		default:
			ok = b.handler(b.ctx)
		}
	}
	// 当chan了满了就会阻塞
	if ok {
		b.q <- value
	}
}

func (b *BlockingQueue[T]) TryPush(value T, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()
	select {
	// 如果chan满了，则会等待timeout时长，如果还是无法push，则返回
	case b.q <- value:
		return true
	case <-ctx.Done():
	}
	return false
}

func (b *BlockingQueue[T]) Poll() T {
	// 如果chan没有数据了，则会阻塞
	ret := <-b.q
	return ret
}

func (b *BlockingQueue[T]) TryPoll(timeout time.Duration) (ret T, ok bool) {
	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()
	select {
	// 如果chan有数据，则直接返回，没有数据的话，select会阻塞timout时长，还是没有数据，则会返回
	case ret = <-b.q:
		return ret, true
	case <-ctx.Done():
	}
	return ret, false
}

func (b *BlockingQueue[T]) size() int {
	return len(b.q)
}

func main() {
	a := NewBlockingQueue[int](context.Background(), 3)
	a.SetRejectHandler(func(ctx context.Context) bool {
		fmt.Println("reject")
		return false
	})
	for i := 0; i < 4; i++ {
		go a.Push(i)
	}
	select {}
}
