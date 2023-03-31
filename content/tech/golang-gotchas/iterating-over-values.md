---
title: Iterating Over Values
draft: false
weight: 3
lastmod: 2023-03-31
summary: Iterating over values can lead to surprising results.
---

Developers might be shocked by the results of this code:

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;:=&nbsp;[]<span class="golang-variable-type">int</span><span class="golang-brace">{</span><span class="golang-primitive-value">3</span>,&nbsp;<span class="golang-primitive-value">9</span>,&nbsp;<span class="golang-primitive-value">2</span>,&nbsp;<span class="golang-primitive-value">6</span>,&nbsp;<span class="golang-primitive-value">4</span><span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;:=&nbsp;<span class="golang-function">make</span><span class="golang-brace">(</span>[]*<span class="golang-variable-type">int</span>,&nbsp;<span class="golang-primitive-value">5</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;i,&nbsp;v&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;x&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y[i]&nbsp;=&nbsp;&v<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;v&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;y&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%d,&nbsp;",&nbsp;*v<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

When run, it prints the following:
```
4, 4, 4, 4, 4, 
```

Why is that?  It seems like it just used the last value in x for everything?
This happens because in the for loop, the variable "v" is updated in its place.
The "y[i]" is pointing to "v", but its value keeps getting overwritten.

There is a workaround for this situation, the value can be copied to a new local
variable within the loop.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;:=&nbsp;[]<span class="golang-variable-type">int</span><span class="golang-brace">{</span><span class="golang-primitive-value">3</span>,&nbsp;<span class="golang-primitive-value">9</span>,&nbsp;<span class="golang-primitive-value">2</span>,&nbsp;<span class="golang-primitive-value">6</span>,&nbsp;<span class="golang-primitive-value">4</span><span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;:=&nbsp;<span class="golang-function">make</span><span class="golang-brace">(</span>[]*<span class="golang-variable-type">int</span>,&nbsp;<span class="golang-primitive-value">5</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;i,&nbsp;v&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;x&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;v2&nbsp;:=&nbsp;v<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y[i]&nbsp;=&nbsp;&v2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;v&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;y&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%d,&nbsp;",&nbsp;*v<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

Now when this runs, it prints what we expected:
```
3, 9, 2, 6, 4, 
```

This problem occurs whenever we have a for loop that iterates over values, not pointers, and within
the loop we capture a pointer to that value.