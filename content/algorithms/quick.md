---
title: Quick Sort
draft: false
weight: 4
summary: How the quick sort works
lastmod: 2023-03-19
---

The quick sort runs faster because it is efficient, and it runs in-place (unlike the merge sort).

![quick-sort](/images/quick-sort.png)

It works by:
* choosing a value to be the "pivot" point.
* In the remaining array, it swaps values between left and right cursors until
  * all values up to and including the left cursor are less than the pivot.
  * all values on or after the right cursor are greater than or equal to the pivot.
* Then the pivot is swapped into the middle, trading places with the right pivot.
* The process gets repeated on the sub-arrays to the right and left of the pivot.

The quick sort is:
* not stable, equal values may be swapped.  This can be worked around by adjusting your
  comparison logic.
* in-place, it uses no significant amount of extra memory.
* Efficient, having O(n * log(n)) performance.
* possible to parallelize.  At step 3 in the diagram, separate coroutines or threads 
  could be created if the sub-arrays are still quite large.
* works with non-numerical types, it uses a comparator.

The quick sort algorithm was made to handle sorting large arrays, but in practice, the
merge sort algorithm tends to be used for larger arrays instead.

Since the quick sort uses recursion, it is possible to have a stack overflow in your program, 
unless the program is written to work around recursion.

# Example In Golang

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;srt<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;quickSort(nums&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;start&nbsp;<span class="golang-variable-type">int</span>,&nbsp;stop&nbsp;<span class="golang-variable-type">int</span>)&nbsp;{<br>
<span class="golang-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;pivot&nbsp;from&nbsp;right</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;start&nbsp;>=&nbsp;stop&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">mid</span>&nbsp;:=&nbsp;partition(nums,&nbsp;start,&nbsp;stop)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;mid&nbsp;>&nbsp;start+1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;quickSort(nums,&nbsp;start,&nbsp;mid-1)<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;mid&nbsp;<&nbsp;stop-1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;quickSort(nums,&nbsp;start+1,&nbsp;stop)<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
}<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;partition(nums&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;start,&nbsp;stop&nbsp;<span class="golang-variable-type">int</span>)&nbsp;<span class="golang-variable-type">int</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;stop&nbsp;<=&nbsp;start&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">pivotIndex</span>&nbsp;:=&nbsp;stop<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">pivot</span>&nbsp;:=&nbsp;nums[pivotIndex]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">j</span>&nbsp;:=&nbsp;start&nbsp;-&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;<span class="golang-variable">i</span>&nbsp;:=&nbsp;start;&nbsp;i&nbsp;<&nbsp;pivotIndex;&nbsp;i++&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;nums[i]&nbsp;<&nbsp;pivot&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;j++<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;nums[i],&nbsp;nums[j]&nbsp;=&nbsp;nums[j],&nbsp;nums[i]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;j++<br>
&nbsp;&nbsp;&nbsp;&nbsp;nums[pivotIndex],&nbsp;nums[j]&nbsp;=&nbsp;nums[j],&nbsp;nums[pivotIndex]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;j<br>
}<br>
{{< /gocode >}}

# Python Example

{{< gocode >}}
<br>
def&nbsp;shift(ar,&nbsp;from_index,&nbsp;to_index):<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;from_index<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">v</span>&nbsp;=&nbsp;ar[from_index]<br>
&nbsp;&nbsp;&nbsp;&nbsp;while&nbsp;i&nbsp;<&nbsp;to_index:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ar[i]&nbsp;=&nbsp;ar[i+1]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;i&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;ar[to_index]&nbsp;=&nbsp;v<br>
<br>
def&nbsp;quick_sort(ar,&nbsp;from_index,&nbsp;to_index):<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;to_index&nbsp;-&nbsp;from_index&nbsp;<&nbsp;1:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">pivot_index</span>&nbsp;=&nbsp;to_index<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">pivot</span>&nbsp;=&nbsp;ar[pivot_index]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">index</span>&nbsp;=&nbsp;from_index<br>
&nbsp;&nbsp;&nbsp;&nbsp;while&nbsp;index&nbsp;<&nbsp;pivot_index:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;ar[index]&nbsp;>&nbsp;pivot:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;shift(ar,&nbsp;index,&nbsp;pivot_index)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;pivot_index&nbsp;-=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;index&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;quick_sort(ar,&nbsp;from_index,&nbsp;pivot_index&nbsp;-&nbsp;1)<br>
&nbsp;&nbsp;&nbsp;&nbsp;quick_sort(ar,&nbsp;pivot_index+1,&nbsp;to_index)<br>
<br>
<span class="golang-control-keyword">if</span>&nbsp;__name__&nbsp;==&nbsp;"__main__":<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">inputs</span>&nbsp;=&nbsp;[5,2,9,6,31,28,63,19,11,64,23]<br>
&nbsp;&nbsp;&nbsp;&nbsp;quick_sort(inputs,&nbsp;0,&nbsp;len(inputs)-1)<br>
&nbsp;&nbsp;&nbsp;&nbsp;print(inputs)<br>

{{< /gocode >}}