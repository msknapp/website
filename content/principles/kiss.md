---
title: KISS and YAGNI
description: How to keep your code simple
summary: Prefer to keep your code as simple as possible.
weight: 1
draft: false
prevlink: ..
---

Most people know that KISS stands for "keep it simple stupid", but might not think about applying this 
to the code they write.  As your code gets increasingly complicated, the chances of having bugs in it
grow exponentially.  It becomes exceedingly difficult to test all of the edge cases.

Ambitious developers take satisfaction from code that is powerful, and this often leads them to 
over complicate things.

For example:
* I want to make this code as fast as possible!  (but it is fast enough to meet 
  requirements as it is).
* I want my code to also support feature X (but nobody asked for it, nor needs it)
* I want to add methods that make this code easy and convenient to use.

# YAGNI

This ties into another software motto, "You ain't gonna need it" (YAGNI).
Sometimes, less is more.  Less code is more maintainable, understandable, and powerful.
Less code is more nights spent sleeping, instead of being paged.

Consider the following methods to keep code simpler:
* Deliberately choose to not support a feature or function if there is little need for it, 
  or an easy workaround.
* Apply default values at the start of a process, so code can safely assume they are set.
* Re-use libraries that already work and are proven.

