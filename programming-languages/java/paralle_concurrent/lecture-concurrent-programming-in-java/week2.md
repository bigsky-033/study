# Concurrent Programming in Java - Week1

Notes for `Concurrent Programming in Java` class in Coursera.

- https://www.coursera.org/learn/concurrent-programming-in-java

## 2.1 Critical Sections

- With critical sections, two block of code that are marked as isolated,say A and B, are guranteed to be executed in mutual exclusion with A executing before B or vice versa.

### Readings for 2.1

- Critical section, https://en.wikipedia.org/wiki/Critical_section
- Atomicity, https://en.wikipedia.org/wiki/Atomicity_(database_systems)

## 2.2 Object Based Isolation (Monitors)

- The fundamental idea behind object-based isolation is that an isolated construct can be extended with a set of objects that indicate the scope of isolation, by using following rules:
  - if two isolated constructs have an empty intersection in their object sets they can execute in parallel, otherwise they must execute in mutual exclusion.
- The relationship between object-based isolation and monitors is that all methods in a monitor object, *M1*, are executed as object-based isolated constructs with a singleton object set, *{M1}*. Similarly, all methods in a monitor object, *M2*, are executed as object-based isolated constructs with a singleton object set, *{M2}* which has an empty intersection with *{M1}*.

### Readings for 2.2

- Monitor, https://en.wikipedia.org/wiki/Monitor_(synchronization)

## 2.3 Concurrent Spanning Tree Algorithm

### Readings for 2.3

- Spanning Tree, https://en.wikipedia.org/wiki/Spanning_tree

## 2.4 Atomic Variables

- If we use *Atomic Integer* variable, below two operations are same.

```
# using atomic integer

cur: Atomic Ineger

j = cur.getAndAdd(1)

# object-based isolation

isolated(cur) {
    j = cur;
    cur = cur + 1;
}
```

- *Atmoic Reference variable*'s *compareAndSet(expected, new)* works similarly.

### Readings for 2.4

- Tutorial on Atomic Integers in Java, https://docs.oracle.com/javase/tutorial/essential/concurrency/atomicvars.html
- Java theory and practice Going atomic, https://www.ibm.com/developerworks/library/j-jtp11234/
- Primitive wrapper class in Java, https://en.wikipedia.org/wiki/Primitive_wrapper_class_in_Java#Atomic_wrapper_classes

## 2.5 Read, Write Isolation

- The main idea behind read-write isolation is to separate read accesses to shared objects from write accesses. This approach enables two threads that only read shared objects to frelly execute in parallel since they are not modifying any shared objects. The need for mutual exclusion only arises when one or more threads attempt to enter an isolated section with write access to a shared object.
- This approach exposes more concurrency than object-based isolation since it allows read accesses to be executed in parallel.

### Readings for 2.5

- Readers-writer lock, https://en.wikipedia.org/wiki/Readers%E2%80%93writer_lock