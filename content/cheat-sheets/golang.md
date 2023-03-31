---
title: Golang
lastmod: 2023-03-24
weight: 1
draft: false
summary: cheat sheet for Golang.
---
{{< cheat "Loop Over List" >}}<span class="golang-comment">//&nbsp;for&nbsp;example:</span><br>
<span class="golang-variable">x</span>&nbsp;:=&nbsp;[]<span class="golang-variable-type">int</span>{1,&nbsp;2,&nbsp;3}<br>
<span class="golang-control-keyword">for</span>&nbsp;_,&nbsp;<span class="golang-variable">y</span>&nbsp;:=&nbsp;<span class="golang-control-keyword">range</span>&nbsp;x&nbsp;{<br>
<span class="golang-comment">&nbsp;&nbsp;&nbsp;&nbsp;//&nbsp;...</span><br>
}<br>


{{< /cheat >}}
{{< cheat "New Map" >}}
<span class="golang-variable">x</span>&nbsp;:=&nbsp;<span class="golang-control-keyword">map</span>[<span class="golang-variable-type">string</span>]<span class="golang-variable-type">int</span>{}<br>
<span class="golang-variable">y</span>&nbsp;:=&nbsp;make(<span class="golang-control-keyword">map</span>[<span class="golang-variable-type">string</span>]<span class="golang-variable-type">int</span>,&nbsp;4)<br>
<span class="golang-variable">z</span>&nbsp;:=&nbsp;<span class="golang-control-keyword">map</span>[<span class="golang-variable-type">string</span>]<span class="golang-variable-type">int</span>{<br>
&nbsp;&nbsp;&nbsp;&nbsp;"one":&nbsp;1,<br>
&nbsp;&nbsp;&nbsp;&nbsp;"two":&nbsp;2,<br>
}<br>
{{< /cheat >}}

{{< cheat "String to Stream" >}}
s&nbsp;:=&nbsp;"foo"<br>
buf&nbsp;:=&nbsp;bytes.NewBuffer([]<span&nbsp;class="golang-variable-type">byte</span>(s))
{{< /cheat >}}
{{< cheat "Split a string" >}}
s&nbsp;:=&nbsp;"foo&nbsp;bar"<br>
parts&nbsp;:=&nbsp;strings.Split(s,&nbsp;"&nbsp;")
{{< /cheat >}}
{{< cheat "Read File to String" >}}
<span&nbsp;class="golang-top-level-keyword">func</span>&nbsp;Ex()&nbsp;(<span&nbsp;class="golang-variable-type">string</span>,&nbsp;error)&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;bts,&nbsp;e&nbsp;:=&nbsp;ioutil.ReadFile("/path/to/file.txt")<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span&nbsp;class="golang-control-keyword">if</span>&nbsp;e&nbsp;!=&nbsp;nil&nbsp;{<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span&nbsp;class="golang-control-keyword">return</span>&nbsp;"",&nbsp;e<br>
&nbsp;&nbsp;&nbsp;&nbsp;}<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span&nbsp;class="golang-control-keyword">return</span>&nbsp;<span&nbsp;class="golang-variable-type">string</span>(bts),&nbsp;nil<br>
}
{{< /cheat >}}
{{< cheat "Unmarshal JSON" >}}

{{< /cheat >}}
{{< cheat "Parse URL" >}}

{{< /cheat >}}
{{< cheat "Reverse a String" >}}

{{< /cheat >}}
{{< cheat "Read a File" >}}

{{< /cheat >}}
