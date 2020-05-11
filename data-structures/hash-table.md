# Hash Tables

- References
  - [WilliamFiset] Hash table hash function, https://youtu.be/2E54GqF0H4s
  - [WilliamFiset] Hash table separate chaining, https://youtu.be/T9gct6Dx-jo
  - [WilliamFiset] Hash table open addressing, https://youtu.be/xIejolxzZS8
  - [WilliamFiset] Hash table linear probing, https://youtu.be/Ma9XOInZJWM
  - [WilliamFiset] Hash table quadratic probing, https://youtu.be/b0858c55TGQ
  - [WilliamFiset] Hash table double hashing, https://youtu.be/H5e9V5x92vI
  - [WilliamFiset] Hash table open addressing removing, https://youtu.be/7eLDTtbzX4M

## What is a Hash table(HT) and what is a hash function

### Hash table(HT)

- A `Hash table`(HT) is a data structure that provides a mapping from keys to values using a technique called `hashing`.
- We refer to these as `key-value pairs`.
- Keys must be unique, but values can be repeated.
- HTs are often used to track item frequencies.
- The key-value pairs in a HT can be any type not just strings and numbers, but also objects. However, the keys needs to be `hashable`.

### Hash function H(x)

- A `hash function` H(x) is a function that maps a key `x` to a whole number in a `fixed range`.

## Properties of hash fuinctions

- If `H(x) = H(y)` then objects x and y `mignt be equal`, but if `H(x) != H(y)` then x and y are `certainly not equal`.
  - This property is very useful for comparing two objects(like files) which needs huge cost to compare every contents. For example, comparing files, we can firstly compare each file's `checksums`. If they are not equal, we do not need to open file.
- A hash function H(x) must be `deterministic`.
  - This means that if H(x) = y then H(x) must always produce y and never another value. This may seen obvious, but it is critical to the functionality of a hash function.
- We try very hard to make `uniform` hash functions to minimize the number of hash collisions.
  - A `hash collision` is when two objects x, y hash to the same value (i.e. H(x) = H(y)).

### What makes a key of type T hashable

- We need our hash functions to be deterministic. To enforce this behavior, we demand that the `keys used in our hash table are immutable` data types. - Hence, if a key of type T is immutable, and we have a hash function H(k) defined for all keys k of type T then we say a key of type T is hashable.

## How does a hash table work

- Ideally we would like to have a very fast insertion, lookup and removal time for the data we are placing within our hash table.
- Remarkably, we can achive all this in `O(1)` time using a `hash function as a way to index into a hash table`.
  - The constant time behavior attributed to hash tables is only true if you have a good `uniform hash function`.

### Complexity analysis

| Operation | Average | Worst |
| --------- | ------- | ----- |
| Insert    | O(1)*   | O(n)  |
| Removal   | O(1)*   | O(n)  |
| Search    | O(1)*   | O(n)  |

- The constant time behavior attributed to hash tables is only true if we have a good `uniform hash function`.

### Separate chaining

- `Separate chaining` is one of many strategies to deal with hash collisions by maintaining a data structure (usually a linked list) to hold all the different values which hashed to a particular value.
  - The data structure used to cache the items which hashed to a particular value is not limited to linked list. Some implementations use one or a mixture of: arrays, binary trees, self balancing trees and etc...

#### How can we maintain O(1) insertion and lookup time complexity

- Q: How do I maintain `O(1)` insertion and lookup time complexity once my HT gets really full and I have long linked list chains?
- A: Once the HT contains a lot of elements you should create a new HT with a larger capacity and rehash all the items inside the old HT and disperse them throughout the new HT at different locations.

### Open addressing

- When using open addressing as a collision resolution technique the `key-value pairs are stored in the table (array) itself` as opposed to a data structure like in separate chaining. This means we need to care a great deal about the size of our hash table and how many elements are currently in the table.
  - Load factor = (items in table) / (size of table)  
    - The `O(1)` constant time behavior attributed to hash tables assumes the load factor(alpha) is kep below a certain fixed value. This means once alpha > *threshold* we need to grow the table size (ideally exponentially, e.g double).

#### Main idea - Open addressing

- When we want to insert a key-value pair (k,v) into the hash table we hash the key and obtain an original position for where this key-value pair belongs, i.e H(k).
- If the position our key hashed to is occupied we try another position in the hash table by offsetting the current position subject to a `probiong sequence P(x)`. We keep doing this until an occupied slot is found.
- Linear probing
  - P(x) = ax + b where a, b are contants
