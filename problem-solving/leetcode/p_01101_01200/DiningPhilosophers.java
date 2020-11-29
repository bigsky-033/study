package p_01101_01200;

// 1226. The Dining Philosophers, https://leetcode.com/problems/the-dining-philosophers/

import java.util.concurrent.Semaphore;

public class DiningPhilosophers {

  private static final int NUMBER_OF_PHILOSOPHERS = 5;
  private final Semaphore[] semaphores;

  public DiningPhilosophers() {
    semaphores = new Semaphore[NUMBER_OF_PHILOSOPHERS];
    for (int i = 0; i < NUMBER_OF_PHILOSOPHERS; i++) {
      semaphores[i] = new Semaphore(1);
    }
  }

  // call the run() method of any runnable to execute its code
  public void wantsToEat(
      int philosopher,
      Runnable pickLeftFork,
      Runnable pickRightFork,
      Runnable eat,
      Runnable putLeftFork,
      Runnable putRightFork)
      throws InterruptedException {
    int indexForRightFork = (philosopher - 1 + NUMBER_OF_PHILOSOPHERS) % NUMBER_OF_PHILOSOPHERS;
    while (true) {
      semaphores[philosopher].acquire();
      if (semaphores[indexForRightFork].tryAcquire()) {
        pickLeftFork.run();
        pickRightFork.run();
        break;
      } else {
        semaphores[philosopher].release();
      }
    }

    eat.run();

    putRightFork.run();
    semaphores[indexForRightFork].release();

    putLeftFork.run();
    semaphores[philosopher].release();
  }
}
