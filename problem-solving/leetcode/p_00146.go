package main

// 146. LRU Cache, https://leetcode.com/problems/lru-cache/

type ListNode struct {
	Key  int
	Val  int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	storage  map[int]*ListNode
	capacity int
	head     *ListNode
	tail     *ListNode
}

func Constructor(capacity int) LRUCache {
	head := new(ListNode)
	tail := new(ListNode)
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		storage:  make(map[int]*ListNode),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	res := -1
	if node, ok := l.storage[key]; ok {
		l.remove(node)
		l.insertToHead(node)
		res = node.Val
	}
	return res
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.storage[key]; ok {
		node.Val = value
		l.remove(node)
		l.insertToHead(node)
	} else {
		if len(l.storage) == l.capacity {
			// evict
			delete(l.storage, l.tail.Prev.Key)
			l.remove(l.tail.Prev)
		}
		newNode := &ListNode{Key: key, Val: value}
		l.insertToHead(newNode)
		l.storage[key] = newNode
	}
}

func (l *LRUCache) remove(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRUCache) insertToHead(node *ListNode) {
	headNext := l.head.Next
	l.head.Next = node
	headNext.Prev = node
	node.Prev = l.head
	node.Next = headNext
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
