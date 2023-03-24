---
title: Code Quality
draft: false
weight: 2
lastmod: 2023-03-22
summary: Code should be automatically evaluated by certain quality metrics.
---

Source code needs to be judged, but people can be highly opinionated and stubborn.
That is why it's better to lean on code quality tools like Sonar, PMD, CPD, dupl, etc.

# Code Duplication

Scans should check for blocks of code that appear to be highly similar to each other, like
they were copied and pasted.  These blocks can usually be combined into a function or 
module with arguments.  Reducing duplication means the code is simpler, and less likely
to hold bugs.  Also if code is duplicated, a bug may be fixed in one place but not 
the other.  In java, the CPD tool is available to check for this.  In golang there is dupl.

# cyclomatic complexity

This is a metric that evaluates if your code is too complicated or not.  If it gets too
high, you should consider splitting up or re-factoring your code to be simpler.

# test coverage

Probably the most important metric to review is code test coverage.  It measures what 
percent of code lines are reached by a unit test.  A suggested minimum is 80%.

# Syntax, formatting

Code should be "linted" or "vetted" during the CICD pipeline, before it is compiled
or built.

# Code smells

Sonar can detect "code smells" which is a generic term for something in the code that
is *not necessarily* a problem, but is usually not a good idea.
