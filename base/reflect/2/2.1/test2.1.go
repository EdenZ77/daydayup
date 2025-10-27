package main

import (
	"fmt"
	"hello/base/reflect/2/2.1/pkgpath"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	MethodByNameTest()
}

func pointerTest() {
	// 声明一个空结构体
	type cat struct {
	}
	// ins 是一个指向 cat 结构体的指针（类型为 *cat）
	// ins 不是结构体本身，而是一个指针（内存地址）。
	ins := &cat{}
	// 获取 ins 的反射类型
	typeOfCat := reflect.TypeOf(ins)
	// Name(): 指针类型（如 *cat）没有名称，所以返回空字符串 ""。
	// Kind(): 指针的种类是 reflect.Ptr（输出为 ptr）。
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kind:'ptr'
	// 获取指针指向的实际类型
	typeOfCat = typeOfCat.Elem()
	// element name: 'cat', element kind: 'struct'
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

func noPointerTest() {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := cat{}
	typeOfCat := reflect.TypeOf(ins)
	// name:'cat' kind:'struct'
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// panic: reflect: Elem of invalid type main.cat
	typeOfCat = typeOfCat.Elem()
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

type Enum int

const (
	Zero Enum = 0
)

func enmuTest() {
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // Enum int
}

func elemMapTest() {
	m := map[string]int{}
	t := reflect.TypeOf(m)
	keyType := t.Key()    // string
	valueType := t.Elem() // int
	fmt.Printf("keyType: %v\n", keyType)
	fmt.Printf("valueType: %v\n", valueType)
}

func elemManyTest() {
	var pp **int
	t := reflect.TypeOf(pp) // **int
	//t.Elem()                // *int (第一层解引用)
	//t.Elem().Elem()         // int  (第二层解引用)
	// elem: *int, elem kind: ptr
	fmt.Printf("elem: %v, elem kind: %v\n", t.Elem(), t.Elem().Kind())
	// elem.elem: int, elem.elem kind: int
	fmt.Printf("elem.elem: %v, elem.elem kind: %v\n", t.Elem().Elem(), t.Elem().Elem().Kind())
}

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}

func SizeofAndAlignofTest() {
	// 使用 unsafe.Sizeof 计算出一个数据类型实例需要占用的字节数。
	fmt.Println(unsafe.Sizeof(Args{})) // 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 8

	// Alignof 方法，可以返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数。
	fmt.Println(unsafe.Alignof(Args{})) // 8
	fmt.Println(unsafe.Alignof(Flag{})) // 4

	fmt.Println("reflect==")
	args := reflect.TypeOf(Args{})
	// Size()与unsafe.Sizeof作用一样
	// Align()与unsafe.Alignof作用一样
	fmt.Printf("%-15s Size:%-2d Align:%d\n", args.Kind(), args.Size(), args.Align())

	flag := reflect.TypeOf(Flag{})
	fmt.Printf("%-15s Size:%-2d Align:%d\n", flag.Kind(), flag.Size(), flag.Align())
}

type MyStruct struct{}

func (m MyStruct) MyMethod() {}

func MethodTest() {
	t := reflect.TypeOf(MyStruct{})
	// 返回类型方法集中指定索引位置的方法（从 0 开始）
	// i int：方法的索引位置（范围必须是 [0, NumMethod()-1]）
	method := t.Method(0)    // 获取第一个方法
	fmt.Println(method.Name) // 输出: MyMethod
}

type MyInterface interface {
	InterfaceMethod(int, string)
}

func MethodByNameTest() {
	t := reflect.TypeOf((*MyInterface)(nil)).Elem()
	// 按名称查找方法集中的方法
	method, found := t.MethodByName("InterfaceMethod")
	if found {
		fmt.Println("Signature:", method.Type) // Signature: func(int, string)
	}
}

type MyStruct2 struct{}

func (m MyStruct2) Method1()    {}
func (m *MyStruct2) Method2()   {}
func (m MyStruct2) unexported() {} // 未导出方法

// NumMethodTest 返回可通过 Method 方法访问的方法数量
func NumMethodTest() {
	tValue := reflect.TypeOf(MyStruct2{})
	fmt.Println(tValue.NumMethod()) // 输出: 1 (只有 Method1)

	tPtr := reflect.TypeOf(&MyStruct2{})
	fmt.Println(tPtr.NumMethod()) // 输出: 2 (Method1 和 Method2)
}

func pkgPathTest() {
	// 返回该类型声明所在包的导入路径
	pkgPath := reflect.TypeOf(mypkgpath.MyType(0)).PkgPath()
	// hello/base/reflect/2/2.1/pkgpath
	fmt.Println(pkgPath)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func fieldTest() {
	t := reflect.TypeOf(User{})

	// 获取第一个字段
	idField := t.Field(0)
	fmt.Println("Field 0:", idField.Name) // 输出: ID

	// 获取字段标签
	jsonTag := idField.Tag.Get("json")
	fmt.Println("JSON Tag:", jsonTag) // 输出: id
}

type Address struct {
	City string
}

type User1 struct {
	Name    string
	Address // 嵌入结构体
}

// 作用：通过索引路径获取嵌套字段
// 参数：索引切片（表示字段的嵌套路径）
// 返回值：嵌套字段的 StructField 信息
// 等效于连续调用 Field()方法，用于访问嵌入结构体中的字段

func fieldByIndexTest() {
	t := reflect.TypeOf(User1{})

	// 获取嵌入的 City 字段
	cityField := t.FieldByIndex([]int{1, 0})
	fmt.Println("Nested Field:", cityField.Name) // 输出: City
}

type User2 struct {
	ID   int
	Name string
}

/*
作用：通过字段名查找字段
参数：字段名称（区分大小写）
返回值：
1.找到的字段信息
2.是否找到的布尔值
*/
func fieldByNameTest() {
	t := reflect.TypeOf(User2{})

	// 查找存在的字段
	_, found := t.FieldByName("Name")
	fmt.Println("Found Name?", found) // true

	// 查找不存在的字段
	_, found = t.FieldByName("Age")
	fmt.Println("Found Age?", found) // false
}

type Contact struct {
	Email string
}

type User3 struct {
	Name    string
	Contact // 嵌入结构体
}

/*
作用：使用自定义匹配函数查找字段
参数：字段名匹配函数
返回值：
1.第一个匹配的字段信息
2.是否找到的布尔值

使用自定义逻辑匹配字段名
在多个匹配字段冲突时返回 false
*/
func fieldByNameFuncTest() {
	t := reflect.TypeOf(User3{})

	// 查找包含 "mail" 的字段
	field, found := t.FieldByNameFunc(func(s string) bool {
		return strings.Contains(strings.ToLower(s), "mail")
	})

	if found {
		fmt.Println("Found field:", field.Name) // 输出: Email
	}
}
