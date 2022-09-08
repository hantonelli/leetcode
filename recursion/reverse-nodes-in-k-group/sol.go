package reverse_nodes_in_k_group

import "github.com/hantonelli/leetcode/ds"

func reverseKGroup(head *ds.ListNode, k int) *ds.ListNode {
	if head == nil {
		return nil
	}
	nodeIdx := 0
	nodes := make([]*ds.ListNode, k)
	currNode := head
	var newHead, prev *ds.ListNode
	fistReverse := true

	for currNode != nil {
		nodes[nodeIdx] = currNode
		nodeIdx++
		currNode = currNode.Next

		if nodeIdx%k == 0 {
			nodeIdx = 0
			if fistReverse {
				fistReverse = false
				newHead = nodes[k-1]
			}
			nodes[0].Next = nodes[k-1].Next
			for i := k - 1; 0 < i; i-- {
				nodes[i].Next = nodes[i-1]
			}
			if prev != nil {
				prev.Next = nodes[k-1]
			}
			prev = nodes[0]
		}
	}

	if newHead != nil {
		return newHead
	}
	return head
}
