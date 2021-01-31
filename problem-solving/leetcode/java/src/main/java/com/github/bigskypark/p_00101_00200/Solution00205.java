package com.github.bigskypark.p_00101_00200;

// 205. Isomorphic Strings, https://leetcode.com/problems/isomorphic-strings/

import java.util.HashMap;
import java.util.Map;

public class Solution00205 {
  public boolean isIsomorphic(String s, String t) {
    if (s == null || t == null || s.length() != t.length()) {
      return false;
    }
    Map<Character, Character> sToT = new HashMap<>();
    Map<Character, Character> tToS = new HashMap<>();

    for (int i = 0; i < s.length(); i++) {
      char charS = s.charAt(i);
      char charT = t.charAt(i);
      Character mappingFromSToT = sToT.get(charS);
      Character mappingFromTToS = tToS.get(charT);

      if (mappingFromSToT == null && mappingFromTToS == null) {
        sToT.put(charS, charT);
        tToS.put(charT, charS);
      } else if ((mappingFromSToT == null || !mappingFromSToT.equals(charT))
          && (mappingFromTToS == null || !mappingFromTToS.equals(charS))) {
        return false;
      }
    }
    return true;
  }
}
