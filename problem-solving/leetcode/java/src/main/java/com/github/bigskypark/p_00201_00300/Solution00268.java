package com.github.bigskypark.p_00201_00300;

// 268. Missing Number, https://leetcode.com/problems/missing-number/

public class Solution00268 {
  public int missingNumber(int[] nums) {
    int sum = 0;
    for (int num : nums) {
      sum += num;
    }
    return (int) (nums.length * (nums.length + 1) * 0.5) - sum;
  }
}
