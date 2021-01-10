package p_01600_01700

// 1669. Merge In Between Linked Lists, https://leetcode.com/problems/merge-in-between-linked-lists/

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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	beforeDelete := list1
	afterDelete := list1
	p2 := list2

	// Move pointer to node before start deletion
	for i := 0; i < a-1; i++ {
		beforeDelete = beforeDelete.Next
		afterDelete = afterDelete.Next
	}
	// Move pointer to node after deletion
	for i := 0; i < (b+1)-(a-1); i++ {
		afterDelete = afterDelete.Next
	}

	// Merge
	beforeDelete.Next = p2
	for p2.Next != nil {
		p2 = p2.Next
	}
	p2.Next = afterDelete
	return list1
}
