package p_00001_00100

// 82. Remove Duplicates from Sorted List II, https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

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
	newHead := getNextNonDuplicate(head)
	if newHead == nil {
		return newHead
	}

	cur := newHead
	for cur != nil {
		cur.Next = getNextNonDuplicate(cur.Next)
		cur = cur.Next
	}
	return newHead
}

func getNextNonDuplicate(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	next := cur.Next
	duplicated := -101 // invalid value

	for next != nil {
		if cur.Val == next.Val {
			duplicated = cur.Val
		}
		if duplicated != cur.Val && cur.Val != next.Val {
			return cur
		}
		cur = next
		next = next.Next
	}

	if cur.Val != duplicated {
		return cur
	}
	return nil
}
