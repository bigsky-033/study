package com.github.bigskypark.p_01500_01600;

// 1551. Minimum Operations to Make Array Equal,
// https://leetcode.com/problems/minimum-operations-to-make-array-equal/

public class Solution01551 {
  public int minOperations(int n) {
    int numberOfOperations = 0;
    for (int i = 0; i < n; i++) {
      int val = (2 * i) + 1;
      if (val > n) {
        break;
      }
      numberOfOperations += (n - val);
    }
    return numberOfOperations;
  }
}
