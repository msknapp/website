---
title: Substitution
draft: false
weight: 6
lastmod: 2023-03-24
summary: Your code should be loosely coupled using abstractions, interfaces, and similar substitution principles.
---
 
This page covers the "L" and "D" parts of the often referenced "SOLID" programming principles.  These
principles are all so similar, I keep them under one page:
* The Liskov Substitution Principle
  * Polymorphism
* Dependency Inversion
* Low Coupling

# Liskov Substitution

In your code, it should be possible to substitute one sub-type of an object for another.
Liskov is often mentioned here because his name fits into the SOLID acronym.

If you write a function, the argument it takes should be the most abstract type that it can
work with.  If possible, the argument type should be an interface that contains the bare
minimum set of functions that the function requires.

This makes it much easier to change implementations in your code later.  You end up having
code that is composable like lego blocks, so it is agile enough to change as needed.

## Polymorphism

This is simply another term for Liskov substitution.  It's used more in java though, like 
for example, it should be possible to swap one implementation of an interface for another
in a function.

# Low Coupling

Low coupling means that a change in one component is unlikely to break other components that 
depend on it.  A component could be a class in the code, or a microservice in your cluster.
You de-couple components by minimizing the contract between them.  That contract could be 
an interface in code, or a Swagger API for microservices.

# Dependency Inversion

Ideally, classes in your code should not depend on specific implementations of
a service, but instead should depend on an interface or abstract definition
of that service.