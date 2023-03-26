---
title: You Build You Own
weight: 20
draft: false
summary: Software developers should also be on the hook to support their applications.
lastmod: 2023-03-24
nextlink: /reliability/
---

In the past, companies often hired one set of engineers to develop applications, 
and another set to support them and be on-call.  The problem is the support staff
did not fully understand how the applications worked, because they didn't write 
them.  So they did not know how to fix them if they were broken.

What's worse, is the software developers were not thinking about how their
applications will be supported in production, because that will not be their 
job.  So common devops themes like logging, metrics, alerts, and controls 
were not given a priority.  This made it even more difficult for support
staff to troubleshoot issues.

# Combining Roles

This lead to the idea "you build it, you own it" (YBY0), which basically means the 
developers are also the support staff.  That motivates the developers to make
their software applications easier to inspect, troubleshoot, and monitor.
Now they will eagerly add logging, metrics, and alerts to their applications
because they know that they are on the hook if it fails.  It also 
encourages them to conduct more thorough testing of their applications, 
because they don't want to be paged in the middle of the night, on a weekend,
to fix their services.

# A Devops Mindset

It is like saying that devops is not a job category or separate position, it 
is a mindset that all engineers are expected to follow.  Following this 
principle generally reduces the number of incidents encountered with 
software systems.