---
title: Bubble Sort
weight: 1
draft: false
summary: How the bubble sort works
lastmod: 2023-03-19
---

The bubble sort is the most intuitive sorting algorithm there is.

![bubble-sort](/images/bubble-sort.png)

It works by:
* iterating over the entire array.
* repeatedly swapping adjacent elements if they are out of order.
* with each iteration, the end index is decremented.

After each iteration, one more item at the end of the array is in 
its final place.

There is an easy way to remember what the bubble sort is.
I like to think of the right side of the array as a growing "bubble"
of sorted elements (maximums if you are sorting in ascending order).

The bubble sort is:
* "stable", if two elements are equal, their 
  relative order will not change.
* "in-place", meaning it does not require extra memory to run.
* not efficient, having O(n^2) performance.
* not possible to parallelize.
* works with non-numerical types, it uses a comparator.

It is possible to adapt the bubble sort so that between each 
iteration, it checks if the array is already sorted, and it could
stop early.

# Example In Java

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;scratch;<br>
<br>
public&nbsp;class&nbsp;BubbleSort&nbsp;{<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;public&nbsp;static&nbsp;void&nbsp;Sort(<span class="golang-variable-type">int</span>[]&nbsp;x)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;(<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;0;&nbsp;i&nbsp;<&nbsp;x.length-1;&nbsp;i++)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SortOnce(x);<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;private&nbsp;static&nbsp;void&nbsp;SortOnce(<span class="golang-variable-type">int</span>[]&nbsp;x)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;(<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;0;&nbsp;i&nbsp;<&nbsp;x.length&nbsp;-&nbsp;1;&nbsp;i++)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;(x[i+1]&nbsp;<&nbsp;x[i])&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;x[i];<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x[i]&nbsp;=&nbsp;x[i+1];<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x[i+1]&nbsp;=&nbsp;tmp;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
{{< /gocode >}}