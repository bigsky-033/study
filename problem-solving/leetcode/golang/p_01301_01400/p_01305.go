package p_01301_01400

// 1305. All Elements in Two Binary Search Trees, https://leetcode.com/problems/all-elements-in-two-binary-search-trees/

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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var e1 []int
	inorderTraversal(root1, &e1)
	var e2 []int
	inorderTraversal(root2, &e2)

	// merge
	var res []int
	i, j := 0, 0
	for i < len(e1) && j < len(e2) {
		if e1[i] < e2[j] {
			res = append(res, e1[i])
			i++
		} else {
			res = append(res, e2[j])
			j++
		}
	}
	for i < len(e1) {
		res = append(res, e1[i])
		i++
	}
	for j < len(e2) {
		res = append(res, e2[j])
		j++
	}
	return res
}

func inorderTraversal(root *TreeNode, elements *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, elements)
	*elements = append(*elements, root.Val)
	inorderTraversal(root.Right, elements)
}
