package edu.coursera.concurrent;

import edu.rice.pcdp.Actor;

import static edu.rice.pcdp.PCDP.finish;

/**
 * An actor-based implementation of the Sieve of Eratosthenes.
 *
 * <p>TODO Fill in the empty SieveActorActor actor class below and use it from countPrimes to
 * determin the number of primes <= limit.
 */
public final class SieveActor extends Sieve {
  /**
   * {@inheritDoc}
   *
   * <p>TODO Use the SieveActorActor class to calculate the number of primes <= limit in parallel.
   * You might consider how you can model the Sieve of Eratosthenes as a pipeline of actors, each
   * corresponding to a single prime number.
   */
  @Override
  public int countPrimes(final int limit) {
    if (limit < 2) {
      return 0;
    }

    final SieveActorActor actor = new SieveActorActor(2);
    finish(
        () -> {
          for (int i = 3; i <= limit; i += 2) {
            actor.send(i);
          }
        });

    int numberOfPrimes = 0;
    SieveActorActor current = actor;
    while (current != null) {
      numberOfPrimes += current.numberOfLocalPrimes;
      current = current.next;
    }
    return numberOfPrimes;
  }

  /** An actor class that helps implement the Sieve of Eratosthenes in parallel. */
  public static final class SieveActorActor extends Actor {
    /**
     * Process a single message sent to this actor.
     *
     * <p>TODO complete this method.
     *
     * @param msg Received message
     */
    private static final int MAX_LOCAL_PRIMES = 2000;

    private final int[] localPrimes = new int[MAX_LOCAL_PRIMES];

    private int numberOfLocalPrimes;
    private SieveActorActor next;

    public SieveActorActor(final int prime) {
      localPrimes[0] = prime;
      numberOfLocalPrimes++;
    }

    @Override
    public void process(final Object msg) {
      int candidate = (Integer) msg;

      if (isLocalPrime(candidate)) {
        if (numberOfLocalPrimes < MAX_LOCAL_PRIMES) {
          localPrimes[numberOfLocalPrimes] = candidate;
          numberOfLocalPrimes++;
        } else if (next == null) {
          next = new SieveActorActor(candidate);
        } else {
          next.send(msg);
        }
      }
    }

    private boolean isLocalPrime(int candidate) {
      for (int i = 0; i < numberOfLocalPrimes; i++) {
        if (candidate % localPrimes[i] == 0) {
          return false;
        }
      }
      return true;
    }
  }
}
