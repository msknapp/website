---
title: Single Responsibility Principle
draft: false
weight: 4
summary: Software components and classes should not be responsible for many concerns.
description: The importance of single responsibility.
---

The Unix motto is "Do one thing, and do it very well."  This is quite similar to the
single responsibility principle, which suggests that any single class
should not be responsible for multiple things.  In other words, we want to keep the
scope of a class small and simple.  So it is in line with the KISS principle.  The
principle of cohesion is also relevant to this.

# High Cohesion

Code is cohesive when its functions are all closely related and aligned to a common purpose
or objective.  It is similar to the single responsibility principle.  It helps you keep the 
scope of classes small, so they are rarely updated (in line with the open/closed principle),
and far less likely to have bugs.  This pattern also means your team is less likely to 
have merge conflicts in your code.

# Example

Let's say that an engineer wrote this class:
{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"errors"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"os"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Service&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;admin&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;factor&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;logTo&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;units&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewService</span><span class="golang-brace">(</span>admin&nbsp;<span class="golang-variable-type">string</span>,&nbsp;factor&nbsp;<span class="golang-variable-type">int</span>,&nbsp;logTo&nbsp;<span class="golang-variable-type">string</span>,&nbsp;units&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;*Service&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Service<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;admin:&nbsp;&nbsp;admin,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;factor:&nbsp;factor,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;logTo:&nbsp;&nbsp;logTo,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>s&nbsp;*Service<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>user&nbsp;<span class="golang-variable-type">string</span>,&nbsp;input&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">(</span><span class="golang-variable-type">string</span>,&nbsp;error<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;user&nbsp;!=&nbsp;s.admin&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;"",&nbsp;<span class="golang-function">errors.New</span><span class="golang-brace">(</span>"unauthorized"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;f,&nbsp;_&nbsp;:=&nbsp;<span class="golang-function">os.Open</span><span class="golang-brace">(</span>s.logTo<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">f.WriteString</span><span class="golang-brace">(</span><span class="golang-function">fmt.Sprintf</span><span class="golang-brace">(</span>"executing&nbsp;%d",&nbsp;input<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;v&nbsp;:=&nbsp;s.factor&nbsp;*&nbsp;input<br>
&nbsp;&nbsp;&nbsp;&nbsp;out&nbsp;:=&nbsp;<span class="golang-function">fmt.Sprintf</span><span class="golang-brace">(</span>"The&nbsp;product&nbsp;is&nbsp;%d&nbsp;%s",&nbsp;v,&nbsp;s.units<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out,&nbsp;nil<br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

Unfortunately, the code above has a single Service that is mixing *many* concerns or responsibilities:
* It has logic to decide if the request is authorized.
* It has logic to control how the request event was logged.
* It holds the logic for a computation.
* It has a built-in, hard coded template for the output text.

Obviously, the code above is going to be very difficult to maintain:
* What if the administrator changes?  What if there are now multiple administrators?
* What if the authorization logic changes to use a different method?  Like maybe ldap groups
  or entitlements will be used instead.
* What if we want the log files to rotate?  Or be shipped to some external service?
* What if we want to use a big integer for the input and output?
* What if we want to reject the request if the input is less than one?
* What if we want to change the template for the output sentence?

We can re-factor the above code to consider the the above problems, and follow the single responsibility 
principle.  The example below is shown in one file, but in reality they would be separate files, which 
lets separate developers work on them while reducing the risk of merge conflicts.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"errors"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"os"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">//&nbsp;in&nbsp;file&nbsp;authorizer.go</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Authorizer&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">ConfirmAuthorized</span><span class="golang-brace">(</span>user&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;error<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;authorizer&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;admin&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewAuthorizer</span><span class="golang-brace">(</span>admin&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;Authorizer&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&authorizer<span class="golang-brace">{</span>admin<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>a&nbsp;*authorizer<span class="golang-brace">)</span>&nbsp;<span class="golang-function">ConfirmAuthorized</span><span class="golang-brace">(</span>user&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;error&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;user&nbsp;!=&nbsp;a.admin&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-function">errors.New</span><span class="golang-brace">(</span>"unauthorized"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;nil<br>
<span class="golang-brace">}</span><br>
<br>
<span class="code-comment">//&nbsp;in&nbsp;file&nbsp;calculation.go</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Calculation&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Calculate</span><span class="golang-brace">(</span>input&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">int</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Multiplier&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;factor&nbsp;<span class="golang-variable-type">int</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewMultiplier</span><span class="golang-brace">(</span>x&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;Calculation&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Multiplier<span class="golang-brace">{</span>factor:&nbsp;x<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>m&nbsp;*Multiplier<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Calculate</span><span class="golang-brace">(</span>input&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;m.factor&nbsp;*&nbsp;input<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Logger&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Log</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;logger&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*logger<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Log</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Fprint</span><span class="golang-brace">(</span>os.Stderr,&nbsp;s<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;OutputFormatter&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Format</span><span class="golang-brace">(</span>value&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;outputFormatter&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;theFormat&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;units&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewOutputFormatter</span><span class="golang-brace">(</span>template&nbsp;<span class="golang-variable-type">string</span>,&nbsp;units&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;OutputFormatter&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;template&nbsp;==&nbsp;""&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;template&nbsp;=&nbsp;"The&nbsp;product&nbsp;is&nbsp;%d&nbsp;%s"<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;units&nbsp;==&nbsp;""&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;units&nbsp;=&nbsp;"units"<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&outputFormatter<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;theFormat:&nbsp;template,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;units:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;units,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>f&nbsp;*outputFormatter<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Format</span><span class="golang-brace">(</span>value&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-function">fmt.Sprintf</span><span class="golang-brace">(</span>f.theFormat,&nbsp;value,&nbsp;f.units<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Service&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;authorizer&nbsp;&nbsp;Authorizer<br>
&nbsp;&nbsp;&nbsp;&nbsp;calculation&nbsp;Calculation<br>
&nbsp;&nbsp;&nbsp;&nbsp;logger&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Logger<br>
&nbsp;&nbsp;&nbsp;&nbsp;formatter&nbsp;&nbsp;&nbsp;OutputFormatter<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">NewService</span><span class="golang-brace">(</span>authorizer&nbsp;Authorizer,&nbsp;calculation&nbsp;Calculation,&nbsp;logger&nbsp;Logger,&nbsp;formatter&nbsp;OutputFormatter<span class="golang-brace">)</span>&nbsp;*Service&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Service<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;authorizer:&nbsp;&nbsp;authorizer,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;calculation:&nbsp;calculation,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;logger:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;logger,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;formatter:&nbsp;&nbsp;&nbsp;formatter,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>s&nbsp;*Service<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>user&nbsp;<span class="golang-variable-type">string</span>,&nbsp;input&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">(</span><span class="golang-variable-type">string</span>,&nbsp;error<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;e&nbsp;:=&nbsp;s.<span class="golang-function">authorizer.ConfirmAuthorized</span><span class="golang-brace">(</span>user<span class="golang-brace">)</span>;&nbsp;e&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;"",&nbsp;e<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;s.<span class="golang-function">logger.Log</span><span class="golang-brace">(</span><span class="golang-function">fmt.Sprintf</span><span class="golang-brace">(</span>"executing&nbsp;%d",&nbsp;input<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;v&nbsp;:=&nbsp;s.<span class="golang-function">calculation.Calculate</span><span class="golang-brace">(</span>input<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;out&nbsp;:=&nbsp;s.<span class="golang-function">formatter.Format</span><span class="golang-brace">(</span>v<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out,&nbsp;nil<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">RunService</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;auth&nbsp;:=&nbsp;<span class="golang-function">NewAuthorizer</span><span class="golang-brace">(</span>"pat"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;calc&nbsp;:=&nbsp;<span class="golang-function">NewMultiplier</span><span class="golang-brace">(</span><span class="golang-primitive-value">3</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;log&nbsp;:=&nbsp;&logger<span class="golang-brace">{}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;formatter&nbsp;:=&nbsp;<span class="golang-function">NewOutputFormatter</span><span class="golang-brace">(</span>"",&nbsp;"things"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;svc&nbsp;:=&nbsp;<span class="golang-function">NewService</span><span class="golang-brace">(</span>auth,&nbsp;calc,&nbsp;log,&nbsp;formatter<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;out,&nbsp;_&nbsp;:=&nbsp;<span class="golang-function">svc.Execute</span><span class="golang-brace">(</span>"pat",&nbsp;<span class="golang-primitive-value">8</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>out<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

The re-factoring above might seem very verbose to new engineers, but it will be far easier to 
maintain in the long term.
* If we need to change the logic for authorizations, calculations, logging, or output formatting,
  the service class does not need to be altered itself.  That means it is more stable.  It is
  in line with the open/closed principle in the next section, because the service class does
  not need to be updated, but it can behave different ways by replacing the services it depends upon.
* If the structs are in separate files, it reduces the chances of merge conflicts.
* We can also write special implementations of the interfaces that wrap existing ones.

Some people might say, why use an interface here?  There is only one implementation.  You might
think it is over-complicating things and going against the KISS principle.  More often
than you first expect, new implementations come around in the future.  It is better to 
de-couple the service from specific implementations, so that it does not need any updates
when the implementation changes.  The interface segregation principle recommends this approach.

Yes, you should frequently have interfaces defined with only one implementation.