package main

// 92. Reverse Linked List II, https://leetcode.com/problems/reverse-linked-list-ii/

/**
 * Definition for singly-linked list.
 */
type ListNodeForP00092 struct {
	Val  int
	Next *ListNodeForP00092
}

func reverseBetween(head *ListNodeForP00092, m int, n int) *ListNodeForP00092 {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	dummy := new(ListNodeForP00092)
	dummy.Next = head
	prev := dummy
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	prev.Next = doReverseBetween(prev.Next, m, n)
	return dummy.Next
}

func doReverseBetween(head *ListNodeForP00092, m int, n int) *ListNodeForP00092 {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	prev, curr, next := new(ListNodeForP00092), head, new(ListNodeForP00092)
	for i := 0; i < n-m+1; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next = next
	return prev
}
