---
title: Writing Effective Operators
weight: 2
draft: false
lastmod: 2023-03-23
summary: How to write Kubernetes Operators that work.
---

The fundamental way operators work is really quite trivial: if a change is made to one 
custom resource, then react by updating other resources.  It sounds so simple, but in 
practice, it tends to get quite complicated.

Here are some guidelines to follow that should make it easier to maintain your operators:
* Prefer to keep their scope limited, so they are simple.
* Avoid having operators manage other Kubernetes resources, these are easily managed from 
  Helm charts instead, with a few exceptions.  Helm charts tend to be more powerful than
  operators when it comes to managing other Kubernetes resources.  One exception is if you
  need this capability to be shared with other teams, or to react to an external state.
* Avoid having an operator manage more than one downstream resource.  Keeping it transactional
  and consistent gets quite complicated.
* Keep the schemas of the resources simple, don't add fields that you expect will be used 
  *eventually*.  The schemas should ideally not cover unrelated, independent concerns.  If 
  it does, that might be an opportunity to split the operator up, along with its custom resources.
* Annotate schema fields with their validation expressions.
* Avoid having a single operator manage more than one custom resource.
* Prefer to have the operator use templates to construct downstream resources.  These 
  templates should be possible to change from outside the application, like from a 
  configuration file.  Usually templates give the engineers far more controls over how
  resources are created than a configuration file.
* The operator should have no assumptions about naming conventions or labels coded into it, 
  these conventions should either be set in configuration, or avoided completely.
* If the operator fails to sync a resource, don't retry that for at least a few seconds, 
  maybe minutes later.
* The operator should update the status of the custom resource it monitors.  Include 
  vital information about downstream resources it created, such as their identifiers.
* The operator should make no assumptions about the current state of downstream resources.
  It should not assume that the status listed in the upstream custom resource is accurate.
  It must check the actual state of downstream resources and see if it is in sync.

# Checking for Sync Status

In Golang, maps have a random order for keys.  That means if you try to compare the desired
state with the actual state of a golang object, it may wrongly consider them different.

To work around this issue, I recommend the following:
* The operator should create a specification for its downstream resource, and this spec
  should be a text document created from a template.
* The downstream specification should be hashed (security is not a concern her), with
  MD5, SHA1, or SHA256 for example.
* The hash should be assigned as an annotation or label on the Kubernetes resource 
  that it creates.
* Then, the operator can compare the hash of the desired state with the hash annotation 
  on the real downstream resource.  If they are the same then there is nothing to 
  synchronize.

# Chaining

It's common to have operators form a chain, where one of them creates a resource that
another operator will react to.  This is a form of loose coupling and is effective and
encouraged.  It means that you can update these stages separately.

# Schema Changes

With Kubernetes operators, schemas present a significant challenge.  The Kubernetes API server
will only store one version of a schema.  So if you remove a field in v2, and make that 
the stored version, then all existing resources permanently lose that information.  If you 
add a new field, and it is required, then you must assign defaults to all existing resources
before you update the operator.  It might be necessary to run a migration script in this process.

Kubernetes schemas are represented by Golang structs.  If you have multiple versions of a 
Kubernetes resource, the operator's code needs to have sufficient structs to represent 
either of them.  That makes the code excessively complex and redundant.  Operators make
it difficult for engineers to follow the DRY principle in their code.

If you update a schema without changing its version, that could result in having an
operator that is not backwards compatible, and impossible to roll back.

Before updating an operator or changing its schemas, you should collect a backup of all
existing resources, because it may be impossible to recover them later if you don't.

# Multi-Cluster Synchronization

A common problem is that operators run in multiple distinct Kubernetes clusters, but they
are attempting to control a common global resource.  They may clobber each other's work, 
and that can result in crashes or failures.

There are a few ways to cope with this problem:
* One controller could be active, and the other passive.  This can be decided with a lock 
  in etcd or ZooKeeper.  Another effective pattern is to have them query for a DNS host, one
  that is latency or failover based.  If the host IP of the DNS record points to the current
  cluster, then that controller is active and may proceed.  Otherwise the controller is
  passively waiting for the moment that the other cluster goes down so that it can take over.
* A global lock could be utilized in DynamoDB.

# Custom Resources or Databases

Teams might choose not to store configuration data in a real database, but instead to 
just have Kubernetes custom resources.  This can work so long as there are not many of
these resources.  There are a few problems with this approach:
* Keeping the resources in sync across clusters or regions can be very difficult.
* The Kubernetes API server and etcd are not nearly as capable as most relational databases.
* The SDKs for interacting with Kubernetes are much more difficult to use than the SDKs 
  for most databases.
* It is not easy to perform complex queries like joins, groups, or aggregation.  These need
  to be written into the code, and that is error prone.  The Kubernetes API server and 
  etcd do not support SQL queries out of the box.
* It can be more difficult to produce dashboards around Kubernetes resources than database
  records.
* Now all the code that accesses these custom resources needs to have Kubernetes service 
  accounts setup and granted access to them.

For these reasons, I recommend having a global datastore be the official source of truth.
Kubernetes custom resources should aspire to be in sync with that source of truth, not 
the other way around.