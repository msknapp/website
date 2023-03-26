---
title: Radix
draft: false
weight: 29
lastmod: 2023-03-25
summary: The radix tree
---

{{< gocode >}}
<br>
def&nbsp;longest_common_base(x:&nbsp;str,&nbsp;y:&nbsp;str)&nbsp;->&nbsp;str:<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;x&nbsp;or&nbsp;not&nbsp;y:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;""<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">i</span>&nbsp;=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;while&nbsp;i&nbsp;<&nbsp;len(x)&nbsp;and&nbsp;i&nbsp;<&nbsp;len(y):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x[i]&nbsp;!=&nbsp;y[i]:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x[:i]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;i&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x[:i]<br>
<br>
<br>
def&nbsp;one_prefixes_other(x:&nbsp;str,&nbsp;y:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;x&nbsp;or&nbsp;not&nbsp;y:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;True<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;first_prefixes_second(x,&nbsp;y)&nbsp;or&nbsp;first_prefixes_second(y,&nbsp;x)<br>
<br>
def&nbsp;first_prefixes_second(x:&nbsp;str,&nbsp;y:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;x:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;True<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;y:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;False<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(y)&nbsp;<&nbsp;len(x):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;False<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;y[:len(x)]&nbsp;==&nbsp;x<br>
<br>
class&nbsp;Node:<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;__init__(self,&nbsp;value:&nbsp;<span class="golang-variable">str</span>&nbsp;=&nbsp;"",&nbsp;leaf:&nbsp;<span class="golang-variable"><span class="golang-variable-type">bool</span></span>&nbsp;=&nbsp;True):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;isinstance(value,&nbsp;str):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;raise&nbsp;"wrong&nbsp;value&nbsp;<span class="golang-control-keyword">type</span>"<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">value</span>&nbsp;=&nbsp;value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">leaf</span>&nbsp;=&nbsp;leaf<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">children</span>&nbsp;=&nbsp;{}<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">size</span>&nbsp;=&nbsp;0<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;is_node(self)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;not&nbsp;self.leaf<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;insert_leaf_node_without_checking(self,&nbsp;new_node):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.leaf:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;raise&nbsp;"cannot&nbsp;insert&nbsp;into&nbsp;a&nbsp;leaf"<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.children[new_node.value]&nbsp;=&nbsp;new_node<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;contains_leaf(self,&nbsp;s:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;s&nbsp;in&nbsp;self.children:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;n:&nbsp;<span class="golang-variable">Node</span>&nbsp;=&nbsp;self.children[s]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;n.leaf<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;child_string&nbsp;in&nbsp;self.children:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;first_prefixes_second(child_string,&nbsp;s):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">sub</span>&nbsp;=&nbsp;s[len(child_string):]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;n:&nbsp;<span class="golang-variable">Node</span>&nbsp;=&nbsp;self.children[child_string]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;n.contains_leaf(sub)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;False<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;insert_and_check(self,&nbsp;new_string:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;<span class="golang-control-keyword">return</span>&nbsp;true&nbsp;<span class="golang-control-keyword">if</span>&nbsp;one&nbsp;word&nbsp;is&nbsp;found&nbsp;to&nbsp;be&nbsp;a&nbsp;prefix&nbsp;of&nbsp;another.<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.leaf:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;raise&nbsp;"cannot&nbsp;insert&nbsp;into&nbsp;a&nbsp;leaf"<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">common_base</span>&nbsp;=&nbsp;""<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;child_string&nbsp;in&nbsp;self.children:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;first_prefixes_second(new_string,&nbsp;child_string):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;True<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">child_node</span>&nbsp;=&nbsp;self.children[child_string]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;child_node.is_node():<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;first_prefixes_second(child_string,&nbsp;new_string):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;This&nbsp;does&nbsp;not&nbsp;mean&nbsp;there&nbsp;is&nbsp;a&nbsp;dupicate,&nbsp;because&nbsp;it&nbsp;is&nbsp;a&nbsp;node.<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">sub</span>&nbsp;=&nbsp;new_string[len(child_string):]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;child_node.insert_and_check(sub)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;it&nbsp;is&nbsp;a&nbsp;leaf.&nbsp;&nbsp;Check&nbsp;<span class="golang-control-keyword">if</span>&nbsp;this&nbsp;leaf&nbsp;is&nbsp;a&nbsp;prefix&nbsp;of&nbsp;the&nbsp;new&nbsp;value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;first_prefixes_second(child_string,&nbsp;new_string):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;True<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">b</span>&nbsp;=&nbsp;longest_common_base(child_string,&nbsp;new_string)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(b)&nbsp;>&nbsp;len(common_base):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">common_base</span>&nbsp;=&nbsp;b<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;len(common_base)&nbsp;>&nbsp;0:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">new_child</span>&nbsp;=&nbsp;Node(common_base,&nbsp;leaf=False)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">to_del</span>&nbsp;=&nbsp;[]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;child_string&nbsp;in&nbsp;self.children:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;node:&nbsp;<span class="golang-variable">Node</span>&nbsp;=&nbsp;self.children[child_string]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;node.is_node():<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">continue</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;first_prefixes_second(common_base,&nbsp;child_string):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;to_del.append(child_string)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;node.<span class="golang-variable">value</span>&nbsp;=&nbsp;node.value[len(common_base):]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;new_child.insert_leaf_node_without_checking(node)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;don't&nbsp;forget&nbsp;to&nbsp;still&nbsp;insert&nbsp;the&nbsp;word&nbsp;that&nbsp;just&nbsp;arrived.<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">sub</span>&nbsp;=&nbsp;new_string[len(common_base):]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;new_child.insert_leaf_node_without_checking(Node(sub,&nbsp;leaf=True))<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;s&nbsp;in&nbsp;to_del:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;del&nbsp;self.children[s]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.children[common_base]&nbsp;=&nbsp;new_child<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.children[new_string]&nbsp;=&nbsp;Node(new_string,&nbsp;leaf=True)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.size&nbsp;+=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;False<br>
<br>
<br>
<br>
class&nbsp;Radix:<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;__init__(self):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">head</span>&nbsp;=&nbsp;Node(<span class="golang-variable">value</span>&nbsp;=&nbsp;"",&nbsp;<span class="golang-variable">leaf</span>&nbsp;=&nbsp;False)<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;insert_and_check(self,&nbsp;v:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;<span class="golang-control-keyword">return</span>&nbsp;true&nbsp;<span class="golang-control-keyword">if</span>&nbsp;one&nbsp;word&nbsp;is&nbsp;found&nbsp;to&nbsp;be&nbsp;a&nbsp;prefix&nbsp;of&nbsp;another.<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.head.insert_and_check(v)<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;size(self)&nbsp;->&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.head.size<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;contains_leaf(self,&nbsp;s:&nbsp;str)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.head.contains_leaf(s)
{{< /gocode >}}