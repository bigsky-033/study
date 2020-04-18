# Time Complexity

## Time Complexity for For Loop

- Reference: Abdul Bari's Youtube
  - https://youtu.be/9TlHvipP5yA
  - https://youtu.be/9SgLBjXqwd4

### 1) Simple loop, i is plused

```
for (i = 1; i < n; i++) {     --- (n + 1)
    stmt;                     ---  (n)
}
```

- => O(n)
- If (i++) is changed to (i = i+k)
  - => f(n) =  n / k
  - => O(n)

### 2) Nested loop

```
for (i = 0; i < n; i++) {         --- (n + 1)
    for (j = 0; j < n; j++) {     --- (n) * (n + 1)
        stmt;                     --- (n) * (n)
    }
}
```

- => O(n^2)

### 3) Nested loop, inner loop's condition dependends on outer loop

```
for (i = 0; i < n; i++) { 
    for (j = 0; j < i; j++) {
        stmt;
    }
}
```

- Trying to count the number of execution of `stmt`

| i   | j       | # stmt |
| --- | ------- | ------ |
| 0   | 0       | 0      |
| 1   | 0       | 1      |
| 2   | 0, 1    | 2      |
| 3   | 0, 1, 2 | 3      |
| ... | ...     | ...    |
| n   | ...     | n      |

- Number of execution of stmt
  - => 1 + 2 + 3 + ... + n
  - => n(n + 1) / 2
  - => O(n^2)

### 4) Loop's terminate condition depends on variable outside of loop

```
p=0;
for (i = 1; p <= n; i++) {
   p = p + i;
}
```

- Trying to find relationship between i & p

| i   | p               |
| --- | --------------- |
| 1   | 0 + 1 = 1       |
| 2   | 1 + 2 = 3       |
| 3   | 3 + 3 = 6       |
| 4   | 4 + 4 = 8       |
| ... | ...             |
| k   | 1 + 2 + ... + k |

- See when it is terminates
  - Assume p > n
  - => Since, p = k(k + 1) / 2
  - => k(k + 1) / 2 > n
  - => (Roughly) k^2 > n
  - => k > root(n)

### 5) Simple loop, i is multiplied

```
for (i = 1; i < n; i = i * 2) {
    stmt;
}
```

- Trying to find how i increases

| i             |
| ------------- |
| 1 = 2^0       |
| 1 * 2 = 2^1   |
| 2^1 * 2 = 2^2 |
| 2^2 * 2 = 2^3 |
| ...           |
| 2 ^ k         |

- See when it is terminates
  - Assume i >= n
  - => Since, i = 2^k
  - => 2^k >= n
  - => (Roughly) 2^k = n
  - => k = log2_n
  - => O(log2_n)
- (Important) When you need to calculate rough number of execution by n and result should be interger value(not float or double), you need to ceil(log2_n)

### 6) Simple loop, i is divided by 2

```
for (i = n; i >= 1; i = i/2) {
    stmt;
}
```

- Trying to find how i increases

| i        |
| -------- |
| n        |
| n / 2    |
| n / 2^2  |
| ...      |
| n / 2 ^k |

- See when it is terminates
  - Assume, i < 1
  - => Since, n / 2^k < 1
  - => (Roughly) n / 2^k = 1
  - => 2^k = n
  - => k = log2_n
  - => O(log2_n)

### 7) Simple loop, check conditon is expression

```
for (i = 0; i * i < n; i++) {
    stmt;
}
```

- See when it is terminates
  - Assume, i * i >= n
  - => (Roughly) i^2 = n
  - i = root(n)

### 8) Two independent loops with same time complexity

```
for (i = 0; i < n; i++) {
    stmt;
}

for (j = 0; j < n; j++) {
    stmt;
}
```

- Each loop's time complexity is O(n)
  - => O(n) + O(n)
  - => O(n)

### 9) Two dependent loops

```
p = 0
for (i = 1; i < n; i = i * 2) {
    p++;
}

for (j = 1; j < p; j = j * 2) {
    stmt;
}
```

- Number of execution of stmt
  - => it will executes log2_(p)
  - => What is p? p is determined by above loop
  - => p is log2_n
  - => O(log2_(log2_n))

### 10) Nested loop with different condition

```
for (i = 0; i < n; i++) {               --- (n)
    for (j = 1; j < n; j = j * 2) {     --- (n) * (log2_n)
        stmt;                           ---  (n) * (log2_n)
    }
}
```

- O(nlog2_n)

### Summary

```
- for (i = 0; i < n; i++)          --- O(n)
- for (i = 0; i < n; i = i+2)      --- n/2 * O(n) => O(n)
- for (i = n; i > 1; i--)          --- O(n)
- for (i = 1; i < n; i = i*2)      --- O(log2_n)
- for (i = 1; i < n; i = i * 3)    --- O(log3_n)
- for (i = n; i > 1; i = i/2)      --- O(log2_n)
```

## Time Complexity for While & If

- Reference: Abdul Bari's Youtube
  - https://youtu.be/p1EnSvS3urU

### 1) Simple while loop, i is plused

```
i = 0;
while (i < n) {    --- (n + 1)
  stmt;            --- (n)
  i++;             --- (n)
}
```

- f(n) = 3n + 2
- => O(n)

### 2) Simple while loop, i is multiplied

```
a = 1;
while (a < b) {
  stmt;
  a = a * 2;
}
```

- Trying to find how a increases

| a             |
| ------------- |
| 1             |
| 1 * 2 = 2     |
| 2 * 2 = 2^2   |
| 2^2 * 2 = 2^3 |
| ...           |
| 2^k           |

