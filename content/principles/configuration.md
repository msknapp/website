---
title: Configuration Levels
draft: false
weight: 9
lastmod: 2023-03-24
summary: You should have ample controls over your code from configuration using expressions, templates, and scripts.
---

This is not a principle often taught in software engineering, it is just something I learned through
experience.  An engineers performance is often judged by how quickly they can deliver results, 
and that in turn will depend a lot on how easy it is for the code to be updated.  One way for 
code to be easy to maintain and update is with loose coupling.  Another is to add ample controls
over how it operates from its configuration, which is what this page discusses.

There is a wide spectrum of control that a team could have over their applications.  Let's go through them.

# Hard Coded

When an application is first written, critical variables are frequently coded right into the spot they are used.
For example, a value could be written inline at the very function it is used.  The situation gets slightly better
as they move this variable to more visible places, like to a class's arguments, or to a file holding constants.
But in any case, the value cannot change unless the program is re-built.  This is really tough to maintain
going forward.

# Environment Variables

One step up, you get a little more control when a variable is set by an environment variable.  At least it can
be changed now without re-building the code, but it still has some challenges.

First, environment variables are globals.  Generally speaking, global variables should be avoided in 
software engineering.  They are difficult to test with because they are mutable, and the order that tests
run in is not guaranteed.

Another problem with environment variables is they are often not documented for users.

# Command Line Arguments

Passing a variable in as a command line argument has at least some advantages over environment variables.
For one thing, it is usually part of the help documentation presented by the command line, so other
engineers will be aware of it.  Also, in the code it is not a global, so you don't need to worry about
it being modified by different places or threads in the code.

It still lacks some control though.  Command line arguments can get quite verbose as they are more 
heavily used.  The values would show up in a "ps" command, and maybe they shouldn't really.  The value
cannot change over the life of the process, and you might want to change it.  It can also be 
difficult to see the whole way the application is configured when command line arguments are used.

# Configuration Files

Using a configuration file brings some pros and cons along with it.

Pros:
* You can update this file while the program is running, and depending on how you wrote the application, it could pick up that change promptly.
* It is easy to see the whole configuration of the application in one place.
* The values are not shown on the "ps" output.

Cons:
* Now this file must be written, and in the correct place on the file system.  The runtime user must be granted permission to read it.
* If it is included in a docker image, that can make it very difficult to change the configuration.  I recommend mounting the configuration into the runtime container.
* In Kubernetes, you would need to write a ConfigMap, and have that properly mounted, for the container and application to use this configuration.

So we have gradually gained more control over this variable with each option.  Still, we basically have a *constant*
variable, that cannot fundamentally change the business logic of the application.  It grants us somewhat
limited control over things.

# Expressions

To gain more control, an expression can be used in the configuration instead of a hard coded value.
There are often a variety of expression languages that can be supported or embedded into configuration.
For Golang, "expr" is a very robust option, but there is also "gval" and others.  Regex can also count.

Switching to expressions means that the value of a variable can depend on the current request being processed.
Another huge benefit is that it can easily be changed without ever re-building the application.

There are some downsides to using expressions:
* Now you have a possible runtime failure instead of compile time.
* Fellow engineers might be unfamiliar with the expression's syntax.
  * They may be unable to write new expressions or explain what it means.
* They may not perform as fast as constants or other options.

It's important to note that, despite not needing to re-build the application, these expressions
*do need to be tested!*

# Templates

Expressions gave us a lot of control over how values can be selected for each request, but they can 
be unwieldly when you have entire objects to define.  Templates give you enough control
to specify how entire objects are built for each request.  So they give you a huge deal more
control than just expressions.  They have similar down-sides as expressions.  They could
also have slow performance, you should weigh that against the needs of the current function.

# Embedded Scripts

Templates give us a ton of control, but they are not exactly turing complete languages.  There
are complex functions we might want to add that templates just cannot help with.

So to gain the utmost control over business logic, applications could add support for embedded
scripting languages.  Lua is a great example, your application configuration could refer to a 
Lua script that might be invoked for each request.  This would make it possible for even 
very high level business logic to change without ever re-building the application.

This practice is related to the Open/Closed principle, which states that software components
should be open to extension, and closed to modification.  In this case, the service permits
extension via the embedded scripting language.

There are some extra down-sides to supporting embedded scripts:
* These tend to perform much more slowly than application source code.  You will need to 
  consider the performance requirements of the service here.
* They also need to be unit tested.
* It is another case of something that might fail at runtime instead of compile time.
* They can be complicated and unfamiliar to other engineers.

Another major drawback and concern of embedded scripts, is that corporate governance might
have an issue with them.  Some companies only permit certain languages be used, because
they have security scanners for them.  It might be impossible to run any security scan
on these scripts.  You could argue that it's actually the lower level language running it,
but who knows if that will be accepted.

# Configurable Endpoints

Similar to the principles above, your service's high level endpoints should be completely
changeable from configuration.  You should be able to add or remove any number of new
endpoints to your service from configuration alone.

This might also be called "configuration driven development", but I'm unsure if it
is the same thing.

# Why Is This Needed?

You might feel like, what's wrong with just updating the source code, re-building, and re-deploying?
There are a few issues:
* Often times your CICD process can take tens of minutes to complete.  If you wait all of that 
  time to discover a failure, that is a very slow iteration cycle.  You will certainly not
  be very productive if you need to wait 20 minutes to test every change you make.
* Peer reviews on pull requests might be hard to get.  Fellow engineers are also working
  on their own stories, and don't want to be distracted trying to understand your
  unfamiliar code.
* Re-deploying can be much more risky than just changing configuration.  Rollback of 
  configuration is likewise easier than rolling back the application version.

Following this pattern of control from configuration means you can develop new
features and changes, and test them, much more rapidly than before.