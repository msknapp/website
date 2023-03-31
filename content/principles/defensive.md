---
title: Defensive Programming
description: Making your code defensive.
summary: Carefully consider the trade-offs of defensive programming patterns.
draft: false
weight: 3
---

Defensive programming is a design pattern where each component in software acts very defensive in nature.
For example:
* Fields become private.
* Classes are often made immutable.
* Functions validate all of their inputs very strictly.
* Code "fails fast" if a problem is encountered.

It might also lead to adopting certain code patterns such as:
* The Singleton.
* The Builder.
* The Factory class.

I would like to argue, from my experience, that:
* Defensive programming is more or less useful depending on the programming language.
* Your number one defense against bad code should be TESTS!  Not defensive coding patterns.

# Immutability

An immutable class is one whose fields cannot be updated once it has been created.  They 
tend to be far safer to use in code, because they basically cannot be corrupted.  Also they
are naturally thread safe, multiple threads can refer to them because none of them can 
alter the object.  Programs that use immutable objects tend to have fewer bugs.  You 
might even say they need less testing.

Immutability adds so much safety to programs that some languages make everything immutable (Scala),
or make it a challenge to let things be mutable (Rust).

The downside of immutability is it tends to complicate some parts of the code, and it 
often results in more objects being created and destroyed than necessary.  That means
garbage cleanup needs to run more often, which can degrade the application's performance.

# Fluent Pattern

The fluent pattern is not a defensive programming pattern, it is included here to draw its 
similarities to others.  Fluent classes are written in a way that you can chain together 
functions that set fields on the class.  The class is mutable.

