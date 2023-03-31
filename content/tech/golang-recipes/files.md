---
title: Working with Files
draft: false
lastmod: 2023-03-31
weight: 4
summary: Working with files.
---

# Read File to String

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"io/ioutil"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">FileToString</span><span class="golang-brace">(</span>filepath&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;bits,&nbsp;e&nbsp;:=&nbsp;<span class="golang-function">ioutil.ReadFile</span><span class="golang-brace">(</span>filepath<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;e&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">panic</span><span class="golang-brace">(</span>e<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">(</span>bits<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
{{< /gocode >}}

# Check if a File Exists

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"os"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">FileExists</span><span class="golang-brace">(</span>filepath&nbsp;<span class="golang-variable-type">string</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">bool</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;e&nbsp;:=&nbsp;<span class="golang-function">os.Stat</span><span class="golang-brace">(</span>filepath<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;e&nbsp;==&nbsp;nil<br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}
