package dataStructures

// // Queue is a queue
// type Queue interface {
// 	Front() *list.Element
// 	Len() int
// 	Add(interface{})
// 	Remove()
// }

// type queueImpl struct {
// 	*list.List
// }

// func (q *queueImpl) Add(v interface{}) {
// 	if v == nil {

// 	}
// 	q.PushBack(v)
// }

// func (q *queueImpl) Pop() interface{} {
// 	if q.Len() == 0 {
// 		return nil
// 	}
// 	e := q.Front()
// 	q.List.Remove(e)
// 	return e
// }

// // New is a new instance of a Queue
// func New() Queue {
// 	return &queueImpl{list.New()}
// }
