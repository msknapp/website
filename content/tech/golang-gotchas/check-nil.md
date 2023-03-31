---
title: Checking for Nil
draft: false
weight: 1
lastmod: 2023-03-31
summary: Programmers get baffled when a variable is considered not nil and then promptly panics because it is nil.
---

Checking if a variable is nil in Golang sounds easy, but there is a common gotcha here.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Named&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Name</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Thing&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;name&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*Thing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Name</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t.name<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;n&nbsp;Named<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;t&nbsp;*Thing<br>
&nbsp;&nbsp;&nbsp;&nbsp;n&nbsp;=&nbsp;t<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;n&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>"n&nbsp;is&nbsp;not&nbsp;nil"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"the&nbsp;name&nbsp;of&nbsp;n&nbsp;is:&nbsp;%s",&nbsp;<span class="golang-function">n.Name</span><span class="golang-brace">())</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

When the code above runs, it prints:
```
n is not nil
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x49edfd]

goroutine 1 [running]:
main.(*Thing).Name(0x0)
	.../main.go:16 +0x1d
main.main()
	.../main.go:26 +0xbb
```

What is going on here?  In golang, a reference to an interface is not nil if the implementation type is known.
There is a workaround for this problem.

{{< gocode >}}
<span class="golang-top-level-keyword">package</span>&nbsp;main<br>
<br>
<span class="golang-top-level-keyword">import</span>&nbsp;<span class="golang-brace">(</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;"fmt"<br>
&nbsp;&nbsp;&nbsp;&nbsp;"reflect"<br>
<span class="golang-brace">)</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Named&nbsp;<span class="golang-control-keyword">interface</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">Name</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-control-keyword">type</span>&nbsp;Thing&nbsp;<span class="golang-control-keyword">struct</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;name&nbsp;<span class="golang-variable-type">string</span><br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-brace">(</span>t&nbsp;*Thing<span class="golang-brace">)</span>&nbsp;<span class="golang-function">Name</span><span class="golang-brace">()</span>&nbsp;<span class="golang-variable-type">string</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;t.name<br>
<span class="golang-brace">}</span><br>
<br>
<span class="golang-top-level-keyword">func</span>&nbsp;<span class="golang-function">main</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;n&nbsp;Named<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">var</span>&nbsp;t&nbsp;*Thing<br>
&nbsp;&nbsp;&nbsp;&nbsp;n&nbsp;=&nbsp;t<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;n&nbsp;!=&nbsp;nil&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>"n&nbsp;is&nbsp;not&nbsp;nil"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;!<span class="golang-function">reflect.ValueOf</span><span class="golang-brace">(</span>n<span class="golang-brace">)</span>.<span class="golang-function">IsNil</span><span class="golang-brace">()</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Printf</span><span class="golang-brace">(</span>"the&nbsp;name&nbsp;of&nbsp;n&nbsp;is:&nbsp;%s",&nbsp;<span class="golang-function">n.Name</span><span class="golang-brace">())</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span>&nbsp;<span class="golang-control-keyword">else</span>&nbsp;<span class="golang-brace">{</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-function">fmt.Println</span><span class="golang-brace">(</span>"but&nbsp;the&nbsp;value&nbsp;of&nbsp;n&nbsp;is&nbsp;nil"<span class="golang-brace">)</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-brace">}</span><br>
<span class="golang-brace">}</span><br>
{{< /gocode >}}

When the above code runs, it prints:
```
n is not nil
but the value of n is nil
```
