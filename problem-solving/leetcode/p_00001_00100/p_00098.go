package p_00001_00100

// 98. Validate Binary Search Tree, https://leetcode.com/problems/validate-binary-search-tree/

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

func isValidBST(root *TreeNode) bool {
	isValid := true
	isFirst := true
	prevValue := 0
	inorderTraversal(root, &isValid, &isFirst, &prevValue)
	return isValid
}

func inorderTraversal(root *TreeNode, isValid *bool, isFirst *bool, prevValue *int) {
	if !(*isValid) || root == nil {
		return
	}
	inorderTraversal(root.Left, isValid, isFirst, prevValue)
	if *isFirst {
		*prevValue = root.Val
		*isFirst = false
	} else {
		if root.Val <= *prevValue {
			*isValid = false
		}
		*prevValue = root.Val
	}
	inorderTraversal(root.Right, isValid, isFirst, prevValue)
}
