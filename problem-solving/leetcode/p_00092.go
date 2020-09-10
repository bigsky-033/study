package main

// 92. Reverse Linked List II, https://leetcode.com/problems/reverse-linked-list-ii/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	prev := dummy
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	prev.Next = doReverseBetween(prev.Next, m, n)
	return dummy.Next
}

func doReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	prev, curr, next := new(ListNode), head, new(ListNode)
	for i := 0; i < n-m+1; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next = next
	return prev
}
