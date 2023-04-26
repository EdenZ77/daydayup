package main

import "fmt"

/*
参考资料：https://mp.weixin.qq.com/s/QFr_Pgt9GzOG3zlgcXEVwQ
*/

type WeekDay int

const (
	Sunday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Monday
)

func main() {
	//fmt.Println(Sunday.Name())
	//fmt.Println(Sunday.Original())
	//fmt.Println(Sunday)
	//fmt.Println(Values())

	//valueOf, err := ValueOf("Monday")
	//if err != nil {
	//	fmt.Printf("%#v\n", err)
	//}
	//fmt.Println(valueOf)
}

// Name 返回枚举值的名称
func (w WeekDay) Name() string {
	if w < Sunday || w > Monday {
		return "Unknown"
	}
	// 让编译器根据初始值的个数自行推断数组的长度
	return [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[w]
}

// Original 由于这个枚举类型的枚举值本质上是整数常量，因此我们可以直接使用枚举值作为它的序号
func (w WeekDay) Original() int {
	return int(w)
}

// 将枚举值转成字符串，便于输出
func (w WeekDay) String() string {
	return w.Name()
}

// Values 返回一个包含所有枚举值的切片 输出：[Sunday Tuesday Wednesday Thursday Friday Saturday Monday]
func Values() []WeekDay {
	return []WeekDay{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
}

// ValueOf 使用 switch 语句来根据名称返回对应的枚举值
func ValueOf(name string) (WeekDay, error) {
	switch name {
	case "Sunday":
		return Sunday, nil
	case "Monday":
		return Monday, nil
	case "Tuesday":
		return Tuesday, nil
	case "Wednesday":
		return Wednesday, nil
	case "Thursday":
		return Thursday, nil
	case "Friday":
		return Friday, nil
	case "Saturday":
		return Saturday, nil
	default:
		return 0, fmt.Errorf("invalid WeekDay name: %s", name)
	}
}
