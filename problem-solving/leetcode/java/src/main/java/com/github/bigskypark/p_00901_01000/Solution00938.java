package com.github.bigskypark.p_00901_01000;

// 938. Range Sum of BST, https://leetcode.com/problems/range-sum-of-bst/

public class Solution00938 {

  private int sum = 0;

  public int rangeSumBST(TreeNode root, int low, int high) {
    inOrderTraversal(root, low, high);
    return sum;
  }

  private void inOrderTraversal(TreeNode root, int low, int high) {
    if (root == null) {
      return;
    }
    inOrderTraversal(root.left, low, high);
    if (root.val >= low && root.val <= high) {
      sum += root.val;
    }
    inOrderTraversal(root.right, low, high);
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
