package main

import (
	"context"
	"fmt"
	"time"
)

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
	if b.handler != nil {
		select {
		case b.q <- value:
			return
		default:
			ok = b.handler(b.ctx)
		}
	}
	if ok {
		b.q <- value
	}
}

func (b *BlockingQueue[T]) TryPush(value T, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()
	select {
	case b.q <- value:
		return true
	case <-ctx.Done():
	}
	return false
}

func (b *BlockingQueue[T]) Poll() T {
	ret := <-b.q
	return ret
}

func (b *BlockingQueue[T]) TryPoll(timeout time.Duration) (ret T, ok bool) {
	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()
	select {
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
