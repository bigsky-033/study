package main

// 2. Add Two Numbers, https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 */
type ListNodeForP00002 struct {
	Val  int
	Next *ListNodeForP00002
}

func addTwoNumbers(l1 *ListNodeForP00002, l2 *ListNodeForP00002) *ListNodeForP00002 {
	dummy := new(ListNodeForP00002)
	curr := dummy
	carry := 0
	var p1, p2 *ListNodeForP00002

	for p1, p2 = l1, l2; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		sum := p1.Val + p2.Val + carry
		carry = sum / 10
		node := &ListNodeForP00002{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}

	for ; p1 != nil; p1 = p1.Next {
		sum := p1.Val + carry
		carry = sum / 10
		node := &ListNodeForP00002{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	for ; p2 != nil; p2 = p2.Next {
		sum := p2.Val + carry
		carry = sum / 10
		node := &ListNodeForP00002{Val: sum % 10, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	if carry != 0 {
		node := &ListNodeForP00002{Val: carry, Next: nil}
		curr.Next = node
		curr = curr.Next
	}
	return dummy.Next
}
