---
title: Avoid Globals
linkTitle: Globals
weight: 30
draft: false
lastmod: 2023-03-26
summary: Global variables create problems, you should avoid using them.
---

Global variables are associated with the current process.  If they are constants, 
like the value of Pi, that should be perfectly fine to keep as a global.  However,
new developers often make far more variables into globals than should be.  This
is due to their simplicity, because developers don't need to worry about passing
the variable around in function calls, or adding it as an object field.

# Assumptions

Globals imply a few assumptions:
* This variable will be the same for all code that uses it.
* The variable will not be updated by any code that has access to it.  Or, the 
  update will be appropriate for all code that uses it.
* If it is updated, then either one thread uses it, or it's updated in a thread
  safe/atomic manner.

Far more often than first expected, the above assumptions become wrong over time.
This is especially true for unit tests, because they often cover edge cases and 
mis-configuration.

Unit tests might be run in sequence or in parallel.  The order that they are run 
in should not be assumed to be constant.  Upgrading the test library or framework
could very easily change the order that they are run in.

So if one unit test updates a global, and is also responsible for setting it back
to its original value when it is done, that is a very unreliable design.  If 
other tests also rely on that global, it is likely that they will break 
periodically and randomly.  When this happens, the developer is often very 
confused and frustrated, because they didn't even change the code that is 
related to the failure.

Alternatively, the enginer might try adding some kind of global synchronization
lock to protect access to the global.  This makes the code extra complicated,
tougher to maintain, and also slower.  It would degrade the performance and 
introduce a bottleneck in the code.  It would be more confusing to other 
engineers.

# Global Services

An engineer might think, it is safe to use a global variable 

# Singletons

The singleton design pattern was very popular in the past.  It is effectively 
using a global service.  Usually that service is immutable once it is built, so
engineers might think it is safe.  However, it still makes it difficult to 
write and manage unit tests, which should test edge cases and different 
configurations of the service.  For this reason, I generally discourage the
use of singletons, because they are globals.

# Environment Variables

Environment variables are visible to an entire process, and mutable, hence they 
are effectively globals.  They suffer from all the problems with globals
listed above.  So I generally discourage using environment variables more
than necessary.

# When Safe

Globals are usually safe to use when:
* They are relatively primitive.  Strings, integers, floats, etc.  Might be 
  safe globals, while things like Clients or API Server classes should *not*
  be globals.
* They represent fundamental constants like the value of Pi.
* They would really never be mutable or updated at runtime, even in edge 
  case unit tests, or different clusters or environments.
* They are only used to avoid duplicating a string or key in your code, like
  if an error message or a log statement is using a format pattern.
* They are relatively private in nature, there are few objects that could 
  even observe the global.

# Advice

For all the above reasons, globals should generally be avoided.  If a class, service, or
object relies on global data, it is better to make that a field in the 
object, and have the value passed as a constructor argument.  You might be 
concerned that this will result in having many constructor arguments, and forcing you
into using a builder or factory.  Usually there is a way to logically group
the variables into fewer objects.  That might also imply that your service 
has too many concerns in its scope.