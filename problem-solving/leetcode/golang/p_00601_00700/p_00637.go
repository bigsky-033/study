package p_00601_00700

// 637. Average of Levels in Binary Tree, https://leetcode.com/problems/average-of-levels-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var averages []float64
	if root == nil {
		return averages
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(n))
	}
	return averages
}
