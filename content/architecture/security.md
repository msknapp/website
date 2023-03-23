---
title: Security
draft: false
weight: 16
lastmod: 2023-03-23
summary: Your system should follow industry best practices for security and implement defense in depth.
---

Adding security to your services can be immensely challenging, and it is very easy to get it wrong.
Generally speaking, you should avoid writing your own custom security solutions, and stick to open
source software that has been battle tested and proven.

# Audit Records

For systems that are critical or access controlled, you should have a record of 
who changed what and when.  One strategy is to leverage git-ops for this purpose, 
but that can become very difficult to automate and also review later.  Instead,
you should treat audit records similar to logs, as streams of events.  In fact,
you could just use logs for your audit records.


# Secret Stores

Source code should never hold secrets, even in past commits.  All secrets should be 
stored in a secure "locker" like Hashicorp Vault.  They should be downloaded at 
boot time.  Using a Kubernetes secret is less than ideal because it doesn't 
actually encrypt anything.

One major advantage of this approach is that secret keys can be rotated independently
and automatically, and the application will just continue functioning.


* access controls, authorization
* Mutual TLS
* Defense in Depth
* Vulnerability scans
  * static
  * runtime
* Firewalls, security groups