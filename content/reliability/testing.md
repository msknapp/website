---
title: Testing
draft: false
weight: 6
lastmod: 2023-03-22
summary: There should be ample automated tests of your system that run routinely.
---

The number one best defense against failures and incidents in your cluster is ample
testing.  This comes in a variety of forms.

# Unit tests

Any code that is mission critical and running in a production environment should have
ample unit tests, which cover low-level functions in the code.  Typically teams should
set a minimum required amount of coverage, around 80%.  The language chosen may 
impact this, for example, Python code could require higher coverage because it 
doesn't require errors be checked.

# Integration Tests

These are higher level, involving at least one external system in their tests.  Ideally
these should be run automatically, like at the end of a CICD stage, or on a cron 
schedule.

# System Tests

These run in a real live cluster, as similar to production as possible.  They involve all
of the services, and are tested the same way a user would interact with the system.
These can also be called "end to end" (e2e) tests.

I do not recommend going past the beta stage of a product unless system tests exist and
are automatically run.  Lacking automated system tests is bound to result in production incidents.

# Performance Tests

These could be run by JMeter for example.  They confirm your system can auto-scale, and can
handle the higher traffic levels it might encounter.  Pair this with tracing and you can
identify bottlenecks in your service easily.

# Chaos Testing

We are trying to build a system that is highly available, but this assumes that
certain failure modes *happen*.  In practice, they are rare, meaning you don't
learn about them until it happens in production and is a major incident, costing 
the company millions of dollars.  That is why you need something to prove the 
system's fault tolerance.

Chaos testing or engineering works by randomly breaking certain aspects,
components, or connections of your service.  For example, it might terminate a node, virtual 
machine, or a single replica in your cluster.  It might block connections to some
external service for a period of time.

If your cluster really is fault tolerant, it should just recover by itself with no intervention.  If 
not, at least you know about it now, before it really impacts anybody.

# Migration Testing

This is not really a "thing" you would read or hear about if you searched for software
testing methods.  I want to call out, that changing your APIs and/or data involves 
risk.  Often times, a developer is only thinking about the *end state* of their
software, and not how to get there in the first place.  In other words, they are 
not thinking about backwards compatibility.

For example, you may have two services and one table they both use.  A developer
may decide to add a new field in the table, and update both of the services to use 
that.  Whenever a new record is created, that new field will be set.

What the developer didn't think about is:
* What about all the records that already exist?
* What if one of the services needs to be rolled back?
* What if the database change needs to be rolled back?

So they should make a procedure for how to update all existing records, and how to 
make the services backwards compatible.  They should each be able to function 
if their peers are rolled back, this is backwards compatibility.  You should
never be required to deploy data updates and multiple service updates all at once,
nor to roll them back all together.

What's often neglected is that, *these procedures need to be tested!*  and 
critically, *that includes roll back procedures!* if they will be done manually, or are
one-time procedures.  The services each need to be tested with different 
versions of their peers too.

This is what I am calling "migration testing".