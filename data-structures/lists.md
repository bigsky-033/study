# Lists

- References
  - [WilliamFiset] Linked Lists Introduction, https://youtu.be/-Yn5DU0_-lw

## About Singly & Doubly Linked Lists

### What is a linked list

- A linked list is a sequential list of nodes that hold data which point to other nodes also containing data.

### Where are linked lists used

- Used in many List, Queue & Stack implementations.
- Greate for creating circular lists.
- Can easily model real world objects such as trains.
- Used in separate chaining, which is present certain Hash Table implementations to deal with hashing collisions.
- Often used in the implementation of adjacency lists for graphs.

### Terminology

- Head: The first node in a linked list
- Tail: The last node in a linked list
- Pointer: Reference to another node
- Node: An object containing data and pointer(s)

### Singly Linked vs Double Linked

- Singly linked lists only hold a reference to the next node.
- Double linked list each node holds a reference to the next and previous node.

- Pros and cons

|        | Pros                                | Cons                                   |
| ------ | ----------------------------------- | -------------------------------------- |
| Singly | Less memory, Simpler implementation | Cannot easily access previous elements |
| Doubly | Can be traversed backwards          | More memory (2x)                       |

## Complexity analysis

| Operation        | Singly | Doubly |
| ---------------- | ------ | ------ |
| Search           | O(n)   | O(n)   |
| Insert at head   | O(1)   | O(1)   |
| Insert at tail   | O(1)   | O(1)   |
| Remove at head   | O(1)   | O(1)   |
| Remove at tail   | O(n)   | O(1)   |
| Remove in middle | O(n)   | O(n)   |

- Remove at tail on Singly Linked List is O(n) even we have pointer because to remove it we need to know previous node. To do that on Singly Linked List, we need to traverse from head.
