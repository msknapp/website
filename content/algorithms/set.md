---
title: Set
draft: false
weight: 14
lastmod: 2023-03-25
summary: Use a set to quickly check for existence.
---

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;collections<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Set[T&nbsp;comparable]&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;data&nbsp;<span class="golang-control-keyword">map</span>[T]<span class="golang-control-keyword">struct</span>{}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewSet[T&nbsp;comparable]()&nbsp;*Set[T]&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Set[T]{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data:&nbsp;make(<span class="golang-control-keyword">map</span>[T]<span class="golang-control-keyword">struct</span>{},&nbsp;2),<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Set[T])&nbsp;Add(x&nbsp;T)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;s.data[x]&nbsp;=&nbsp;<span class="golang-control-keyword">struct</span>{}{}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Set[T])&nbsp;Has(x&nbsp;T)&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;<span class="golang-variable">ex</span>&nbsp;:=&nbsp;s.data[x]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;ex<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Set[T])&nbsp;Size()&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;len(s.data)<br>
}<br>
{{< /gocode >}}