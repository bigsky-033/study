package com.github.bigskypark.p_01600_01700;

// 1672. Richest Customer Wealth, https://leetcode.com/problems/richest-customer-wealth/

public class Solution01672 {
  public int maximumWealth(int[][] accounts) {
    int max = Integer.MIN_VALUE;
    for (int[] account : accounts) {
      int wealth = 0;
      for (int money : account) {
        wealth += money;
      }
      max = Math.max(max, wealth);
    }
    return max;
  }
}
