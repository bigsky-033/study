package com.github.bigskypark.p_01301_01400;

// 1302. Deepest Leaves Sum, https://leetcode.com/problems/deepest-leaves-sum/

import java.util.ArrayDeque;
import java.util.Deque;

public class Solution01302 {
  public int deepestLeavesSum(TreeNode root) {
    if (root == null) {
      return 0;
    }

    Deque<TreeNode> queue = new ArrayDeque<>();
    queue.add(root);
    int sum = 0;
    while (!queue.isEmpty()) {
      sum = 0;
      int size = queue.size();
      for (int i = 0; i < size; i++) {
        TreeNode node = queue.removeFirst();
        sum += node.val;

        if (node.left != null) {
          queue.addLast(node.left);
        }
        if (node.right != null) {
          queue.addLast(node.right);
        }
      }
    }
    return sum;
  }

  static class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {}

    TreeNode(int val) {
      this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
      this.val = val;
      this.left = left;
      this.right = right;
    }
  }
}
