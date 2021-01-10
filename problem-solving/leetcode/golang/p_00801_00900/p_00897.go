package p_00801_00900

// 897. Increasing Order Search Tree, https://leetcode.com/problems/increasing-order-search-tree/

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

func increasingBST(root *TreeNode) *TreeNode {
	var values []int
	inOrderTraversal(root, &values)
	dummy := &TreeNode{Val: 0}
	cur := dummy
	for _, v := range values {
		cur.Right = &TreeNode{Val: v}
		cur = cur.Right
	}
	return dummy.Right
}

func inOrderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, values)
	*values = append(*values, root.Val)
	inOrderTraversal(root.Right, values)
}
