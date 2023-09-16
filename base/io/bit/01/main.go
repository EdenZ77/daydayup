package main

import "fmt"

/*

 */

type ByteSize float64

const (
	B  ByteSize = 1 << (10 * iota) // 1<<(10*0)
	KB                             // 1<<(10*1) 左移动10位  2的10次方=1024
	MB                             // 1<<(10*2)
	GB                             // 1<<(10*3)
	TB                             // 1<<(10*4)
	PB                             //  1<<(10*5)
)

/*
这里 Like 为 1，Collect 为 2，Comment 为 4。
转为二进制以后分别为 001, 010, 100。作为 2 的 n次方，我们保证了只有一位是 1，其他都是 0。
也就意味着，当我把它们【或】起来的时候，在一个最终的结果里，实际上包含了他们的每个标志位的信息。
*/
const (
	Like = 1 << iota
	Collect
	Comment
)

func main() {
	// 看代码就一目了然了，原来的 001, 010, 100 三个数字进行【或运算】之后，得到了 111 （按十进制的话是 7）。每一个原始为 1 的位，就是它们自己的标志位。
	ability := Like | Collect | Comment

	fmt.Printf("%b\n", ability) // 111

	// 所以，当我们用 ability & Like 时，由于 Like 的非标志位都是 0，标志位是 1，
	// 也就意味着，若 ability 具备了 Like 的标志位能力，则这一次 & 过后，结果必然和 Like 是完全相等的。
	// 这就是为什么，当我们将三个枚举都【或】进 ability 时，再去【与】操作，会发现相等。
	fmt.Println((ability & Like) == Like)       // true
	fmt.Println((ability & Collect) == Collect) // true
	fmt.Println((ability & Comment) == Comment) // true

	// 假如我们稍微做一点改动，ability 不把 Collect 或进去了，你会发现结果也变了：
	ability2 := Like | Comment

	fmt.Printf("%b\n", ability2) // 101

	fmt.Println((ability2 & Like) == Like)       // true
	fmt.Println((ability2 & Collect) == Collect) // false
	fmt.Println((ability2 & Comment) == Comment) // true

	// ok，或进去了，就代表有能力了。那我如果想把某个能力下掉呢？
	// 比如原来 ability 是同时包含了 Like, Comment，现在这个用户不喜欢这篇文章了，把点赞去掉了，那我 ability 应该怎么搞？
	// 思考一下，此时 aibility 应该是 111，而 Like 是 001，我的目标是把 ability 变成 110 （这样就去掉了 Like 对应的标志位）。
	/*
		111
		001
		---
		110
	*/
	// 发现了么？这不就是【异或】嘛。若位相同则为0，若不同则为 1。 此时我们只需要执行一次
	// ability = ability ^ Like 或者用简洁的写法：ability ^= Like

	/*
		这样就够了。总结一下：

		定义 int 枚举值时注意二进制错位，每个枚举有一个自己的标志位；
		用 | 操作添加能力；
		用 ^ 操作下掉能力；
		用 & 操作校验是否具备对应的能力。
	*/

}

func testDelete() {
	ability := Like | Comment

	fmt.Printf("%b\n", ability) // 101

	fmt.Println((ability & Like) == Like)       // true
	fmt.Println((ability & Collect) == Collect) // false
	fmt.Println((ability & Comment) == Comment) // true

	ability = ability ^ Like
	fmt.Printf("%b\n", ability)           // 100
	fmt.Println((ability & Like) == Like) // false
}
