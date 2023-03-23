---
title: Loose Coupling
draft: false
weight: 17
lastmod: 2023-03-23
summary: Microservices in the system should make few assumptions about peers, and have limited scopes.  They must be backwards compatible.
---

The linux motto is:

```
Do one thing, and do it very well.
```

Likewise, there is a software principle of "Segregation of Responsibilities" which can also be 
extended to microservices.

Essentially, system components should:
* Be responsible for one small and simple function, not many.
* Have well documented contracts for how they interact with external services.  For example, they should
  have a Swagger API with ample documentation about the inputs and outputs of each of their endpoints.
  They should have schemas that explain the input and output record formats.
* Keep their APIs and contracts as minimal as possible.  This aligns with the "keep it simple stupid"
  (KISS) principle of software development.  Schemas should not contain fields that are not used.
  Schemas of inputs and outputs should be small.  Service endpoints should be few, or labelled as 
  unstable, experimental, or beta endpoints.
* Make very few assumptions about how their peers are operating.  For example, they should not assume
  that a separate service is using a certain naming convention for resources it creates.
* Be backwards compatible with older versions of peer micro-services.  This is critical, because
  *you do not want to be forced to upgrade or rollback more than one service at a time in your system.*

# Monoliths

Often times when software systems are being developed, engineers tend to throw all functions into a
single API server.  That is because:
* It already compiles and has a CICD pipeline.
* They are already familiar with it.
* Setting up network connections and security between services is scary and confusing.
* They don't want to think about what should happen if one of them is unavailable.
* They think that making a separate application or service will over-complicate their cluster and 
  make it more difficult to maintain.

Eventually the source code of that API server is over 10k lines long, nobody understands all of it, 
and any kind of update is extremely risky.  If an engineer wants to make an update to one feature,
they are also jeopardizing a dozen others at the same time.  Incidents and roll-backs are bound
to occur.

Having worked in complex systems for years, I consider the notion that "more distinct services is
more difficult to maintain" as simply **wrong.**  In my experience, having more system components,
that are each simpler, tends to be far easier to maintain than a monolith.
That is because you can upgrade each individual component separately, and test them separately.
Go ahead and separate things, you won't regret it.

# Backwards Compatibility

We said earlier that applications should be backwards compatible, but what does that really mean?

* If an object schema has changed, the service endpoint should accept either the old or the 
  new schema.  It should not fail if it encounters an object using an older schema.
* Endpoints should not change the schema they return when an application is updated.
  Instead, endpoints should have a version in their path, so clients can start using the newer
  version when they are ready.  Endpoints should not be removed until it has been proven they
  are not needed for a long period, like months.
* If the service depends on an external API, it should be able to function with older versions
  of that external API.  That way, if the external service is rolled back, the current one
  does not need to also roll back.

If you find yourself in a situation where updating or rolling back one service or table forces 
you to also update or roll back others, then you have a "tightly coupled" system that is 
more difficult to maintain than it needs to be.

# How to Split Services

There are two ways that you might think about splitting up a service:
* By Function: You would split up a monolith into micro-services, where each is concerned with just
  one common functionality or resource.  For example, with a news site, one micro-service handles
  requests from customers to get news pages.  A separate micro-service could handle requests from 
  article publishers to create or update news articles.  Another example: One microservice 
  handles CRUD operations on table A, and another handles CRUD on table B.
* By Concern or Technology: You split the monolith into micro-services, where each deals with 
  a common concern among all endpoints, like querying, filtering, authorization, monitors, or 
  integration with a database.  For example, you could have one micro-service that handles all
  data queries, another that handles data updates, and yet another that handles data validation.

Generally speaking, it is better to split up a monolith by functionality, not technology.  There
could be an exception to this rule if the micro-services follow a configuration driven development
strategy, whereby they have no built-in schemas, and their logic is controlled by expressions in
configuration files.

# Re-Factoring a Monolith to Micro-Services

When a team decides that a monolith application needs to be split up, they are forced to consider
how.  If the team decides to re-create the entire monolith as distinct micro-services before 
switching:
* That involves a lot of risk.  Rolling back means shifting everything back to the monolith.
* That will take a huge deal of time.  Stakeholders will often become impatient and unsatisfied
  with the seemingly slow progress of the team.

Instead, it is better to split out just a few features at a time, in a more piece-wise manner.
Pick one or two vital functions in the monolith, write a new micro-service app for them, and 
then carefully integrate those into the system.  You will have more gradual improvements, 
which is in line with the Agile methodology.