package p_00001_00100

// 2. Add Two Numbers, https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0
	var p1, p2 *ListNode

	for p1, p2 = l1, l2; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		sum := p1.Val + p2.Val + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}

	for ; p1 != nil; p1 = p1.Next {
		sum := p1.Val + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	for ; p2 != nil; p2 = p2.Next {
		sum := p2.Val + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	if carry != 0 {
		node := &ListNode{Val: carry, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	return dummy.Next
}
