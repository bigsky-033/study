package main

// 24. Swap Nodes in Pairs, https://leetcode.com/problems/swap-nodes-in-pairs/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmp := head.Next
	if tmp != nil {
		nextHead := tmp.Next
		tmp.Next = head
		head.Next = swapPairs(nextHead)
		head = tmp
	}
	return head
}
