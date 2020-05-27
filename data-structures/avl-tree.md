# AVL Tree

- References
  - [WilliamFiset] Balanced binary search tree rotations, https://youtu.be/q4fnJZr8ztY
  - [WilliamFiset] AVL tree insertion, https://youtu.be/1QSYxIKXXP4
  - [WilliamFiset] AVL tree removals, https://youtu.be/g4y2h70D6Nk

## Balanced Binary Search Tree

### What is a BBST

A `Balanced Binary Search Tree (BBST)` is a *self-balancing* binary search tree. This type of tree will adjust itself in order to maintain a low (logarithmic) height allowing for faster operations such as insertions and deletions.

### Time complexity of Binary Search Trees

| Operation | Average   | Worst |
| --------- | --------- | ----- |
| Insert    | O(log(n)) | O(n)  |
| Delete    | O(log(n)) | O(n)  |
| Remove    | O(log(n)) | O(n)  |
| Search    | O(log(n)) | O(n)  |

### Time complexity of Balanced Binary Search Trees

| Operation | Average   | Worst     |
| --------- | --------- | --------- |
| Insert    | O(log(n)) | O(log(n)) |
| Delete    | O(log(n)) | O(log(n)) |
| Remove    | O(log(n)) | O(log(n)) |
| Search    | O(log(n)) | O(log(n)) |

### Tree Rotations

- A `tree invariant` is a property you impose on your tree that it must meet after every operation. To ensure that the invariant is always satisfied a series of tree rotations are normally applied.
- It does not matter what the structure of the tree looks; all we care about is that the BST invariant holds. This means we can shuffle/transform/rotate the values and nodes in the tree as we please as long as the BST invariant remains satisfied.

## ABout AVL Tree

### AVL Tree Intro

An `AVL Tree` is one of many types of *Balanced Binary Search Trees (BBSTs)* which allow for logarithmic *O(log(n))* insertion, deletion and search operations.

### AVL Tree Invariant

The property which keeps an AVL tree balanced is called the *Balanced Factor (BF)*.

- BF(node) = H(node.right) - H(node.left)
  - H(x) is the height of node x. Recall that H(x) is calculated as the *number of edges* between x and the furthest leaf.
- Invariant in the AVL which forces it to remain balanced is the requirement that the balance factor is always either -1, 0 or +1.

### Node Information to Store

- The actual value we're storing in the node.
  - Value must be comparable.
- A value storing this node's *balance factor*.
- The *height* of this node in the tree.
- Pointer to the *left/right child nodes*.

### How to resolve inbalance

If a node's BF does not included in [-1, 0, +1] then the BF of that node is +-2 which can be adjusted using *tree rotations*.

- Left Left Case: Right rotation
- Left Right Case: Left rotation -> Right rotation
- Right Right Case: Left rotation
- Right Left Case: Right rotation -> Left rotation

### How to Insert

```
# public insert
function insert(value):
  if value == null:
    return false

  # Only insert unique values
  if !contains(root, value):
    root = insert(root, value)
    nodeCount = nodeCount + 1
    return true

  # Value already exists in tree
  return false
```

```
# private insert
function insert(node, value):
  if node == null: return Node(value)

  cmp := compare(value, node.value)

  if cmp < 0:
    node.left = insert(node.left, value)
  else:
    node.right = insert(node.right, value)

  # update balance factor and height values
  update(node)

  # rebalance
  return balance(node)
```

```
# update
function update(node):

  # Variables for left/right subtree heights
  lh := -1
  rh := -1
  if node.left != null: lh = node.eft.height
  if node.right != null: rh = node.right.height

  node.height = 1 + max(lh, rh)

  node.bf = rh - lh
```

```
function balance(node):
  # Left heavy subtree
  if node.bf == -2:
    if node.left.bf <= 0:
      return leftLeftCase(node)
    else:
      return leftRightCase(node)

  # Right heavy subtree
  else if node.bf == +2:
    if node.right.bf >= 0:
      return rightRightCase(node)
    else:
      return rightLeftCase(node)
  
  # Already balanced
  return node
```

```
function leftLeftCase(node):
  return rightRotation(node)

function leftRightCase(node):
  node.left = leftRotation(node.left)
  return leftLeftCase(node)

function rightRightCase(node):
  return leftRotation(node)

function rightLeftCase(node):
  node.right = rightRotation(node.right)
  return rightRightCase(node)
```

```
function rightRotate(A):
  B := A.left
  A.left = B.right
  B.right = A

  # After rotation update balance factor and height values
  update(A)
  update(B)
  return B
```

### How to Remove

Removing element in AVL Tree is almost identical with removing element in BST

#### Removing Elements from a BST

- Find the element we wish to remove
- If exists, Replace the node we want to remove with its successor to maintain the BST invariant

#### Augmenting BST Removal Algorithm for AVL Tree

```
function remove(node, value):
  // ...
  // Code for BST item removal here
  // ...

  # Update balance factor
  update(node)

  # Rabalance tree
  return balance(node)
```
