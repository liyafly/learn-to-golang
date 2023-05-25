package linkedlist

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 头结点不变，设置一个哨兵节点
	dummy := &ListNode{Next: head}

	pre := dummy
	// 找到left前的一个节点，即（left-1）节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	// 反转区间内的节点
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	// 返回链表头
	return dummy.Next
}
