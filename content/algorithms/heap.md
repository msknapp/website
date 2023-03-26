---
title: Heap Datastructure
weight: 10
draft: false
summary: The heap structure explained.
lastmod: 2023-03-23
nextlink: /tech/
---

The heap datastructure is a clever design that makes an array behave like a binary 
tree, and the root of the tree will always be the min or max value.

For any index `i` in the heap:
* its left child will be at `2*i`
* its right child will be at `2*i + 1`
* its children will always be greater or lesser than the current node.

# Example

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;misspos<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;heap&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;data&nbsp;&nbsp;[]<span class="golang-variable-type">int</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;index&nbsp;<span class="golang-variable-type">int</span><br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewHeap(n&nbsp;<span class="golang-variable-type">int</span>)&nbsp;*heap&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&heap{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data:&nbsp;make([]<span class="golang-variable-type">int</span>,&nbsp;n),<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;Empty()&nbsp;<span class="golang-variable-type">bool</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;h.index&nbsp;<&nbsp;1<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;Add(x&nbsp;<span class="golang-variable-type">int</span>)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.data[h.index]&nbsp;=&nbsp;x<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.climb(h.index)<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.index++<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;Peek()&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;h.data[0]<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;Pop()&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">v</span>&nbsp;:=&nbsp;h.data[0]<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.data[0]&nbsp;=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.lift()<br>
&nbsp;&nbsp;&nbsp;&nbsp;h.index--<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;v<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;climb(i&nbsp;<span class="golang-variable-type">int</span>)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">p</span>&nbsp;:=&nbsp;(i&nbsp;-&nbsp;1)&nbsp;/&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;i&nbsp;>&nbsp;0&nbsp;&&&nbsp;h.data[p]&nbsp;>&nbsp;h.data[i]&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;h.data[p],&nbsp;h.data[i]&nbsp;=&nbsp;h.data[i],&nbsp;h.data[p]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;p<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">p</span>&nbsp;=&nbsp;(i&nbsp;-&nbsp;1)&nbsp;/&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;(h&nbsp;*heap)&nbsp;lift()&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;<span class="golang-variable">i</span>&nbsp;:=&nbsp;0;&nbsp;i&nbsp;<&nbsp;len(h.data);&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">j</span>&nbsp;:=&nbsp;2*i&nbsp;+&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;j&nbsp;>=&nbsp;h.index&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;i&nbsp;<&nbsp;h.index-1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;h.data[i],&nbsp;h.data[h.index-1]&nbsp;=&nbsp;h.data[h.index-1],&nbsp;h.data[i]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;h.climb(i)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">break</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;j+1&nbsp;<&nbsp;h.index&nbsp;&&&nbsp;h.data[j]&nbsp;>&nbsp;h.data[j+1]&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;j++<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;h.data[i],&nbsp;h.data[j]&nbsp;=&nbsp;h.data[j],&nbsp;h.data[i]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;j<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;FirstMissing(x&nbsp;[]<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">h</span>&nbsp;:=&nbsp;NewHeap(len(x))<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">missing</span>&nbsp;:=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;<span class="golang-variable">v</span>&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;x&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;v&nbsp;<&nbsp;missing&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;v&nbsp;>&nbsp;missing&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;h.Add(v)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;missing&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;!h.Empty()&nbsp;&&&nbsp;missing&nbsp;>=&nbsp;h.Peek()&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">missing</span>&nbsp;=&nbsp;h.Pop()&nbsp;+&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;missing<br>
}<br>


{{< /gocode >}}