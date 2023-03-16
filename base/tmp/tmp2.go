package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	fmt.Println("CapReader=====before")
	n, err := r.r.Read(p)
	fmt.Println("CapReader=====after")
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func main() {
	r := strings.NewReader("hello, gopher!\n")
	r1 := CapReader(LimitReaderX(r, 4))
	//if _, err := io.Copy(os.Stdout, r1); err != nil {
	//	log.Fatal(err)
	//}
	strArr := make([]byte, 8)
	n, err := r1.Read(strArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count=%d, str=%s", n, strArr)
}

func LimitReaderX(r io.Reader, n int64) io.Reader { return &LimitedReaderX{r, n} }

type LimitedReaderX struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

func (l *LimitedReaderX) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	fmt.Println("LimitedReadx=====before")
	n, err = l.R.Read(p)
	fmt.Println("LimitedReadx=====after")
	l.N -= int64(n)
	return
}
