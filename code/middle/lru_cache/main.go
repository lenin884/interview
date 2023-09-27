package lru_cache

type ListNode struct {
	Prev  *ListNode
	Next  *ListNode
	Key   int
	Value int
}

type LRUCache struct {
	hashmap map[int]*ListNode
	cap     int
	left    *ListNode
	right   *ListNode
}

func Constructor(capacity int) LRUCache {
	left, right := &ListNode{}, &ListNode{}
	left.Next, right.Prev = right, left
	return LRUCache{
		hashmap: make(map[int]*ListNode, capacity),
		left:    left,
		right:   right,
		cap:     capacity,
	}
}

func (l *LRUCache) remove(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRUCache) insert(node *ListNode) {
	node.Prev = l.left
	node.Next = l.left.Next
	l.left.Next.Prev = node
	l.left.Next = node
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hashmap[key]; ok {
		l.remove(node)
		l.insert(node)
		return node.Value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.hashmap[key]; ok {
		l.remove(node)
	}

	l.hashmap[key] = &ListNode{
		Key:   key,
		Value: value,
	}
	l.insert(l.hashmap[key])

	if len(l.hashmap) > l.cap {
		delete(l.hashmap, l.right.Prev.Key)
		l.remove(l.right.Prev)
	}

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
