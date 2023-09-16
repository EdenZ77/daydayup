package lru_cache

/*
LRU的算法思路是真的简单，概括下： 使用定长链表来保存所有缓存的值，并且最老的值放在链表最后面
当访问的值在链表中时： 将找到链表中值将其删除，并重新在链表头添加该值（保证链表中 数值的顺序是从新到旧）
当访问的值不在链表中时：
	当链表已满：删除链表最后一个值，将要添加的值放在链表头
	当链表未满：直接在链表头添加

我们也可以通过数组实现LRU算法
*/

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev *lruNode
	next *lruNode

	key   int // lru key
	value int // lru value

	hnext *lruNode
}
