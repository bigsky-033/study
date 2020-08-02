# Parallel Programming in Java - Week3

Notes for `Parallel Programming in Java` class in Coursera.

- <https://www.coursera.org/learn/parallel-programming-in-java>

## 3.1 Parallel Loops

- Parallel loop as an async task

```
finish {
    for (p = head; p != null; p = p.next) async compute(p)
}
```

- Parallel loop for *counted-for loops*

```

forall (i : [0:n-1]) a[i] = b[i] + c[i]

or

a = IntStream.rangeClosed(0, N-1).parallel().toArray(i -> b[i] + c[i]);
```

### Readings for 3.1

- Tutorial on Executing Streams in Parallel, <https://docs.oracle.com/javase/tutorial/collections/streams/parallelism.html#executing_streams_in_parallel>

## 3.2 Parallel Matrix Multiplication

- Matrix Multiplication pseudocode

```
for (i : [0:n-1]) {
    for (j : [0:n-1]) {
        c[i][j] = 0
        for (k : [0:n-1]) {
            c[i][j] = c[i][j] + a[i][k] * b[j][j]
        }
    }
}
```

- Q: Which part can be executed in parallel?
  - A: It is safe to convert for-i and for-j into *forall* loops, but for-k must remain a sequential loop to avoid data races.

### Readings for 3.2

- Matrix multiplication algorithm, <https://en.wikipedia.org/wiki/Matrix_multiplication_algorithm>

## 3.3 Barriers in Parallel Loops

Barriers extend a parallel loop by dividing its execution into a sequence of *phases*. While it may be possible to write a separate *forall* loop for each phase, it is both more convenient and more efficient to instead insert barriers in a single *forall loop.

## 3.4 Parallel One-Dimensional Iterative Averaging

### Readings for 3.3

- Stencil codes, <https://en.wikipedia.org/wiki/Stencil_code>

## 3.5 Iteration Grouping / Chunking in Parallel Loops

- It is wasteful when number of parallel tasks is much larger than the number of available processor cores. To address this problem, there is a common tactic used in practice that is referred to as *loop chunking* or *iteration grouping*, and focuses on reducing the nubmber of tasks created to be closer to the number of processor cores, so as to reduce the overhead of parallel execution.
- Example
  
```
forall (i : [0:n-1]) a[i] = b[i] + c[i]

=>
// g: group

forall (g : [:ng-1])
    for (i : mygroup(g, ng, [0:n-1])) a[i] = b[i] + c[i]
```

- There are two well known approaches for iteration groups: *block* and *cyclic*.
  - Block: maps consecutive iterations to the same group
  - Cyclic: maps iteration in the same congruence class (mod *ng*) to the same group.
