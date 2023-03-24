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

# Service

Classes that are services tend to be created when the process boots, and live
for as long as the process is running.

# Builder

# Factory

# Fluent

# 