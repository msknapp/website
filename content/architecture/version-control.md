---
title: Version Control
draft: false
weight: 1
lastmod: 2023-03-22
summary: Code and configuration should be tracked in a version control system.
---

The code in your environment should always be in some kind of version control.  Git is by 
far the most common version control software in use today, because it has a powerful 
branching feature and is open source.

It's not just code that should be version controlled, but configuration too.  This permits
a git-ops CICD pattern.  It also makes it possible to troubleshoot issues in your clusters
later.  Frequently, we need to look back in time to understand why a failure occurred, in 
a post mortem.  You need to be able to identify *what code* was running at the time.

# Tags and Releases

The code you run in a cluster should never be linked to a certain branch of your git repository,
but instead, it should always be tied to a tag or release of the code.  Tags or releases,
once made, should never change.  The reason is that, by using tags, you can be absolutely
sure about what code is running in the environment.  If you use a branch, then it is 
difficult to be certain what code is running, because branches can change over time.

When it comes time to conduct a post-mortem, if you did not use tags or releases in your
cluster, you will regret it.  It will be difficult to be 100% certain about what code 
was running.

## Releases versus Tags

Releases add extra metadata to git tags.  They are not necessary to use, but they can add
value.  You can upload binaries or other artifacts for a release, but that is not 
possible with a tag alone.  Releases are more important if you intend on distributing
executables from Github instead of Artifactory or Nexus.

## Semantic Versioning

The version you assign to tags or releases means something.  In semantic versioning,
you have:
* major version
* minor version
* patch version

These get concatenated, like `v1.2.3`.  Each major version greater than 0 implies that
it is not backwards compatible with previous versions.  `1` is the exception, because
major version `0` is considered to be in flux or unstable, version `1` may or may
not be backwards compatible with `0`.

A minor version is backwards compatible, but adds new features or capabilities.

A patch version increment usually means that bugs or insignificant issues are being fixed.
Users should not notice a difference.

Typically, git tags for Golang projects include a `v` prefix.  Docker images usually 
do not have a `v` prefix.

Creating temporary or experimental tags is encouraged, but they should be distinguished from 
more stable tags by having some kind of extra suffix.  My habit is to append a hyphen, 
my initials, and then a number.  That way, other engineers will know where the tag came from.

## Rollback Plans

You should never deploy a change to production without also having a plan to roll back the 
change.  The plan should be documented or obvious.  The rollback plan must also be tested
if it is a manual procedure.

Typically releases are made using a CICD pipeline, and they simply state that the cluster
should now synchronize itself with a certain tag or release of the infrastructure code.
This pattern is highly effective because rollback is made obvious, all you do is run 
the same CICD pipeline again, but with the previous tag.

# Artifact Management

Your project artifacts (executables, jars, docker images, etc.) should be stored in a
reliable repository.  Artifactory is distributed and replicated, making it highly
available and reliable.
