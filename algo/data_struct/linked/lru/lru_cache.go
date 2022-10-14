package lru_cache

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev *lruNode
	next *lruNode

	key   int
	value int

	hnext *lruNode
}

type LRUCache struct {
	node []lruNode

	head *lruNode
	tail *lruNode

	capacity int
	used     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (l *LRUCache) Get(key int) int {
	return -1
}
