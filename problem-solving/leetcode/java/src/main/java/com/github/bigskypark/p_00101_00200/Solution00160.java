package com.github.bigskypark.p_00101_00200;

// 160. Intersection of Two Linked Lists,
// https://leetcode.com/problems/intersection-of-two-linked-lists/

public class Solution00160 {

  public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
    if (headA == null || headB == null) {
      return null;
    }

    ListNode p1 = headA;
    ListNode p2 = headB;

    while (p1 != p2) {
      p1 = p1.next;
      p2 = p2.next;

      if (p1 == p2) {
        break;
      }

      if (p1 == null) {
        p1 = headB;
      }
      if (p2 == null) {
        p2 = headA;
      }
    }
    return p1;
  }

  //    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
  //        Set<ListNode> nodes = new HashSet<>();
  //        ListNode current = headA;
  //        while (current != null) {
  //            nodes.add(current);
  //            current = current.next;
  //        }
  //
  //        current = headB;
  //        while (current != null) {
  //            if (nodes.contains(current)) {
  //                return current;
  //            }
  //            current = current.next;
  //        }
  //        return null;
  //    }

  static class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
      val = x;
      next = null;
    }
  }
}
