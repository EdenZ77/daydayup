package main

import (
	"fmt"
	"math"
)

// 打印 int32 的二进制表示（32位）
func printBinary(num int32) {
	for i := 31; i >= 0; i-- {
		if (num & (1 << i)) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Println()
}

func main() {
	// 非负数
	var a int32 = 78
	fmt.Println(a)
	printBinary(a)
	fmt.Println("===a===")

	// 负数
	var b int32 = -6
	fmt.Println(b)
	printBinary(b)
	fmt.Println("===b===")

	// 二进制字面量
	c := int32(0b1001110)
	fmt.Println(c)
	printBinary(c)
	fmt.Println("===c===")

	// 十六进制字面量
	d := int32(0x4e)
	fmt.Println(d)
	printBinary(d)
	fmt.Println("===d===")

	// ~、相反数
	fmt.Println(a)
	printBinary(a)
	printBinary(^a) // Go 的 ^ 是取反运算符
	e := ^a + 1
	fmt.Println(e)
	printBinary(e)
	fmt.Println("===e===")

	// int32 最小值
	f := int32(math.MinInt32)
	fmt.Println(f)
	printBinary(f)
	fmt.Println(-f)
	printBinary(-f) // 需要转换类型
	fmt.Println(^f + 1)
	printBinary(^f + 1)
	fmt.Println("===f===")

	// | & ^
	g := int32(0b0001010)
	h := int32(0b0001100)
	printBinary(g | h)
	printBinary(g & h)
	printBinary(g ^ h) // Go 的 ^ 作为二元运算符时是异或
	fmt.Println("===g、h===")

	// 逻辑运算符测试

	// <<
	i := int32(0b0011010)
	printBinary(i)
	printBinary(i << 1)
	printBinary(i << 2)
	printBinary(i << 3)
	fmt.Println("===i << ===")

	// >> 和 >>>
	printBinary(i)
	printBinary(i >> 2)
	printBinary(int32(uint32(i) >> 2)) // 无符号右移
	fmt.Println("===i >> >>>===")

	// 负数移位
	// 修复：使用正确的十六进制表示法
	j := int32(-0x10000000) // 相当于二进制 11110000000000000000000000000000
	printBinary(j)
	printBinary(j >> 2)
	printBinary(int32(uint32(j) >> 2)) // 无符号右移
	fmt.Println("===j >> >>>===")

	// 非负数 << 1，等同于乘以2
	// 非负数 << 2，等同于乘以4
	// 非负数 << 3，等同于乘以8
	// 非负数 << i，等同于乘以2的i次方
	// ...
	// 非负数 >> 1，等同于除以2
	// 非负数 >> 2，等同于除以4
	// 非负数 >> 3，等同于除以8
	// 非负数 >> i，等同于除以2的i次方
	// 只有非负数符合这个特征，负数不要用

	// 移位与乘除法
	k := int32(10)
	fmt.Println(k)
	fmt.Println(k << 1)
	fmt.Println(k << 2)
	fmt.Println(k << 3)
	fmt.Println(k >> 1)
	fmt.Println(k >> 2)
	fmt.Println(k >> 3)
	fmt.Println("===k===")
}
