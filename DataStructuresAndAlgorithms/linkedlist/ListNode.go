package linkedlist

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

// NewListNode Create new node.
func NewListNode(val int) *ListNode {
	return &ListNode{val, nil, nil}
}
