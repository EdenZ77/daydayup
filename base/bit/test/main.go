package main

import "fmt"

// 1. 左右移等价于 2 的 n 次方运算=================
/*
i << n 相当于 i 乘以 2 的 n次方
i >> n 相当于 i 除以 2 的 n 次方
*/

type ByteSize float64

const (
	B  ByteSize = 1 << (10 * iota) // 1<<(10*0) = 1
	KB                             // 1<<(10*1) 左移动10位  2的10次方=1024
	MB                             // 1<<(10*2)
	GB                             // 1<<(10*3)
	TB                             // 1<<(10*4)
	PB                             //  1<<(10*5)
)

// 2. 判断奇偶数=================
/*
奇偶数的本质区别就是看是否为 2 的倍数，直白点讲，在二进制的世界里，看最低位是 0 还是 1 即可。一个 & 就能解决问题：
奇数：i&1 == 1
偶数：i&1 == 0
*/

// 3. 异或的特性=================
/*
i ^ i == 0
i ^ 0 == i
两个相同的数异或的结果是 0，一个数和 0 异或的结果是它本身。
实际上，我们经常用的 staticcheck 就内置了这里的规则，如果你写出了类似 i ^ i 的代码，会得到
identical expressions on the left and right side of the '^' operator (SA4000)
*/

// 4. 枚举=================重点！！！
/*

这样就够了。总结一下：

定义 int 枚举值时注意二进制错位，每个枚举有一个自己的标志位；
用 | 操作添加能力；
用 ^ 操作下掉能力；
用 & 操作校验是否具备对应的能力。
*/

const (
	Like    = 1 << iota // 1的0次方，1
	Collect             // 2
	Comment             // 4
)

func main() {
	ability := Like | Comment

	fmt.Printf("%b\n", ability) // 101

	fmt.Println((ability & Like) == Like)       // true
	fmt.Println((ability & Collect) == Collect) // false
	fmt.Println((ability & Comment) == Comment) // true

	ability = ability ^ Like
	fmt.Printf("%b\n", ability)           // 100
	fmt.Println((ability & Like) == Like) // false
}

/*
负数以补码形式表示
123 (decimal) = 01111011 (binary)
-123 (decimal) = 10000101 (binary, 8-bit two's complement)

<<：左移操作，将二进制数向左移动，右边补0。
>>：右移操作，通常是算术右移（Arithmetic Right Shift），保留符号位，左边补符号位。
>>> 表示逻辑右移（Logical Right Shift），左边补0。


01111011 << 1 = 11110110 (binary)
10000101 << 1 = 00001010 (binary, ignoring sign extension)

01111011 >> 1 = 00111101 (binary)
10000101 >> 1 = 11000010 (binary, sign extension)

01111011 >>> 1 = 00111101 (binary)
10000101 >>> 1 = 01000010 (binary, ignoring sign extension)

*/
