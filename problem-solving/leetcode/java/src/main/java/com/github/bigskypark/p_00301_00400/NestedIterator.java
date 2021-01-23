package com.github.bigskypark.p_00300_00400;

// 341. Flatten Nested List Iterator, https://leetcode.com/problems/flatten-nested-list-iterator/

import java.util.Deque;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

import static java.util.Objects.requireNonNull;

public class NestedIterator implements Iterator<Integer> {

  private final Deque<NestedInteger> stack = new LinkedList<>();

  public NestedIterator(final List<NestedInteger> nestedIntegers) {
    requireNonNull(nestedIntegers);
    for (int i = nestedIntegers.size() - 1; i >= 0; i--) {
      stack.push(nestedIntegers.get(i));
    }
    flatten();
  }

  @Override
  public Integer next() {
    Integer ret = stack.pop().getInteger();
    flatten();
    return ret;
  }

  @Override
  public boolean hasNext() {
    return !stack.isEmpty();
  }

  private void flatten() {
    while (!stack.isEmpty()) {
      NestedInteger top = stack.peek();
      // Ensure next element is integer
      if (top.isInteger()) {
        return;
      }

      stack.pop();
      List<NestedInteger> nestedIntegers = top.getList();
      for (int i = nestedIntegers.size() - 1; i >= 0; i--) {
        stack.push(nestedIntegers.get(i));
      }
    }
  }

  interface NestedInteger {

    // @return true if this NestedInteger holds a single integer, rather than a nested list.
    boolean isInteger();

    // @return the single integer that this NestedInteger holds, if it holds a single integer
    // Return null if this NestedInteger holds a nested list
    Integer getInteger();

    // @return the nested list that this NestedInteger holds, if it holds a nested list
    // Return null if this NestedInteger holds a single integer
    List<NestedInteger> getList();
  }
}
