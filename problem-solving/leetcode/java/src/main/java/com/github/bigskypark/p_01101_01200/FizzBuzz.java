package com.github.bigskypark.p_01101_01200;

// 1195. Fizz Buzz Multithreaded, https://leetcode.com/problems/fizz-buzz-multithreaded/

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.function.IntConsumer;

class FizzBuzz {
  private static final int STOP_SIGNAL = -1;
  private int n;

  private final BlockingQueue<Integer> channelForFizz = new ArrayBlockingQueue<>(1);
  private final BlockingQueue<Integer> channelForBuzz = new ArrayBlockingQueue<>(1);
  private final BlockingQueue<Integer> channelForFizzBuzz = new ArrayBlockingQueue<>(1);
  private final BlockingQueue<Integer> channelForNumber = new ArrayBlockingQueue<>(1);

  public FizzBuzz(int n) {
    this.n = n;
    dispatch(1);
  }

  // printFizz.run() outputs "fizz".
  public void fizz(Runnable printFizz) throws InterruptedException {
    while (true) {
      int i = channelForFizz.take();
      if (i == STOP_SIGNAL) {
        break;
      }
      printFizz.run();
      dispatch(i + 1);
    }
  }

  // printBuzz.run() outputs "buzz".
  public void buzz(Runnable printBuzz) throws InterruptedException {
    while (true) {
      int i = channelForBuzz.take();
      if (i == STOP_SIGNAL) {
        break;
      }
      printBuzz.run();
      dispatch(i + 1);
    }
  }

  // printFizzBuzz.run() outputs "fizzbuzz".
  public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
    while (true) {
      int i = channelForFizzBuzz.take();
      if (i == STOP_SIGNAL) {
        break;
      }
      printFizzBuzz.run();
      dispatch(i + 1);
    }
  }

  // printNumber.accept(x) outputs "x", where x is an integer.
  public void number(IntConsumer printNumber) throws InterruptedException {
    while (true) {
      int i = channelForNumber.take();
      if (i == STOP_SIGNAL) {
        break;
      }
      printNumber.accept(i);
      dispatch(i + 1);
    }
  }

  private void dispatch(int i) {
    if (i > n) {
      channelForFizzBuzz.offer(STOP_SIGNAL);
      channelForFizz.offer(STOP_SIGNAL);
      channelForBuzz.offer(STOP_SIGNAL);
      channelForNumber.offer(STOP_SIGNAL);
      return;
    }

    if (i % 3 == 0 && i % 5 == 0) {
      channelForFizzBuzz.offer(i);
    } else if (i % 3 == 0) {
      channelForFizz.offer(i);
    } else if (i % 5 == 0) {
      channelForBuzz.offer(i);
    } else {
      channelForNumber.offer(i);
    }
  }
}
