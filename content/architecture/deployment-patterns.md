---
title: Deployment Patterns
weight: 11
draft: false
lastmod: 2023-03-22
summary: You should be able to update your applications safely with zero down-time.
---

For mission critical systems, the largest risks are generally encountered during deployments
of new software.  Updating your applications can be accomplished by several methods.

# Big Bang Deployments

Basically, the older service is stopped all at once, and then the new service is started 
all at once.  This is the most primitive and risky deployment pattern.  There *will* 
be downtime, a period where requests from users cannot be processeed.  If you have a 
mission critical service, this deployment pattern is not an option.

# Rolling Updates

In this pattern, one replica of the new service is started.  When it is healthy and 
ready for processing requests, then one replica of the old service is stopped.  This
repeats until only the new service is running.  You can configure how many replicas
may be stopped at once.  The pattern ensures that there is no downtime during the deployment.

In Kubernetes, this is managed using Pod Disruption Budgets (PDBs), which can state that,
for example, at most two replicas can be down at once, or at least three must be 
available at all times.  Using PDBs can greatly reduce the risk of having your 
application be overloaded during a deployment.

The problem with rolling updates is there is usually little to no testing or 
monitoring in the process.  The health checks on containers usually just return 
a 200 status code without actually checking anything.  The application could 
be failing every single request except health checks, and the rolling update
would continue.

The rollout can happen very fast, like in a matter of a few seconds.  However, 
recognizing a failure and choosing to roll back could take much longer, in the 
order of minutes or hours.  At that point, there could have already been so
many failures that a severe incident has occurred, substantial money has been 
lost, and the company's reputation has been tarnished.

# Blue Green Deployments

In this strategy, the new version of the service is brought up completely, then tested,
then traffic is shifted entirely to the new service before the old one is torn down.
Traffic is usually shifted all at once by altering what a domain name points to.
It is less efficient than other methods, because for a period of time, there are 
two running versions of the application.  These take up twice the resources in the 
cluster.

In some environments, there are limits on the total number of nodes, replicas,
virtual machines, pods, etc.  A blue green deployment might run into these limits, 
and at that point it would fail.  So there is an added risk involved with this 
strategy.

Some applications are stateful in nature, and the total number of replicas is 
assumed to be constant.  Apache NiFi or Kafka are examples like this.  For these
applications, a rolling update might not be possible, but a blue-green deployment
is manageable.

Its major benefit is that it is simple to perform, test, and roll back if necessary.
You can even pause to do more extensive tests of the new service.  For this reason
it is safer than a rolling update.

# Canary Deployments

The canary deployment pattern is the safest and most complicated of the options.
A new version of the service is started.  A service mesh is utilized to gradually 
shift web traffic to the newer version.  Typically a small amount of traffic 
is started with, like 1%, and the increments grow much faster.  There is a 
period of time between each increment, of about a few minutes.

While the deployment is in progress, metrics are being monitored of each service.
For example, the number of failed requests, or the number of requests that are slow
to respond might be monitored.  If the metric goes out of the acceptable range,
the canary deployment automatically rolls back.  So for example, if 1% of the 
http requests sent to the new version are failing, it might automatically shift
all traffic back to the original stable version.

The wait period between increments is so that metrics can be monitored over time.
Some problems don't arise until an application has been running for a while under
load.

An example schedule might be 1%, 5%, 10%, 25%, 50%, 90%, and then 100%, with five 
minutes between each interval.

Canary rollout traffic shifting is usually not done manually, but instead by some
automated orchestrator like Argo Rollouts.  It integrates with a service mesh 
like Istio or LinkerD, and a metrics service like Prometheus, NewRelic, or 
DataDog.

One of the challenges of canary rollouts is setting up the metrics that need 
to be monitored.  The applications need to export these somewhere.

## Low Traffic

If a web application receives a low amount of traffic, canary rollouts generally 
won't do you any good.  For a canary rollout to detect a problem, it needs
metrics related to failed or passing requests.  You could generate artificial
traffic to the service, but that may not accurately represent real requests.

So my advice is for applications with low traffic to use one of the simpler
deployment methods, and leverage system tests to confirm it once it's done.

# A/B Testing

At first glance, people might confuse A/B testing with Canary rollouts.  A/B 
testing *also* involves sending some traffic to a new application version, and 
some to the older one.

Here are the differences:
* A/B testing tends to be more concerned with the UI for users, and if the changes
  are acceptable to the users.
* A/B testing tends to run for longer periods of time, because what you are waiting
  for is to see if the end user dislikes the changes.  For example, a UI change 
  might be shown to a small subset of users for a few days before it is fully 
  utilized or rolled back.
* The users in A/B testing get a consistent experience, they either always get
  the behavior of A, or of B.  That is unlike canary rollouts, which may not 
  have sticky sessions, and where consecutive requests by the same user could 
  be handled by different versions of the application.


* rollback plans