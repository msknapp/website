---
title: Queue
draft: false
weight: 7
lastmod: 2023-03-25
summary: Use queues for faster push and pop operations.
---

A queue is a first-in, first-out data structure.  There are a variety of ways to represent them:
* As arrays
* As circular arrays
* As linked lists

# Array Backed Queues

The problem with arrays is that, as items come and go, the array may need to be re-built, which
can involve copying all of its elements.  This won't perform as well as linked lists, which have
O(1) append or pop operations.

For instance, one way a queue (backed by an array) could support a pop operation is by shifting all elements
one index to their left.  This could be a very expensive operation.

Another way an array-backed queue can support the pop operation, is by maintaining a pointer to the start
of the array.  As items are popped, the pointer just shifts right one index.  Then you still have 
O(1) pop operations.  A push onto the array might still perform slowly though.  If the underlying
array does not have space for another element, then it might have to build a new array and copy all
elements to that, which can be expensive.

You could also design the Queue so that after a certain number of operations, if cleans itself, shifting
all elements to the front of the underlying array.  This will give better performance most of the time, 
but also uses more memory.  It also doesn't mean the array will not need a resize.

If the programmer can safely and reasonably assume that the Queue will never exceed a certain size, 
then it might be safe to use an array for a Queue.

# Circular Arrays

A circular array maintains two pointers, one for the head, and one for the tail.  These can circle back
to the front of the underlying arrays.  So for example, if the array is six items long, and the tail 
points to index 5 (the last element), and the head points to index 1 (the second element), then this
can still support pushing another element to it.  The tail would then point to index 0.  This is what
I mean by circular, the pointers wrap around the array in this fashion.

The circular array might have a size limit, and if code tries pushing to the collection when it is
full, that could either:
* overwrite the head (probably not what you want).
* block the thread until space is available (which can lead to deadlock).
* throw an error.

Alternatively, the circular array could choose to resize in this situation.  Which means there is still
a chance that a push operation can be expensive, involving the copying of the array into a 
newer larger one.

# Linked Lists

A more common pattern is to use a linked list.  The Queue would then maintain pointers to the head and
tail of the list.  The advantages here are:
* Push and pop operations are always O(1).
* There is no size limit to the Queue.

The downside is that Linked Lists can generate more garbage, which might degrade the application performance.
Linked Lists are usually the best choice for unbounded Queues.

# Golang Generic Example

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