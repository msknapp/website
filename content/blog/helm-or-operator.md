---
title: Helm Chart or Operator
summary: Which one should I use?
weight: 1
draft: false
lastmod: 2023-03-22
---

If you are trying to customize your own Kubernetes cluster, you may eventually be 
faced with the choice between writing your own operator, or writing a helm chart.

Having worked with Kubernetes for years, and writing several operators, I can 
verify that *it is tough!*  The concept sounds simple and brilliant, but in practice
it has been consistently challenging to get right and test too.  Updating 
operators can also be a significant problem.

I will discuss the pros and cons of each approach below.

# External Resources

Kubernetes operators are able to manage resources external from the cluster they 
are running in.  To do the same thing with a Helm chart, you would need to create 
a job or deployment running a script.  It would certainly not be ideal.

If you need to manage a resource external from the cluster, an operator is 
the ideal solution.

# Reacting to Lifecycle events

Kubernetes operators can listen to external events and automatically be triggered
to run a task.  With Helm charts, this might not be possible.

However, in my experience, it is rare that you have a situation like this.

Let's give an example.  Your company has an enterprise wide service, which reports
if it is in a "stable" period.  During that time, certain processes must not run.
An operator could listen to that API and react to it.  For a Helm chart to do 
the same thing, you would need to add a job or deployment running a script or 
application with the logic.  It would be ugly.

# Reuse

Helm charts can be re-used, but the author needs to keep other use cases in mind.
That is often not the case.  Usually there are many variables in a Helm chart, 
and new engineers would be confused by them.  

Operators are easier to re-use among many people than helm charts.

# Ease of Creation

Writing operators involves programming and testing.  It can be very time consuming.
Helm charts only need some templates to be written.  It is far easier to create
Helm charts than operators.

# Ability to Change

Since operators involve writing code, they tend to be far more difficult to change
or update than Helm charts.  They require having thorough unit tests.  Operators
also have rigid object schemas built into them.  When you want to change the 
schema, it can be a major hassle.

Helm charts, on the other hand, are **very** easy to change or update, with little
fuss.  Adding new controls to them tends to be much easier than updating an operator.

# Loose Coupling

Helm charts tend to be far more loosely coupled from services around them than
operators.  This means over time, they are easier to maintain.

# Some Examples

Let's consider some example scenarios

## Creation of Lambdas

Let's say you need AWS Lambda functions to be created at will.  Since this is a
service external to the cluster, writing an operator is a better option than
some very odd Helm chart.

## Creation of a Deployment

Let's say you need a Kubernetes Deployment, Service, PDB, and HPA to all be created
in response to the creation of a custom resource.  You *can* do this with an 
operator.  The advantages you get here are:
* it would be easy to hand this operator to other teams and expect them to use it.
* the operator could react to external events.

However, in my experience, it is a much better idea to use a Helm chart in this 
scenario, because all of the resources being created are internal to the cluster.

If you go the route of writing an operator, what tends to happen is engineers
code in the resource schemas.  Then, if they decide to change them later, they 
need to update their code, their unit tests, and go through the entire testing
and deployment process.  This is a slow iteration process and will result in 
unproductive teams.

A keen developer might decide to just use golang templates in the operator,
which get used to generate the other resources.  Congratulations!  You just
invented a much weaker and brittler version of Helm.

That is why I recommend using Helm charts for creation of Kubernetes resources
within the current cluster.

## Creation of Lambdas and a Deployment

In this scenario, some event happens and the system must create one external
resource (Lambda), and one internal resource (a Kubernetes Deployment).

You could write one operator to handle both, but what will you do later, when
it's decided that a Kubernetes Service is also needed?  Update the operator?
Or decide to just make a Helm chart for that?

In this scenario, I recommend writing one operator specifically for the 
single external resource, and one Helm chart that wraps everything together.
