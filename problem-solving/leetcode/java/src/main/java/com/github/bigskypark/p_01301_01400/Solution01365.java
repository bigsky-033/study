package com.github.bigskypark.p_01301_01400;

// 1365. How Many Numbers Are Smaller Than the Current Number,
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Solution01365 {
  public int[] smallerNumbersThanCurrent(int[] nums) {
    Map<Integer, Integer> inverted = new HashMap<>();
    {
      int[] copied = Arrays.copyOf(nums, nums.length);
      Arrays.sort(copied);
      for (int i = 0; i < copied.length; i++) {
        inverted.putIfAbsent(copied[i], i);
      }
    }

    int[] result = new int[nums.length];
    for (int i = 0; i < nums.length; i++) {
      result[i] = inverted.get(nums[i]);
    }
    return result;
  }
}
