---
title: Configuration Management
weight: 4
draft: false
lastmod: 2023-03-22
summary: It should be possible to reliably re-create your infrastructure from code.
---

A professionally made software system is not setup manually.  There must be code that is
able to reliably re-create the system and all of its environment.  This might consist
of an Infrastructure as Code (IAC) solution, or a CICD pipeline, or a combination of these
things.  IAC can also be called configuration management.

# Infrastructure as Code

IaC gives us a means to reliably reproduce the state of a cluster or system.  Typically
these IaC configuration files will describe a "desired" state, but not the exact
method by which it should be reached.  Examples of functions done in IaC:
* Create managed cloud resources (EC2s, ASGs, ELBs, etc.)
* Install certain libraries, packages, or applications on a set of nodes.
* Run certain daemons or programs on the nodes.
* Download secret values.

Some examples of Infrastructure as code, or configuration management:
* Terraform
* Ansible
* Chef
* AWS CloudFormation
* ArgoCD, Helm