- See when it is terminates
  - Assume, a >= b
  - => Since a = 2^k
  - => 2^k >= b
  - => (Roughly) 2^k = b
  - => k = log2_b
  - => O(log2_b)

### 3) Simple while loop, i is divided

```
i = n;
while (i > 1) {
  stmt;
  i = i / 2;
}
```

- O(log2_n)

### 4) Complex while loop

```
i = 1;
k = 1;
while (k < n) {
  stmt;
  k = k + i;
  i++;
}

// above while loop is same with below for loop

for (k = 1, i = 1; k < n; i++) {
  stmt;
  k = k + i;
}
```

- Trying to find how i & k increases

| i   | k                                |
| --- | -------------------------------- |
| 1   | 1 + 1                            |
| 2   | 1 + 1 + 2                        |
| 3   | 1 + 1 + 2 + 3                    |
| 4   | 1 + 1 + 2 + 3 + 4                |
| ... | ...                              |
| m   | 1 + 1 + ... m = 1 + m(m + 1) / 2 |

- See when it is terminates
  - Assume k >= n
  - => Since, k = (Roughly) m(m + 1) / 2
  - => m(m + 1) / 2 >= n
  - => (Roughly) m (m + 1) / 2 = n
  - => (Roughly) m^2 = n
  - => m = root(n)
  - => O(root(n))

### 5) Finding GCD of two numbers m and n

```
while (m != n) {
  if (m > n) {
    m = m - n;
  } else {
    n = n - m;
  }
}

// above while loop is same with below for loop
for (; m != n; ) {
  if (m > n) {
    m = m - n;
  } else {
    n = n - m;
  }  
}
```

- Trying to find how m & n changes

| m   | n   |
| --- | --- |
| 6   | 3   |
| 3   | 3   |

=> 2 time

| m   | n   |
| --- | --- |
| 5   | 5   |

=> 1 time (minimum)

| m   | n   |
| --- | --- |
| 16  | 2   |
| 14  | 2   |
| 10  | 2   |
| 8   | 2   |
| 6   | 2   |
| 4   | 2   |
| 2   | 2   |

=> 7 time
=> Almost (16 / 2) times
=> (m / 2) times
=> O(n) (maximum)

### 6 - 1) What happens If is exists

```
Function(n):
    if (n < 5) {
      printf(n);                    --- (1)
    } else {
      for (i = 0; i < n; i++) {
        print(i);                   --- (n)
      }
    }
```

- Time complexity dependes on conditional statement
- If condition is true
  - O(1): best
- If conditon is false
  - O(n): worst

### 6 - 2) What happens If is exists

```
Function(n):
    if (n > 5) {
      for (i = 0; i < n; i++) {
        print(i);                   --- (n)
      }
    }
```

- Just O(n)

## Time functions

- Reference: Abdul Bari's Youtube
  - https://youtu.be/w7t4_JUUTeg
  - https://youtu.be/5v-tKX2uRAk

### Types of time functions

| function |             |
| -------- | ----------- |
| O(1)     | constant    |
| O(log_n) | logarithmic |
| O(n)     | linear      |
| O(n^2)   | quadratic   |
| O(n^3)   | cubic       |
| O(m^n)   | exponential |

### Compare time functinos

1 < log_n < root(n) < n < nlogn < n^2 < n^3 < ... < 2^n < 3^n < n^n

## Asymptotic Notations

- References: Abdul Bari's Youtube
  - https://youtu.be/A03oI0znAoc
  - https://youtu.be/NI4OKSvGAgM

- Big O: Uppder bound
- Big Omega: Lower bound
- Theta: Average bound

### Big O

- Definition
  > The function f(n) = O(g(n)), if and only if there exists positive constant c and n_0 such that f(n) <= c*g(n) for all n >= n_0

- When you write Big O, try to use the closest function

### Big Omega

- Definition
  > The function f(n) = Omega(g(n)), if and only if there exists positive constant c and n_0 such that f(n) >= c*g(n) for all n >= n_0

- When you write Big Omega, try to use the closest function

### Theta

- Definition
  > The functionf f(n) = Theta(g(n)), if and only if there exists positive contants c-1, c_2 and n_0, such that c_1 * g(n) <= f(n) <= c_2 * g(n)

- When ever you have any function and if you are able to represent it as theta annotation, that's better. If it i not possible, then consider Big O or Big Omega
- Always theta is prefable if can find theta for any functions, this is better

### Don't mix this concepts with best case & worst case

- Big O is worst case? no
- Big Omega is best case? no
- We can use any notation for best case, any notation for worst case

### Properteis of Asymptotic Notations

- General Properties
  - If f(n) is O(g(n)), then a * f(n) is also O(g(n))
  - If f(n) is Omega(g(n)), then a * f(n) is also Omega(g(n))
  - If f(n) is Theta(g(n)), then a * f(n) is also Theta(g(n))
- Reflexive
  - If f(n) is given, then f(n) is O(f(n))
- Transitive
  - If f(n) is O(g(n)) and g(n) is O(h(n)), then f(n) = O(h(n))
- Symmetrics
  - If f(n) is Theta(g(n)), then g(n) is Theta(f(n))
    - Only for Theta
- Transpose Symmetric
  - If f(n) = O(g(n)), then g(n) is Omega(f(n))
- If f(n) is O(g(n)) and f(n) is Omega(g(n)), then f(n) = Theta(g(n))
- If f(n) is O(g(n)) and d(n) is O(e(n)), then f(n) + d(n) = O(max(g(n), e(n)))
- If f(n) is O(g(n)) and d(n) is O(e(n)), then f(n) * d(n) = O(g(n) * e(n))
