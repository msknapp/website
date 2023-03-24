---
title: Defensive Programming
description: Making your code defensive.
summary: Carefully consider the trade-offs of defensive programming patterns.
draft: false
weight: 3
---

Defensive programming is a design pattern where each component in software acts very defensive in nature.
For example:
* Fields become private.
* Classes are often made immutable.
* Functions validate all of their inputs very strictly.
* Code "fails fast" if a problem is encountered.

It might also lead to adopting certain code patterns such as:
* The Singleton.
* The Builder.
* The Factory class.

I would like to argue, from my experience, that:
* Defensive programming is more or less useful depending on the programming language.
* Your number one defense against bad code should be TESTS!  Not defensive coding patterns.

# Immutability

An immutable class is one whose fields cannot be updated once it has been created.  They 
tend to be far safer to use in code, because they basically cannot be corrupted.  Also they
are naturally thread safe, multiple threads can refer to them because none of them can 
alter the object.  Programs that use immutable objects tend to have fewer bugs.  You 
might even say they need less testing.

Immutability adds so much safety to programs that some languages make everything immutable (Scala),
or make it a challenge to let things be mutable (Rust).

The downside of immutability is it tends to complicate some parts of the code, and it 
often results in more objects being created and destroyed than necessary.  That means
garbage cleanup needs to run more often, which can degrade the application's performance.

# Fluent Pattern

The fluent pattern is not a defensive programming pattern, it is included here to draw its 
similarities to others.  Fluent classes are written in a way that you can chain together 
functions that set fields on the class.  The class is mutable.

# Builder Pattern

The builder pattern involves having one class that is mutable, and its purpose is to build 
a different immutable class.  Usually the builder has a no-argument constructor, fluent 
methods to set values, and a build method.

There are two ways to write a builder:
* The builder has duplicated fields for all or most of what the final class has.
  This is less likely to produce a class that can be mutated accidentally, but it also 
  makes it harder to maintain because fields have been duplicated.  This method is thread safe.
* The builder holds its own private instance of the final class, which only it can modify.
  The problem with this approach is you must be careful to not permit updating that 
  instance once it has been built.  The builder is suddently not thread safe.

In practice, builders may or may not be re-used.  The downsides of using builders are:
* They add more complexity to your code, and hence more opportunities for bugs to creep in,
  and a tougher time for co-workers to understand it.
* They must also be unit tested.
* They create more objects, which can add strain to garbage collectors, and hence it 
  requires more memory and can slow down performance.

# Factory Pattern

The factory pattern is very similar to the builder pattern, but with a few differences:
* Factorys tend to be re-used more than builders, and are longer lived.  Builders tend
  to be ephemeral.  Consequently, factory's lend themselves to application bootstrapping, 
  while builders are more commonly used at request time.
* Once the factory is configured, it usually does not change its values.
* Typically the factory's configuration is held in static variables (in java).

Factory's also have similar downsides as builders:
* Added complexity
* More required unit testing
* If the factory uses globals (static variables), that is more risky and unsafe.

If you are considering a pattern listed above, you should weigh the pros and cons, and consider
the problem you have.

# Immutable Constructors

It is worth calling out here, that classes can be immutable without being produced by 
a builder or factory.  They can simply have a constructor that sets their fields, and
no other methods are able to mutate them.

The only downside is that some immutable objects have many fields.  Somewhere between 
five and seven arguments to a method, it starts to get too confusing and difficult 
to maintain.  After seven arguments, a fluent pattern, a builder, or a factory pattern
becomes easier to understand.

You might prefer immutability when:
* The object is used by many parallel processes or threads.
* The code is getting very complicated and you want to reduce risk.
* The lifecycle of the object is quite long, spanning many requests.

You might prefer mutability when:
* The object lifecycle is short lived and not shared by many functions.
* Memory and garbage collection are becoming a concern.
* You need your code to perform faster.
* You are willing to write more unit tests to be certain everything is safe.

