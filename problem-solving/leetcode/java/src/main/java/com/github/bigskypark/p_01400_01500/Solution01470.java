package com.github.bigskypark.p_01400_01500;

// 1470. Shuffle the Array, https://leetcode.com/problems/shuffle-the-array/

public class Solution01470 {
  public int[] shuffle(int[] nums, int n) {
    int[] res = new int[2 * n];
    for (int i = 0; i < res.length; i++) {
      if (i % 2 == 0) {
        res[i] = nums[i / 2];
      } else {
        res[i] = nums[n + (i / 2)];
      }
    }
    return res;
  }
}
