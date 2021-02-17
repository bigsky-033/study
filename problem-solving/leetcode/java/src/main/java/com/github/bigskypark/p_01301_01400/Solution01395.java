package com.github.bigskypark.p_01301_01400;

// 1395. Count Number of Teams, https://leetcode.com/problems/count-number-of-teams/

public class Solution01395 {
  public int numTeams(int[] rating) {

    int numberOfTeams = 0;
    for (int i = 0; i < rating.length; i++) {
      for (int j = i + 1; j < rating.length; j++) {
        for (int k = j + 1; k < rating.length; k++) {
          if ((rating[i] < rating[j] && rating[j] < rating[k])
              || (rating[i] > rating[j] && rating[j] > rating[k])) {
            numberOfTeams++;
          }
        }
      }
    }

    return numberOfTeams;
  }
}
