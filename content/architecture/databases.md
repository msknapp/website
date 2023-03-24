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

## Choosing Data Storage

When it comes time to choose what tech should serve as your database, consider
the following:
* How many requests do you expect over time?
* How much data is there to store?
* How fast of a response do you need?
* How will the data be queried?
* Is eventual consistency acceptable?
* Does your data have a strict schema?

### Number of Requests

If your data store will rarely receive requests, then a serverless solution will
probably serve your needs while costing far less than running a database instance.
In AWS, DynamoDB can satisfy the needs here, it is serverless.  If you have a 
Dynamo table with no data and no transactions, it costs nothing.

The only exception might be if you have strictly relational data and you want
transactional guarantees that it will be consistent.  Dynamo is a NoSQL store,
so it cannot perform transactions.  Keep in mind that in NoSQL stores can always
choose to just keep all information in a single record, even if it has some
redundancy.

### Data Volume

Relational databases tend to have the most features, but they may not scale as 
well as NoSQL stores.  Somewhere between 10GB and 100GB, you might want to consider
switching to a NoSQL datastore.  These can be sharded, partitioned, and hence
parallelized between multiple servers.

* data replication
  * RTO
  * RPO