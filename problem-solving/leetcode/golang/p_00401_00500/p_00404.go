package p_00401_00500

// 404. Sum of Left Leaves, https://leetcode.com/problems/sum-of-left-leaves/

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

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	traverse(root.Left, res)
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			*res = *res + root.Left.Val
		}
	}
	traverse(root.Right, res)
}
