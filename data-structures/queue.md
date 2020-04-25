# Queue

- References
  - [WilliamFiset] Queue Introduction, https://youtu.be/KxzhEQ-zpDc

## Discussion about Queues

### What is queue

- A queue is a linear data structure which models real world queues by having two primary operations, namely `enqueue` and `dequeue`.

### Terminology

- Enqueue == Adding == Offering
- Dequeue == Polling

### When and where is a queue used

- Any waiting line models a queue, for example a lineup at a movie theatre.
- Can be used to efficiently keep track of the x most recently added elements.
- Web server request management where you want first come first serve.
- Breadth first search(BFS) graph traversal.

### Complexity Analysis

| Operation | Time Complexity |
| --------- | --------------- |
| Enqueue   | O(1)            |
| Dequeue   | O(1)            |
| Peeking   | O(1)            |
| Contains  | O(n)            |
| Removal   | O(n)            |
| Is Empty  | O(1)            |

### Queue Example - BFS

```
Let Q be a Queue
Q.enqueue(starting_node)
starting_node.visited = true

while Q is not empty Do
    node = Q.dequeue()

    For neighbour in neighbours(node):
        If neighbour has not been visited:
            neighbour.visited = true
            Q.enqueue(neighbour)
```