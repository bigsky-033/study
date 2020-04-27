# Priority Queue

- References
  - [WilliamFiset] Priority Queue Introduction, https://youtu.be/wptevk0bshY
  - [WilliamFiset] Priority Queue Min Heaps and Max Heaps, https://youtu.be/HCEr35qpawQ
  - [WilliamFiset] Priority Queue Inserting Elements, https://youtu.be/QOJ-CmQiXko
  - [WilliamFiset] Priority Queue Removing Elements, https://youtu.be/eVq8CmoC1x8

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

### Ways of Implementing a Priority Queue

- Priority Queues are usually implemented with heaps since this gives them the best possible time complexity.
- The Priority Queue is an Abstract Data Type(ADT), hence heaps are not only way to implement PQs. As an example, we could use an unsorted list, but this would not give us the best possible time complexity.

#### Binary Heap

- A binary heap is a binary tree that supports that heap invariant. In a binary tree very node has exactly two children.
- A complete binary tree is a tree in which at every level, except possibly the last is completely filled and all the nodes are as far left as possible.
- Binary heap array
  - Let i be the parent node index
  - Left child index: 2i + 1
  - Right child index: 2i + 2

### Adding elements to Priprity Queue

- Add at last and trying to satisfy heap invariant.
  - Compare with parent & if parent is bigger(for mean heap), swap it.

### Removing (polling) elements from Priority Queue

- Poll: Swap last node & root node & trying to satisfy heap invariant.
  - Complexity: O(log(n))
  - Look both children & swap with smaller one. Repeat this until satisfy heap invariant.
- Remove some node - simple
  - Complexity: O(n)
  - Search target value
  - Swap with last node
  - Satisfy heap invariant

#### Removing elements from binary heap in O(log(n))

- The inefficiency of the removal algorithm comes from the fact that we have to perform a linear search to find out where an element is index at. What if instead we did a lookup using a `Hashtable` to find out where a node is indexed at?
- A hashtable provides a constant time lookup and update for a mapping from a key to a value.

- Caveat: What if there are two or more nodes with the same value? What problems would that cause?
  - Dealing with the multiuple value problem:
    - Instead of mapping one value to one position we will map one value to multuple positions. We can maintain a `Set` or `TreeSet` of indexes for which a particular node value maps to.
- Question: If we want to remove a repeated node in our heap, which node do we remove and does it matter which one we pick?
  - Answer: No it doesn't matter which node we remove as long as we satisfy the heap invariant in the end.

### Heap sinking and swimming

- To satisfy heap invariant, after inserting node at last & sinking & swimming.
- Bubbling up & down
