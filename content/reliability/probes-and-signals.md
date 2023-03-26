---
title: Application States and Signals
weight: 9
draft: false
lastmod: 2023-03-22
summary: Your applications should expose health information and be able to gracefully shut down.
---

The applications in your cluster should be able to convey the state they are in, and be able 
to respond to Linux signals.

# Probes

In Kubernetes, there are liveness probes and readiness probes.  Liveness means that the 
application is running, while readiness means that the application can safely accept 
requests, like HTTP traffic.  Robust software systems *support* these probes.

Typically the liveness probe just checks if a port is open, while a readiness probe actually
submits an HTTP request to a "health" endpoint, and expects a 200 response code.

These checks are critical for Kubernetes, because it ensures that Gateway servers do not
direct web traffic to replicas that are unhealthy.  It also regulates how deployments 
can proceed.

# Signals

Robust applications will also be able to respond to Linux process signals.  Specifically,
they should respond to an instruction to terminate gracefully.  At that point, the 
readiness probe should not return a 200, it should be considered unhealthy, so that 
no future requests are sent to the replica.  All requests that are already being 
processed should continue.  When they are finished, the application should shut down.
That is a graceful shutdown, and it is necessary for rolling updates.
