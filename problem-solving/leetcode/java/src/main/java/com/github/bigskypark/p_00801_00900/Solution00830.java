package com.github.bigskypark.p_00801_00900;

// 830. Positions of Large Groups, https://leetcode.com/problems/positions-of-large-groups/

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Solution00830 {
  public List<List<Integer>> largeGroupPositions(String s) {
    if (s == null || s.length() < 3) {
      return Collections.emptyList();
    }

    List<List<Integer>> groupPositions = new ArrayList<>();
    int right = 0;
    int left = 0;

    for (; right < s.length(); right++) {
      if (isNewGroup(s, right)) {
        if (right - left + 1 > 3) {
          groupPositions.add(Arrays.asList(left, right - 1));
        }
        left = right;
      }
    }

    if (right - left + 1 > 3) {
      groupPositions.add(Arrays.asList(left, right - 1));
    }
    return groupPositions;
  }

  private boolean isNewGroup(String s, int index) {
    return index <= 0 || (s.charAt(index - 1) != s.charAt(index));
  }
}
