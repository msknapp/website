---
title: Data Storage
draft: false
weight: 3
lastmod: 2023-03-23
summary: You should choose the right data storage system depending on the system requirements.
prevlink: ..
---

Earlier we said that to have a scaleable system, you should prefer to have
stateless applications.  Obviously you *need* to have data *somewhere*, and
that will also need to be scaleable.  That brings us to databases.

There are several ways to classify data storage technologies:
* Filesystem
* Object storage
* Data lakes
* Caches, key-value stores
* Relational (RDBMS), SQL
* NoSQL
  * columnar
  * document
  * consistent
  * responsive
* Data WareHouses
* Time Series
* Geospatial Information Systems

# Choosing Data Storage

When it comes time to choose what tech should serve as your database, consider
the following:
* How many requests do you expect over time?
* How much data is there to store?
* How fast of a response do you need?
* How will the data be queried?
* Is eventual consistency acceptable?
* Does your data have a strict schema?

## Number of Requests

If your data store will rarely receive requests, then a serverless solution will
probably serve your needs while costing far less than running a database instance.
In AWS, DynamoDB can satisfy the needs here, it is serverless.  If you have a 
Dynamo table with no data and no transactions, it costs nothing.

The only exception might be if you have strictly relational data and you want
transactional guarantees that it will be consistent.  Dynamo is a NoSQL store,
so it cannot perform transactions.  Keep in mind that in NoSQL stores can always
choose to just keep all information in a single record, even if it has some
redundancy.

## Data Volume

Relational databases tend to have the most features, but they may not scale as 
well as NoSQL stores.  Somewhere between 10GB and 100GB, you might want to consider
switching to a NoSQL datastore.  These can be sharded, partitioned, and hence
parallelized between multiple servers.

* data replication
  * RTO
  * RPO

## Response Latency and Throughput

If your database needs to support *very* fast reads and writes, or *very* high
throughput, then an in-memory data store is required.  Redis and Memcached are
the most common options here.

Using in-memory storage can be much more expensive than disk storage, but can 
respond much faster.  These are usually chosen for data caches.  Redis can 
also be configured as a least recently used (LRU) cache, whereby the items
that are queried the least tend to get purged, and the most frequently used
items tend to stay in memory.

If you need a very large cache, there are services that can manage sharding
Redis or Memcached separately.

Note that other databases often have some cache support too.

## Query Modes

There are several patterns by which data can be queried, and these suggest
a certain database type.

### Random Access

If the most frequent access pattern is random reads of single records, this
lends itself to relational databases, or NoSQL data stores, which excel in
looking up single records.

### Analytical Queries

In this case, large scale trends and aggregations are the most common way 
that data is queried.  For example, you want to answer business questions 
like, what percent of the users come from each region.

For this data access pattern, Data Warehouses tend to be the answer. 
Amazon Redshift, or SnowFlake are good examples.  They store data in 
columnar data files, which make them perform much better with queries.
They also are more efficiently stored.

### Batch Processing.

In this case, the entire table tends to be accessed and processed in 
one large batch job.  For this, a distributed data lake or data store
tends to be the best answer.  S3, HDFS, or Minio are good examples.
Data locality can also help speed up the processing.

# ACID

The acronym "ACID" describes transactions and consistency guarantees in a relational database.
It stands for:
* Atomic: all of the transaction or none of the transaction occurs
* Consistent: Readers will not be given stale values.
* Isolated: While the transaction is being processed, concurrent queries will not see partial updates.
  Also called Serializable.
* Durable: The changes will survive a crash.

# CAP Theorem

The CAP theorem states there is a relationship between these three aspects
of databases:
* Consistency (also called Linearizable)
* Availability
* Partitionability (also called sharding)

Specifically, a database can only guarantee two of the three aspects, because
there will always be a possibility of network failures between nodes.

A data-store is consistent if a read operation cannot return a stale value,
nor a new value that has not been committed.  Consistency does not require
the full ACID compliance that relational databases offer.  It only 
requires that the datastore is guaranteed to return the latest current 
value in any query.

A data-store is available if it responds to requests within a reasonable
time.  Availability here is not concerned with the possibility that
a replica could crash.  This might be due to the fact that many relational
databases also support having a hot backup standing by.  The aspect
of availability is more concerned with slow network connections between
nodes.  If the network is too slow, the database will either need to 
respond with what it knows (which might be stale data), or wait long
enough that it times out (which means it is not available).

A data-store is tolerant of partitions if replicas can be made of the 
service and they can inter-operate.  All the database operations can 
still function in the presence of these partitions.

## Relational Databases

