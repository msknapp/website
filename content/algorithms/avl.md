---
title: AVL Tree
weight: 25
draft: false
lastmod: 2023-03-25
summary: A binary search tree that stays balanced.
---


# Python Example

{{< gocode >}}
<br>
def&nbsp;right_rotation(node):<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">x</span>&nbsp;=&nbsp;node.left<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">t2</span>&nbsp;=&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;x:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">t2</span>&nbsp;=&nbsp;x.right<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;x.<span class="golang-variable">right</span>&nbsp;=&nbsp;node<br>
&nbsp;&nbsp;&nbsp;&nbsp;node.<span class="golang-variable">left</span>&nbsp;=&nbsp;t2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;x<br>
<br>
def&nbsp;left_rotation(node):<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">y</span>&nbsp;=&nbsp;node.right<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">t2</span>&nbsp;=&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;y:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">t2</span>&nbsp;=&nbsp;y.left<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;y.<span class="golang-variable">left</span>&nbsp;=&nbsp;node<br>
&nbsp;&nbsp;&nbsp;&nbsp;node.<span class="golang-variable">right</span>&nbsp;=&nbsp;t2<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;y<br>
<br>
<br>
class&nbsp;Node:<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;__init__(self,&nbsp;value,&nbsp;less,&nbsp;equal):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">value</span>&nbsp;=&nbsp;value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">less</span>&nbsp;=&nbsp;less<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">equal</span>&nbsp;=&nbsp;equal<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">left</span>&nbsp;=&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">right</span>&nbsp;=&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">height</span>&nbsp;=&nbsp;1<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;get(self,&nbsp;value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.equal(value,&nbsp;self.value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.less(value,&nbsp;self.value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.left.get(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.right.get(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;None<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;has(self,&nbsp;value:&nbsp;<span class="golang-variable-type">int</span>)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.get(value)&nbsp;is&nbsp;not&nbsp;None<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;get_size(self)&nbsp;->&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">s</span>&nbsp;=&nbsp;1<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;s&nbsp;+=&nbsp;self.right.get_size()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;s&nbsp;+=&nbsp;self.left.get_size()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;s<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;traverse(self):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;x&nbsp;in&nbsp;self.left.traverse():<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;yield&nbsp;x<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;yield&nbsp;self.value<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;x&nbsp;in&nbsp;self.right.traverse():<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;yield&nbsp;x<br>
&nbsp;&nbsp;&nbsp;&nbsp;<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;insert(self,&nbsp;value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.equal(value,&nbsp;self.value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.less(value,&nbsp;self.value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">left</span>&nbsp;=&nbsp;self.left.insert(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">left</span>&nbsp;=&nbsp;Node(value,&nbsp;less=self.less,&nbsp;equal=self.equal)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">right</span>&nbsp;=&nbsp;self.right.insert(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">right</span>&nbsp;=&nbsp;Node(value,&nbsp;less=self.less,&nbsp;equal=self.equal)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.update_height()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.rebalance()<br>
&nbsp;&nbsp;&nbsp;&nbsp;<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;update_height(self):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">x</span>&nbsp;=&nbsp;0<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">x</span>&nbsp;=&nbsp;self.right.height<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left&nbsp;and&nbsp;self.left.height&nbsp;>&nbsp;x:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">x</span>&nbsp;=&nbsp;self.left.height<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">height</span>&nbsp;=&nbsp;x+1<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;height_factor(self)&nbsp;->&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hr</span>&nbsp;=&nbsp;height_of(self.right)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hl</span>&nbsp;=&nbsp;height_of(self.left)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;hl&nbsp;-&nbsp;hr<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;rebalance(self):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">left</span>&nbsp;=&nbsp;self.left.rebalance()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">right</span>&nbsp;=&nbsp;self.right.rebalance()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">hf</span>&nbsp;=&nbsp;self.height_factor()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;hf&nbsp;>&nbsp;-2&nbsp;and&nbsp;hf&nbsp;<&nbsp;2:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;self<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;hf&nbsp;<&nbsp;-1:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;right&nbsp;sub-tree&nbsp;is&nbsp;higher&nbsp;than&nbsp;left<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.right.height_factor()&nbsp;<&nbsp;0:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;RR&nbsp;<span class="golang-control-keyword">case</span>,&nbsp;left&nbsp;rotation<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;left_rotation(self)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;RL&nbsp;<span class="golang-control-keyword">case</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">right</span>&nbsp;=&nbsp;right_rotation(self.right)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;left_rotation(self)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;hf&nbsp;>&nbsp;1:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;left&nbsp;sub-tree&nbsp;is&nbsp;higher&nbsp;than&nbsp;right<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.left.height_factor()&nbsp;>&nbsp;0:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;LL&nbsp;<span class="golang-control-keyword">case</span>&nbsp;->&nbsp;right&nbsp;rotation<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;right_rotation(self)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#&nbsp;LR&nbsp;<span class="golang-control-keyword">case</span><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">left</span>&nbsp;=&nbsp;left_rotation(self.left)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-variable">tmp</span>&nbsp;=&nbsp;right_rotation(self)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.update_height()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;tmp<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<br>
<br>
def&nbsp;height_of(nd:&nbsp;Node)&nbsp;->&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;nd:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;nd.height<br>
&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;0<br>
<br>
class&nbsp;Tree:<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;__init__(self,&nbsp;less,&nbsp;equal):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">head</span>&nbsp;=&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">less</span>&nbsp;=&nbsp;less<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">equal</span>&nbsp;=&nbsp;equal<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;insert(self,&nbsp;value):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;self.head:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">head</span>&nbsp;=&nbsp;Node(value,&nbsp;less=self.less,&nbsp;equal=self.equal)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">else</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;self.<span class="golang-variable">head</span>&nbsp;=&nbsp;self.head.insert(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;get(self,&nbsp;value:&nbsp;<span class="golang-variable-type">int</span>)&nbsp;->&nbsp;Collatz:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.head:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.head.get(value)<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;None<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;has(self,&nbsp;value:&nbsp;<span class="golang-variable-type">int</span>)&nbsp;->&nbsp;<span class="golang-variable-type">bool</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.get(value)&nbsp;is&nbsp;not&nbsp;None<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;get_size(self)&nbsp;->&nbsp;<span class="golang-variable-type">int</span>:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;self.head:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;self.head.get_size()<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;0<br>
<br>
&nbsp;&nbsp;&nbsp;&nbsp;def&nbsp;traverse(self):<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">if</span>&nbsp;not&nbsp;self.head:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">return</span>&nbsp;None<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="golang-control-keyword">for</span>&nbsp;x&nbsp;in&nbsp;self.head.traverse():<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;yield&nbsp;x<br>
<br>
{{< /gocode >}}