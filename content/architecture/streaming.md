---
title: Streaming
draft: false
weight: 4
lastmod: 2023-03-24
summary: Add stream processing for async needs.
---

New engineers might feel like throwing all processing into their monolithic web server.
This doesn't scale well:
* Web servers typically need to respond in a reasonable timeout.
  * Ideally four seconds should be the limit, but ten seconds may do.
  * If it takes more than ten seconds, that could impact the scaling operations of the app.
* The extra processing might prevent the server from responding promptly to other
  requests.  It is vulnerable to a denial of service attack because of this.

Hence, eventually we need to separate processing from the request, so that the 
web server's availability and response time are not jeopardized.  The server then
accepts the request and simply throws it onto a queue for later processing, by a
separate service.

That brings us to streaming, which acts as the queue for these asynchronous
transactions.  Examples are Kafka, Kinesis, SQS, JMS (RabbitMQ), etc.

# Loose Coupling

There is a second benefit of streaming, when you separate the data processing
from the web server, it becomes de-coupled.  You can update either one without
changing the other.  That makes the system easier to maintain.

They can also scale independently.  You could have the processors use more
CPU or memory than the web servers.

Lastly, you get an extra layer of security.  The processors can run with
tighter fire-walls and network controls than the web server.  For instance,
they might have no or reduced access to the internet.

# When to Use Streaming

Generally, you should consider stream processing when:
* A request could take more than a few seconds to process.
* These requests could happen frequently.
* There can be spikes in traffic, where many of these requests arrive in 
  a short time frame.
* You want to be able to update the processor logic separately from
  the web server or the queue.
* The web server and the transaction processing have different resource needs
  and scaling needs.
* You want stricter isolation of the processing for security reasons.