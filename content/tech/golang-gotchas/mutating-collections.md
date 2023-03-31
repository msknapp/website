---
title: Mutating Collections
draft: false
weight: 3
lastmod: 2023-03-31
summary: You might be surprised by the results of a function mutating a collection.
---


{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">//&nbsp;PrintSlice&nbsp;will&nbsp;print&nbsp;all&nbsp;the&nbsp;values&nbsp;of&nbsp;the&nbsp;slice&nbsp;to&nbsp;standard&nbsp;output.</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">PrintSlice</span><span class="golang-brace">(</span>x&nbsp;[]<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;y&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;x&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"%d,&nbsp;",&nbsp;y<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>"\n----"<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="code-comment">//&nbsp;Add&nbsp;attempts&nbsp;to&nbsp;append&nbsp;a&nbsp;value&nbsp;to&nbsp;the&nbsp;slice.</span><br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">Add</span><span class="golang-brace">(</span>x&nbsp;[]<span class="golang-variable-type">int</span>,&nbsp;y&nbsp;<span class="golang-variable-type">int</span><span class="golang-brace">)</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;<span class="golang-function">append</span><span class="golang-brace">(</span>x,&nbsp;y<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;:=&nbsp;[]<span class="golang-variable-type">int</span><span class="golang-brace">{</span><span class="golang-primitive-value">1</span><span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Add</span><span class="golang-brace">(</span>x,&nbsp;<span class="golang-primitive-value">2</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">PrintSlice</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;expand&nbsp;the&nbsp;slice&nbsp;to&nbsp;its&nbsp;full&nbsp;capacity.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;x[:<span class="golang-function">cap</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span>]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">PrintSlice</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;restart,&nbsp;this&nbsp;time&nbsp;using&nbsp;make&nbsp;and&nbsp;a&nbsp;pre-chosen&nbsp;capacity.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;<span class="golang-function">make</span><span class="golang-brace">(</span>[]<span class="golang-variable-type">int</span>,&nbsp;<span class="golang-primitive-value">0</span>,&nbsp;<span class="golang-primitive-value">2</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;<span class="golang-function">append</span><span class="golang-brace">(</span>x,&nbsp;<span class="golang-primitive-value">1</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Add</span><span class="golang-brace">(</span>x,&nbsp;<span class="golang-primitive-value">2</span><span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">PrintSlice</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
<br>
<span class="code-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;expand&nbsp;the&nbsp;slice&nbsp;to&nbsp;its&nbsp;full&nbsp;capacity.</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;x&nbsp;=&nbsp;x[:<span class="golang-function">cap</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span>]<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">PrintSlice</span><span class="golang-brace">(</span>x<span class="golang-brace">)</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

Prints:
```
1, 
----
1, 
----
1, 
----
1, 2, 
```

What on earth is going on here?

So a slice in Golang is actually a pointer to an array with a fixed size.  The slice records a virtual
start and end index of the array.  The array has a fixed size.

When you have a slice append a value, the behavior depends on if the underlying array is large
enough to accept the new value.

When the "append" operation is called, and the underlying array lacks the capacity for a new item:
* It creates a new array, with a larger capacity
* Then it copies all values from the old array to the new one
* Then it inserts the new value that was just appended.
* Then it has the slice point to that new array.

When a slice is passed to a function, you are passing the slice by value, which means it is copied in,
*but* the slice itself holds a pointer (reference) to the underlying array.  So this means the 
operations performed by the called function may or may not be visible to higher functions.

Operations visible to higher code on the stack:
* altering a value in the slice

Not visible:
* appending a value

So let's step through the example.

**Step 1**
```
x := []int{1}
Add(x, 2)
```

When the "Add" function is called, it tries to append to the slice it has.  The underlying
array of the slice is fixed to size 1, so the append operation creates a new array with 
at least a capacity of 2, and adds the new value.  This only altered the slice that is
known inside the "Add" function, which was a copy.  Consequently, higher functions have
no insight into the change made by the "Add" function, and it effectively does nothing.

**Step 2**

```
x = x[:cap(x)]
```

In the main function, the slice "x" was actually unaltered by the "Add" function, so its
local x still has a capacity of just 1.  Expanding it here does nothing, the change
made by the "Add" function before was made to a copy of the underlying array.

**Step 3**
```
x = make([]int, 0, 2)
x = append(x, 1)
Add(x, 2)
```

This time, we make a slice whose underlying array has a capacity of 2.  In the local
main function, we append the first value, and this alters the slice.  Essentially,
the slice x holds a start index of 0, and an end index of 1.

When the Add method is called, it receives a copy of the slice x, and this copy
holds a pointer to the underlying array of size 2.  Since the underlying array has
a capacity exceeding the length of the slice, the append operation in Add does not
need to make a new copy of the array.  It simply sets the value of the next item
in the underlying array, adds one to the length of its local slice, and returns.

However, the Add function had a *copy* of the slice, so the calling main code
was not affected by this.  main still has a slice that considers its length to 
be just 1, but it does point to the same underlying array.

**Step 4**
```
x = x[:cap(x)]
```

This time, the change made by the Add function becomes visible, because it altered
the underlying array of x.  The full capacity of the slice was 2 from the start,
so when the main code expands the length of the slice x, it actually sees the 
change made by the Add method before.