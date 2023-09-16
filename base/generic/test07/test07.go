package main

/*
从方法集(Method set)到类型集(Type set)
*/

/*
上面的例子中，我们学习到了一种接口的全新写法，而这种写法在Go1.18之前是不存在的。
如果你比较敏锐的话，一定会隐约认识到这种写法的改变这也一定意味着Go语言中 接口(interface) 这个概念发生了非常大的变化。

是的，在Go1.18之前，Go官方对 接口(interface) 的定义是：接口是一个方法集(method set)
就如下面这个代码一样， ReadWriter 接口定义了一个接口(方法集)，这个集合中包含了 Read() 和 Write() 这两个方法。所有同时定义了这两种方法的类型被视为实现了这一接口。
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
但是，我们如果换一个角度来重新思考上面这个接口的话，会发现接口的定义实际上还能这样理解：
我们可以把 ReaderWriter 接口看成代表了一个 类型的集合，所有实现了 Read() Writer() 这两个方法的类型都在接口代表的类型集合当中
通过换个角度看待接口，在我们眼中接口的定义就从 方法集(method set) 变为了 类型集(type set)。而Go1.18开始就是依据这一点将接口的定义正式更改为了 类型集(Type set)

你或许会觉得，这不就是改了下概念上的定义实际上没什么用吗？是的，如果接口功能没变化的话确实如此。但是还记得下面这种用接口来简化类型约束的写法吗：
type Float interface {
    ~float32 | ~float64
}
type Slice[T Float] []T
这就体现出了为什么要更改接口的定义了。用 类型集 的概念重新理解上面的代码的话就是：
接口类型 Float 代表了一个 类型集合， 所有以 float32 或 float64 为底层类型的类型，都在这一类型集之中
而 type Slice[T Float] []T 中， 类型约束 的真正意思是：
类型约束 指定了类型形参可接受的类型集合，只有属于这个集合中的类型才能替换形参用于实例化


================接口实现(implement)定义的变化
既然接口定义发生了变化，那么从Go1.18开始 接口实现(implement) 的定义自然也发生了变化：
当满足以下条件时，我们可以说 类型 T 实现了接口 I ( type T implements interface I)：
	T 不是接口时：类型 T 是接口 I 代表的类型集中的一个成员 (T is an element of the type set of I)
	T 是接口时： T 接口代表的类型集是 I 代表的类型集的子集(Type set of T is a subset of the type set of I)
并集我们已经很熟悉了，之前一直使用的 | 符号就是求类型的并集( union )
type Uint interface {  // 类型集 Uint 是 ~uint 和 ~uint8 等类型的并集
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
接口可以不止书写一行，如果一个接口有多行类型定义，那么取它们之间的 交集
type AllInt interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint32
}
type Uint interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
type A interface { // 接口A代表的类型集是 AllInt 和 Uint 的交集
    AllInt
    Uint
}
type B interface { // 接口B代表的类型集是 AllInt 和 ~int 的交集
    AllInt
    ~int
}
上面这个例子中
	接口 A 代表的是 AllInt 与 Uint 的 交集，即 ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	接口 B 代表的则是 AllInt 和 ~int 的交集，即 ~int


除了上面的交集，下面也是一种交集：
type C interface {
    ~int
    int
}
很显然，~int 和 int 的交集只有int一种类型，所以接口C代表的类型集中只有int一种类型

当多个类型的交集如下面 Bad 这样为空的时候， Bad 这个接口代表的类型集为一个空集：
type Bad interface {
    int
    float32
} // 类型 int 和 float32 没有相交的类型，所以接口 Bad 代表的类型集为空
没有任何一种类型属于空集。虽然 Bad 这样的写法是可以编译的，但实际上并没有什么意义


上面说了空集，接下来说一个特殊的类型集——空接口 interface{} 。因为，Go1.18开始接口的定义发生了改变，所以 interface{} 的定义也发生了一些变更：空接口代表了所有类型的集合
所以，对于Go1.18之后的空接口应该这样理解：
	1、虽然空接口内没有写入任何的类型，但它代表的是所有类型的集合，而非一个 空集
	2、类型约束中指定 空接口 的意思是指定了一个包含所有类型的类型集，并不是类型约束限定了只能使用 空接口 来做类型实参
因为空接口是一个包含了所有类型的类型集，所以我们经常会用到它。于是，Go1.18开始提供了一个和空接口 interface{} 等价的新关键词 any ，用来使代码更简单：
type Slice[T any] []T // 代码等价于 type Slice[T interface{}] []T
实际上 any 的定义就位于Go语言的 builtin.go 文件中（参考如下）， any 实际上就是 interaface{} 的别名(alias)，两者完全等价
所以从 Go 1.18 开始，所有可以用到空接口的地方其实都可以直接替换为any，如：
var s []any // 等价于 var s []interface{}
var m map[string]any // 等价于 var m map[string]interface{}

comparable(可比较) 和 可排序(ordered)
对于一些数据类型，我们需要在类型约束中限制只接受能 != 和 == 对比的类型，如map：
// 错误。因为 map 中键的类型必须是可进行 != 和 == 比较的类型
type MyMap[KEY any, VALUE any] map[KEY]VALUE
所以Go直接内置了一个叫 comparable 的接口，它代表了所有可用 != 以及 == 对比的类型：
type MyMap[KEY comparable, VALUE any] map[KEY]VALUE // 正确
comparable 比较容易引起误解的一点是很多人容易把他与可排序搞混淆。可比较指的是 可以执行 != == 操作的类型，并没确保这个类型可以执行大小比较（ >,<,<=,>= ）。


================接口两种类型
我们接下来再观察一个例子，这个例子是阐述接口是类型集最好的例子：
type ReadWriter interface {
    ~string | ~[]rune

    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
接口类型 ReadWriter 代表了一个类型集合，所有以 string 或 []rune 为底层类型，并且实现了 Read() Write() 这两个方法的类型都在 ReadWriter 代表的类型集当中
如下面代码中，StringReadWriter 存在于接口 ReadWriter 代表的类型集中，而 BytesReadWriter 因为底层类型是 []byte（既不是string也是不[]rune） ，所以它不属于 ReadWriter 代表的类型集
// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string
func (s StringReadWriter) Read(p []byte) (n int, err error) {
    // ...
}
func (s StringReadWriter) Write(p []byte) (n int, err error) {
 // ...
}

//  类型BytesReadWriter 没有实现接口 Readwriter
type BytesReadWriter []byte
func (s BytesReadWriter) Read(p []byte) (n int, err error) {
 ...
}
func (s BytesReadWriter) Write(p []byte) (n int, err error) {
 ...
}
你一定会说，啊等等，这接口也变得太复杂了把，那我定义一个 ReadWriter 类型的接口变量，然后接口变量赋值的时候不光要考虑到方法的实现，还必须考虑到具体底层类型？
心智负担也太大了吧。是的，为了解决这个问题也为了保持Go语言的兼容性，Go1.18开始将接口分为了两种类型
	基本接口(Basic interface)
	一般接口(General interface)

接口定义中如果只有方法的话，那么这种接口被称为基本接口(Basic interface)。这种接口就是Go1.18之前的接口，用法也基本和Go1.18之前保持一致。
基本接口大致可以用于如下几个地方：
1、最常用的，定义接口变量并赋值
type MyError interface { // 接口中只有方法，所以是基本接口
    Error() string
}
// 用法和 Go1.18之前保持一致
var err MyError = fmt.Errorf("hello world")
2、基本接口因为也代表了一个类型集，所以也可用在类型约束中
// io.Reader 和 io.Writer 都是基本接口，也可以用在类型约束中
type MySlice[T io.Reader | io.Writer]  []Slice

如果接口内不光只有方法，还有类型的话，这种接口被称为 一般接口(General interface) ，如下例子都是一般接口：
type Uint interface { // 接口 Uint 中有类型，所以是一般接口
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
type ReadWriter interface {  // ReadWriter 接口既有方法也有类型，所以是一般接口
    ~string | ~[]rune

    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
一般接口类型不能用来定义变量，只能用于泛型的类型约束中。所以以下的用法是错误的：
type Uint interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
var uintInf Uint // 错误。Uint是一般接口，只能用于类型约束，不得用于变量定义
这一限制保证了一般接口的使用被限定在了泛型之中，不会影响到Go1.18之前的代码，同时也极大减少了书写代码时的心智负担

// =======================================泛型接口
type DataProcessor[T any] interface {
    Process(oriData T) (newData T)
    Save(data T) error
}
type DataProcessor2[T any] interface {
    int | ~struct{ Data interface{} }

    Process(data T) (newData T)
    Save(data T) error
}
因为引入了类型形参，所以这两个接口是泛型类型。而泛型类型要使用的话必须传入类型实参实例化才有意义。所以我们来尝试实例化一下这两个接口。
因为 T 的类型约束是 any，所以可以随便挑一个类型来当实参(比如string)：
DataProcessor[string]
// 实例化之后的接口定义相当于如下所示：
type DataProcessor[string] interface {
    Process(oriData string) (newData string)
    Save(data string) error
}
经过实例化之后就好理解了， DataProcessor[string] 因为只有方法，所以它实际上就是个 基本接口(Basic interface)，
这个接口包含两个能处理string类型的方法。像下面这样实现了这两个能处理string类型的方法就算实现了这个接口：
type CSVProcessor struct {
}
// 注意，方法中 oriData 等的类型是 string
func (c CSVProcessor) Process(oriData string) (newData string) {
    ....
}
func (c CSVProcessor) Save(oriData string) error {
    ...
}

// CSVProcessor实现了接口 DataProcessor[string] ，所以可赋值
var processor DataProcessor[string] = CSVProcessor{}
processor.Process("name,age\nbob,12\njack,30")
processor.Save("name,age\nbob,13\njack,31")

// 错误。CSVProcessor没有实现接口 DataProcessor[int]
var processor2 DataProcessor[int] = CSVProcessor{}

再用同样的方法实例化 DataProcessor2[T] ：
DataProcessor2[string]
// 实例化后的接口定义可视为
type DataProcessor2[T string] interface {
    int | ~struct{ Data interface{} }

    Process(data string) (newData string)
    Save(data string) error
}
DataProcessor2[string] 因为带有类型并集所以它是 一般接口(General interface)，所以实例化之后的这个接口代表的意思是：
	1、只有实现了 Process(string) string 和 Save(string) error 这两个方法，并且以 int 或 struct{ Data interface{} } 为底层类型的类型才算实现了这个接口
	2、一般接口(General interface) 不能用于变量定义只能用于类型约束，所以接口 DataProcessor2[string] 只是定义了一个用于类型约束的类型集
// XMLProcessor 虽然实现了接口 DataProcessor2[string] 的两个方法，但是因为它的底层类型是 []byte，所以依旧是未实现 DataProcessor2[string]
type XMLProcessor []byte
func (c XMLProcessor) Process(oriData string) (newData string) {}
func (c XMLProcessor) Save(oriData string) error {}

// JsonProcessor 实现了接口 DataProcessor2[string] 的两个方法，同时底层类型是 struct{ Data interface{} }。所以实现了接口 DataProcessor2[string]
type JsonProcessor struct {
    Data interface{}
}
func (c JsonProcessor) Process(oriData string) (newData string) {}
func (c JsonProcessor) Save(oriData string) error {}

// 错误。DataProcessor2[string]是一般接口不能用于创建变量
var processor DataProcessor2[string]

// 正确，实例化之后的 DataProcessor2[string] 可用于泛型的类型约束
type ProcessorList[T DataProcessor2[string]] []T

// 正确，接口可以并入其他接口
type StringProcessor interface {
    DataProcessor2[string]
    PrintString()
}

// 错误，带方法的一般接口不能作为类型并集的成员(参考6.5 接口定义的种种限制规则
type StringProcessor interface {
    DataProcessor2[string] | DataProcessor2[[]byte]

    PrintString()
}

// =================接口定义的种种限制规则

5、带方法的接口(无论是基本接口还是一般接口)，都不能写入接口的并集中：
type _ interface {
    ~int | ~string | error // 错误，error是带方法的接口(一般接口) 不能写入并集中
}
type DataProcessor[T any] interface {
    ~string | ~[]byte

    Process(data T) (newData T)
    Save(data T) error
}
// 错误，实例化之后的 DataProcessor[string] 是带方法的一般接口，不能写入类型并集
type _ interface {
    ~int | ~string | DataProcessor[string]
}
type Bad[T any] interface {
    ~int | ~string | DataProcessor[T]  // 也不行
}


*/

