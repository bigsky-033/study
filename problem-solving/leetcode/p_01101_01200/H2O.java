package p_01101_01200;

// 1117. Building H2O, https://leetcode.com/problems/building-h2o/

import java.util.concurrent.Semaphore;

class H2O {

  private final Semaphore lockForHydrogen = new Semaphore(2);
  private final Semaphore lockForOxygen = new Semaphore(0);

  // If a hydrogen thread arrives at the barrier when no other threads are present,
  // it has to wait for an oxygen thread and another hydrogen thread.
  public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
    lockForHydrogen.acquire(1);
    // releaseHydrogen.run() outputs "H". Do not change or remove this line.
    releaseHydrogen.run();
    lockForOxygen.release(1);
  }

  // If an oxygen thread arrives at the barrier when no hydrogen threads are present,
  // it has to wait for two hydrogen threads.
  public void oxygen(Runnable releaseOxygen) throws InterruptedException {
    lockForOxygen.acquire(2);
    // releaseOxygen.run() outputs "O". Do not change or remove this line.
    releaseOxygen.run();
    lockForHydrogen.release(2);
  }
}
