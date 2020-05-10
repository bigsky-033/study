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

- Some key differences between future tasks and regular tasks in the ForkJoin framework are as follows:
  - A future task extends the RecursiveTask class in the FJ framework, instead of RecursiveAction as in regular tasks.
  - The *compute()* method of a future task must have a non-void return type, whereas it has a void return type for regular tasks.
  - A method call like *left.join()* waits for the task referred to by object *left* in both cases, but also provides the task's return value in the case of future tasks.

## 2.3 Memoization

- `Memoization` is the idea which is to remember results of function calls *f(x)* as follows:
  - Create a data structure that stores the set *{(x_1, y__1 = f(x_1)), ((x_2, y_2 = f(x_2)), ...}* for each call f(x_i) that returns y_i.
  - Perform lookups in that data structure when processing calls of the form f(x_') when x_' equals one of the x_i inputs for which f(x_i) has already been computed.
- The memoization patter lends itself easily to parallelization using futures by modifying the momoized data structure to store futures. The loopup operation can then be replaced by *get()* operation on the future value, if a future has already been created for the result of a given input.

### Readings for 2.3

- Memoization, https://en.wikipedia.org/wiki/Memoization

## 2.4 Java Streams

From the view point of `Parallel Programming in Java`, an important benefit of using Java streams whjen possible is that the pipeline can be made to execute in parallel by designating the source to be a *parallel* stream, i.e, by simply replacing *collection.stream()* by *collection.parallelStream()* or *Stream.of(collection).parallel()*.

### Readings for 2.4

- Article on "Processing Data with Java SE 8 Streams", https://www.oracle.com/technical-resources/articles/java/ma14-java-se-8-streams.html
- Tutorial on specifying Aggregate Operations using Java streams, https://docs.oracle.com/javase/tutorial/collections/streams/
- Documentation on java.util.stream.Collectors class for performing reductions on streams, https://docs.oracle.com/javase/8/docs/api/java/util/stream/Collectors.html

## 2.5 Data Races and Determinism

- Functional deterministic: if it always computes the same answer when given the same input.
- Structurally deterministic: if it always same computation graph, when given the same input.
- The presence of data races often leads to functional and/or structural nondeterminism because a parallel program with data races may exhibit different behaviors for the same input, depending on the ralative scheduling and timing of memory accesses involved in a data race.
- `Benign` nondeterminism: for programs with data races in which different executions with the same input may generate different ouptuts, but all the outputs may be acceptable in the context of the application.
