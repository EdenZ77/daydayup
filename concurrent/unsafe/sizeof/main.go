package main

import (
	"fmt"
	"unsafe"
)

func main() {
	func_example()
}

type User struct {
	Name   string
	Age    uint32
	Gender bool // 男:true 女：false 就是举个例子别吐槽我这么用。。。。
}

func func_example() {
	// sizeof
	fmt.Println(unsafe.Sizeof(true))                 // 1
	fmt.Println(unsafe.Sizeof(int8(0)))              // 1
	fmt.Println(unsafe.Sizeof(int16(10)))            // 2
	fmt.Println(unsafe.Sizeof(int(10)))              // 8
	fmt.Println(unsafe.Sizeof(int32(190)))           // 4
	fmt.Println(unsafe.Sizeof("asong"))              // 16,这是因为string的结构体中有两个字段，指针和长度
	fmt.Println(unsafe.Sizeof([]int{1, 3, 4, 5, 6})) // 24,这是因为slice的结构体中有三个字段，指针，长度，容量
	fmt.Println("============================")
	// Offsetof
	user := User{Name: "Asong", Age: 23, Gender: true}
	userNamePointer := unsafe.Pointer(&user)

	nNamePointer := (*string)(userNamePointer)
	*nNamePointer = "Golang梦工厂"

	nAgePointer := (*uint32)(unsafe.Pointer(uintptr(userNamePointer) + unsafe.Offsetof(user.Age)))
	*nAgePointer = 25

	nGender := (*bool)(unsafe.Pointer(uintptr(userNamePointer) + unsafe.Offsetof(user.Gender)))
	*nGender = false

	// u.Name: Golang梦工厂, u.Age: 25,  u.Gender: false
	fmt.Printf("u.Name: %s, u.Age: %d,  u.Gender: %v\n", user.Name, user.Age, user.Gender)
	fmt.Println("============================")
	// Alignof
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64
	var f32 float32
	var s string
	var m map[string]string
	var p *int32

	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(i8))
	fmt.Println(unsafe.Alignof(i16))
	fmt.Println(unsafe.Alignof(i64))
	fmt.Println(unsafe.Alignof(f32))
	fmt.Println(unsafe.Alignof(s))
	fmt.Println(unsafe.Alignof(m))
	fmt.Println(unsafe.Alignof(p))
}
