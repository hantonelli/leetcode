package insertdeletegetrandom

import "math/rand"

type RandomizedSet struct {
	items map[int]int
	list  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		items: map[int]int{},
		list:  []int{},
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.items[val]; ok {
		return false
	}
	r.items[val] = len(r.list)
	r.list = append(r.list, val)
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	deleteListIdx, ok := r.items[val]
	if !ok {
		return false
	}
	if len(r.list) == 1 || val == r.list[len(r.list)-1] {
		r.list = r.list[:len(r.list)-1]
		delete(r.items, val)
		return true
	}

	// Get the last value in the list
	listLastValue := r.list[len(r.list)-1]
	// Remove last element from the list
	r.list = r.list[:len(r.list)-1]

	// Put last value where the deleted element is
	r.list[deleteListIdx] = listLastValue
	// Update the last element value to the position where the deleted element was
	r.items[listLastValue] = deleteListIdx
	// remove deleted value from the map
	delete(r.items, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.list[rand.Intn(len(r.list))]
}
