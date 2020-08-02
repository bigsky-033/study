# Concurrent Programming in Java - Week1

Notes for `Concurrent Programming in Java` class in Coursera.

- <https://www.coursera.org/learn/concurrent-programming-in-java>

## 1.1 Threads

Learn about basics of `Thread`.

### Readings for 1.1

- Threads, <https://en.wikipedia.org/wiki/Thread_(computing)>
- Tutorial on Java threads, <https://docs.oracle.com/javase/tutorial/essential/concurrency/runthread.html>
- Documentation on Thread class in Java 8, <https://docs.oracle.com/javase/8/docs/api/java/lang/Thread.html>

## 1.2 Structured Locks

- *Structured locks* can be implemented using *synchronized* statements and methods in java.
- Structured locks can be used to enforce *mutual exclusion* and avoid *data races*.
- Major benefit of structured locks is that their *acquire* and *release* operations are implicit, since these operations are automatically performed by the java runtime environment when entering and exiting the scope of *synchronized* statement or method, even if exception is thrown in the middle.
- *wait()*, *notify()* can be used to block and resume threads that need to wait for specific conditions.

### Readings for 1.2

- Tutorial on Intrinsic Locks and Synchronization in Java, <https://docs.oracle.com/javase/tutorial/essential/concurrency/locksync.html>
- Tutorial on Guarded Blocks in Java, <https://docs.oracle.com/javase/tutorial/essential/concurrency/guardmeth.html>
- Monitor (synchronization), <https://en.wikipedia.org/wiki/Monitor_(synchronization)>

## 1.3 Unstructured Locks

- In java, we can obtain *unstructured locks* by creating instances of *ReentrantLock*.
- Major operations
  - *lock(), unlock()* can be used to support a *hand-over-hand* locking pattern that implements non-nested pairing of lock/unlock operations which cannot be achieved with synchronized statements/methods.
  - *tryLock()* operations in unstructured locks can enable a thread to check the availability of a lock, and thereby acquire it if it is available or do something else if it is not.
  - *read-write locks* - This can be obtained in java by creating *ReentrantReadWriteLock* - whereby multiple threads are permitted to acquire a lock L in "read mode", *L.readLock().lock*, but only one thread is permitted to acquire the lock in "write mode", *L.writeLock().lock().

### Readings for 1.3

- Tutorial on Lock Objects in Java, <https://docs.oracle.com/javase/tutorial/essential/concurrency/newlocks.html>
- Documentation on Java’s Lock interfaces, <https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/locks/Lock.html>

## 1.4 Liveness

- Studied three cases in which a parallel program may enter a state in which it stops making forward progress.
  - `Deadlock`: all threads are blocked indefinitely.
  - `Livelock`: all threads repeatedly perform an interaction that prevents forward progress.
  - `Starvation`: At least one thread is prevented from making any foward progress.

### Readings for 1.4

- Deadlock example with synchronized methods in Java, <https://docs.oracle.com/javase/tutorial/essential/concurrency/deadlock.html>
- Starvation and Livelock examples in Java, <https://docs.oracle.com/javase/tutorial/essential/concurrency/starvelive.html>
- Deadlock and Livelock, <https://en.wikipedia.org/wiki/Deadlock>

## 1.5 Dining Philosophers

- Using structured locks, can lead to a `deadlock scenario`(but not livelock).
- Using unstructured locks with *tryLock()* and *unlock()* operations that never block, can lead to a `livelock scenario`(but not deadlock).

### Readings for 1.5

- Dining Philosophers Problem, <https://en.wikipedia.org/wiki/Dining_philosophers_problem>