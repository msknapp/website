---
title: Premature Optimization
draft: false
weight: 39
lastmod: 2023-03-29
summary: Avoid optimizing code before proving it is necessary.
---

There is an old saying in software engineering that "premature optimization is the root of all evil."
It's a bit dramatic isn't it?

The point is that our efforts to optimize code often:
* Greatly complicate the code.
* Don't actually improve performance.
* Are based on bad assumptions.
* Are not even needed.

This can make the code difficult to maintain or add features to.  We generally want the 
code to be as intuitive as possible, until that cannot meet the performance requirements.

# Measuring Performance

When optimizing performance of applications, the decisions should be guided by measurable facts.
You should be able to point at benchmark test results, profiling data, or request
traces, to prove that performance is an issue.

When you improve the code's performance, profiling or benchmark results should be 
used to prove that the change you made actually has improved performance.  Your code
changes should be guided by metrics, not hand-waving or opinions.

Sometimes slow algorithms in code are acceptable if they are rarely called, use small
amounts of data, or are relatively inconsequential to the application's purpose.

Hence the old saying: "If it's not broken, then don't fix it."