# Fenwick Tree

- References
  - [WilliamFiset] Fenwick Tree range queries, https://youtu.be/RgITNht_f4Q
  - [WilliamFiset] Fenwick Tree point updates, https://youtu.be/B-BkW9ZpKKM
  - [WilliamFIset] Fenwick Tree construction, https://youtu.be/BHPez138yX8

## What is a Fenwick tree

A `Fenwick Tree`(also called Binary Indexed Tree) is a data structure that supports sum range queries as well as setting values in a static array and getting the value of the prefix sum up some index efficiently.

## Complexity analysis

| Operation      | Complexity |
| -------------- | ---------- |
| Construction   | O(n)       |
| Point Update   | O(log(n))  |
| Range Sum      | O(log(n))  |
| Range Update   | O(log(n))  |
| Adding Index   | N/A        |
| Removing Index | N/A        |

## Implementation details

### Range query

- Unlike a regular array, in a Fenwick tree a specific cell is responsible for other cells as well.
- The position of the `least significant bit` (LSB) determines the range of responsibility that cell has to the cells below itself.
- In a Fenwick tree we may compute the `prefix sum` up to a certain index, which ultimately lets us perform range sum queries.
  - Idea: Suppose you want to find the prefix sum of [1, i], then you `start at i and cascade downnwards` until you reach zero adding the value at each of the indices you encounter.
    - Ex) Find the prefix sum up to index 7
      - sum = A[7] + A[6] + A[4]
    - Ex) Find the prefix sum up to index 11
      - sum = A[11] + A[10] + A[8]
    - Ex) Compute the interval sum between [11, 15]
      - sum = (prefix sum of *[1,15]*) - (prefix sum of *[1,11)*)
        - prefix sum of [1,15] => A[15] + A[14] + A[12] + A[8]
        - prefix sum of [1,11] => A[10] + A[8]

#### Range query algorithm

To do a range query from [i,j] both inclusive a Fenwick tree of size N:

```
function prefixSum(i):
    sum := 0
    while i != 0:
        sum = sum + tree[i]
        i = i = LSB(i)
    return sum


function rangeQuery(i, j):
    return prefixSum(j) - prefixSum(i - 1)

Where LSB returns the value of the least significant bit.
```

### Point updates

- Instead of querying a range to find the interval sum we want to update a cell in our array. For this, we can cascaded down from the current index by *continuously removing the LSB*.
- Point updates are the opposite of this, we want to *add the LSB* to propagate the value up to the cells responsible for us.
- Point update algorithm
  - Update the cell at index i in the Fendwick tree of size N

```
function add(i, x):
    while i < N:
        tree[i] = tree[i] + x
        i = i + LSB(i)

where LSB returns the value of the least significant bit.
```

### Fenwick tree construction

#### Naive Construction

Let A be an array of values. For each element in A at index i do a point update on the Fenwick tree with a value of A[i]. There are *n* elements and each point update takes O(log(n)) for a total of *O(nlog(n))*.

#### Linear Construction

- Idea: Add the value in the current cell to the `immediate cell` that is responsible for us. This resembles what we did for point updates but only one cell at a time.
- This will make the 'cascading' effect in range queries possible by propagating the value in each cell throughout the tree.
  - Let i be the current index
  - The immediate cell above us is a position j given by:
    - j := i + LSB(i), where LSB is the *Least Significant Bit* of i
- Constuction algorithm

```
# Make sure values is 1-based

function construct(values):
    N := length(values)

    tree = deepCopy(values)

    for i = 1...N:
        j := i + LSB(i)
        if j < N:
            tree[j] = tree[j] + tree[i]

    return tree
```
