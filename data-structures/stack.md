# Stack

- References
  - [WilliamFiset] Stack Introduction, https://youtu.be/L3ud3rXpIxA

## Discussion about Stacks

### What is a Stack

- A stack is a one-ended linear data structure which models a real world stack by having two primary operatinos, namely `push` and `pop`.

### When and where is a Stack used

- Used by undo mechanisms in text editors.
- Used in compiler syntax checking for matching brackets and braces.
- Can be used to model a pile of books or plates.
- Used behind the scenes to support recursion by keeping track of previous function calls.
- Can be used to do a Depth First Search(DFS) on a graph.

### Complexity Analysis

| Operation | Time Complexity |
| --------- | --------------- |
| Push      | O(1)            |
| Pop       | O(1)            |
| Peek      | O(1)            |
| Search    | O(n)            |
| Size      | O(1)            |
