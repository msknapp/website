---
title: API Development
draft: false
weight: 79
lastmod: 2023-03-29
summary: Writing effective and maintainable APIs.
---

API servers give engineering teams total control over a service, and are widely re-usable,
but making them right can be a challenge.  This is a terse summary of the best practices for APIs.

# Choose the Right Method

Endpoints should use the appropriate HTTP method:
* GET:
  * The request only retrieves information but does not alter data.
  * The request parameters do not contain any sensitive information.
* POST:
  * The request is *creating* data, or
  * It is getting data, but the input contains sensitive information, like passwords or tokens.  or
  * The request should be callable from an html form, and may alter the backend.
* PUT: for requests that update existing information by submitting the whole new resource.
* PATCH: for requests that update existing information by submitting the delta or change to apply to the existing resource.
* DELETE: For requests to delete information.
* OPTIONS: To inform clients, such as browsers, what methods are supported.  This is needed for CORS.

If you want the API endpoint to be invokable from an HTML form, then you must choose either GET or POST as the method.
Otherwise, some amount of javascript will be needed to call it from a web page.  Only GET methods can be 
bookmarked in browsers.

It is wise to support PATCH because:
* If clients use PATCH, they may not need to be updated when the server's schema changes.
  * Otherwise, updating the schema used on the server (like adding a new field), means also
    updating all clients.  That may be difficult, if not impossible, to enforce.
* It can reduce the amount of network traffic from client to server, by reducing the payload size.
* It reduces the risk of altering fields that the client never meant to alter.

# Use the Right Parameter Type

For API endpoints, inputs can be provided by three different types of parameters.

## Query Parameters

These are supplied after the URL path, after a "?", and keys are separated from values using an equals sign.
* These can be saved in browser bookmarks.
* These can be copied/pasted and shared between multiple users.
* These are usually only used with GET requests.
* Usually these should not alter the response schema.
* These can be easily controlled from html forms.
* These can accept multiple values, like lists or maps of data.

Logically, query parameters should not change what resource is being accessed, they just alter
the view of that resource.  If a web-page is being written, and a form is being created, the 
form fields map to query parameters.  So a good rule of thumb is to use query parameters for
variables that the end user might supply.

Examples of use:
* Filter a collection by a certain attribute.
* Only return results within a range, like results 20 to 40.  Support pagination.
* Filter by a time frame.

Wrong Ways to use query parameters:
* Provide the current user's ID.  Use the Authorization header instead, or a token.
* Provide a secret password.  Use the Authorization header instead, or a token.
* Alter the response format, like requesting text, csv, excel, pdf, or yaml in the response.
  * This should be accomplished using the "Accept" header instead.
* Change the resource type being retrieved.  Use a path parameter instead.
* Selecting a specific one of the items in the collection.  Use a path parameter instead.

Preferably, changing a query parameter's value should *not* change the response schema.

## Path Parameters

These are supplied in the URL's path, after the host name and before any "?" for query parameters.
* These *can* logically distinguish different resource types.
* These can be saved in browser bookmarks.
* These can be copied/pasted and shared between multiple users.
* These can be used with any request method.
* These can be easily controlled from html forms.
* These cannot represent multiple values, like slices, lists, or maps.

In HTML forms, the path can be set in the form's action field.  Hence it tends to represent
things that are less likely to change, or not selected by the end user directly.  They should
be used for variables that are constant for a single page or resource.

URL paths, including path parameters, should always be nouns.  Endpoints should use plural
nouns if multiple values could be returned.

Examples of use:
* Select a subset of resources.
* Select an individual item in a resource.

Wrong Ways to use path parameters:
* Provide the current user's ID.  Use the Authorization header instead, or a token.
* Provide a secret password.  Use the Authorization header instead, or a token.
* Selecting a range of responses, or pagination, like items 20 through 40.
* Filtering results by an attribute.

## Header Parameters

Header parameters:
* Are not part of URLs.
* Will not be part of bookmarks.
* Will not be shared if a user shares a URL with another user.
* Tend to be used for metadata, sensitive information, or to change connection behavior.
* These cannot be set by html forms.  Javascript would be needed to alter them.

Usually headers are used more by browsers, clients, and servers sitting between the client and backend.
They are usually not set by the end-user directly.

