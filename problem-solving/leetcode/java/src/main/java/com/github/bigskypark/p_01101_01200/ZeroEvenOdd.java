package com.github.bigskypark.p_01101_01200;

// 1116. Print Zero Even Odd, https://leetcode.com/problems/print-zero-even-odd/

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.function.IntConsumer;

class ZeroEvenOdd {
  private static final int STOP_SIGNAL = -1;

  private final int n;
  private final BlockingQueue<Integer> channelForZero = new ArrayBlockingQueue<>(1);
  private final BlockingQueue<Integer> channelForOdd = new ArrayBlockingQueue<>(1);
  private final BlockingQueue<Integer> channelForEven = new ArrayBlockingQueue<>(1);

  public ZeroEvenOdd(final int n) {
    this.n = n;
  }

  // printNumber.accept(x) outputs "x", where x is an integer.
  public void zero(final IntConsumer printNumber) throws InterruptedException {
    if (n < 0) {
      return;
    }
    int counter = 1;
    channelForZero.offer(0);

    while (counter <= n) {
      printNumber.accept(channelForZero.take());
      if (counter % 2 == 0) {
        channelForEven.offer(counter);
      } else {
        channelForOdd.offer(counter);
      }
      counter++;
    }
    channelForZero.take();
    channelForOdd.offer(STOP_SIGNAL);
    channelForEven.offer(STOP_SIGNAL);
  }

  public void even(final IntConsumer printNumber) throws InterruptedException {
    while (true) {
      int given = channelForEven.take();
      if (given == STOP_SIGNAL) {
        break;
      }

      printNumber.accept(given);
      channelForZero.offer(0);
    }
  }

  public void odd(final IntConsumer printNumber) throws InterruptedException {
    while (true) {
      int given = channelForOdd.take();
      if (given == STOP_SIGNAL) {
        break;
      }
      printNumber.accept(given);
      channelForZero.offer(0);
    }
  }
}
