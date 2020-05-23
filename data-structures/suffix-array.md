# Suffix Array

- References
  - [WilliamFiset] Suffix Array introduction, https://youtu.be/zqKlL3ZpTqs
  - [WilliamFiset] Longest Common Prefix (LCP) array, https://youtu.be/53VIWj8ksyI

## What is a suffix?

- A `suffix` is a substring at the end of a string of characters. For our purposes suffixes are non empty.
- Example) HORSE
  - E, SE, RSE, ORSE, HORSE => are all suffixes

## What is a Suffix Array(SA)

- A `suffix array` is an array which contains all the *sorted* suffixes of a string.
- The actual `suffix array` is the array of sorted indices. This provides a compressed representation of the sorted suffixes without actually needing to store the suffixes.
- The suffix array provides a space efficient alternative to a `suffix tree` which itself is a compressed version of a `trie`.
  - suffix arrays can do everything suffix trees can, with some additional information such as a Longest Common Prefix (LCP) array.

## What is a Longest Common Prefix(LCP) array

- Example) 'ABABBAB'

| Sorted Index | LCP value | Suffix  |
| ------------ | --------- | ------- |
| 5            | 0         | AB      |
| 0            | 2         | ABABBAB |
| 2            | 2         | ABBAB   |
| 6            | 0         | B       |
| 4            | 1         | BAB     |
| 1            | 3         | BABBAB  |
| 3            | 1         | BBAB    |

- The LCP array is an array in which every index tracks how *many characters two sorted adjacent suffixes have in common*.
- By convention, LCP[0] is undefined, but for most purposes it's fine to set it to zero.
- There exists many methods for efficiently constructing the LCP array in O(nlog(n)) and O(n).