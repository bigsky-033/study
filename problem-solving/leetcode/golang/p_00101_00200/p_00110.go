package p_00101_00200

// 110. Balanced Binary Tree, https://leetcode.com/problems/balanced-binary-tree/

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	b, _ := helper(root)
	return b
}

func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lb, lh := helper(root.Left)
	rb, rh := helper(root.Right)

	if !lb || !rb || diff(lh, rh) > 1 {
		return false, -1
	}
	if lh > rh {
		return true, lh + 1
	} else {
		return true, rh + 1
	}
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
