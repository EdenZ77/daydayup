package main

type ReadWriter interface {
	~string | ~[]rune

	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string

func (s StringReadWriter) Read(p []byte) (n int, err error) {
	//...
	return 0, err
}

func (s StringReadWriter) Write(p []byte) (n int, err error) {
	//...
	return 0, err
}

// 类型BytesReadWriter 没有实现接口 Readwriter
type BytesReadWriter []byte

func (s BytesReadWriter) Read(p []byte) (n int, err error) {
	//...
	return 0, err
}

func (s BytesReadWriter) Write(p []byte) (n int, err error) {
	//...
	return 0, err
}

func main() {

}
