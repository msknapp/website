---
title: Stateless
weight: 10
draft: false
lastmod: 2023-03-23
summary: Your applications should ideally be stateless.
---

In a highly available system, you should assume that replicas could be abruptly stopped at any time.
This is one of the 12-factor app principles.

When an application is stateless:
* It can be scaled up or down at will to meet traffic demands.
* It can be bounced at will if it is mis-behaving.
* There is no need to recover its data if it crashes.
* The load can easily be balanced among its replicas in a round-robin
  manner.  Sticky sessions are not needed.

This is not to say that applications are truly stateless, they may hold some cache of information.
That is acceptable because the application does not depend on that cache, it is simplay a 
performance optimization.

When applications need data, they get that from a database or data-store instead of by holding
and replicating the data themselves.  The only exception is if you are writing a data-store application.
Separating services from data storage permits them to be updated and managed independently.