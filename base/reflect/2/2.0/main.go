package main

import (
	"fmt"
	"reflect"
)

// IUser 定义一个接口
type IUser interface {
	SayHello()
}

// EmptyInterface 定义一个空接口（所有类型都实现空接口）
type EmptyInterface interface{}

// User 定义一个结构体类型
type User struct {
	Name string
}

// SayHello 为 User 类型定义一个方法
func (u *User) SayHello() {
	fmt.Printf("Hello, my name is %s!\n", u.Name)
}

// UpdateName 为 User 类型定义另一个方法
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	// 示例 1: 使用 MethodByName 获取并调用方法
	u := User{Name: "Alice"}

	// 获取 u 的动态类型 (User)
	userType := reflect.TypeOf(u)

	// 查找名为 "SayHello" 的方法
	method, found := userType.MethodByName("SayHello")
	if found {
		fmt.Printf("Found method: %s\n", method.Name)

		// 调用该方法。Method.Func 是一个 reflect.Value，代表该函数。
		// 调用 Call 时，第一个参数必须是方法的接收者 (u)，后面是方法的参数 (SayHello 没有参数)。
		method.Func.Call([]reflect.Value{reflect.ValueOf(u)}) // 输出: Hello, my name is Alice!
	} else {
		fmt.Println("Method 'SayHello' not found on type User")
	}

	// 尝试查找一个不存在的方法
	_, found = userType.MethodByName("NonExistentMethod")
	fmt.Println("Found 'NonExistentMethod'?", found) // 输出: false

	// 注意：MethodByName 只查找 receiver 是值类型 (User) 的方法。
	// UpdateName 的 receiver 是指针 (*User)，所以通过 User 的 Type 找不到它。
	_, found = userType.MethodByName("UpdateName")
	fmt.Println("Found 'UpdateName' on User?", found) // 输出: false

	// 获取 *User 的类型
	userPtrType := reflect.TypeOf(&u)
	// 现在可以在 *User 上找到 UpdateName
	_, found = userPtrType.MethodByName("UpdateName")
	fmt.Println("Found 'UpdateName' on *User?", found) // 输出: true

	// 示例 2: 使用 Implements 判断类型是否实现接口
	// 获取接口 IUser 的 reflect.Type
	// 注意：我们需要一个指向接口的指针，然后取其 Elem() 来获取接口类型本身
	iUserType := reflect.TypeOf((*IUser)(nil)).Elem()
	emptyInterfaceType := reflect.TypeOf((*EmptyInterface)(nil)).Elem()

	// 检查 User 类型是否实现了 IUser 接口
	implementsIUser := userType.Implements(iUserType)
	fmt.Printf("Does User implement IUser? %t\n", implementsIUser) // 输出: true (因为 User 有 SayHello 方法)

	// 检查 *User 类型是否实现了 IUser 接口
	implementsIUser = userPtrType.Implements(iUserType)
	fmt.Printf("Does *User implement IUser? %t\n", implementsIUser) // 输出: true (*User 也能调用 SayHello)

	// 检查 User 类型是否实现了空接口 (总是 true)
	implementsEmpty := userType.Implements(emptyInterfaceType)
	fmt.Printf("Does User implement EmptyInterface? %t\n", implementsEmpty) // 输出: true

	// 检查 int 类型是否实现了 IUser 接口
	intType := reflect.TypeOf(0)
	implementsIUser = intType.Implements(iUserType)
	fmt.Printf("Does int implement IUser? %t\n", implementsIUser) // 输出: false
}
