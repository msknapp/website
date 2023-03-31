---
title: Creating Unit Tests
draft: false
lastmod: 2023-03-31
weight: 5
summary: How to create unit tests by various forms.
---

# Simple Unit Test

The following is how you could start a very simple unit test:
{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"testing"<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;"github.com/stretchr/testify/assert"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">TestSomeFunction</span><span class="golang-brace">(</span>t&nbsp;*testing.T<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;tester&nbsp;:=&nbsp;<span class="golang-function">assert.New</span><span class="golang-brace">(</span>t<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">tester.Nil</span><span class="golang-brace">(</span>nil<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

# Test Cases

The following can serve as an example for unit tests with a list of cases:

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"math"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"testing"<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;"github.com/stretchr/testify/assert"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">IsSquare</span><span class="golang-brace">(</span>x&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">bool</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;y&nbsp;:=&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">(</span><span class="golang-function">math.Sqrt</span><span class="golang-brace">(</span><span class="golang-variable-type">float64</span><span class="golang-brace">(</span>x<span class="golang-brace">)))</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x&nbsp;==&nbsp;y*y<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">TestAListOfCases</span><span class="golang-brace">(</span>t&nbsp;*testing.T<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;tester&nbsp;:=&nbsp;<span class="golang-function">assert.New</span><span class="golang-brace">(</span>t<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">type</span>&nbsp;myCase&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;value&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;expected&nbsp;<span class="golang-variable-type">bool</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;tc&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;[]myCase<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-<span class="golang-primitive-value">4</span>,&nbsp;<span class="golang-primitive-value">false</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-primitive-value">4</span>,&nbsp;<span class="golang-primitive-value">true</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-primitive-value">64</span>,&nbsp;<span class="golang-primitive-value">true</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-primitive-value">63</span>,&nbsp;<span class="golang-primitive-value">false</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;tc.expected&nbsp;!=&nbsp;<span class="golang-function">IsSquare</span><span class="golang-brace">(</span>tc.value<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">tester.Fail</span><span class="golang-brace">(</span>"<span class="golang-function">IsSquare</span><span class="golang-brace">(</span>%d<span class="golang-brace">)</span>&nbsp;did&nbsp;not&nbsp;get&nbsp;%t",&nbsp;tc.value,&nbsp;tc.expected<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}