// ===============测试验证一般接口不能用来定义变量
type ReadWriterx interface {
	~string | ~[]rune

	Readx(p []byte) (n int, err error)
	Writex(p []byte) (n int, err error)
}

type StringReadWriterx string

func (s StringReadWriterx) Readx(p []byte) (n int, err error) {
	return 0, nil
}
func (s StringReadWriterx) Writex(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	/*
		如果接口内不光只有方法，还有类型的话，这种接口被称为 一般接口(General interface)，如下例子：
			type Uint interface { // 接口 Uint 中有类型，所以是一般接口
				~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
			}

			type ReadWriter interface {  // ReadWriter 接口既有方法也有类型，所以是一般接口
				~string | ~[]rune

				Read(p []byte) (n int, err error)
				Write(p []byte) (n int, err error)
			}

		一般接口类型不能用来定义变量，只能用于泛型的类型约束中。所以以下的用法是错误的：
			type Uint interface {
			    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
			}

			var uintInf Uint // 错误。Uint是一般接口，只能用于类型约束，不得用于变量定义
	*/
	// 这一限制保证了一般接口的使用被限定在了泛型之中，不会影响到Go1.18之前的代码，同时也极大减少了书写代码时的心智负担

	//var xx ReadWriterx = StringReadWriterx("xx") // × Interface includes constraint elements '~string', '~[]rune', can only be used in type parameters
}

// type MyInt int

type MyStruct struct {
	name string
}

type xx struct{ Data interface{} }

type _ interface {
	//~xx // 错误
	//~MyStruct // 错误
	~[]byte // 正确
	~struct{ Data interface{} }
	//~MyInt // 错误，~后的类型必须为基本类型
	//~error // 错误，~后的类型不能为接口
}
