---
title: Interface Segregation Principle
draft: false
weight: 6
summary: Function arguments should be minimal interfaces.
description: explaining the interface segregation principle
---

The interfaces in your code should be minimal.  If needed, there can be a hierarchy of interfaces,
each adding a few more methods.  Functions should declare argument types that use the most
minimal and abstract interface possible for their needs.  The advantage here is that the 
calling code has a lower burden to meet if it wants to use that function.

# Software Component Cardinality

Through my programming career, one trend in my code has become obvious: I keep splitting
services and components up into smaller ones.  I have increasingly more very small things.

It sounds like it defies the KISS principle, how can you have more things and call that 
simpler?  It's because:
* Each of these small components is very simple in its own right.
* These small components change very rarely.
* It is so easy for me to create tests for each small and simple component.

Having monolithic services has always proven to be unmaintainable in my experience.
Developer teams typically decide to write one server, and throw everything into it.
Whenever a new feature is required, they furiously code it into the same application.

Why not have a new application?  Well then they would need to figure out how to setup 
all of the infrastructure for that too, and to manage network connections for it.
Most programmers are in their comfort zone writing their applications.  They are 
less comfortable setting up all the infrastructure around it, so they don't.

