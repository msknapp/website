---
title: Queue
draft: false
weight: 9
lastmod: 2023-03-25
summary: Use queues for faster push and pop operations.
---

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;coll<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Node[T&nbsp;any]&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;value&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;next&nbsp;&nbsp;*Node[T]<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewNode[T&nbsp;any](value&nbsp;T)&nbsp;*Node[T]&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Node[T]{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;value:&nbsp;value,<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Queue[T&nbsp;any]&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;head&nbsp;*Node[T]<br>
&nbsp;&nbsp;&nbsp;&nbsp;tail&nbsp;*Node[T]<br>
&nbsp;&nbsp;&nbsp;&nbsp;size&nbsp;<span class="golang-variable-type">int</span><br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewQueue[T&nbsp;any]()&nbsp;*Queue[T]&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Queue[T]{}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(q&nbsp;*Queue[T])&nbsp;Add(value&nbsp;T)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">nd</span>&nbsp;:=&nbsp;NewNode[T](value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;q.tail&nbsp;!=&nbsp;nil&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;q.tail.<span class="golang-variable">next</span>&nbsp;=&nbsp;nd<br>
&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;q.<span class="golang-variable">head</span>&nbsp;=&nbsp;nd<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;q.<span class="golang-variable">tail</span>&nbsp;=&nbsp;nd<br>
&nbsp;&nbsp;&nbsp;&nbsp;q.size&nbsp;+=&nbsp;1<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(q&nbsp;*Queue[T])&nbsp;IsEmpty()&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;q.size&nbsp;<&nbsp;1<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(q&nbsp;*Queue[T])&nbsp;IsNotEmpty()&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;q.size&nbsp;>&nbsp;0<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(q&nbsp;*Queue[T])&nbsp;PopLeft()&nbsp;T&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;v&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;q.head&nbsp;!=&nbsp;nil&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">v</span>&nbsp;=&nbsp;q.head.value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;q.<span class="golang-variable">head</span>&nbsp;=&nbsp;q.head.next<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;q.size&nbsp;-=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;v<br>
}<br>
{{< /gocode >}}