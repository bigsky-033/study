package com.github.bigskypark.p_00100_00200;

// 235. Lowest Common Ancestor of a Binary Search Tree,
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

public class Solution00235 {

  private boolean found = false;
  private TreeNode answer = null;

  public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
    find(root, p, q);
    return answer;
  }

  private TreeNode find(TreeNode root, TreeNode p, TreeNode q) {
    if (root == null || found) {
      return null;
    }

    TreeNode left = find(root.left, p, q);
    TreeNode right = find(root.right, p, q);

    if ((root == p) && (left == q)
        || (root == p) && (right == q)
        || (root == q) && (left == p)
        || (root == q) && (right == p)
        || (left == p) && (right == q)
        || (left == q) && (right == p)) {
      found = true;
      answer = root;
      return root;
    } else if (left == p || right == p || root == p) {
      return p;
    } else if (left == q || right == q || root == q) {
      return q;
    } else {
      return null;
    }
  }

  static class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
      val = x;
    }
  }
}
