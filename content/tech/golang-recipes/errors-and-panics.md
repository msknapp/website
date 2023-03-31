---
title: Errors and Panics
weight: 9
lastmod: 2023-03-31
draft: false
summary: Working with errors, panics, and recovery.
---

# Ignore Errors

Many functions return an error as their second argument, and you may know that the
error won't be encountered for some reason.  In that case, ignoring the error is nice,
but the first return type might vary, so use a generic to ignore errors like this:

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"strconv"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;IgnoreError[T&nbsp;<span class="golang-variable-type">any</span>]<span class="golang-brace">(</span>out&nbsp;T,&nbsp;_&nbsp;error<span class="golang-brace">)</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;i0&nbsp;:=&nbsp;<span class="golang-function">IgnoreError</span><span class="golang-brace">(</span><span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"<span class="golang-primitive-value">3</span>"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;i1&nbsp;:=&nbsp;<span class="golang-function">IgnoreError</span><span class="golang-brace">(</span><span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"M"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;b0&nbsp;:=&nbsp;<span class="golang-function">IgnoreError</span><span class="golang-brace">(</span><span class="golang-function">strconv.ParseBool</span><span class="golang-brace">(</span>"<span class="golang-primitive-value">true</span>"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;b1&nbsp;:=&nbsp;<span class="golang-function">IgnoreError</span><span class="golang-brace">(</span><span class="golang-function">strconv.ParseBool</span><span class="golang-brace">(</span>"banana"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"i0=%d,&nbsp;i1=%d,&nbsp;b0=%t,&nbsp;b1=%t\n",&nbsp;i0,&nbsp;i1,&nbsp;b0,&nbsp;b1<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;e0&nbsp;:=&nbsp;<span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"M"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>e0.<span class="golang-function">Error</span><span class="golang-brace">())</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;e1&nbsp;:=&nbsp;<span class="golang-function">strconv.ParseBool</span><span class="golang-brace">(</span>"banana"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>e1.<span class="golang-function">Error</span><span class="golang-brace">())</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

When run, the above code prints:
```
i0=3, i1=0, b0=true, b1=false
strconv.Atoi: parsing "M": invalid syntax
strconv.ParseBool: parsing "banana": invalid syntax
```

# Panic!  Or Don't

The following example shows how to do all of the following:
* promote errors to panics
* ignore panics
* reduce panics to errors

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"strconv"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">//&nbsp;PanicOnError&nbsp;returns&nbsp;the&nbsp;first&nbsp;input,&nbsp;or&nbsp;panics&nbsp;if&nbsp;the&nbsp;second&nbsp;argument&nbsp;is</span><br>
<span class="code-comment">//&nbsp;a&nbsp;non-nil&nbsp;error.</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;PanicOnError[T&nbsp;<span class="golang-variable-type">any</span>]<span class="golang-brace">(</span>out&nbsp;T,&nbsp;err&nbsp;error<span class="golang-brace">)</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;err&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">panic</span><span class="golang-brace">(</span>err<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
<span class="golang-brace">}</span><br>
<br>
<span class="code-comment">//&nbsp;IgnorePanic&nbsp;runs&nbsp;the&nbsp;function&nbsp;given,&nbsp;and&nbsp;ignores&nbsp;any&nbsp;panic&nbsp;thrown.</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">IgnorePanic</span><span class="golang-brace">(</span>whinyFunction&nbsp;<span class="golang-top-level-keyword">func</span><span class="golang-brace">())</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">defer</span>&nbsp;<span class="golang-top-level-keyword">func</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span>&nbsp;<span class="golang-function">recover</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">}()</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">whinyFunction</span><span class="golang-brace">()</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">ErrorOnPanic</span><span class="golang-brace">(</span>whinyFunction&nbsp;<span class="golang-top-level-keyword">func</span><span class="golang-brace">())</span>&nbsp;<span class="golang-brace">(</span>err&nbsp;error<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">defer</span>&nbsp;<span class="golang-top-level-keyword">func</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;p&nbsp;:=&nbsp;<span class="golang-function">recover</span><span class="golang-brace">()</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;p&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;err&nbsp;=&nbsp;<span class="golang-function">fmt.Errorf</span><span class="golang-brace">(</span>"%v",&nbsp;p<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}()</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">whinyFunction</span><span class="golang-brace">()</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;x,&nbsp;y&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">IgnorePanic</span><span class="golang-brace">(</span><span class="golang-top-level-keyword">func</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;<span class="golang-function">PanicOnError</span><span class="golang-brace">(</span><span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"<span class="golang-primitive-value">3</span>"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;=&nbsp;<span class="golang-function">PanicOnError</span><span class="golang-brace">(</span><span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"M"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">})</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"x=%d,&nbsp;y=%d\n",&nbsp;x,&nbsp;y<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;e0&nbsp;:=&nbsp;<span class="golang-function">ErrorOnPanic</span><span class="golang-brace">(</span><span class="golang-top-level-keyword">func</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">PanicOnError</span><span class="golang-brace">(</span><span class="golang-function">strconv.Atoi</span><span class="golang-brace">(</span>"M"<span class="golang-brace">))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">})</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>e0.<span class="golang-function">Error</span><span class="golang-brace">())</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

When the above code runs, it prints:
```
x=3, y=0
strconv.Atoi: parsing "M": invalid syntax
```
