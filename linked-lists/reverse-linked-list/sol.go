package reverse_linked_list

import "github.com/hantonelli/leetcode/ds"

func reverseList(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nextNode, prevNode *ds.ListNode
	currNode := head
	for currNode.Next != nil {
		nextNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	currNode.Next = prevNode
	return currNode
}
