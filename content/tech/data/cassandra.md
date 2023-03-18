---
title: "Cassandra"
draft: false
weight: 1
summary: Notes on Cassandra
description: Notes on Cassandra
lastmod: 2023-03-14
date: 2023-03-14
---

start cassandra:
$ cassandra -f
the -f means print logs to the foreground.

type 'cqlsh' to open the cassandra shell.
$ cqlsh

in the shell:
> HELP
> HELP <COMMAND>
> DESCRIBE CLUSTER;
> DESCRIBE KEYSPACES;
> SHOW VERSION;


no need for user, password, etc.




the book is wrong about commit_log_directory, that is not a yaml property.

uses port 9160
