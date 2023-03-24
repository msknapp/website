---
title: Single Responsibility Principle
draft: false
weight: 4
summary: Software components and classes should not be responsible for many concerns.
description: The importance of single responsibility.
---

The Unix motto is "Do one thing, and do it very well."  This is quite similar to the
single responsibility principle, which suggests that any single class
should not be responsible for multiple things.  In other words, we want to keep the
scope of a class small and simple.  So it is in line with the KISS principle.  The
principle of cohesion is also relevant to this.

# High Cohesion

Code is cohesive when its functions are all closely related and aligned to a common purpose
or objective.  It is similar to the single responsibility principle.  It helps you keep the 
scope of classes small, so they are rarely updated (in line with the open/closed principle),
and far less likely to have bugs.  This pattern also means your team is less likely to 
have merge conflicts in your code.
