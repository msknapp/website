---
title: Batch Processing
draft: false
weight: 5
lastmod: 2023-03-24
summary: Add batch processing for big data.
nextlink: /algorithms/
---

When a company has large volumes of data, processing all of that on a single instance can
take an enormous amount of time.  You need to compare the size of the data with how
soon the results are needed.  Typically we should expect jobs to finish in a few hours.
You don't want to wait a whole day to discover nothing was processed right.

So we process data in parallel, this is batch processing.  Some examples
are Hadoop MapReduce, Spark, or AWS Glue.  A common use case is to support
"Extract, Transform, and Load" (ETL) jobs.