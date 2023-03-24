---
title: Open Closed Principle
draft: false
weight: 5
summary: Design software classes and components to be easily extended but rarely modified.
lastmod: 2023-03-17
---

There is a famous saying:

```
If it ain't broke, don't fix it!
```

The open/closed principle is similar to the motto above.  Once you have a working 
component of software, it is preferable that it is not changed directly.  Instead,
it should be possible to change its behavior without changing its code.  It's similar
to another software design principle:

```
Favor composition over inheritance.
```

Your software components should be open to being extended, but closed to being modified.
There are several ways they can be extended.  The most common or obvious is to have
a constructor argument that accepts a service or interface.  Then, that implementation
can easily be swapped for one with more new capabilities.

Another way to extend software components is to support expressions or scripts within them.
For example, the code could accept a Lua script which defines some custom behavior.