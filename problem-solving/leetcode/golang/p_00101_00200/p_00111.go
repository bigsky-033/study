package p_00101_00200

// 111. Minimum Depth of Binary Tree, https://leetcode.com/problems/minimum-depth-of-binary-tree/

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)

	if ld == 0 && rd != 0 {
		return rd + 1
	} else if ld != 0 && rd == 0 {
		return ld + 1
	} else if ld < rd {
		return ld + 1
	} else {
		return rd + 1
	}
}
