---
title: Merge Sort
draft: false
weight: 5
summary: How the merge sort works
lastmod: 2023-03-19
---

The merge sort is among the fastest sorting algorithms, but it requires extra memory to run.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;srt<br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;MergeSort[T&nbsp;any](x&nbsp;[]T,&nbsp;c&nbsp;Comparer[T])&nbsp;[]T&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(x)&nbsp;<&nbsp;2&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">middle</span>&nbsp;:=&nbsp;len(x)&nbsp;/&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">left</span>&nbsp;:=&nbsp;x[:middle]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">right</span>&nbsp;:=&nbsp;x[middle:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">leftSorted</span>&nbsp;:=&nbsp;MergeSort(left,&nbsp;c)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">rightSorted</span>&nbsp;:=&nbsp;MergeSort(right,&nbsp;c)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">out</span>&nbsp;:=&nbsp;make([]T,&nbsp;len(x))<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;<span class="golang-variable">i</span>&nbsp;:=&nbsp;0;&nbsp;i&nbsp;<&nbsp;len(x);&nbsp;i++&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(leftSorted)&nbsp;<&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out[i]&nbsp;=&nbsp;rightSorted[0]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">rightSorted</span>&nbsp;=&nbsp;rightSorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(rightSorted)&nbsp;<&nbsp;1&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out[i]&nbsp;=&nbsp;leftSorted[0]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">leftSorted</span>&nbsp;=&nbsp;leftSorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;c.Less(leftSorted[0],&nbsp;rightSorted[0])&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out[i]&nbsp;=&nbsp;leftSorted[0]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">leftSorted</span>&nbsp;=&nbsp;leftSorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}&nbsp;<span class="golang-control-keyword">else</span>&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out[i]&nbsp;=&nbsp;rightSorted[0]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">rightSorted</span>&nbsp;=&nbsp;rightSorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
}<br>
{{< /gocode >}}

# Example In Python

{{< gocode >}}
<br>
def&nbsp;merge_sort(data):<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">li</span>&nbsp;=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">ri</span>&nbsp;=&nbsp;len(data)-1<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;ri&nbsp;==&nbsp;li:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;data<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">middle</span>&nbsp;=&nbsp;(len(data)&nbsp;-&nbsp;li)&nbsp;//&nbsp;2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">left</span>&nbsp;=&nbsp;data[:middle]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">right</span>&nbsp;=&nbsp;data[middle:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">left_sorted</span>&nbsp;=&nbsp;merge_sort(left)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">right_sorted</span>&nbsp;=&nbsp;merge_sort(right)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">out</span>&nbsp;=&nbsp;[]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;i&nbsp;in&nbsp;<span class="golang-control-keyword">range</span>(len(data)):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(left_sorted)&nbsp;<&nbsp;1:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out.append(right_sorted[0])<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">right_sorted</span>&nbsp;=&nbsp;right_sorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(right_sorted)&nbsp;<&nbsp;1:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out.append(left_sorted[0])<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">left_sorted</span>&nbsp;=&nbsp;left_sorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(left_sorted)&nbsp;>&nbsp;0&nbsp;and&nbsp;left_sorted[0]&nbsp;<&nbsp;right_sorted[0]:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out.append(left_sorted[0])<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">left_sorted</span>&nbsp;=&nbsp;left_sorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;out.append(right_sorted[0])<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">right_sorted</span>&nbsp;=&nbsp;right_sorted[1:]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
<br>
<br>
<span class="golang-control-keyword">if</span>&nbsp;__name__&nbsp;==&nbsp;"__main__":<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">data</span>&nbsp;=&nbsp;[3,9,2,&nbsp;16,38,&nbsp;25,&nbsp;91,6,&nbsp;34,8,&nbsp;59,&nbsp;62,&nbsp;7,&nbsp;5,&nbsp;46]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">out</span>&nbsp;=&nbsp;merge_sort(data)<br>
&nbsp;&nbsp;&nbsp;&nbsp;print(out)
{{< /gocode >}}