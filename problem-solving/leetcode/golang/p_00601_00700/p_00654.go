package p_00601_00700

// 654. Maximum Binary Tree, https://leetcode.com/problems/maximum-binary-tree/

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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	i := maxIdx(nums)
	node := TreeNode{Val: nums[i]}
	node.Left = constructMaximumBinaryTree(nums[:i])
	node.Right = constructMaximumBinaryTree(nums[i+1:])
	return &node
}

func maxIdx(nums []int) int {
	max := nums[0]
	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	return maxIdx
}
