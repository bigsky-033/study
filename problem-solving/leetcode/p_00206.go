package main

// 206. Reverse Linked List, https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reverseListRecursive(head)
}

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
