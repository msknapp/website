---
title: Quick Sort
draft: false
weight: 4
summary: How the quick sort works
lastmod: 2023-03-19
---

The quick sort runs faster because it is efficient, and it runs in-place (unlike the merge sort).

![quick-sort](/images/quick-sort.png)

It works by:
* choosing a value to be the "pivot" point.
* In the remaining array, it swaps values between left and right cursors until
  * all values up to and including the left cursor are less than the pivot.
  * all values on or after the right cursor are greater than or equal to the pivot.
* Then the pivot is swapped into the middle, trading places with the right pivot.
* The process gets repeated on the sub-arrays to the right and left of the pivot.

The quick sort is:
* not stable, equal values may be swapped.  This can be worked around by adjusting your
  comparison logic.
* in-place, it uses no significant amount of extra memory.
* Efficient, having O(n * log(n)) performance.
* possible to parallelize.  At step 3 in the diagram, separate coroutines or threads 
  could be created if the sub-arrays are still quite large.
* works with non-numerical types, it uses a comparator.

The quick sort algorithm was made to handle sorting large arrays, but in practice, the
merge sort algorithm tends to be used for larger arrays instead.

Since the quick sort uses recursion, it is possible to have a stack overflow in your program, 
unless the program is written to work around recursion.