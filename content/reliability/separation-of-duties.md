---
title: Separation of Duties
draft: false
weight: 3
lastmod: 2023-03-22
summary: Code needs to be peer reviewed as a requirement.
---

There is a saying, that one can be "judge, jury, and executioner", which is obviously 
a terrible idea.  Likewise, in software systems, you don't want a single engineer to 
be able to commit a change, merge it to main, and then deploy it all by themselves.
Your git repositories should be configured in a way that:
* Engineers cannot push directly to main, they must create pull requests.
* Pull requests can only be accepted and merged if a peer engineer has reviewed it 
  and approves.  You can require more than one peer review too.

Along the same lines, when code is deployed to production, that should also require
review or approval from a stakeholder, except if the service being changed is not
mission critical.
