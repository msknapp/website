---
title: Kafka
draft: false
weight: 50
lastmod: 2023-03-26
summary: How Kafka Works
---

Kafka is a message queue.  It's designed to scale to higher volume and velocity of data, and it
does this using replication.

# Brokers

Kafka clusters contain brokers.  A "Broker" is a node running Kafka.
There can be many brokers working together in a Kafka cluster.

# Topics

Streams are called "topics".  A cluster can have many topics.  Topics are partitioned.

# Partitions

* Partitions are composed of segments, which is like a rolling file.
  It is durable and persistent, stored on disk.
* Records get assigned a partition key, which decides which partition they go to
  on the topic.
* Partitions are replicated by a factor, to other brokers.  This lets Kafka be fault
  tolerant.
* Partitions get allocated to brokers.
  * There is strict ordering on partitions.
  * There is not strict ordering when reading from a topic.
* re-balancing of partitions is not built in to Kafka, Confluent adds it.
* partitions are also called logs.  They must be written in sequence, and
  are immutable once written.  You cannot edit or remove past records.
  They can expire.
* Every partition has its own offset space.

Choice of partition
* Producers can specify the partition
* it can by chosen by hashing the key and taking that mod the number of partitions.
* can just be round-robin.
* You can provide a custom partitioner.

# Segments

# Durability

Events are written and stored on disk.

# Retention

Reading a message does not remove it.

Retention is based on time.  1 week is default.  Each topic can have
its own retention.  It will only purge entire segments at a time, and
only if all of that segment has expired.

Different policies can be set for Queue record deletion.  They can be 
deleted if they are over a certain age.

# Records

Records have:
* headers
* key - can be anything serialized, like bytes or a string or integer.
  * keys do not need to be explicitly set, the broker can randomly assign one.
* value
* timestamp - can be set by producer, or by broker.

# Keys

Messages can have the same key, that is allowed.  Then they would go to 
the same partition, and hence the order would be guaranteed.  So giving 
messages the same key is a means of guaranteeing the records are processed
in order.  For example, choosing the customer ID as the key would guarantee
that for each customer, their transactions are processed in order.

# Producers

* Producers send events to Kafka.
* Events are produced and sent to the Queue.

Producers and consumers need to know or declare an object type for 
the key and value, so that it can be marshalled correctly.

There are many different SDKs for producers and consumers in different languages.

## Acknowledgement

* Brokers can acknowledge receipt of events.
There are three different ack policies.
* No acknowledgement.  Low latency but possibility of data loss.
* leader acknowledges. 
* all ack, including followers.

Ack means it is written on a disk now, not just memory.

# Consumers

There can be many consumers.
Consumers can subscribe to as many or as few topics as they want.
Consumers poll.  Consumer poll by a timespan they choose, usually in milliseconds.
Kafka keeps track of the "offset" that each consumer is at.

Newer versions of Kafka allow consumers to read from followers.

Frequently a consumer is also a producer, sending to a different topic.

## Consumer Groups

Consumers can be grouped, which means they will not be given duplicate records.

Each consumer offset is tracked within the kafka cluster in its own topic.
This means that consumers could crash and recover, and pick up where they
left off.

In a consumer group, records from a topic will be read by one of the 
consumers.

# Replication

Followers are trying to keep up with the leader by pulling from them.
Leaders do not push records.  When a record is written, that can 
return before it is replicated.

* Each partition is replicated to other brokers. 
* For each partition, one broker is the leader, others are followers.
* You can configure how many replicas there are for each topic.

# Fault Tolerance

# Scaling

# Loose Coupling

Enables loose coupling between producer and consumer.  They can scale
separately, they can be updated separately.  If one is slow, it will 
have no impact on the others.

# Distributed Consensus

* Kafka uses ZooKeeper Ensembles to keep consensus on the state of things.
  * distributed consensus manager.
* ZooKeeper is used to elect a new leader when a broker dies.
* KIP: kafka improvement proposals 500, wants to remove ZooKeeper.

# Delivery Guarantees

* At most once.  Some are lost.
* At least once.  Duplicates sometimes occur.
* Exactly Once

# Patterns

There are certain patterns or scenarios to be considered.
* Exactly once processing
* Replaying from past.
* Guaranteeing order of processing.

How is Kafka Highly Available?
* Add more replicas to topics.

Can you scale up kafka brokers at runtime?