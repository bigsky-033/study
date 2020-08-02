# Parallel Programming in Java - Week1

Notes for `Parallel Programming in Java` class in Coursera.

- <https://www.coursera.org/learn/parallel-programming-in-java>

## 1.1 Task Creation and Termination (Async, Finish)

- async [stmt]: Causes parent task to create a new child task to execute the body of the `async`([stmt]) asynchronously with the raminder of the parent task.
- finish [stmt]: Causes the parent task to execute [stmt], and then wait until [stmt] and all sync tasks created within [stmt] have completed.

```
finish {
  async S1; // asynchronously compute sum of the lower half of the array
  S2;       // compute sum of the upper half of the array in parallel with S1
}
S3; // combine the two partial sums after both S1 and S2 have finished
```

### Readings for 1.1

- Asynchronous method invocation, <https://en.wikipedia.org/wiki/Asynchronous_method_invocation>

## 1.2 Creating Tasks in Java's Fork/Join Framework

### Readings for 1.2

- Tutorial on Java’s Fork/Join framework, <https://docs.oracle.com/javase/tutorial/essential/concurrency/forkjoin.html>
- Documentation on Java’s RecursiveAction class, <https://docs.oracle.com/javase/7/docs/api/java/util/concurrent/RecursiveAction.html>

## 1.3 Computation Graphs, Work, Span, Ideal Parallelism

### Computation Graphs (CGs)

- Computation Graphs: Model the executhion of a parallel program as a partially ordered set.
- For `fork-join` programs, it is useful to partition the edges into three caes
  - Continue: Capture sequencing of steps within a task.
  - Fork: Connect a fork operation to the first step of child tasks.
  - Join: Connect the last step of a task to all join operations on that task.
- Data race

> CGs can be used to define data races, an important class of bugs in parallel programs. We say that a data race occurs on location L in a computation graph, G, if there exist steps S1 and S2 in G such that there is no path of directed edges from S1 to S2 or from S2 to S1 in G, and both S1 and S2 read or write L (with at least one of the accesses being a write, since two parallel reads do not pose a problem).

- `Ideal parallelism` of a parallel program
  - WORK(G): Sum of the execution times of all nodes in CG G.
  - SPAN(G): Length of a longest path in G, when adding up the execution times of all nodes in the path. The longest paths are known as critical paths, so SPAN also represents the `critical path length`(CPL) of G.
  - Ideal parallelism of Computation Graph G as the ration, `WORK(G) / SPAN(G)`.
    - This is the upper limit on the speedup factor that can be obtained from parallel execution of nodes in computation graph G.

### Readings for 1.3

- Partailly ordered set, <https://en.wikipedia.org/wiki/Partially_ordered_set>
- Analysis of parallel algorithms, <https://en.wikipedia.org/wiki/Analysis_of_parallel_algorithms>

## 1.4 Multiprocessor Scheduling, Parallel Speedup

- T_p: execution time of a `CG` on `P` processors, and observed that
  - T_infinity <= T_p <= T_1
  - T_p could be different how schedules on the same CG
- Speedup(p): parallel speedup for a given schedule of a CG on P processors, and observed that
  - Speedup(p) = T_1 / T_p
  - Speedup(p) <= # of processors P
  - Speedup(p) <= WORK/SPAN (ideal parallelism)

## 1.5 Amdahl's Law

> if q ≤ 1 is the fraction of WORK in a parallel program that must be executed sequentially, then the best speedup that can be obtained for that program for any number of processors, P , is Speedup(P) ≤ 1/q.

### Readings for 1.5

- Amdahl's law, <https://en.wikipedia.org/wiki/Amdahl%27s_law#Speedup_in_a_serial_program>