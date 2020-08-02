# Parallel Programming in Java - Week4

Notes for `Parallel Programming in Java` class in Coursera.

- <https://www.coursera.org/learn/parallel-programming-in-java>

## 4.1 Split-phase Barriers with Java Phasers

- Basic of Java's Phaser class
  - Phaser::arriveAndAwaitAdvance
    - Can be used to implement a barrier through phaser object phaser.
  - To facilitate the *split-phase- barrier*(also known as a *fuzzy barrier*), use two seprate APIs from Java Phaser class - Phaser::arrive, Phaser::awaitAdvance.

### Readings for 4.1

- Documentation on Java’s Phaser class, <https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/Phaser.html>

## 4.2 Point-to-Point Synchronization with Phasers

- Example of getting benefit using Phasers

| Task 0                  | Task 1                     | Task 2                  |
| ----------------------- | -------------------------- | ----------------------- |
| 1a: X = A() // cost = 1 | 1b: Y = B() // cost = 2    | 1c: Z = C() // cost = 3 |
| 2a: ph0.arrive()        | 2b: ph1. arrive()          | 2c: ph2.arrive()        |
| 3a: ph1.awaitAdvance(0) | 3b: ph0.awaitAdvance(0)    | 3c: ph1.awaitAdvance(0) |
| 4a: D(X, Y) // cost = 3 | 4b: ph2.awaitAdvance(0)    | 4c: F(Y, Z) // cost= 1  |
|                         | 5b: E(X, Y, Z) // cost = 2 |                         |

- Without phaser, CPL(critical path length) = 6
- With phaser, CPL = 5

## 4.3 One-Dimensional Iterative Averaging with Phasers

## 4.4 Pipeline Parallelism

### Basics

- n: number of input items
- p: number of stages in the pipeline
- WORK = n * p
- CPL = n + p - 1
- PAR = WORK / CPL = np / (n + p - 1)
  - p = 1, the ideal parallelism degenerates to PAR = 1
    - computation is sequential when only one stage is available
  - n = 1, the ideal parallelism degenerates to PAR = 1
  - when n is much larger than p (n >> p), then the ideal parallelism approaches PAR = 1
- How synchronization is implemented in pipeline
- Using phasers by allocating an array of phasers, such that phaser `ph[i]` is signalled in iteration i by a call to `ph[i].arrive()` as follows:

```
while (there is an input to be processed) {
    if (i > 0) ph[i-1] awaitAdvance();

    process input;

    ph[i].arrive();
}
```

### Readings for 4.4

- Pipeline (computing), <https://en.wikipedia.org/wiki/Pipeline_(computing)>

## 4.5 Data Flow Parallelism

### Readings for 4.5

- Data flow parallelism model is, to specify parallel program as computation graph.

- Data flow diagram, <https://en.wikipedia.org/wiki/Data-flow_diagram>