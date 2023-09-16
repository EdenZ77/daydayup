package main

// 这段代码把类型约束给单独拿出来，写入了接口类型 IntUintFloat 当中。需要指定类型约束的时候直接使用接口 IntUintFloat 即可。
//type IntUintFloat interface {
//	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
//}
//
//type Slice[T IntUintFloat] []T
//==============================================

// 不过这样的代码依旧不好维护，而接口和接口、接口和普通类型之间也是可以通过 | 进行组合：
// 下面的代码中，我们分别定义了 Int, Uint, Float 三个接口类型，并最终在 Slice[T] 的类型约束中通过使用 | 将它们组合到一起。
//type Int interface {
//	int | int8 | int16 | int32 | int64
//}
//
//type Uint interface {
//	uint | uint8 | uint16 | uint32
//}
//
//type Float interface {
//	float32 | float64
//}
//
//type Slice[T Int | Uint | Float | bool] []T // 使用 '|' 将多个接口类型组合
// ===========================================

// 同时，在接口里也能直接组合其他接口，所以还可以像下面这样：
//type SliceElement interface {
//	Int | Uint | Float | string // 组合了三个接口类型并额外增加了一个 string 类型
//}
//
//type Slice[T SliceElement] []T
// ==============================================

// 上面定义的 Slie[T] 虽然可以达到目的，但是有一个缺点：
// var s1 Slice[int] // 正确
//
// type MyInt int
// var s2 Slice[MyInt] // ✗ 错误。MyInt类型底层类型是int但并不是int类型，不符合 Slice[T] 的类型约束
// 这里发生错误的原因是，泛型类型 Slice[T] 允许的是 int 作为类型实参，而不是 MyInt （虽然 MyInt 类型底层类型是 int ，但它依旧不是 int 类型）。
// 为了从根本上解决这个问题，Go新增了一个符号 ~ ，在类型约束中使用类似 ~int 这种写法的话，就代表着不光是 int ，所有以 int 为底层类型的类型也都可用于实例化。

// 限制：使用 ~ 时有一定的限制：
//	~后面的类型不能为接口
//	~后面的类型必须为基本类型
//type MyInt int
//
//type _ interface {
//	~[]byte        // 正确
//	~MyInt         // 错误，~后的类型必须为基本类型
//	~error         // 错误，~后的类型不能为接口
//	~struct{ int } // 正确
//}

func main() {

}
