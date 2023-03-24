---
title: Binary Search
weight: 9
draft: false
summary: How to use the binary search algorithm.
---

A binary search can be used to find the index of an item in an array that is
already sorted.  It works by repeatedly cutting the array in half until
the item is found.


```
package srt

func BinarySearch[T comparable](x []T, y T, c Comparer[T]) int {
	max := len(x) - 1
	min := 0
	index := max / 2
	for {
		if c.Equals(x[min], y) {
			return min
		}
		if c.Equals(x[max], y) {
			return max
		}
		if c.Greater(x[min], y) || c.Less(x[max], y) {
			return -1
		}
		if c.Equals(x[index], y) {
			return index
		}
		if max-min <= 1 {
			return -1
		}
		if c.Less(x[index], y) {
			min = index
		} else {
			max = index
		}
		if max == min {
			return -1
		}
		index = min + (max-min)/2
	}
}

type Comparer[T any] interface {
	Equals(x T, y T) bool
	Less(x T, y T) bool
	Greater(x T, y T) bool
}
```