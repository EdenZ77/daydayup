package main

import (
	"bytes"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	// 将内容清空，这很重要，后面调用者使用Get得到的才是干净的
	buf.Reset()
	buffers.Put(buf)
}

func main() {
	//sync.Pool{}
}
