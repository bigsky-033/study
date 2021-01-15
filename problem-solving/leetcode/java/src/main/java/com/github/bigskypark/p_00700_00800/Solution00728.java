package com.github.bigskypark.p_00700_00800;

import java.util.LinkedList;
import java.util.List;

public class Solution00728 {
  public List<Integer> selfDividingNumbers(int left, int right) {
    List<Integer> numbers = new LinkedList<>();
    for (int i = left; i <= right; i++) {
      if (i % 10 == 0) {
        continue;
      }

      boolean dividing = true;
      char[] chars = String.valueOf(i).toCharArray();
      for (char c : chars) {
        int d = c - '0';
        if (d == 0 || i % d != 0) {
          dividing = false;
          break;
        }
      }

      if (dividing) {
        numbers.add(i);
      }
    }
    return numbers;
  }
}