- Quadratic probing
  - P(x) = ax^2 + bx + c, where a,b,c are constants
- Double hashing
  - P(k, x) = x * H_2(k), where H_2(k) is a secondary hash function
- Psudo random number generator
  - P(k, x) = x * RNG(H(k), x), where RNG is a random number generator function seeded with H(k).

- General insertion method for open addressing on a table of size N goes as follows:

```
x := 1
keyHash := H(k)
index := keyHash

while table[index] != null:
    index = (keyHash + P(k,x)) mod N
    x = x + 1

insert (k,v) at table[index]

Where H(k) is the hash for the key k and P(k,x) is the probing function
```

#### Chaos with cycles

Most randomly selected probing sequences modulo N will produce a cycle shorter than the table size.
This becomes problematic when you are trying to insert a key-value pair and all the buckets on the cycle are occupied because you will get stuck in an `infinite loop`!

- In general the consensus is that we don't handle this issue, instead we avoid it altogether by restricting our domain of probing functions to those which produce a cycle of exactly length N*.

#### Removing elements

- If we just remove key-value pair when remove operation executed(naive way), there can be a problem. Sometimes the key cleary exists in our table, due to the naive remove we cannot search that key.
- The solution is to place a `unique marker` called `tombstone` instead of null to indicate that a (k,v) pair has been deleted and that the bucket should be skipped during a search.
- Q: I have a lot of tombstones cluttering my HT how do I get rid of them?
- A: Tombstones count as filled slots in the HT so they increase the load factor and will be removed when the table is resized. Additionally, when inserting a new (k, v) pair you can replace buckets with tombstones with the new key-value pair.
- An optimization we can do it replace the earliest tombstone encountered with the value we did a lookup for. The next time we lookup the key it'll be found much faster. We call this `lazy deletion`.

#### What is Linear Probing(LP)

- LP is a `probing moethod` which probes according to a linear formula, specifically:
  - P(x) = ax + b where (a != 0), b are constants

##### Chaos with cycles - LP

- Q: Which value(s) of the constant a in P(x) = ax produce a full cycle modulo N?
- A: This happens when a and N are *relatively prime*. Two numbers are relatively prime if there `Greatest Common Denominator (GCD)` is equal to one. Hence, when GCD(a, N) = 1 the probing function P(x) be able to generate a complete cycle and we will always be able to find an empty bucket.
  => To avoid cycle, we need to chosse *a* which is relatively prime with *N*.

#### What is Quadratic Probing(QP)

- QP is `probing method` which probes according to a quadratic formula, specifically:
  - P(x) = ax^2 + bx + c where a, b, c are constants and a != 0

##### Chaos with cycles - QP

- Randomly selected QP functions have the issue that they easily produce short cycles.
- Q: How do we pick a probing function we can work with?
  - A: There are numerous ways, but three of the most popular approaches are:
    - 1) Let P(x) = x^2, keep the table size a prime number > 3 and also keep alpha <= 1/2
    - 2) Let P(x) = (x^2 + x) / 2 and keep the table size a power of two
    - 3) Let P(x) = (-1_^x) * x^2 and keep the table size a prime N where N = 3 mod 4

#### What is Double Hashing(DH)

- DH is a `probing method` which probes according to a constant multiple of another hash function, specifically:
  - P(k, x) = x * H_2(k), where H_2(k) is a second hash function
  - H_2(k) must hash the same type of keys as H_1(k)

##### Chaos with cycles - DH

- Since DH reduces to linear probing at runtime we may end up with cyclesame with LP.
- To fix the issue of cycles pick the table size to be a prime number and also compute the value of delta.
  - delta = H_2(k) mod N
  - If delta = 0, then we are guaranteed to be stuck in a cycle, so when this happens just set delta = 1.
- Notice that 1 <= delta <= N and GCD(delta, N) = 1 since `N is prime`. Hence, with these conditions we know that modulo N the sequence *H_1(k), H_1(k) + delta, H_1(k) + 2delta,...* is certain to have order N.

##### Constructing H_2(k)

- Suppose the key k has type T. Whenever we want to use double hashing as a collision resolution method we need to fabricate a new function H_2(k) that knows how to hash keys of type T.
- It would be nice to have a systematic way to able to effectively produce a new hash function every time we need one.
- Frequently the hash functions selected to compose H_2(k) are picked from a pool of hash functions called `universal hash functions` which generally operate on one fundamental data type.
