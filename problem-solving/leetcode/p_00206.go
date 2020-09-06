package main

// 206. Reverse Linked List, https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 */
type ListNodeForP00448 struct {
	Val  int
	Next *ListNodeForP00448
}

func reverseList(head *ListNodeForP00448) *ListNodeForP00448 {
	return reverseListRecursive(head)
}

func reverseListIterative(head *ListNodeForP00448) *ListNodeForP00448 {
	var prev *ListNodeForP00448
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseListRecursive(head *ListNodeForP00448) *ListNodeForP00448 {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
