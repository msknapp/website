---
title: CICD
weight: 5
draft: false
lastmod: 2023-03-22
summary: You should have an automated way to build, test, and deploy your applications.
---

CICD stands for continuous integration, continuous delivery or deployment.  It basically
automates how code is built, tested, published, and deployed.

Some examples of CICD:
* Jenkins
* CircleCI
* Travis CI
* Tekton

Continuous integration means that as code is being changed, a service will automatically
build it (compile it), and test it.  Continuous delivery means the code is published 
somewhere, like Nexus or Artifactory.  Continuous deployment means that the code is 
automatically installed and run in a cluster.

Typically the delivery or deployment steps would only run for certain branches of code, 
like the main (or the old-school "master") branch.  If you are using a gitflow pattern,
you might have merges into the "develop" branch result in a deployment to your QA cluster.

Your CICD pipelines should be in code and committed to a git repository, so they are 
version controlled.  Previously, teams would often define their pipelines in a UI, and
there was no history of changes.  That made it impossible to roll them back if there was
an issue.