---
title: Class Stereotypes
weight: 20
draft: false
lastmod: 2023-03-24
summary: Design your classes to follow effective patterns or stereotypes.
---

After writing enough code, and dealing with performance issues and bugs, engineers
gradually adopt certain patterns of writing code, and they might not even overtly
realize they are doing this.  These patterns work, which is why we use them.

# LifeCycles

The pattern used for a class often depends greatly on its life-cycle.  If a class
will live for the entire life of the process, you might design it differently than
a class that is created and destroyed with every single request.

# Domain

Domain classes simply hold the state of something.  They usually have little 
business logic in them, or very trivial business logic.  Making them immutable
could result in a safer application with fewer bugs, or it could just 
complicate things and reduce performance.  It depends on the situation.

## Fluent

A fluent class is one that is mutable, and its setter methods all return the 
current instance, so that they can be chained together.  Unlike the builder pattern,
these do not build separate instances.

Their advantages are:
* They are very easily understood and worked with by fellow developers.
* They can easily have very many fields that need to be set.

They have the usual disadvantages of mutable objects, thread safety concerns, 
garbage collection, performance impacts, etc.
The pros and cons should be evaluated for the current problem, sometimes fluent
classes are great, sometimes not.

## Builder

A builder is a class that is mutable, and is responsible for creating instances
of another class.  Usually what it builds is immutable once made.  Builders
tend to be used when the object being built has many parameters.  Somewhere 
between five and seven parameters is when a builder might be useful.  If 
the object being created has fewer parameters, then using an immutable constructor
is a better idea.  Usually the methods of the builder are designed for 
chaining.  So the return value of the setters is the builder itself.

Some advantages of builders:
* You can validate the inputs before the product is built.
  Hence a builder can provide some consistency guarantees for its output.
* You can easily assign default values.
* When an engineer reads how the builder is used, it is intuitive and easily understood.
  This is because of the fluent style chaining of its setters.

Some disadvantages of builders:
* They add to code complexity.
* They also need to be tested.
* It takes a bit more time to write them.
* They result in creating more objects than necessary, which in turn leads to more
  garbage build-up on the heap, and that can degrade the application's performance.

Builders tend to be created per request.  Values are unlikely to be re-used.

## Factory

A factory is similar to a builder, but with a few exceptions:
* It tends to be long lived itself.  It is created at bootstrap, not per request.
* Its values tend to change rarely or never.
* It is more likely to have static or global variables.

The way I like to remember this is that real-world factories tend to stay in 
operation longer than the employees (the builders) working in them.

# Service

Classes that are services tend to be created when the process boots, and live
for as long as the process is running.  Their fields usually do not change throughout
the life of the application.  The one exception is if their values are tied to 
a configuration that could be updated from outside the application.

For extra safety, they *might* be made immutable, but it is not strictly necessary.

# Configuration

An effective coding pattern is to have a set of mutable classes that represent the application's
configuration.  This means a separate configuration file can be written in a common
format like yaml, json, toml, or a simple properties file.  Then, at boot time, the
file can be unmarshaled directly into these configuration classes.

There are also libraries that help with this by adding logic to:
* Validate the configuration, requiring values to be consistent.
* Automatically re-load the configuration on a schedule.
* Apply default values to missing fields in the configuration.
* Merge the configuration from different sources, like environment variables,
  multiple config files, and command line arguments.

These configuration classes tend to have all public fields and are mutable.

Generally your service classes should not have direct dependencies or knowledge of 
these configuration classes, because they are mutable.  Also, the configuration
classes are likely to change more often.

Instead, the service classes could have references to the configuration variables
that they need or depend upon.  If the reference is not atomic, that field could
be changed to a "supplier" type of class or function.

# Bootstrap

A bootstrap class or package is responsible for wiring together the high level
service that represents the application.  It usually has access to all of the 
configuration and service class packages.  Frameworks like Spring or J2EE serve
the "bootstrapping" function too.  They support "dependency injection", which 
means it is easy to change the application's behavior without re-building it.

**My personal opinion:** Spring adds more complexity and problems than writing your
own bootstrapper class.  If you write your own bootstrapper, you will have far 
more control over how things are wired, and it will be much easier to troubleshoot
and fix later.  I have only tried Spring core though, and it was 
many years ago, maybe Spring Boot is better, I don't know.

# Package Dependency Tree

In languages like Golang, there must not be a circular dependency in packages.
In other words, if package A depends on something in package B, then nothing in 
package B can depend on anything in package A.

For this reason it's best to establish a convention for dependency order.

* Packages for domain objects should have few or no dependencies on higher level logic.
* Business Logic packages may depend on domain objects.
* Service classes and packages may depend on domain or business logic packages.
* configuration classes and packages may depend on domain or business logic packages.
* Bootstrap packages may depend on all of the above.