Common Uses:
* Authorization credentials
* Choosing a desired response format with "Accept".
* Specifying the input format with "Content-Type" for POST and PUT requests.
* Describing desired cache and connection behavior.
* Passing Cookies, which might contain configuration settings, and default values, specific to the client.
* Compression methods.

How you might use them:
* If you have a Gateway or reverse-proxy server, it could extract or derive metadata, and add that as a header for the backend.
* Delivering metadata for troubleshooting purposes, like a link to logs, or a request identifier.
* Additional sensitive metadata about the request.
* From an authorization micro-service, a header could be set to instruct the backend to filter or limit the results
  returned to the user.

How not to use header parameters:
* Selecting a different type of resource.
* Selecting pages of a response.
* Filtering results, unless an authorization micro-service is adding that header to limit what the user sees.

# Use the Right Response Code

The response code returned gives fundamental clues to the client about what happened.
Avoid returning an error response with a 200 status code.

Most Used Status Codes:
* 200 - OK.  For get, put, patch, or delete requests.  For get requests, there should be a response body.
* 201 - Created.  For post requests to indicate success.
* 202 - Accepted.  The request is being processed asynchronously, the user should check later for its status.
  An identifier should be provided for them to poll later.  If that is lost, there should be some means to get 
  that identifier separately.
* 400 - Bad Request.  Something in the user's input was invalid.
* 401 - Not Authenticated.  The credentials were not provided, or were incorrect.
* 403 - Not Authorized.  The user was authenticated, but lacks permission to perform the operation.
* 404 - Not Found.  The server does not know about the resource requested.
* 409 - Conflict.  There is some conflict with the request and the current backend state.
* 500 - Internal Server Failure.  The request was received but the backend server had an error processing it.

You are unlikely to add this into your own API, but are likely to encounter this:
* 503 - Server Unavailable.  A gateway received the request, but the backend server is not running.

# Support Different Formats

## Content Type

With POST and PUT requests, it is useful if different content types are supported.  For example, supporting
json, yaml, and form-data are recommended.

While it might be easiest for an API developer to support json or yaml, that is not easily used by
web page developers and forms.  The HTML form is limited in terms of what content types it can post.

Forms can only submit the following content types:
* application/x-www-form-urlencoded
* multipart/form-data - for submitting files
* text/plain

If your API endpoint does not support one of those, then it could be more difficult for web page
developers to use your endpoint, because they will need to use some custom javascript to invoke it.
For this reason, I recommend supporting the form url encoded content type, in addition to json.

You should generally **design your APIs so that web pages can use them.**

## Accept Type

It is useful if your API supports multiple output formats.  JSON is the most easy to work with, so
should be supported.  Returning yaml, csv, or plain text might also be useful.  You might even 
consider xml, excel, pdf, etc.  This enables a variety of end users, with different levels
of experience, to all use your API.

# Schemas

Your API endpoints should have well defined schemas for their inputs and outputs.  One challenge
with APIs, is keeping them backwards compatible.  Adding new fields is usually safe, but 
removing them involves risk.  You don't want to break your clients.  That is why I 
suggest that you should not add fields unless certain they are:
* needed
* ready for use
* going to be supported for the long term.

You should be conservative with adding new fields to schemas, avoid adding more than necessary.
Keep your schemas concise, simple, and cohesive.  Avoid mixing multiple different concerns
into a single resource or endpoint, it can get difficult to make changes to that later.
Don't be afraid to create separate endpoints for separate concerns.  It is easier to 
maintain many simple endpoints than a few complex ones.

The fields in your schemas should not include composites, where one field is capturing
multiple separable bits of information.  When you include composites like this, it means
that client and servers need some business logic to parse them.  The conventions used
to make the composite can change over time, and that would result in breaking business
logic in clients and servers.  For this reason, you should keep inseparable information
as distinct fields in your schemas.

Field names in your schemas should not be abbreviated, because you want them to be as
easily understood by end users as possible.  Follow a consistent convention for how 
they are named, like camelCase, snake_case, or kebab-case.  Likewise, use plural
field names consistently.

The output schema for a single endpoint may change depending on the response code.  If
there is a failure, a standard json response is returned by convention.
The industry has differing opinions on what that output format is, but usually it
includes an "error" field, a "message", and a "code" somewhere.  For example:

