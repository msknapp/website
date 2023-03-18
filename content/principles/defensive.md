---
title: Defensive Programming
description: Making your code defensive.
summary: Making your code defensive.
draft: false
weight: 2
---

Defensive programming is a design pattern where each component in software acts very defensive in nature.
For example:
* Fields become private.
* Classes are often made immutable.
* Functions validate all of their inputs very strictly.

It might also lead to adopting certain code patterns such as:
* The Singleton.
* The Builder.
* The Factory class.

I would like to argue, from my experience, that:
* Defensive programming is more or less useful depending on the programming language.
* Your number one defense against bad code should be TESTS!  Not defensive coding patterns.