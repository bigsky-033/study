# Priority Queue

- References
  - [WilliamFiset] Priority Queue Introduction, https://youtu.be/wptevk0bshY
  - [WilliamFiset] Priority Queue Min Heaps and Max Heaps, https://youtu.be/HCEr35qpawQ
  - [WilliamFiset] Priority Queue Inserting Elements, https://youtu.be/QOJ-CmQiXko

## Discussion & Examples of Priority Queue

### What is a Priority Queue

- Priority Queue is an Abstract Data Type that operates similar to a normal queue except that `each element has a certain priority`. The property of the elements in the priority queue determine the order in which elements are removed from the PQ.
- Priority Queue only supports `comparable data`, meaning the data inserted into the priority queue must be able to be ordered in some way either from least to greatest or greatest to least. This is so that we are able to assign relative priorities to each element.

### What is a Heap

- A heap is a `tree` based Data Structure that satisfies the `heap invariant`(also called heap property):
  - If A is a parent node of B then A is ordered with respect to B for all nodes A, B in the heap.

### When and where is a Priority Queue used

- Used in certain implementations of Dijkstar's Shortest Path algorithm.
- Anytime you need the dynamically fetch `next best` or `next worst` element.
- Used in Huffman coding(which is often used for lossless data compression).
- Best First Search (BFS) algorithms such A* use PQs to continuously grab the next most promising node.
- Used by Minimum Spanning Tree(MST) algorithms.

### Complexity Analysis

| Operation                         | Time Complexity |
| --------------------------------- | --------------- |
| Binary Heap Construction          | O(n)            |
| Poll                              | O(log(n))       |
| Peek                              | O(1)            |
| Add                               | O(log(n))       |
| Naive Remove                      | O(n)            |
| Advanced removing with hash table | O(log(n))       |
| Naive contains                    | O(n)            |
| Contains with hash table          | O(1)            |

- Using a hash table to help optimize operations does take up linear space and also adds some ovverhead to the binary heap implementation.

## Priority Queue Min Heaps and Max Heaps

### Why is it important

- Problem: Often the standard library of most programming languages only provides a min PQ which sorts by smallest elements first, but sometimes we need a MAX PQ.
- Since elements in a priority queue are comparable they implement some sort of comparable interface which we can simply negate to achieve a Max heap.

### How to - Numbers

- Let x, y be numbers in the PQ. For a min PQ, if `x <= y` then x comes out of the PQ before y, so the negation of this is `x >= y` (not x > y) then y comes out before x.
- An alternative method for numbers is to negate the numbers as you insert them into the PQ and negate them again when they are taken out. THis has the same effect as negating the comparator.

### How to - Strings

- Suppose `lex` is a comparator for strings which sorts strings in lexicographic order (the default in most programming languages). Then let `nlex` be the negation of `lex`, and also set s_1, s_2 be strings
  - lex(s_1, s_2) = -1 if s_1 < s_2 lexicographically
  - lex(s_1, s_2) = 0 if s_1 = s_2 lexicographically
  - lex(s_1, s_2) = +1 if s1 > s_2 lexicographically
  - nlex(s_1, s_2) = -(-1) = +1 if s_1 < s_2 lexicographically
  - nlex(s_1, s_2) = -(0) = 0 if s_1 = s_2 lexicographically
  - nlex(s_1, s_2) = -(+1) = -1 if s1 > s_2 lexicographically

- By adding all strings to the PQ with the `nlex` comparator

## Binary heap Priority Queue Implementation Details

### Heap sinking and swimming

### Adding elements to Priority Queue

### Removing (polling) elements from Priority Queue

## Code Implementation