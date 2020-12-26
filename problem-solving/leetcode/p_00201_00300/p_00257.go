package p_00201_00300

// 257. Binary Tree Paths, https://leetcode.com/problems/binary-tree-paths/

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	dfs(root, &res, "")
	return res
}

func dfs(root *TreeNode, res *[]string, path string) {
	if root == nil {
		return
	}
	if len(path) != 0 {
		path += "->"
	}
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	dfs(root.Left, res, path)
	dfs(root.Right, res, path)
}
