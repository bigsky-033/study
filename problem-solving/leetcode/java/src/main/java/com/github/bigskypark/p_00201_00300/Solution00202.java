package com.github.bigskypark.p_00201_00300;

// 202. Happy Number, https://leetcode.com/problems/happy-number/

import java.util.HashSet;
import java.util.Set;

public class Solution00202 {
  public boolean isHappy(int n) {
    char[] chars = String.valueOf(n).toCharArray();
    Set<Integer> visited = new HashSet<>();
    for (; ; ) {
      int sum = 0;
      for (char c : chars) {
        sum += (int) Math.pow(c - '0', 2);
      }
      if (visited.contains(sum)) {
        return false;
      }
      visited.add(sum);
      if (sum == 1) {
        return true;
      }
      chars = String.valueOf(sum).toCharArray();
    }
  }
}
