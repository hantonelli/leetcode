package reorder_list

import "github.com/hantonelli/leetcode/ds"

func reorderListFast(head *ds.ListNode) {
	if head == nil {
		return
	}
	dummy := &ds.ListNode{
		Next: head,
	}
	// find half
	back := dummy
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		back = back.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	back.Next = nil

	if fast == slow {
		return
	}

	// revert
	var last *ds.ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = last
		last = slow
		slow = next
	}

	// merge
	fast = dummy.Next
	fastTail := fast
	for fast != nil && last != nil {
		fastNext := fast.Next
		lastNext := last.Next

		fast.Next = last
		last.Next = fastNext
		fastTail = last

		fast = fastNext
		last = lastNext
	}

	// deal with tail
	fastTail.Next = last

	head = dummy.Next
}
