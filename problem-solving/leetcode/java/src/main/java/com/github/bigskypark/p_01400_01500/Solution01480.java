package com.github.bigskypark.p_01400_01500;

// 1480. Running Sum of 1d Array, https://leetcode.com/problems/running-sum-of-1d-array/

public class Solution01480 {
  public int[] runningSum(int[] nums) {
    int sum = 0;
    for (int i = 0; i < nums.length; i++) {
      int temp = nums[i];
      nums[i] = nums[i] + sum;
      sum += temp;
    }
    return nums;
  }
}
