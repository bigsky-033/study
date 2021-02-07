package com.github.bigskypark.p_01000_01100;

// 1038. Binary Search Tree to Greater Sum Tree,
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

public class Solution01038 {

  private int total = 0;
  private int current = 0;

  public TreeNode bstToGst(TreeNode root) {
    calculateTotalSum(root);
    toGreaterSumTree(root);
    return root;
  }

  private void calculateTotalSum(TreeNode root) {
    if (root == null) {
      return;
    }
    calculateTotalSum(root.left);
    total += root.val;
    calculateTotalSum(root.right);
  }

  private void toGreaterSumTree(TreeNode root) {
    if (root == null) {
      return;
    }
    toGreaterSumTree(root.left);
    current += root.val;
    root.val += (total - current);
    toGreaterSumTree(root.right);
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
