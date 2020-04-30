# Union Find

- References
  - [WilliamFiset] Union Find Introduction, https://youtu.be/ibjEGG7ylHk
  - [WilliamFiset] Union Find Kruskal's Algorithm, https://youtu.be/JZBQLXgSGfs
  - [WilliamFiset] Union Find - Union and Find Operations, https://youtu.be/0jNmHPfA_yE
  - [WilliamFiset] Union Find Path Compression, https://youtu.be/VHRhJWacxis

## Discussion & Example

### What is Union Find

`Union Find` is a data structure that keeps track of elements which are split into one or more disjoint sets. It has two primary operations: `find` and `union`.

- Find: given an element the Union Find will tell you what group that element belongs to
- Union: merges two groups together

### When and where is a Union Find used

- Kruskal's minimum spanning tree algorithm
- Grid percolation
- Network connectivity
- Least common ancestor in trees
- Image processing

### Complexity analysis

| Operation          | Time Complexity |
| ------------------ | --------------- |
| Construction       | O(n)            |
| Union              | alpha(n)        |
| Find               | alpha(n)        |
| Get component size | alpha(n)        |
| Check if connected | alpha(n)        |
| Count components   | O(1)            |

- alpha(n): amortized constant time, almost constant time

### Kruskal's minimum spanning tree algorithm

#### What is minumum spanning tree

Given a graph G = (V, E) we want to find a `Minumum Spanning Tree` in the graph (it may not be unique). A minimum spanning tree is a subset of the edges which connect all the vertices in the graph with the minimal total edge cost.

#### How does work

1. Sort edges by ascending edge weight
2. Walk through the sorted edges and look at the two nodes the edge belongs to, if the nodes are already unified we don't include this edge(prevent cycle), otherwise we include it and unify the nodes.
3. The algorithm terminates when every edge has been processed or all the vertices have been unified

### Operation details

#### Creating Union Find

- To begin using Union Find, first construct a `bijection` (a mapping) between your objects and the integers in the range [0, n).
  - NOTE: This step is not necessary in general, but it wll allow us to construct an array-based union find.
- Store Union Find information in an array. Each index has an associated object (like letter) we can lookup through our mapping.

#### Find Operation

To `find` which component a particular element belongs to find the root of that component by following the parent nodes until a self loop is reached (a node who's parent is itself).

#### Union Operation

To `unify` two elements find which are the root nodes of each component and if the root nodes are different make one for the root nodes be the parent of the other.

### Path compression

Ensure every node in the same group except root, points root directly.
