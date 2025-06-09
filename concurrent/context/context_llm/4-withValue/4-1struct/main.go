package main

import (
	"context"
	"fmt"
)

/*
使用空结构体作为 Context 的 key，最关键的步骤，其实是要基于空结构体定义一个新的类型。
我们使用这个新类型的实例对象作为 key，而不是直接使用空结构体变量作为 key，这二者是有本质区别的。
*/
type emptyKey struct{}
type anotherEmpty struct{}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, emptyKey{}, "empty struct data")

	fmt.Printf("empty data: %s\n", ctx.Value(emptyKey{}))

	ctx = context.WithValue(ctx, anotherEmpty{}, "another empty struct data")

	fmt.Printf("another empty data: %s\n", ctx.Value(anotherEmpty{}))

	// 再次查看 emptyKey 对应的 value
	fmt.Printf("empty data: %s\n", ctx.Value(emptyKey{}))
}
