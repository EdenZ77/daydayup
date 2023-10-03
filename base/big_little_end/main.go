package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"unsafe"
)

func main() {
	testVar()
}

func testWrite1() {
	num := int64(0x1122334455667788)
	//声明两个Writer的实现buffer，长度8,因为长整型int是64位，也就是8字节
	bigEndianBuffer := bytes.NewBuffer(make([]byte, 0))
	littleEndianBuffer := bytes.NewBuffer(make([]byte, 0))
	//将变量写入流中 一个使用高字端，一个使用低字端
	_ = binary.Write(bigEndianBuffer, binary.BigEndian, num)
	_ = binary.Write(littleEndianBuffer, binary.LittleEndian, num)

	fmt.Println(bigEndianBuffer.Bytes())
	//[17 34 51 68 85 102 119 136] 其实对应的16进制就是
	//[0x11 0x22 0x33 0x44 0x55 0x66 0x77 0x88]

	fmt.Println(littleEndianBuffer.Bytes())
	//[136 119 102 85 68 51 34 17] 与上面结果正好相反，对应的16进制就是
	//[0x88 0x77 0x66 0x55 0x44 0x33 0x22 0x11]
}

func testRead() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi) // 3.141592653589793
}

func testWrite2() {
	buf := new(bytes.Buffer)
	pi := math.Pi
	fmt.Println(pi) // 3.141592653589793
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	// 输出的是十进制： [24 45 68 84 251 33 9 64]
	// 对应的十六进制就是：0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40
	fmt.Println(buf.Bytes())
}

func testSize() {
	var a int
	p := &a
	b := [10]int64{1}
	s := "adsa"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  // -1
	fmt.Println(binary.Size(p))  // -1
	fmt.Println(binary.Size(b))  // 80
	fmt.Println(binary.Size(s))  // -1
	fmt.Println(binary.Size(bs)) // 10
}

func testPut() {
	// write
	v := uint32(500)
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)
	// [0 0 1 244]
	// 十六进制对应：0x0, 0x0, 0x1, 0xf4
	fmt.Println(buf)

	// read
	x := binary.BigEndian.Uint32(buf)
	fmt.Println(x) // 500
}

func testRead2() {
	var v uint32
	b := []byte{0x0, 0x0, 0x1, 0xf4}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, &v)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(v) // 500
}

const INT_SIZE int = int(unsafe.Sizeof(0))

// 判断我们系统中的字节序类型
func systemEdian() {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
}

func testVar() {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []int64{-65, 1, 2, 127, 128, 255, 256} {
		n := binary.PutVarint(buf, x)
		fmt.Print(x, "输出的可变长度为：", n, "，十六进制为：")
		fmt.Printf("%x\n", buf[:n])
	}
}

/*
-65输出的可变长度为：2，十六进制为：8101
1输出的可变长度为：1，十六进制为：02
2输出的可变长度为：1，十六进制为：04
127输出的可变长度为：2，十六进制为：fe01
128输出的可变长度为：2，十六进制为：8002
255输出的可变长度为：2，十六进制为：fe03
256输出的可变长度为：2，十六进制为：8004
*/

func testVar1() {
	inputs := [][]byte{
		[]byte{0x81, 0x01},
		[]byte{0x7f},
		[]byte{0x03},
		[]byte{0x01},
		[]byte{0x00},
		[]byte{0x02},
		[]byte{0x04},
		[]byte{0x7e},
		[]byte{0x80, 0x01},
	}
	for _, b := range inputs {
		x, n := binary.Varint(b)
		if n != len(b) {
			fmt.Println("Varint did not consume all of in")
		}
		fmt.Println(x) // -65,-64,-2,-1,0,1,2,63,64,
	}
}
