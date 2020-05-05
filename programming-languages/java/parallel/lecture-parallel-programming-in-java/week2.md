# Parallel Programming in Java - Week2

Notes for `Parallel Programming in Java` class in Coursera.

- https://www.coursera.org/learn/parallel-programming-in-java

## 2.1 Futures: Tasks with Return Values

- Future tasks: Tasks with return values.
- Future: Object is a handle for accessing a task's return value.
  - Key operations that can be performaed on a future object *A*:
    - Assignment: *A* can be assigned a reference to future object returned by task of the futre, *future {< task-with-return-value >}*. The content of the future object is constrained to be single assignment, and cannot be modified after the future task has returned.
    - Blocking: the operation, *A.get()*, waits until the task associated with future object A has completed, and then propagates the task's return value as the value returned by *A.get()*. Any statement, S, executed after *A.get()* can be assured that the task associated with future object A must have completed before S starts execution.

### Readings for 2.1

- Futures and promises, https://en.wikipedia.org/wiki/Futures_and_promises

## 2.2 Futures in Java's Fork/Join Framework

## 2.3 Memoization

## 2.4 Java Streams

## 2.5 Data Races and Determinism
