---
title: Sorting Slices
draft: false
lastmod: 2023-03-31
weight: 3
summary: How to sort a slice and do a binary search.
---

# Sort a Slice

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;"sort"<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Person&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Name&nbsp;<span class="golang-variable-type">string</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Age&nbsp;&nbsp;<span class="golang-variable-type">int</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">New</span><span class="golang-brace">(</span>name&nbsp;<span class="golang-variable-type">string</span>,&nbsp;age&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;*Person&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&Person<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Name:&nbsp;name,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Age:&nbsp;&nbsp;age,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">SortPeople</span><span class="golang-brace">(</span>people&nbsp;[]*Person<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">sort.Slice</span><span class="golang-brace">(</span>people,&nbsp;<span class="golang-top-level-keyword">func</span><span class="golang-brace">(</span>i,&nbsp;j&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-variable-type">bool</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;people[i].Name&nbsp;==&nbsp;people[i].Name&nbsp;<span class="golang-brace">{</span><br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;This&nbsp;is&nbsp;descending&nbsp;order&nbsp;by&nbsp;age.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;people[i].Age&nbsp;>&nbsp;people[j].Age<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;ascending&nbsp;order&nbsp;by&nbsp;name</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;people[i].Name&nbsp;<&nbsp;people[j].Name<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">})</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Run</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;people&nbsp;:=&nbsp;[]*Person<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">New</span><span class="golang-brace">(</span>"bob",&nbsp;<span class="golang-primitive-value">30</span><span class="golang-brace">)</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">New</span><span class="golang-brace">(</span>"jane",&nbsp;<span class="golang-primitive-value">24</span><span class="golang-brace">)</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">New</span><span class="golang-brace">(</span>"bob",&nbsp;<span class="golang-primitive-value">17</span><span class="golang-brace">)</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">New</span><span class="golang-brace">(</span>"jane",&nbsp;<span class="golang-primitive-value">46</span><span class="golang-brace">)</span>,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">SortPeople</span><span class="golang-brace">(</span>people<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}
