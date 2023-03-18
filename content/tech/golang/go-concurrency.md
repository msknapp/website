---
title: Golang Concurrency
weight: 2
draft: false
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

# Summary Notes

* use "go" keyword to make a function run in parallel.
* When main completes, it will not wait for other goroutines to finish, 
  the program stops.
* channels allow communication between goroutines.  make(chan <type>[, capacity])
  Also just declaring a channel makes it.
  * make(chan <type>)
  * by default, it is bi-directional, and blocking.  Sending to a channel will
    suspend the current goroutine until something reads from the channel.
  * these are also called "unbuffered" or "synchronous" channels.
* write to channel: ch <- v
* read from channel: v <- ch
* read but discard from channel: <- ch
* deadlock is possible if all goroutines are asleep.  It can happen if the 
  send and receive operations are not equal in number.
* Channels can be cast, even implicitly, to send-only or receive only channels:
  * send-only: chan<- <type>
  * read-only: <-chan <type>
  * read-only channels cannot be closed, only the write channels can be closed.
* when a channel is closed, it can still be read from.
* unbuffered channels have nowhere to store their data. 
* buffered channels will not block when written to, unless they are full.  
  They are made by giving a capacity to the make command: 
  * make(chan <type>, <capacity>)
  * example: make(chan int, 3)
* for loops can read from channels: for x := range ch {...}
* the "cap" function can determine the capacity of a channel.  Likewise, "len"
  can determine how many items are waiting in it. 
* they do not recommend using a channel as a queue data type, except when used
  with multiple goroutines.  Use a slice as a queue instead.


# Sync Package:
* WaitGroup makes a routine pause until other coroutines have finished.
  * wg := new(sync.WaitGroup)
    wg.Add(2)
  * in other goroutines, they call wg.Done() when they are finished. 
    Typically using defer. 