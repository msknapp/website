---
title: "ZooKeeper"
weight: 3
draft: true
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

Start zookeeper:
$ cd /share/usr/zookeeper/bin
$ ./server.sh start

See if it's running:
$ jps
michael@michael-GA-890GPA-UD3H /share/usr/zookeeper/bin $ jps
14753 QuorumPeerMain
14773 Jps

The "QuorumPeerMain" is zookeeper.

It writes logs to /share/usr/zookeeper/bin/zookeeper.out
The zk config is in /share/usr/zookeeper/conf/zoo.cfg

The zoo.cfg file specifies a directory for zookeeper to use:
dataDir=/share/var/zookeeper

so we are writing to /share/var/zookeeper

The default port it listens to is 2181
