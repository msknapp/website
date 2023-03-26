---
title: Binary Search
weight: 9
draft: false
summary: How to use the binary search algorithm.
---

A binary search can be used to find the index of an item in an array that is
already sorted.  It works by repeatedly cutting the array in half until
the item is found.

# Example

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;search<br>
<br>
<span class="golang-comment">//&nbsp;BinarySearch&nbsp;takes&nbsp;sorted&nbsp;values,&nbsp;a&nbsp;desired&nbsp;value,&nbsp;and&nbsp;the&nbsp;index.</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;BinarySearch(values&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;targetValue&nbsp;<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
<span class="golang-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;already&nbsp;sorted</span><br>
<span class="golang-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;in&nbsp;middle,&nbsp;if&nbsp;you&nbsp;get&nbsp;to&nbsp;match,&nbsp;or&nbsp;go&nbsp;left,&nbsp;right,</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">lower</span>&nbsp;:=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">upper</span>&nbsp;:=&nbsp;len(values)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;upper&nbsp;>&nbsp;lower&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">index</span>&nbsp;:=&nbsp;lower&nbsp;+&nbsp;((upper&nbsp;-&nbsp;lower)&nbsp;/&nbsp;2)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">actualValue</span>&nbsp;:=&nbsp;values[index]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;actualValue&nbsp;==&nbsp;targetValue&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;actualValue&nbsp;>&nbsp;targetValue&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">upper</span>&nbsp;=&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;<span class="golang-control-keyword">if</span>&nbsp;lower&nbsp;==&nbsp;index&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;lower&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">lower</span>&nbsp;=&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
}<br>
{{< /gocode >}}

# Generic Example

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;srt<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;BinarySearch[T&nbsp;comparable](x&nbsp;[]T,&nbsp;y&nbsp;T,&nbsp;c&nbsp;Comparer[T])&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">max</span>&nbsp;:=&nbsp;len(x)&nbsp;-&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">min</span>&nbsp;:=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">index</span>&nbsp;:=&nbsp;max&nbsp;/&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Equals(x[min],&nbsp;y)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;min<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Equals(x[max],&nbsp;y)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;max<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Greater(x[min],&nbsp;y)&nbsp;||&nbsp;c.Less(x[max],&nbsp;y)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Equals(x[index],&nbsp;y)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;max-min&nbsp;<=&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Less(x[index],&nbsp;y)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">min</span>&nbsp;=&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">max</span>&nbsp;=&nbsp;index<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;max&nbsp;==&nbsp;min&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">index</span>&nbsp;=&nbsp;min&nbsp;+&nbsp;(max-min)/2<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Comparer[T&nbsp;any]&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;Equals(x&nbsp;T,&nbsp;y&nbsp;T)&nbsp;<span class="golang-variable-type">bool</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Less(x&nbsp;T,&nbsp;y&nbsp;T)&nbsp;<span class="golang-variable-type">bool</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Greater(x&nbsp;T,&nbsp;y&nbsp;T)&nbsp;<span class="golang-variable-type">bool</span><br>
}<br>

{{< /gocode >}}

# Example Closest Index

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;bisect<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;MODE&nbsp;<span class="golang-variable-type">int</span><br>
<br>
<span class="golang-control-keyword">const</span>&nbsp;(<br>
&nbsp;&nbsp;&nbsp;&nbsp;EQUAL&nbsp;<span class="golang-variable">MODE</span>&nbsp;=&nbsp;iota<br>
&nbsp;&nbsp;&nbsp;&nbsp;BEFORE<br>
&nbsp;&nbsp;&nbsp;&nbsp;AFTER<br>
)<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;IndexOf(x&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;v&nbsp;<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;i,&nbsp;<span class="golang-variable">ok</span>&nbsp;:=&nbsp;closestIndex(x,&nbsp;v)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;ok&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;i<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;Contains(x&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;v&nbsp;<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;_,&nbsp;<span class="golang-variable">ok</span>&nbsp;:=&nbsp;closestIndex(x,&nbsp;v)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;ok<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;IndexOfOrBefore(x&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;v&nbsp;<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;i,&nbsp;<span class="golang-variable">_</span>&nbsp;:=&nbsp;closestIndex(x,&nbsp;v)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;i&nbsp;==&nbsp;-2&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;len(x)<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;i&nbsp;==&nbsp;-1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;i<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;closestIndex(x&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;v&nbsp;<span class="golang-variable-type">int</span>)&nbsp;(<span class="golang-variable-type">int</span>,&nbsp;<span class="golang-variable-type">bool</span>)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x&nbsp;==&nbsp;nil&nbsp;||&nbsp;len(x)&nbsp;<&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1,&nbsp;false<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x[0]&nbsp;==&nbsp;v&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;0,&nbsp;true<br>
&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x[0]&nbsp;>&nbsp;v&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1,&nbsp;false<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x[len(x)-1]&nbsp;==&nbsp;v&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;len(x)&nbsp;-&nbsp;1,&nbsp;true<br>
&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x[len(x)-1]&nbsp;<&nbsp;v&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-2,&nbsp;false<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">lo</span>&nbsp;:=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hi</span>&nbsp;:=&nbsp;len(x)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;mid&nbsp;<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;hi-lo&nbsp;>&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">mid</span>&nbsp;=&nbsp;(lo&nbsp;+&nbsp;hi)&nbsp;/&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">mv</span>&nbsp;:=&nbsp;x[mid]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;v&nbsp;==&nbsp;mv&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;mid,&nbsp;true<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;v&nbsp;>&nbsp;mv&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">lo</span>&nbsp;=&nbsp;mid<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;v&nbsp;<&nbsp;mv&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hi</span>&nbsp;=&nbsp;mid<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;hi,&nbsp;false<br>
}<br>


{{< /gocode >}}

# Example In Java

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;search;<br>
<br>
public&nbsp;class&nbsp;Binary&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;public&nbsp;static&nbsp;<span class="golang-variable-type">int</span>&nbsp;Search(<span class="golang-variable-type">int</span>[]&nbsp;nums,&nbsp;<span class="golang-variable-type">int</span>&nbsp;value,&nbsp;<span class="golang-variable-type">int</span>&nbsp;start,&nbsp;<span class="golang-variable-type">int</span>&nbsp;stop)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">lo</span>&nbsp;=&nbsp;start;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">hi</span>&nbsp;=&nbsp;stop&nbsp;-&nbsp;1;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;(nums[lo]&nbsp;==&nbsp;value)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;lo;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;(nums[hi]&nbsp;==&nbsp;value)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;hi;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;while&nbsp;(hi&nbsp;-&nbsp;lo&nbsp;>&nbsp;1)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">mid</span>&nbsp;=&nbsp;(hi&nbsp;+&nbsp;lo)&nbsp;/&nbsp;2;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">v</span>&nbsp;=&nbsp;nums[mid];<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;(v&nbsp;==&nbsp;value)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;v;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;(v&nbsp;<&nbsp;value)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">lo</span>&nbsp;=&nbsp;mid;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hi</span>&nbsp;=&nbsp;mid;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1;<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>

{{< /gocode >}}