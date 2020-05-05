# Binary Search Tree

- References
  - [WilliamFiset] Binary Search Tree Introduction, https://youtu.be/JfSdGQdAzq8
  - [WilliamFiset] Binary Search Tree Insertion, https://youtu.be/LwpLXm3eb6A
  - [WilliamFiset] Binary Search Tree Removal, https://youtu.be/8K7EO7s_iFE
  - [WilliamFiset] Binary Search Tree Traversals, https://youtu.be/k7GkEbECZK0

## Discussion & examples

### What is Tree

A `tree` is an `undirected graph` which satisfies any of the following definitinos:

- An acyclic connected graph.
- A connected graph with N nodes and N-1 edges.
- An graph in which any two vertices are connected by `exactly` one path.

A `child` is a node extending from another node. A `parent` is the inverse of this.

A `leaf node` is a node with no children.

A `sub tree` is a tree entirely contained within another.

- Subtrees may consist of a sinlge node

### What is a Binary Tree(BT)

A `binary tree` is a stree for which every node has at most two child nodes.

### What is a Binary Search Tree(BST)

A `binary search tree` is a binary tree that satisfies the `BST invariant`:

- Left subtree has smaller elements and right subtree has larger elements.

### Where are BTs and BSTs used

- Binary Search Trees(BSTs)
  - Implementation of some map and set ADTs
  - Red Black Trees
  - AVL Trees
  - Splay Trees
  - Etc...
- Used in the implementation of binary heaps
- Syntax trees

## Complexity Analysis

| Operation | Average   | Worst |
| --------- | --------- | ----- |
| Insert    | O(log(n)) | O(n)  |
| Delete    | O(log(n)) | O(n)  |
| Remove    | O(log(n)) | O(n)  |
| Search    | O(log(n)) | O(n)  |

## Implementations

Binary Search Tree(BST) elements must be `comparable` so that we can order them inside the tree.

### Insert

When inserting an element we want to compare its value to the value stored in the current node we're considering to decide on one of the following:

- Recurse down left subtree (< case)
- Recurse down right subtree (> case)
- Handle finding a duplicate value (= case)
- Create a new node (found a null leaf)

### Remove

Removing elements from a Binary Search Tree(BST) can be seen as a two step process.

- Find the element we want to remove(if it exists).
- Replace the node we want to remove with its successor (if any) to maintain the BST invariant.

#### Find phase

When searching our BST for a node with a particular value one of four things will happen:

- We hit a `null node` at which point we know the value does not exist within our BST
- Comparator value `eqaul to 0` (found)
- Comparator value `less than 0` (the value, if it exists, is in the left subtree)
- Comparator value `greater than 0` (the value, if it exists, is in the right subtree)

#### Remove phase

- Node to remove is a leaf node.
  - If the node we want to remove is a left node, then we amy do without side effect.
- Node to remove has a right subtree but no left subtree.
  - The successor of the node we are trying to remove will be the root node of the `right subtree`.
- Node to remove has a left subtree but no right subtree.
  - The successor of the node we are trying to remove will be the root node of the `left subtree`.
- Node to remove has a both a left subtree and a right subtree.
  - Q: In which subtree will be the successor of the node we are trying to remove bew?
  - A: The answer is both! The successor can either be the `largest value in the left subtree` OR the `smallest value in the right subtree`.
  - Justification for why there could be more than one successor is:
    - The `largset value` in the `left subtree` satisfies the BST invariant since it:
      - Is larger than everything in the left subtree. This follows immediately from the definition of being the largest.
      - Is smaller than everything in right subtree because it was found in the left subtree.
    - The `smallest value` in the `right subtree` satisfies the BST invariant (similar with above).

### Treversal

#### Preoreder

Preorder do what you want for node before the recursive calls.

```
preorder(node):
    if node == null: return
    do(node)
    preorder(node.left)
    preorder(node.right)
```

#### Inorder

Ineorder do what you want for node between the recursive calls.

```
inorder(node):
    if node == null: return
    inorder(node.left)
    do(node)
    inorder(node.right)
```

#### Postorder

Postorder do what you want for node after the recursive calls.

```
postorder(node):
    if node == null: return
    postorder(node.left)
    postorder(node.right)
    do(node)
```

#### Level order

In a level order traversal we want to print the nodes as they appear one layer at a time.

To obtain this ordering we want to do a `Bread First Search (BFS)` from the root node down to the leaf nodes.

To do a BFS we will need to maintain a `Queue` of the nodes left to explore.
