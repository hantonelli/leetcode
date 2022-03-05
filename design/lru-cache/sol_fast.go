package lrucache

type node struct {
	key, val   int
	prev, next *node
}

type LRUCacheFast struct {
	capacity, current int
	head, tail        *node
	cache             map[int]*node
}

func ConstructorFast(capacity int) LRUCacheFast {
	return LRUCacheFast{
		capacity: capacity,
		cache:    make(map[int]*node),
	}
}

func (this *LRUCacheFast) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		this.remove(n)
		this.append(n)
		return n.val
	} else {
		return -1
	}
}

func (this *LRUCacheFast) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		n.val = value
		this.remove(n)
	} else {
		n = &node{key: key, val: value}
		this.cache[key] = n
		if this.capacity == this.current {
			delete(this.cache, this.head.key)
			this.remove(this.head)
		} else {
			this.current++
		}
	}
	this.append(n)
}

func (this *LRUCacheFast) remove(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if this.head == n {
		this.head = n.next
	}
	if this.tail == n {
		this.tail = n.prev
	}
	n.prev = nil
	n.next = nil
}

func (this *LRUCacheFast) append(n *node) {
	if this.tail == nil {
		this.tail = n
		this.head = n
	} else {
		this.tail.next = n
		n.prev = this.tail
		this.tail = n
	}
}
