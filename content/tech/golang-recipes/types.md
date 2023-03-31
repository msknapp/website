---
title: Working with Types
draft: false
weight: 2
lastmod: 2023-03-31
summary: Checking variables types and acting accordingly.
---

# Act Based on Type

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>t&nbsp;<span class="golang-variable-type">any</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">switch</span>&nbsp;x&nbsp;:=&nbsp;t.<span class="golang-brace">(</span><span class="golang-control-keyword">type</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">case</span>&nbsp;<span class="golang-variable-type">string</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">case</span>&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%d\n",&nbsp;x<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">case</span>&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%t\n",&nbsp;x<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">case</span>&nbsp;<span class="golang-variable-type">float64</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%f\n",&nbsp;x<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">DoIt</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;things&nbsp;:=&nbsp;[]<span class="golang-variable-type">any</span><span class="golang-brace">{</span>"foo",&nbsp;<span class="golang-primitive-value">1</span>,&nbsp;<span class="golang-primitive-value">true</span>,&nbsp;<span class="golang-primitive-value">6.4</span><span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;t&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;things&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>t<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

# Check Object Type

As a function:
{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">IsString</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">any</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">bool</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;ok&nbsp;:=&nbsp;s.<span class="golang-brace">(</span><span class="golang-variable-type">string</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;ok<br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

For example:
{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>t&nbsp;<span class="golang-variable-type">any</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;s,&nbsp;ok&nbsp;:=&nbsp;t.<span class="golang-brace">(</span><span class="golang-variable-type">string</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;ok&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>s<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">DoIt</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;things&nbsp;:=&nbsp;[]<span class="golang-variable-type">any</span><span class="golang-brace">{</span>"foo",&nbsp;<span class="golang-primitive-value">1</span>,&nbsp;<span class="golang-primitive-value">true</span>,&nbsp;<span class="golang-primitive-value">6.4</span><span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;t&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;things&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Execute</span><span class="golang-brace">(</span>t<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}
