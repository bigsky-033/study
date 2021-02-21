package com.github.bigskypark.p_00500_00600;

// 559. Maximum Depth of N-ary Tree, https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

import java.util.List;

public class Solution00559 {
  public int maxDepth(Node root) {
    if (root == null) {
      return 0;
    }
    int max = 0;
    for (Node child : root.children) {
      max = Math.max(max, maxDepth(child));
    }
    return max + 1;
  }

  static class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
      val = _val;
    }

    public Node(int _val, List<Node> _children) {
      val = _val;
      children = _children;
    }
  }
}
