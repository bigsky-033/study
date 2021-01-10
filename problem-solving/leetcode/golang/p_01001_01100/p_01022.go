package p_01001_01100

// 1022. Sum of Root To Leaf Binary Numbers, https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

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

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	dfs(root, &sum, 0)
	return sum
}

func dfs(root *TreeNode, sum *int, cur int) {
	if root == nil {
		return
	}
	cur = cur << 1
	cur += root.Val

	if root.Left == nil && root.Right == nil {
		*sum += cur
		return
	}

	dfs(root.Left, sum, cur)
	dfs(root.Right, sum, cur)
}
