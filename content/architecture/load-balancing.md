---
title: Load Balancing
weight: 13
draft: false
lastmod: 2023-03-22
summary: Traffic to your services should be balanced among your replicas.
---

We just established that for an app to be highly available, there must be multiple
replicas, and in multiple availability zones.  This would not help much if all
the web traffic went to just one of them. 

Load balancing ensures that requests are split evenly among all replicas, so that
no single instance is overloaded with requests.  Cloud providers have load
balancers for connections to EC2 instances.

In Kubernetes, a "Service" also provides some simple load balancing.  Istio supports
more complex load balancing logic too.

The principle applies to databases and stream processing too, with sharding and replication.
