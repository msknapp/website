---
title: Monitoring
weight: 15
draft: false
lastmod: 2023-03-22
summary: You should have ample means to investigate, troubleshoot, and observe your system.  You should be alerted if it is unhealthy.
---

A robust software system should have ample monitors in place.
Metrics should be exported to a separate system such as:
* DataDog
* NewRelic
* Prometheus and Grafana
* OpenTSDB

# System Monitors

Metrics should be collected for basic system metrics like CPU, memory, network traffic, 
and disk utilization.  This applies to nodes, virtual machines, pods, and containers 
in the cluster.

# Application Events

These are more like primitives for monitoring, meaning they do not hold aggregate data 
for time periods.  They could also be called transactions.  An example of an application event is:
* A web server receives an HTTP request.
* A data processing server reads a record from a streaming service like Kafka, SQS, or Kinesis.
* A Kubernetes operator synchronizes a resource.

## Tags or Attributes

Typically during an application event, you want to capture certain details.  For example, the 
request method is "GET" or "POST", the response time was 100 milliseconds, or that the user 
calling the service is "David".  These are attributes or tags that are associated with the event.

Application events should include ample tags or attributes so that troubleshooting later is possible.
Nothing is more annoying than when a user wants to know why something failed, but you just cannot
separate that data in your dashboards, and hence it is impossible to answer them.

One common practice is to combine attributes into one, for example, "cluster" could be the key, 
and it is set to "qa-east".  Here the value is representing two facts simultaneously.  I generally
discourage this practice, because it can complicate how we write queries later.  If you want
to see metrics on all of the QA clusters together, writing that as a query is more complicated now.

For some metrics systems, the number of keys and values may be limited.  So for example, having 
a key for request ID would be a bad idea, it would always be unique and certainly have more than
a few hundred values.  In this case, either make that value a metric by itself, assign it to 
a bucket, or don't include that attribute.

## Suggested Tags

When it comes to tagging, you should ask yourself how you might troubleshoot things later, what
fields would you want to filter on?  Here are some recommended attributes:
* response code
* environment
* region
* availability zone
* user ID
* host names
* query parameters
* input or path parameters
* container name
* version of the application server
* values of any critical variables

# Application Performance Metrics

Application performance metrics are aggregates of events that happened over an interval of time.
Most time series databases have some logic to combine these gradually as metrics get older.

One common practice is to have application servers expose an endpoint for metrics to be 
collected from, in the Prometheus format.  This is a recommended practice, because these 
metrics could be used for scaling later.

# Traces

For critical systems in production, being able to determine how long each transaction took 
can be crucial for troubleshooting.  This is called tracing.  HTTP requests are a perfect
example, if there are three steps in handling the request, the trace would capture how long
each step took, and it would display this information in a burndown chart.

Traces are vital when it is discovered that a particular event is taking a very long time
to execute.  You need to be able to identify the culprit.  Without traces you may waste
time trying to solve the wrong problem.

# Logs

An old-school practice is to have your application write logs to a file on the disk of the 
current server.  The problem with this pattern is:
* Now you must have a hard disk on the server.
* What if that server crashes?  You would lose the logs.
* It can run out of disk space, which leads to errors and crashing servers.

Libraries like log4j would even manage zipping and rotating these log files for users.  That
is truly not part of the main responsibility of that server, but it is consuming CPU cycles
none the less, and hence jeopardizing performance of its core functions.  Having applications
write logs to file should not be done.

The better practice is to ship them off to an external service.  This aligns with a 12-factor
app principle, that logs should be treated as streams of events.  Typically applications just
write to standard output now, and that is forwarded to a service like FluentD.  FluentD
is able to add tags to logs, and filter them too, but within limits.  It is able to down-sample
the logs as well, but I would only consider that if the number of log records is getting 
very high and impacting the performance of FluentD, Splunk, or ELK.

Generally you should add ample tags or attributes to your log records, as if they were metrics.
Hence logs are often written as json records, so that they can easily be ingested, indexed, 
processed, and queried on later.
Usually there is no limit on the number of possible values of an attribute key, but there may
be a limit to the number of attributes.

Examples of log aggregation services are ELK and Splunk.

# Alerting

Critical systems in production should have alerts created for them.  These might send text
messages to engineers, slack messages, emails, or even phone calls.  Alerts might be sent
to PagerDuty, so that the person on-call can rotate on a schedule.

Ambitious new engineers might put alerts on everything that might be wrong.  The reasoning 
is that they will treat metrics and alerts as their system's "validation".  There are 
several problems with this practice:
* Alert fatigue is a serious problem.  If engineers are getting paged frequently, then eventually they
  will just ignore the alerts.  If you are going to send an alert, make sure that *it is a serious issue*.
* As engineers develop changes, the expected value of metrics changes too.  This is especially 
  true during maintenance windows.

My advice is: be aggressive with dashboard creation, but conservative with alerting.
Sometimes more custom logic is used too:
* For the QA environment, don't send alerts on nights or weekends.
* Alerts could have priorities assigned to them.
* For certain problems, only send an email or slack message.