```
{
  "error": {
    "code": 404,
    "message": "foo bar"
  }
}
```

Often times the code matches the http response code, but I don't see the value in this.
Knowing the code should assist the support staff in finding the root cause of the issue.

# Backwards Compatibility

If you decide to alter your schemas in a way that is not backwards compatible, then your
API should support both formats for a period of days, weeks, months, or even years, 
depending on how many users you have.  You can support multiple versions by a few methods:
* Have a separate path in your API.
  * For example, /api/v1 and /api/v2 can both be supported.
* Use header parameters to select the API version.

I recommend the first option more, having a separate path, because it is very easy 
for end users and clients to adopt the new version at their own pace.  They are not
forced to quickly switch to the new API.

You should only remove older versions of API endpoints after the newer stable ones
have been proven stable for a long period, like months.

You could add a response header to the older versions that delivers a notification
to the end users, that the endpoint should change.

Separate paths for new versions are needed if the input or output schemas change.
If the schemas do not change, a header to select the version can be supported.
The advantage of headers is the API owners can default users onto newer or older
versions of APIs.  The end users would have to elect a different version on their own.
Headers could be more useful for canary endpoints.

# Documentation

Your API should provide swagger documentation, or something similar.
* Endpoints should have descriptions about what they do.
* Edge cases should be documented.  Will fields be empty, omitted, etc.  Can they be?
* All response codes possible should be listed, and what conditions leads to them.
* Input parameters should be documented, along with acceptable values for them.
* Document if fields are mandatory or not.
* Document default values used for fields if they are not provided.
* Document input and output schemas, and what each field does in them.
* Document expected response time, SLAs, maximum payload sizes, and the maximum
  transactions per second that will be supported.

# Choice of Identifiers

Your choice of resource identifiers can have a serious impact on the long term 
maintainability of your API.

Using a monotonically increasing number for an ID is risky, attackers might be 
able to determine the next ID and use that as part of an attack.  Also it 
can be difficult for separate servers to coordinate this global counter.  Even
letting the backend database choose the ID by this pattern could limit its
ability to be partitioned or scale up.

Using composites of other fields is also risky, for several reasons:
* The convention you use to make these identifiers might change over time.
  Then you would need to manage some kind of migration script to update all
  the existing data.  That can be risky.
* The combination of other fields might exceed constraints on variable length.
* You might need to write custom logic to parse fields out of the composite, and
  that is error prone.  The logic might need to be shared among several applications.
  If it is wrong, then many will need to be simultaneously updated.
* This pattern often forces the sub-fields to be further restricted in terms of
  acceptable characters.  For example, if a hyphen "-" is used to delimit fields
  in the identifier, then that means the sub-fields cannot contain hyphens
  themselves.
* In a NoSQL store, composite IDs like this might result in having single servers
  take on far more traffic or workloads than their peers.  This leads to 
  performance bottlenecks.

Using a hash of fields is slightly better:
* Now you don't have to worry about exceeding the length limits of the identifier.
* If storing data in a NoSQL store, it will naturally be load balanced.
* This naturally prevents and avoids having duplicate records.

However, it still has an issue:
* The fields used in the hash might need to be changed, then the hash would be
  inconsistent with the record.

For these reasons, it is often best to use some large random identifier, like a
UUID, for records.  It will be the least difficult to maintain over time.  Views
can be created for engineering teams as needed.

The last problem here is that you might want to prevent having duplicates.  You
should block duplicates by a separate means.  You can have constraints in your
database server that would block it.  The unique fields will need to be indexed.
For Kubernetes, you can leverage an admission controller like GateKeeper or
Kyverno (recommended) to prevent duplication.

# Separate Static Content

There seem to be two approaches to hosting websites:
* Having all static content.
* Having a web server using the MVC pattern.

Normally, a static website could be difficult to maintain, because there will be 
a lot of duplication between pages.  Updating the view would mean manually 
updating dozens or even hundreds of pages.  However, there are a lot of programs
out there that can generate static websites from templates now.  This website
is generated using Hugo.

You might think that a static website *cannot* perform dynamic operations.  That
is certainly not true.  First, javascript sources are still static files.
Secondly, there can still be an API server, just hosted separately.

The problem with having a web server using MVC, is that *now you must be running that server*.
The server will need some kind of virtual machine (EC2) to host it, and that can
get expensive.  It requires resources that get billed by the hour.

