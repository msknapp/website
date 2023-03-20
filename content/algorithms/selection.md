---
title: Selection Sort
weight: 3
draft: false
summary: How the selection sort works
lastmod: 2023-03-19
---

The selection sort is also relatively intuitive, but may out-perform
bubble and insertion sorts because it makes fewer swaps on average.

![selection-sort](/images/selection-sort.png)

It works by:
* maintaining a sorted section on the left side of the array.
* the left side will always contain the smallest items in the array
  (assuming an ascending sort).
* A cursor iterates from left to right.
* In each iteration, the algorithm "selects" the next smallest item
  in the remaining unsorted array.  Then it swaps that with the cursor.

The selection sort is:
* stable, equal elements maintain their relative order
* in-place, it uses no significant extra memory to run.
* inefficient, having O(n^2) performance.
* not possible to run in parallel.
* works with non-numerical types, it uses a comparator.

The main advantage of the selection sort is it requires fewer swaps
to run than insertion sort, so it can out-perform that.