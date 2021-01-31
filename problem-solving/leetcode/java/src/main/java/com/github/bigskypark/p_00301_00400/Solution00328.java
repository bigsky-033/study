package com.github.bigskypark.p_00301_00400;

// 328. Odd Even Linked List, https://leetcode.com/problems/odd-even-linked-list/

public class Solution00328 {
  public ListNode oddEvenList(ListNode head) {
    if (head == null || head.next == null || head.next.next == null) {
      return head;
    }

    ListNode dummyForOdd = new ListNode(-1);
    ListNode dummyForEven = new ListNode(-1);
    dummyForOdd.next = head;
    dummyForEven.next = head.next;

    ListNode currentOdd = dummyForOdd.next;
    ListNode currentEven = dummyForEven.next;

    ListNode current = head;
    int nodeNumber = 1;
    while (current != null) {
      if (nodeNumber > 2 && nodeNumber % 2 == 0) {
        currentEven.next = current;
        currentEven = currentEven.next;
      } else if (nodeNumber > 2 && nodeNumber % 2 != 0) {
        currentOdd.next = current;
        currentOdd = currentOdd.next;
      }
      nodeNumber++;
      current = current.next;
    }
    currentEven.next = null;
    currentOdd.next = dummyForEven.next;
    return dummyForOdd.next;
  }

  static class ListNode {
    int val;
    ListNode next;

    ListNode() {}

    ListNode(int val) {
      this.val = val;
    }

    ListNode(int val, ListNode next) {
      this.val = val;
      this.next = next;
    }
  }
}