Static websites can often be hosted in a serverless manner.  For example, they
can just be placed into S3.  If they are never accessed, the website owners only 
pay a few meager cents per month to store it.  You can still embed *some*
dynamic operations in web assembly or javascript.  Where you are limited is
for data that needs to be stored on the backend for later, or shared between users.

When a static website needs operations only a server can support, you don't
necessarily need a full server, you could leverage a managed API Gateway,
and serverless lambda functions to support that.  This will cost far less than
running your own server, like with EC2 instances.  You should only really
consider running your own server if the level of requests will be considerably
high, to the point where lambdas would be invoked thousands of times per hour.

Any operations that *can* be *safely* run client side should be, so the server
is used as little as possible.  Then you might get away with a cheap serverless
option, or with fewer running servers.

One down-side to serverless solutions, is they might be difficult to switch
to a different cloud provider.

Using a static website with a separate API can be one of the most affordable
and scalable options.  You just need to make sure the API supports CORS, so 
the static site can use it.

Another advantage of static sites is that cache controls often come by default.
You don't need to program your server to set any cache-control headers.

I generally recommend using static sites as much as possible.  Consequently,
**I recommend always supporting CORS and OPTIONS requests in APIs.**

# Monitors

Like any application, APIs should be easily monitored.
* Log events, and include ample tags on parameters.
* Trace responses.
* Export metrics, along with ample tags.

# Support Asynchronous Operations

For any operation that might take over four to five seconds, it is ideal if the web API server
is not processing that locally.  It could make it difficult to scale the server up or down.
Instead, it would be wise to publish the event to a separate stream.  For instance, the server
could write an event to Kafka, Kinesis, SQS, etc.  Some managed databases like DynamoDB also
permit lambdas to run after a record is written, which could in turn process that record, 
or queue up a separate job to process it.

Don't have the web server handle long requests locally, that can result in slow responses and
scalability issues.  Having separate processors means they can be scaled on their own, or 
updated independently.

Returning asynchronous results comes in two forms:
* Return a 202.  Give an ID for them to retrieve from.  Clients will need to poll for results periodically.
* Utilize web sockets.

Using a 202 response means the clients must poll for results.  This might impact the scalability
of the server.  If there are thousands of clients all polling for results simultaneously, that can
bring the API server to a halt, or cause it to scale so much it is exceedingly expensive.

When there are potentially many clients that could poll for results simultaneously, web sockets
might work better.  The server can push results to the client when they are ready.  These also
enable more prompt results when the client needs them.

# Support Troubleshooting

When a request fails, there should be information delivered to the end user about why it failed.
The only exception is for unauthorized requests, you should not inform the user about what 
permission they lacked, if the user ID was correct, etc.

You might return information like a request ID, an error code, error message, or links to a 
log query, or links to metrics/traces that can be viewed.  If the request was invalid, the error
message should indicate why.  A common practice I have used is to return a link to relevant
Splunk logs in a response header.

Avoid returning stack traces to end users, it can reveal information that an attacker might
be able to take advantage of.

# Security

* Use mtls
* Use pop tokens.
* Use only trusted headers for sensitive decisions.
  * PoP tokens indicate which headers have been signed.
* Don't expose sensitive information that might help attackers.
* Sanitize inputs, do not trust them.
* Keep authorization logic separate.
* Leverage tokens for improved security.
* Block DDoS attacks.
* Don't permit selection of variables they should not control.
  * Avoid path or query parameters that a user should not be able to set or change.

# Performance

* Support GRPC, it can cut response time in half.
* Consider cache controls.
* Consider using a L1 and/or L2 cache.  Monitor cache-hit and miss rates.
  * Avoid caching items that are not requested repeatedly.
* Consider connection keep-alives
* limit input sizes
* Rate Limit clients.
* Keep stateless, except maybe for a cache.

# Other Recommendations
* Support standard Headers
* Don't combine everything into one monolithic server.  Keep scope narrow and cohesive.
* Use middleware.
* have a health check endpoint.
* Support concurrency.
* Support GraphQL
* Don't let your paths prevent future additions to the API.
* Return troubleshooting information in the event of failures.
* avoid having multiple ways to do the same thing.
* Use a Gateway and Load Balancer.
  * Leverage URL path re-writing