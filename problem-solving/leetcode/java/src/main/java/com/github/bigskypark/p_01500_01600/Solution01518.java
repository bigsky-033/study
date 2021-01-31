package com.github.bigskypark.p_01500_01600;

// 1518. Water Bottles, https://leetcode.com/problems/water-bottles/

public class Solution01518 {
  public int numWaterBottles(int numBottles, int numExchange) {
    int res = 0;
    int numEmpty = 0;

    while (numBottles + numEmpty >= numExchange) {
      res += numBottles;

      numEmpty += numBottles;

      numBottles = numEmpty / numExchange;
      numEmpty -= (numBottles) * numExchange;
    }
    res += numBottles;
    return res;
  }
}
