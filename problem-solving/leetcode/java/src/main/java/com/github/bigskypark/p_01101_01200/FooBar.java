package com.github.bigskypark.p_01101_01200;

// 1115. Print FooBar Alternately, https://leetcode.com/problems/print-foobar-alternately/

import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

class FooBar {
  private int n;
  private final AtomicInteger counter = new AtomicInteger();

  public FooBar(int n) {
    this.n = n;
  }

  public void foo(Runnable printFoo) throws InterruptedException {

    for (int i = 0; i < n; i++) {

      // printFoo.run() outputs "foo". Do not change or remove this line.
      while (true) {
        if (counter.get() % 2 == 0) break;
      }
      printFoo.run();
      counter.incrementAndGet();
    }
  }

  public void bar(Runnable printBar) throws InterruptedException {

    for (int i = 0; i < n; i++) {

      // printBar.run() outputs "bar". Do not change or remove this line.
      while (true) {
        if (counter.get() % 2 != 0) break;
        TimeUnit.NANOSECONDS.sleep(1L);
      }
      printBar.run();
      counter.incrementAndGet();
    }
  }
}