Here is an example of a class following the fluent pattern in golang:

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;"fmt"<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;MyFluentThing&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Value&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Text&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewFluentThing</span><span class="golang-brace">()</span>&nbsp;*MyFluentThing&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&MyFluentThing<span class="golang-brace">{}</span><br>
<span class="golang-brace">}</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*MyFluentThing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">WithValue</span><span class="golang-brace">(</span>v&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;*MyFluentThing&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;t.Value&nbsp;=&nbsp;v<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*MyFluentThing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">WithText</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;*MyFluentThing&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;t.Text&nbsp;=&nbsp;s<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Run</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;:=&nbsp;<span class="golang-function">NewFluentThing</span><span class="golang-brace">()</span>.<span class="golang-function">WithText</span><span class="golang-brace">(</span>"foo"<span class="golang-brace">)</span>.<span class="golang-function">WithValue</span><span class="golang-brace">(</span><span class="golang-primitive-value">3</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%s,&nbsp;%d",&nbsp;x.Text,&nbsp;x.Value<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

# Builder Pattern

The builder pattern involves having one class that is mutable, and its purpose is to build 
a different immutable class.  Usually the builder has a no-argument constructor, fluent 
methods to set values, and a build method.

There are two ways to write a builder:
* The builder has duplicated fields for all or most of what the final class has.
  This is less likely to produce a class that can be mutated accidentally, but it also 
  makes it harder to maintain because fields have been duplicated.  This method is thread safe.
* The builder holds its own private instance of the final class, which only it can modify.
  The problem with this approach is you must be careful to not permit updating that 
  instance once it has been built.  The builder is suddently not thread safe.

In practice, builders may or may not be re-used.  The downsides of using builders are:
* They add more complexity to your code, and hence more opportunities for bugs to creep in,
  and a tougher time for co-workers to understand it.
* They must also be unit tested.
* They create more objects, which can add strain to garbage collectors, and hence it 
  requires more memory and can slow down performance.

Here is an example of the builder pattern in golang:

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;"errors"<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Thing&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;text&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;number&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;array&nbsp;&nbsp;[]<span class="golang-variable-type">float64</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*Thing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Text</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t.text<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*Thing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Number</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t.number<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;ThingBuilder&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;text&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;number&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;array&nbsp;&nbsp;[]<span class="golang-variable-type">float64</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">StartSomething</span><span class="golang-brace">()</span>&nbsp;*ThingBuilder&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&ThingBuilder<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;array:&nbsp;<span class="golang-function">make</span><span class="golang-brace">(</span>[]<span class="golang-variable-type">float64</span>,&nbsp;<span class="golang-primitive-value">3</span><span class="golang-brace">)</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>b&nbsp;*ThingBuilder<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Text</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;*ThingBuilder&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;b.text&nbsp;=&nbsp;s<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;b<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>b&nbsp;*ThingBuilder<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Number</span><span class="golang-brace">(</span>x&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;*ThingBuilder&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;b.number&nbsp;=&nbsp;x<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;b<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>b&nbsp;*ThingBuilder<span class="golang-brace">)</span>&nbsp;<span class="golang-function">AddValue</span><span class="golang-brace">(</span>x&nbsp;<span class="golang-variable-type">float64</span><span class="golang-brace">)</span>&nbsp;*ThingBuilder&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;b.array&nbsp;=&nbsp;<span class="golang-function">append</span><span class="golang-brace">(</span>b.array,&nbsp;x<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;b<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>b&nbsp;*ThingBuilder<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Build</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">(</span>*Thing,&nbsp;error<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;b.text&nbsp;==&nbsp;""&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;nil,&nbsp;<span class="golang-function">errors.New</span><span class="golang-brace">(</span>"text&nbsp;is&nbsp;not&nbsp;set"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;b.number&nbsp;<&nbsp;<span class="golang-primitive-value">1</span>&nbsp;<span class="golang-brace">{</span><br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;example&nbsp;of&nbsp;setting&nbsp;a&nbsp;default.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;b.number&nbsp;=&nbsp;<span class="golang-primitive-value">1</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;defensive&nbsp;copies&nbsp;of&nbsp;slices.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;c&nbsp;:=&nbsp;<span class="golang-function">make</span><span class="golang-brace">(</span>[]<span class="golang-variable-type">float64</span>,&nbsp;<span class="golang-function">len</span><span class="golang-brace">(</span>b.array<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">copy</span><span class="golang-brace">(</span>c,&nbsp;b.array<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Thing<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;text:&nbsp;&nbsp;&nbsp;b.text,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;number:&nbsp;b.number,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;array:&nbsp;&nbsp;c,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>,&nbsp;nil<br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

In the code above, we can see that "Thing" is immutable to any code outside its package.  This is 
accomplished by making its fields have lowercase first characters, meaning they are package private.
It exposes public methods to get its values, but not to mutate them.

The builder has mutable fields, but once it has built the "Thing", any update to the builder 
will not alter the instance.  The slice it holds for the "array" is cloned, which is called
a defensive copy.  The build method can also validate inputs or assign defaults.  In the example,
the validation is that the text is not empty, and the default is that the number will take
the value 1 if it was not set.

Calling code would use this to create the "Thing":

```
instance, err := golang.StartSomething().Text("foo").Number(8).AddValue(2.3).AddValue(8.4).Build()
```

In golang, the builder pattern is not needed nearly as often as it is in other languages, because
Golang lets you pass structs by value.  When a struct is passed by value, all of its fields are
copied to the function being called.  If the function changes something, it doesn't affect the
value seen by higher functions.  So then the only reason you might use a builder is if the 
struct has *many* fields in it.

# Factory Pattern

The factory pattern is very similar to the builder pattern, but with a few differences:
* Factorys tend to be re-used more than builders, and are longer lived.  Builders tend
  to be ephemeral.  Consequently, factory's lend themselves to application bootstrapping, 
  while builders are more commonly used at request time.
* Once the factory is configured, it usually does not change its values.
* Typically the factory's configuration is held in static variables (in java).

Factory's also have similar downsides as builders:
* Added complexity
* More required unit testing
* If the factory uses globals (static variables), that is more risky and unsafe.

If you are considering a pattern listed above, you should weigh the pros and cons, and consider
the problem you have.

# Immutable Constructors

It is worth calling out here, that classes can be immutable without being produced by 
a builder or factory.  They can simply have a constructor that sets their fields, and
no other methods are able to mutate them.

The only downside is that some immutable objects have many fields.  Somewhere between 
five and seven arguments to a method, it starts to get too confusing and difficult 
to maintain.  After seven arguments, a fluent pattern, a builder, or a factory pattern
becomes easier to understand.

You might prefer immutability when:
* The object is used by many parallel processes or threads.
* The code is getting very complicated and you want to reduce risk.
* The lifecycle of the object is quite long, spanning many requests.

You might prefer mutability when:
* The object lifecycle is short lived and not shared by many functions.
* Memory and garbage collection are becoming a concern.
* You need your code to perform faster.
* You are willing to write more unit tests to be certain everything is safe.

