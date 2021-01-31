package com.github.bigskypark.p_01101_01200;

// 1189. Maximum Number of Balloons, https://leetcode.com/problems/maximum-number-of-balloons/

import java.util.HashMap;
import java.util.Map;

public class Solution01189 {
  public int maxNumberOfBalloons(String text) {
    if (text == null || text.length() < "balloon".length()) {
      return 0;
    }
    Map<Character, Integer> countsByChar = new HashMap<>();
    char[] chars = text.toCharArray();
    for (char c : chars) {
      if (isBalloonChar(c)) {
        countsByChar.merge(c, 1, Integer::sum);
      }
    }

    int min = Integer.MAX_VALUE;
    for (Map.Entry<Character, Integer> entry : countsByChar.entrySet()) {
      Character key = entry.getKey();
      if (key.equals('l') || key.equals('o')) {
        min = Math.min(min, entry.getValue() / 2);
      } else {
        min = Math.min(min, entry.getValue());
      }
    }
    return min;
  }

  private boolean isBalloonChar(char c) {
    return c == 'b' || c == 'l' || c == 'o' || c == 'a' || c == 'n';
  }
}
