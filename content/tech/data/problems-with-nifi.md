---
title: "Problems With NiFi"
draft: false
weight: 2
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

0. It has a limited degree of fault tolerance.  Flow files are written to one disk, they are not replicated.
   If that one disk crashes, the flow file is gone permanently.  We work around this by using EBS (which replicates 
   within an AZ automatically), but we shouldn't have to.  You can also work around this by leveraging kafka.
1. While NiFi is fault tolerant (to a small degree), it is not highly available.  If one node in a cluster dies, the other nodes will not pick up its work.
   We have hacked our way around this by leveraging Kubernetes, but really it should have been highly available out of the box.
2. There is one content repository per flow.  Really, you should get to pick and choose.  For any connection between processors, you should
   have the option to say "this must only be stored in RAM" or "this can spill to disk" or "this must all be written to disk".  You should
   even get to choose a replication factor.  
   There should have been an option to change content repositories with each connection.
   This makes it much harder for us to stay compliant, because we may have sensitive data in RAM but not on disk.  If we use the volatile 
   content repository, the entire flow must use that, and you risk running out of RAM.  If you have separate NiFi flows for this, you have just 
   doubled the amount of maintenance you must perform.
3. Since it relies heavily on the disk, it is reducing the processing speed.  Streaming systems could process the data faster.
4. Configuration is a nightmare.  There are a million things to configure and they are not well documented.
5. There is no native way to re-distribute flow files amongst the cluster.  One must rely upon kafka or http methods with an external load balancer.
6. NARs are very heavy weight and redundant.  Every nar has its own set of dependencies.  If two nars have the same dependency, it is still
   represented twice in NiFi.  The structure for a nar is very complicated and if you get it wrong you can lose days trying to fix it.
7. Very difficult to learn how to extend.
8. Clustering on kubernetes is a hassle.
9. The whole "process group" heirarchy, and how it relates to controllers is very confusing.  The way that controller groups you define in the root are not visible 
   to child groups is super confusing to people.  I feel this was a big mistake on the NiFi developer's side.
10. You often need to take an extra two or three steps In NIFI purely because a more direct translation of data is not supported out of the box.
   For example, to go from psv to parquet you must first translate the file to avro in between.
11. Support for table style data is still limited.

# The advantages it has
* Start/stop sections of the flow at runtime
  - we are not supposed to leverage this because we have an "immutable architecture" policy in production.
* Easily re-configure your flow at runtime, even add/remove whole sections
  - the "immutable architecture" prevents us from really leveraging this.
* Visual representation of the flow.  This is still hard to reproduce without it.
* realtime metrics.  This is easily produced with JMX, graphana, and/or cloudwatch.  However,
  the metrics will not be shown visually next to a node/processor on a visual graph.

Where NiFi really shines is with large (megabytes), semi-structured documents that don't require low latency processing.
