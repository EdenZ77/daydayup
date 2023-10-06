package main

import (
	"github.com/bytedance/gopkg/util/gopool"
)

// 参考资料：https://juejin.cn/post/7086443265309818894#heading-1
func main() {
	gopool.Go(func() {
		/// do your job
	})
}
