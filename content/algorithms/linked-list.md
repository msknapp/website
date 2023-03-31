---
title: Linked List
draft: false
weight: 6
lastmod: 2023-03-30
summary: Use a linked list for O(1) append, push, or pop operations.
---

A linked list is a collection where nodes form a chain.  Each node knows the next node in the linked
list.  That is a singly linked list.  If each node knows its previous node as well, that is called
a doubly-linked list, or a double-ended queue.  It's often called a Dequeue (pronounced "Deck") for short.

A linked list has pros and cons.

**Pros:**
* Inserts into the head or tail of the list are O(1), they happen almost immediately.
* Reads from the head or tail also happen immediately.

**Cons:**
* Iterating over a linked list can be slow.
* Searching for a single item in the list is O(n), even if the items are sorted.

A linked list is often the base for a Stack or Queue collection.  Very often these are the most
efficient collection type for scenarios like:
* binary tree searches:
  * a breadth first search of a binary tree uses a Queue.
  * a depth first search of a binary tree uses a Stack.
* backtracking algorithms.

# Generic Golang Example

The following is a double-ended queue that supports generic value types in Golang.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;golang<br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;node[T&nbsp;<span class="golang-variable-type">any</span>]&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;Value&nbsp;&nbsp;&nbsp;&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;*node[T]<br>
&nbsp;&nbsp;&nbsp;&nbsp;Previous&nbsp;*node[T]<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;newNode[T&nbsp;<span class="golang-variable-type">any</span>]<span class="golang-brace">(</span>value&nbsp;T<span class="golang-brace">)</span>&nbsp;*node[T]&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&node[T]<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Value:&nbsp;value,<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;LinkedList[T&nbsp;<span class="golang-variable-type">any</span>]&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;head&nbsp;*node[T]<br>
&nbsp;&nbsp;&nbsp;&nbsp;tail&nbsp;*node[T]<br>
&nbsp;&nbsp;&nbsp;&nbsp;size&nbsp;<span class="golang-variable-type">int</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;NewLinkedList[T&nbsp;<span class="golang-variable-type">any</span>]<span class="golang-brace">()</span>&nbsp;*LinkedList[T]&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;&LinkedList[T]<span class="golang-brace">{}</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Push</span><span class="golang-brace">(</span>value&nbsp;T<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;n&nbsp;:=&nbsp;<span class="golang-function">newNode</span><span class="golang-brace">(</span>value<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;l.size++<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.head&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.head&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.tail&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;n.Previous&nbsp;=&nbsp;l.tail<br>
&nbsp;&nbsp;&nbsp;&nbsp;l.tail.Next&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;l.tail&nbsp;=&nbsp;n<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">PushToFront</span><span class="golang-brace">(</span>value&nbsp;T<span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;n&nbsp;:=&nbsp;<span class="golang-function">newNode</span><span class="golang-brace">(</span>value<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;l.size++<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.head&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.head&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.tail&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;l.head.Previous&nbsp;=&nbsp;n<br>
&nbsp;&nbsp;&nbsp;&nbsp;n.Next&nbsp;=&nbsp;l.head<br>
&nbsp;&nbsp;&nbsp;&nbsp;l.head&nbsp;=&nbsp;n<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Pop</span><span class="golang-brace">()</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;out&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.tail&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;out&nbsp;=&nbsp;l.tail.Value<br>
&nbsp;&nbsp;&nbsp;&nbsp;l.tail&nbsp;=&nbsp;l.tail.Previous<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.tail&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.tail.Next&nbsp;=&nbsp;nil<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;l.size--<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">PopFront</span><span class="golang-brace">()</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;out&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.head&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;l.size--<br>
&nbsp;&nbsp;&nbsp;&nbsp;out&nbsp;=&nbsp;l.head.Value<br>
&nbsp;&nbsp;&nbsp;&nbsp;l.head&nbsp;=&nbsp;l.head.Next<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.head&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;l.head.Previous&nbsp;=&nbsp;nil<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Size</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">int</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;l.size<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Peek</span><span class="golang-brace">()</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.tail&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;out&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;l.tail.Value<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>l&nbsp;*LinkedList[T]<span class="golang-brace">)</span>&nbsp;<span class="golang-function">PeekFront</span><span class="golang-brace">()</span>&nbsp;T&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;l.head&nbsp;==&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;out&nbsp;T<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;out<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;l.head.Value<br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}