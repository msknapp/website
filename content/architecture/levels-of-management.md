---
title: Levels of Management
weight: 8
draft: false
lastmod: 2023-03-22
summary: You should choose between serverless, managed, and custom solutions depending on your system's usage.
---

A common design decision involves choosing between:
* Serverless options
* Managed services
* custom solutions

Each could be the solution depending on the situation.

# Serverless

Obviously, there *is* a server, but not one that your team sees or needs to 
manage directly.  You trust the cloud to perform maintenance on the real servers.
The perfect example is AWS Lambda, you just give it a function, a brief program
that defines what it should do when invoked.  Your team is not charged for the 
time it is not running.

Serverless solutions require far less maintenance time than other options.
They also tend to cost less.  They may be limited by run duration, for example,
Lambda functions cannot run more than fifteen minutes (it was five).
They tend to be the best option for services with lower traffic and short
process times.

# Managed Services

A "managed" service is one that *does* have a server, and your team can see it, 
interact with it, etc.  Your team could be on the hook for any vulnerabilities in it,
and some maintenance issues.  However, the vast majority of the maintenance
for the service is managed by the cloud.  Amazon RDS is a good example, a database
could easily be replicated, backed up, patched, and scaled by the cloud provider.

Managed services require less maintenance than custom solutions.  They are much
more expensive than serverless options, because you are being charged by the hour,
even if the service is not doing anything.

For many use cases, managed services will be adequate, but some will need 
a newer feature that is not possible with managed services.

# Custom Solutions

Essentially, this involves spinning up your own virtual machine (EC2 instance),
and installing the software yourself exactly as needed.  This obviously 
requires the most work to setup and maintain.  It will be more expensive
than serverless options.

The cost of running the service will be less than managed services, unless
you add the cost of engineering time/labor, then it would be more expensive.

Making your own custom solution might be the answer if:
* You have high traffic levels.
* You have the engineering expertise to handle it.
* You need custom features and capabilities that are not available from 
  serverless or managed solutions.
