---
title: Authorization Architecture Mistakes
draft: false
lastmod: 2023-03-17
weight: 1
description: Common pitfalls people make when adding authorization to their apps.
summary: Common pitfalls people make when adding authorization to their apps.
---

In my experience, adding authorization logic to applications is very challenging.  I have seen
multiple teams in large organizations all try to solve the problem themselves, and they each 
struggle.  Here are some common mistakes.

* Writing authorization logic into your server's code.
* Not using pre-made solutions.
* Establishing a team secret.