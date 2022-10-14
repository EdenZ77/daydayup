package test02

type Slice[T int | float32 | float64] []T

var aSlice Slice[int] = []int{1, 2, 3}

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

var a MyMap[string, float64] = map[string]float64{
	"jack_score": 9.6,
	"bob_score":  8.4,
}

type MyStruct[T int | string] struct {
	Name string
	Data T
}

type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

type MyChan[T int | string] chan T

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

type NewType[T interface{ *int }] []T

type NewType2[T interface{ *int | *float64 }] []T

type NewType3[T *int,] []T

type Wow[T int | string] int

var aWow Wow[int] = 123
var bWow Wow[string] = 123

// 编译错误，因为“hello”不能赋值给底层类型int
//var cWow Wow[string] = "hello"

func testMap() {

}
