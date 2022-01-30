package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	newHead := &Node{}
	currNew := newHead
	mapping := map[*Node]*Node{}

	currOld := head
	for currOld != nil {
		currNew.Next = &Node{
			Val: currOld.Val,
		}
		mapping[currOld] = currNew.Next
		currNew = currNew.Next
		currOld = currOld.Next
	}

	currOld = head
	currNew = newHead.Next
	for currOld != nil {
		if currOld.Random != nil {
			currNew.Random = mapping[currOld.Random]
		}
		currNew = currNew.Next
		currOld = currOld.Next
	}

	return newHead.Next
}
