---
title: High Availability
draft: false
lastmod: 2023-03-22
weight: 12
summary: Your critical applications should be replicated across availability zones and regions.  They should automatically scale based on demand.
---

# Auto Scaling

Robust system architectures always include some kind of auto-scaler. Essentially
the service will add more instances or replicas when the load on the system increases,
and may remove instances when the load drops.

There are several kinds of auto-scaling available:
* Virtual Machine, node auto-scaling.  For example, Amazon Auto-Scaling Groups (ASGs)
* In Kubernetes, Horizontal Pod Autoscalers (HPA) can scale based on CPU or memory
* Cluster autoscalers in Kubernetes can scale up the entire cluster.

# Replication

If your service is mission critical, then there should be more than one replica of
the service.  Ideally it should be stateless, so new replicas can come and go safely
without needing to re-shard data.  Prefer to use well established data storage 
software for your state, like Redis, Postgres, DynamoDB, MongoDB, Cassandra, etc.

# Multi-Region

Cloud services tend to be split into regions, each region is far enough from 
other regions that it should not be impacted by world events happening in
one of them.  They tend to be countries or states apart, with hundreds of 
miles between them.  For example, if an earthquake brought down a west
region, the east region is probably unaffected, and will continue functioning.

Availability Zones are within a region, there tends to be several of them.
Availability zones are separated enough that they don't depend on the same
power source.  Hence a power outage in one is unlikely to also affect another.

Availability Zones have high bandwidth between them, so the network speed is 
quite fast.  Between regions though, the latency of requests could be significant.
It could even make atomic transactions across regions slow enough that they
are impractical to use.

Important applications should always be replicated across availability zones.
This tends to not be a problem.

Mission critical applications should also be replicated to more than one region,
and that is where things get complicated.  There are several patterns to consider.

## Active/Passive

One region is on and serving all requests.  A different region is on and 
ready to serve requests, but is not receiving any traffic.

## Active/Active

All/both regions are on and serving all requests.  Any single request is sent
to both regions via mirroring.

## Active "and" Active

All/both regions are on and serving at least some requests.  Any single request is
sent to a single region for processing.  The region selected might be based on latency
between the host and client.

## Synchronization

When services are replicated across regions, a new problem arises: keeping the 
regions in sync with each other.  This tends to be more challenging than it sounds.

As a rule of thumb, there should be only a single source of truth for information.
For example, we could use a DynamoDB global table to store state.  Only one region
would be able to update it though, otherwise the consistency guarantees fall apart.

Along the same lines, your application servers should be stateless, then there is
no need to worry about synchronizing their runtime state across regions.
