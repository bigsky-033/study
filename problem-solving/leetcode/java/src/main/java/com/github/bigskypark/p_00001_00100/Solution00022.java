package com.github.bigskypark.p_00001_00100;

// 22. Generate Parentheses, https://leetcode.com/problems/generate-parentheses/

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution00022 {
  public List<String> generateParenthesis(int n) {
    List<String> parenthesises = new ArrayList<>();
    if (n == 0) {
      return parenthesises;
    }
    Set<String> visited = new HashSet<>();
    helper(parenthesises, visited, "", n, 0, 0);

    return parenthesises;
  }

  private void helper(
      List<String> parenthesises,
      Set<String> visited,
      String parenthesis,
      int n,
      int leftCnt,
      int rightCnt) {
    if (n == leftCnt && n == rightCnt) {
      parenthesises.add(parenthesis);
      return;
    }

    if (visited.contains(parenthesis)) {
      return;
    }
    visited.add(parenthesis);

    for (int i = 0; i < n - leftCnt; i++) {
      helper(parenthesises, visited, parenthesis + "(", n, leftCnt + 1, rightCnt);
    }

    for (int i = 0; i < leftCnt - rightCnt; i++) {
      helper(parenthesises, visited, parenthesis + ")", n, leftCnt, rightCnt + 1);
    }
  }
}
