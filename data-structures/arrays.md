# Arrays

- References
  - [WilliamFiset] Dynamic and Static Arrays, https://youtu.be/PEnFFiQe1pM

## What is an static Array

- A static array is a fixed length container containing n elements indexale from the range [0, n-1].
- Indexable is that each slot/index in the array can be referenced with a number.
- Static arrays are given an contiguous chunks of memory.

## When and where is an Array used

1. Storing and accessing sequential data
2. Temporarily storing objects
3. Used by IO routines as buffers
4. Lookup tables and inverse lookup tables
5. Can be used to return multiple values from a function
6. Used in dynamic programming to cache answers to subproblems

## Complexity

| Operation | Static Array | Dynamic Array |
| --------- | ------------ | ------------- |
| Access    | O(1)         | O(1)          |
| Search    | O(n)         | O(n)          |
| Insert    | N/A          | O(n)          |
| Append    | N/A          | O(1)          |
| Delete    | N/A          | O(n)          |

## Dynamic Array

- Dynamic array can do all the similar get set operation static array can do. But unlikely the static array it grows inside as dynamically as needed.
- Q. How can we implement a dynamic array?
  - A. One way is to use a static array!
    1. Create a static array with an initial capacity
    2. Add elements to the underlying static array, keeping track of the number of elements.
    3. If adding another element will exceed the capacity, then create a new static array with twice the capacity and copy the original elements into it.
