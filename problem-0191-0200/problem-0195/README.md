This problem was asked by Google.

Let `A` be an `N` by `M` matrix in which every row and every column is sorted.

Given `i`<sub>`1`</sub>, `j`<sub>`1`</sub>, `i`<sub>`2`</sub>, and `j`<sub>`2`</sub>, compute the number of elements of `A` smaller than `A[i`<sub>`1`</sub>`, j`<sub>`1`</sub>`]` and larger than `A[i`<sub>`2`</sub>`, j`<sub>`2`</sub>`]`.

For example, given the following matrix:

```
[[1, 3, 7, 10, 15, 20],
 [2, 6, 9, 14, 22, 25],
 [3, 8, 10, 15, 25, 30],
 [10, 11, 12, 23, 30, 35],
 [20, 25, 30, 35, 40, 45]]
```

And `i`<sub>`1`</sub>`= 1`, `j`<sub>`1`</sub>`= 1`, `i`<sub>`2`</sub>`= 3`, `j`<sub>`2`</sub>`= 3`, return `15` as there are `15` numbers in the matrix smaller than `6` or greater than `23`.
