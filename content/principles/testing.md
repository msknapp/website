---
title: Software Testing
draft: false
weight: 8
summary: Your software should have ample automated tests across the board.
description: The importance of software testing.
---

When it comes to having safe code, I'm going to make this assertion:

```
Code testing is by far the most important and effective means of producing robust code.
```

It is more important than using defensive patterns like builders, factories, etc.  It 
is more important than having documented code, loosely coupled code, logging, metrics,
and even alerts.  Professional software engineers need to place a heavy priority on
getting tests automated.

There are several types of testing for software, which are discussed below.

# Unit Tests

These are by far the easiest tests to write, they go in your code, and usually their
scope is kept limited.  They typically just test one class, function, or method (in line with the 
single responsibility principle).  Sometimes they may test the combined behavior of 
a collection of classes, or even the highest level service in the application.

Unit tests should be possible to run from any developers machine or from CICD servers
without needing to start external services like a database, server, or infrastructure.

# Benchmark Tests

These are similar to unit tests in that they run on local machines.  They are usually 
not run during CICD pipelines, they are just used by developers to guide performance
improvements.

Depending on the language, benchmark tests might be unreliable.  For instance, in java,
the "just-in-time" compiler can make performance appear to be very slow at first, but
much faster later.  So your benchmark test should run the function under evaluation
many times before it draws any conclusions.

# Integration Tests

An integration test runs at a higher level, and covers more things.  It starts to pull in 
external services.  For example, integration tests may involve running your code with a 
real database, or with real infrastructure.  It might involve running a Kafka server,
Apache, or creating infrastructure.  It might involve running docker containers or 
virtual machines.

Integration tests start to cover configuration as well as the code.  However, integration 
tests do not go so far as to run in a real runtime environment, nor to cover everything.

Integration tests might not be easily run from any person's machine, or from CICD.  Your 
team should consider it important though, to have them run routinely, either by a 
cron schedule or after a deployment.

# System Tests

These are the highest level tests, they run in a real cluster or runtime environment, 
with real databases and everything.  They could also be called end-to-end testing.
The goal here is to cover every aspect of the system, and to simulate interacting 
with it the way a real user would.  It is critical to maintain system tests and 
to automate running them either by a schedule or post release.

System tests are still primarily concerned with functionality, does the system behave
as expected.

# Performance Tests

These could also be called stress tests.  They typically run in a non-prod environment and 
work by making many requests to the service in parallel, for an extended period of time.
Metrics are gathered as it runs, for things like response time, throughput, and the 
count of failures.  JMeter is the typical test software.

# User Acceptance Tests

These tests differ from the previous ones in a few ways:
* they are often written by product owners instead of engineers.
* The emphasis is more on if the system meets the product requirements, over how it behaves.
* It evaluates if the service really is something the users can utilize.

These tend to not be automated, and are usually just run manually after major milestones.