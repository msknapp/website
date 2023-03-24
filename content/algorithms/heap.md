---
title: Heap Datastructure
weight: 10
draft: false
summary: The heap structure explained.
lastmod: 2023-03-23
nextlink: /tech/
---

The heap datastructure is a clever design that makes an array behave like a binary 
tree, and the root of the tree will always be the min or max value.

For any index `i` in the heap:
* its left child will be at `2*i`
* its right child will be at `2*i + 1`
* its children will always be greater or lesser than the current node.

