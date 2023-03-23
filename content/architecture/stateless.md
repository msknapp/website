---
title: Stateless
weight: 10
draft: false
lastmod: 2023-03-23
summary: Your applications should ideally be stateless.
---

In a highly available system, you should assume that replicas could be abruptly stopped at any time.

When an application is stateless:
* It can be scaled up or down at will to meet traffic demands.
* It can be bounced at will if it is mis-behaving.
* There is no need to recover its data if it crashes.
* The load can easily be balanced among its replicas in a round-robin
  manner.  Sticky sessions are not needed.