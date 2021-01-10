package p_00001_00100

// 83. Remove Duplicates from Sorted List, https://leetcode.com/problems/remove-duplicates-from-sorted-list/

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	next := cur.Next
	for next != nil {
		if cur.Val != next.Val {
			cur.Next = next
			cur = next
			next = next.Next
		} else {
			next = next.Next
		}
	}
	cur.Next = next
	return head
}
