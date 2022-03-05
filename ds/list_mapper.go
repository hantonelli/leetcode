package ds

func ArrOfArrToArrLinkedList(aa [][]int) []*ListNode {
	if len(aa) == 0 {
		return []*ListNode{}
	}
	arrLL := make([]*ListNode, len(aa))
	for i := range aa {
		arrLL[i] = ArrToLinkedList(aa[i])
	}
	return arrLL
}

func ArrToLinkedList(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	root := ListNode{Val: a[0]}
	var prev *ListNode = &root
	var next *ListNode
	for i := 1; i < len(a); i++ {
		next = &ListNode{
			Val: a[i],
		}
		prev.Next = next
		prev = next
	}
	return &root
}

func LinkedListToArr(l *ListNode) []int {
	if l == nil {
		return []int{}
	}
	var i int
	var last *ListNode
	for last != nil {
		i++
		last = last.Next
	}
	arr := make([]int, 0, i)
	i = 0
	for last != nil {
		arr[i] = last.Val
		i++
	}
	return arr
}
