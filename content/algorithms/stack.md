---
title: Stack
weight: 8
draft: false
lastmod: 2023-03-25
summary: Use a stack for LIFO operations.
---

# Example

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;collatz<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Stack&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;data&nbsp;[]<span class="golang-variable-type">int</span><br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewStack(cap&nbsp;<span class="golang-variable-type">int</span>)&nbsp;*Stack&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;cap&nbsp;<&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">cap</span>&nbsp;=&nbsp;1024<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Stack{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data:&nbsp;make([]<span class="golang-variable-type">int</span>,&nbsp;0,&nbsp;cap),<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;push(x&nbsp;<span class="golang-variable-type">int</span>)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;s.<span class="golang-variable">data</span>&nbsp;=&nbsp;append(s.data,&nbsp;x)<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;pop()&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(s.data)&nbsp;<&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">v</span>&nbsp;:=&nbsp;s.data[len(s.data)-1]<br>
&nbsp;&nbsp;&nbsp;&nbsp;s.<span class="golang-variable">data</span>&nbsp;=&nbsp;s.data[:len(s.data)-1]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;v<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;empty()&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;len(s.data)&nbsp;<&nbsp;1<br>
}<br>
{{< /gocode >}}

# Generic Example

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;collections<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Stack&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;items&nbsp;[]TempDate<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewStack()&nbsp;*Stack&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Stack{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;items:&nbsp;make([]TempDate,&nbsp;0,&nbsp;3),<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;Push(x&nbsp;TempDate)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;s.<span class="golang-variable">items</span>&nbsp;=&nbsp;append(s.items,&nbsp;x)<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;Pop()&nbsp;TempDate&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">x</span>&nbsp;:=&nbsp;s.items[len(s.items)-1]<br>
&nbsp;&nbsp;&nbsp;&nbsp;s.<span class="golang-variable">items</span>&nbsp;=&nbsp;s.items[:len(s.items)-1]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;Peek()&nbsp;TempDate&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;s.items[len(s.items)-1]<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(s&nbsp;*Stack)&nbsp;Empty()&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;len(s.items)&nbsp;<&nbsp;1<br>
}<br>
{{< /gocode >}}