Relational databases can achieve these aspects:
* Consistency: They are able to support ACID transactions, guaranteeing they will
  return the correct value.
* Availability: Since they are hosted by single machines, they can also respond
  promptly to requests.  They will not be waiting for network connections 
  because they hold all of the information locally.

They cannot *practically* be partitioned among multiple nodes.  There are projects
which have **sharded** relational databases among multiple nodes, effectively 
partitioning them, but some capabilities are impossible to retain like this.
For instance, keeping foreign keys between tables is no longer possible to 
guarantee.

## Consistent NoSQL Stores

These databases achieve the following:
* Partition Tolerance: The data is sharded among multiple nodes, and can be
  re-balanced as needed.
* Consistency: These databases achieve consistency using a quorum strategy.
  They could use ZooKeeper, etcd, or have a quorum built-in to their source 
  code.

Availability is not guaranteed.  If nodes in the quorum crash and are 
unavailable, you can reach a point where reads cannot respond in time, because
they are waiting for information from a crashed node.

## Eventually Consistent NoSQL Stores

These databases achieve the following:
* Partition Tolerance: The data is sharded among multiple nodes, and can be
  re-balanced as needed.
* Availability: Nodes may attempt to check with peers for the most recent
  information, but within a time limit.  It may return stale data.

Consistency is not guaranteed because the data can be stale.

# Data Store Types

The types of databases can often fit into several categories.

## Cache

These concern themselves simply with having the fastest response time
possible.  They usually retain all information in memory to achieve this.

Examples:
* memcached: a simple key-value store
* Redis: Holds a lot more features and special data types than mem-cached.

## TimeSeries

These store data in a special format that makes it easy to query for 
information over a time frame.  They are usually used for metrics and 
monitoring.

Examples:
* Prometheus
* 

## Search

These are designed to make searching text documents easy and efficient, 
even in the presence of imperfect terms.

Examples:
* Solr, Lucene

## Graph

These are specially designed for storing graph data, edges and vertices.
They make is easy to perform a variety of graph queries against them.

Example:
* Neo

## Quorums

These are designed to help other applications synchronize with each other.
Keeping multiple nodes in sync is a very challenging problem to solve in its
own right, so using a Quorum can make that easier.

Examples:
* ZooKeeper
* etcd

## Relational

These run on a single server, and are able to support very complex queries,
including SQL queries.  They tend not to scale very easily.  Many of them
support read-replicas.

Examples:
* SQL Server: for Microsoft fans, and integrating with other Microsoft products.
* MySQL, MariaDB: Open source, GPL licensed, easier to use than others.
* Postgres: Open source, considered more full featured and performant than others.
* Oracle: The most powerful database but is closed source and is not free.

## Document

These specialize in storing "documents" which are generic objects that can have
nested information and lists.  The schema is not rigidly enforced, and they 
support queries on fields that might not even exist on all records.  Documents
are usually conveyed as JSON.

Examples:
* MongoDB
* DynamoDB

Note that some relational databases also support these, but it's not their main feature.
Most document oriented datastores are NoSQL and partitioned.

## Big Tables

These are able to hold tables of data, with rigid schemas, among partitioned nodes.
Hence these scale very well, but they cannot perform ACID compliant transactions.
They cannot provide foreign key restrictions.  They do not support SQL queries, 
nor joins.

Examples:
* Cassandra
* HBase
* Accumulo
* BigTable

## WareHouses

These are designed more for queries about trends over the entire dataset.  The 
data is structured and pre-processed.  Often the data is stored in a column-oriented
format like Parquet.  Avro is row oriented so woudln't be used.

Examples:
* Redshift
* SnowFlake

## Eventually Consistent

These databases do not guarantee read operations will be accurate.  Usually this is 
so they can respond faster.

Examples:
* Cassandra

They tend to be used when the consequences of stale information are low.

## Data Lakes

These are more like very large file systems in the cloud.  They are also called
"Object Stores".  The data they hold *may* be raw, unprocessed.

Examples:
* S3
* Minio
* HDFS

# Recovery

Recovery Time Objective is the maximum amount of time that may elapse when an 
incident occurs before the system is operational again.

Recovery Point Objective is the maximum amount of time that may elapse where
data is lost.  It is like the interval that database backups are taken.

As each gets shorter, the cost to support it gets exponentially greater.

Achieving a short RTO means you have some kind of automated recovery, or
the engineers must be very fast to respond when paged.

Achiving a short RP means database backups must be taken more often, and 
stored in a very highly reliable data store, like a data lake.  When backed
up, they may be kept as deltas, change logs, write-ahead logs, or similar.
It's stored in a way that it can be recovered, but is not easily queried
directly.