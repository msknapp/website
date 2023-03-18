---
title: "Sudo"
weight: 2
draft: false
description: passwordless sudo
summary: How to setup passwordless sudo.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---
instead of editing /etc/sudoers, it's better to add a file in 
/etc/sudoers.d

if they want sudo access with a password add:

```
<user> ALL=(ALL) ALL
```
if you want access without any password prompts:
```
<user> ALL=(ALL) NOPASSWD: ALL
```
if you want everybody to have sudo access with their own password prompt:
```
ALL  ALL=(ALL) ALL
```
if you want everybody to have sudo access without password prompts:
```
ALL  ALL=(ALL) NOPASSWD: ALL
```
The file should be read only to the user and group:
```
chmod 440 /etc/sudoers.d/<myfile>
```
on some linux distributions it should be 400 (no group reads).


