---
title: Very Basic Recipes
draft: false
lastmod: 2023-03-31
weight: 1
summary: Ternary, working with strings and streams, etc.
---

# Ternary

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;Ternary[T&nbsp;<span class="golang-variable-type">any</span>]<span class="golang-brace">(</span>condition&nbsp;<span class="golang-variable-type">bool</span>,&nbsp;ifTrue&nbsp;T,&nbsp;ifFalse&nbsp;T<span class="golang-brace">)</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;condition&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;ifTrue<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;ifFalse<br>
<span class="golang-brace">}</span><br>
<br>
{{< /gocode >}}

# Convert String To Stream and Back

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"bytes"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"io"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"io/ioutil"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">StringToStream</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;io.Reader&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-function">bytes.NewBuffer</span><span class="golang-brace">(</span>[]<span class="golang-variable-type">byte</span><span class="golang-brace">(</span>s<span class="golang-brace">))</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">StreamToString</span><span class="golang-brace">(</span>r&nbsp;io.Reader<span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;r&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;""<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;bts,&nbsp;e&nbsp;:=&nbsp;<span class="golang-function">ioutil.ReadAll</span><span class="golang-brace">(</span>r<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;e&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">panic</span><span class="golang-brace">(</span>e<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">(</span>bts<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

# Reverse a String

This is a frequent coding interview challenge.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Reverse</span><span class="golang-brace">(</span>s&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;:=&nbsp;[]<span class="golang-variable-type">byte</span><span class="golang-brace">(</span>s<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;i&nbsp;:=&nbsp;<span class="golang-primitive-value">0</span>;&nbsp;i&nbsp;<&nbsp;<span class="golang-function">len</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span>/<span class="golang-primitive-value">2</span>;&nbsp;i++&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x[i],&nbsp;x[<span class="golang-function">len</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span>-i-<span class="golang-primitive-value">1</span>]&nbsp;=&nbsp;x[<span class="golang-function">len</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span>-i-<span class="golang-primitive-value">1</span>],&nbsp;x[i]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}
