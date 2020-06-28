# Concurrent Programming in Java - Week4

Notes for `Concurrent Programming in Java` class in Coursera.

## 4.1 Optimistic Concurrency

- Basic idea of of *AtomicInteger::getAndAdd*
  - Allow multiple threads to read the existig value of the shared opject(*curVal*) without synchronization.
  - Compute its new value after the addition (*newVal*) without synchronization.
  - These computations are performeed `optimistically` under the assumption that no interference will occur with other threads during the period between reading *curVal* and computing *newVal*.
  - It is necessary for each thread to confirm this assumption by using the *compareAndSet* methods.
- A.compareAndSet(curVal, newVal)
  - Checks that the value in A still equals *curVal*, and, if so, updates A's value to *newVal* before returning *true*.
  - Otherwise, the method simply returns *false* without updating *A*.
  - `compareAndSet` method is guaranteed to be performed atomically, as if it was in ad object-based isolated statement with respect to object *A*.
    - If two threads call compareAndSet method with the same *curVal* that matches A's current value, only one of them will succeed in updating A with their *newVal*.
    - Furthermore, each thread will invoke an operation like *compareAndSet* repeatedly in a loop until the operation succeeds.
    - There are no blocking operation => no deadlock.
    - Guarantee to eventually succeed => there cannot be livelock.

### Readings for 4.1

- Optimistic concurrency control, https://en.wikipedia.org/wiki/Optimistic_concurrency_control
- Atomic Integer, https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/atomic/AtomicInteger.html

## 4.2 Concurrenct Queue

- Concurrent queue ensure that call to *enq* and *deq* maintain the correct semantics, even if the calls are invoked concurrently from different threads.
- Common approach for implementation is to replace an object reference like *tail* by an *AtomicReference*.

### Readings for 4.2

- Atomic Reference, https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/atomic/AtomicReference.html
- Concurrent Linked Queue, https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/ConcurrentLinkedQueue.html

## 4.3 Linearizability

- *Linearizability* explains that what return values are permissible when multiple threads perform these operations in paralle, taking into account what we know about the expected return values from those operations when they are performed sequentially.
- Goal for any implementation of a concurrent data structure is to ensure that all its executions are linearizable by using whatever combination of constructs (e.g., locks, isolated, actors, optimistic concurrency) is deemed appropriate to ensure correctness while giving the maximum performance.

### Readings for 4.3

- Linearizability, https://en.wikipedia.org/wiki/Linearizability

## 4.4 Concurrent Hash Map

### Readings for 4.4

- Concurrent Hash Map, https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/ConcurrentHashMap.html
- Java ConcurrentMap, https://en.wikipedia.org/wiki/Java_ConcurrentMap

## 4.5 Concurrent Minimum Spanning Tree Algorithm

### Readings for 4.5

- Borůvka's algorithm, https://en.wikipedia.org/wiki/Bor%C5%AFvka%27s_algorithm