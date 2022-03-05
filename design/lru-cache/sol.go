package lrucache

import (
	"container/list"
)

// Great implementations
// https://github.com/hashicorp/golang-lru

// entry is used to hold a value in the evictList
type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	evictList *list.List
	items     map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		evictList: list.New(),
		items:     make(map[int]*list.Element),
	}

}

func (l *LRUCache) Get(key int) int {
	el, ok := l.items[key]
	if !ok {
		//we should be retuning an ok value, but this it the signature that need to be implemented
		return -1 // required return value when not found
	}
	l.evictList.MoveToFront(el)
	// the value can't be nil, so just casting it
	return el.Value.(*entry).value
}

func (l *LRUCache) Put(key int, value int) {
	if el, ok := l.items[key]; ok {
		// Update value
		el.Value.(*entry).value = value
		l.evictList.MoveToFront(el)
		return
	}

	// Add new value
	if l.capacity == len(l.items) {
		elR := l.evictList.Back()
		delete(l.items, elR.Value.(*entry).key)
		l.evictList.Remove(elR)
	}
	ent := &entry{key: key, value: value}
	el := l.evictList.PushFront(ent)
	l.items[key] = el
}
