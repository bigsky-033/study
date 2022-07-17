package p00876middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	size := 0
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}

	numMoves := size / 2
	curr = head
	for numMoves > 0 {
		curr = curr.Next
		numMoves--
	}

	return curr
}
