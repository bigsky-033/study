package com.github.bigskypark.p_01101_01200;

// 1114. Print in Order, https://leetcode.com/problems/print-in-order/

import java.util.concurrent.CountDownLatch;

class Foo {

  CountDownLatch latchForSecond = new CountDownLatch(1);
  CountDownLatch latchForThird = new CountDownLatch(1);

  public Foo() {}

  public void first(Runnable printFirst) throws InterruptedException {

    // printFirst.run() outputs "first". Do not change or remove this line.
    printFirst.run();
    latchForSecond.countDown();
  }

  public void second(Runnable printSecond) throws InterruptedException {

    // printSecond.run() outputs "second". Do not change or remove this line.
    latchForSecond.await();
    printSecond.run();
    latchForThird.countDown();
  }

  public void third(Runnable printThird) throws InterruptedException {

    // printThird.run() outputs "third". Do not change or remove this line.
    latchForThird.await();
    printThird.run();
  }
}
