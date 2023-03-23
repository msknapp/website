---
title: Boot and Shutdown Times
draft: false
weight: 10
lastmod: 2023-03-23
summary: Applications should boot and shutdown promptly for scaling purposes.
---

In a distrubuted, auto-scaling cluster, it is critical that applications can boot
promptly, so that services can automatically scale up and start serving requests
quickly.  Ideally they should boot within ten seconds.  This is one of the 12-factor
app principles.

Likewise, if the application is sent the kill signal, it should be able to shut
down promptly.  Ten seconds is a good rule of thumb here too.  If they do not, 
they could be consuming cluster resources, and blocking other applications from
being scheduled on the node.  This could delay cluster maintenance or scaling of
other services.

When dealing with languages like Java or Python, and frameworks like Spring, 
Hibernate, and Anaconda, it is not unusual to have hundreds of dependencies, 
and boot times measured in minutes.  These often create problems in clusters.
That is one reason lower level system languages are sometimes preferred, they
are more light-weight and even complicated apps tend to boot in under ten seconds.