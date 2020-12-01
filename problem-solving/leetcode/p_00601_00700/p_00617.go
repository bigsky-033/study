package p_00601_00700

// 617. Merge Two Binary Trees, https://leetcode.com/problems/merge-two-binary-trees/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return merge(t1, t2)
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	node := &TreeNode{Val: t1.Val + t2.Val}
	node.Left = merge(t1.Left, t2.Left)
	node.Right = merge(t1.Right, t2.Right)
	return node
}
