package main

import (
	"github.com/bytedance/gopkg/util/gopool"
)

func main() {
	gopool.Go(func() {
		/// do your job
	})
}
