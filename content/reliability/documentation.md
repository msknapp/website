---
title: Documentation
weight: 90
draft: false
lastmod: 2023-03-24
summary: Thoroughly document your code and services for fellow engineers and end users.
nextlink: /architecture/
---

Code documentation is usually the last thing that engineers ever worry about.  Failing to 
document your code and services means that:
* You must be called or paged if it fails.  You cannot realistically expect fellow co-workers to
  be able to fix problems in it.
* If fellow co-workers are assigned a story on that application, they cannot quickly get up
  to speed on it and contribute improvements to it.
* If you are sick, on vacation, or unavailable, there is a great risk that the application 
  could go down and cause a severe incident.  The team might blame you for the issue, because
  they did not understand how to fix it and you were unavailable.
* There is a greater chance that other engineers will change it in ways that break the
  service, are not backwards compatible, or are not in line with your future plans or intentions
  to update the service.
* The team is more likely to abandon or replace that service.
* There is practically no chance the application could be re-used by other teams in the company.
* Competing solutions are bound to be developed, and one day you might find other teams 
  insisting that your service gets replaced with theirs.
* Management and product owners are unlikely to understand the application, and it can be
  very difficult to convince them that certain changes are needed.  It becomes a pain
  to try justifying stories related to the application.

There are a variety of ways to document code and services on a team.

# In Code

Confusing or un-intuitive business logic should have some comments in the code explaining it.
Ideally code should be intuitive enough that comments are not needed, but there are 
often exceptions.  Links to external docs or requirements are encouraged.

# Function Documentation

It is a good idea to add comments related to individual functions in the code.  Most IDEs
today have keyboard shortcuts that will show those function comments automatically, 
which could be immensely helpful for engineers who didn't write it, or can't remember
how it works.

Be sure to document how the function behaves in certain edge cases, and any requirements
it has on the inputs.  Document what the output would look like, if it can return 
null values, and under what conditions it throws an error or exception.

# Unit Tests

There is also a belief that unit tests should serve as documentation too.  The unit
tests should demonstrate how components get wired together.  They serve as examples
of how to use the code.  However, these still might be un-clear to readers, plus
the users might struggle to find the right unit test.  I do not recommend 
relying on them for documentation alone.

# README Documentation

This is probably more important than the previous documentation places.  When fellow
engineers, product owners, and teams take a look at your code in github, the
README is often the only thing they take a look at.

In your README, be sure to have:
* A general description of what the application does.
* An explanation of what is in scope or not.
* A diagram or two to aid in the explanation.
* A thorough description of the configuration for the app, what values can be 
  used in each field of the configuration.
* An outline of packages in the repo, major high level services, and how
  they interact with each other.

The README is designed for other software engineers.

# User Manuals

For end users, a less technical user manual should be provided.  It can be a text
document, a Confluence page, or a static website.  It's more important that the 
manual is written than that it is in a nice medium or format.

Confluence is useful for documentation that might need to be updated by less 
technical employees, or management.  The managers might not have the time, 
access, or experience to update static website code.

Having a poor user manual will result in a low net-promote score, low adoption
of your service, many support questions sidetracking engineers, and probably
a worse performance rating for the engineers that made it.

User manuals should be easy to search through, and not overwhelm the user all
at once.

# Playbooks and Runbooks

This is more of a support plan for the application.  These outline how the team
should respond to incidents, what things they should check for troubleshooting
purposes, and what controls they might need to toggle.  These are a bit more 
likely to change over time, so using a medium that is easily updated by co-workers
is recommended.  This might suggest using Confluence over making static
websites.

## Playbooks

Playbooks tend to be more high level, focusing on company policies and procedures.
For example, they document who should be called, how to setup a bridge, how
the incident severity is decided, etc.

## Runbooks

These are more about what controls should be switched in the event of specific
failures.  They can involve administrative scripts, but don't necessarily need to.