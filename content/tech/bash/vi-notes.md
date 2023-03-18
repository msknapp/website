---
title: "Vi Notes"
draft: false
weight: 6
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

# Review of vi Commands

* 'a' will go into insert mode, but one character after the curser, it's great for appending to the end of a line.
* 'o' will go into insert mode, but it automatically inserts a new line beneath the cursor, and moves the cursor there.
* 'd$' will delete to the end of the line
* 'dw' will delete to the end of the current word.
* '0' will move the cursor to the start of the current line.
* '$' will move the cursor to the end of the current line.
* 'u' will undo the last operation.
* 'ctrl+r' will redo.
* 'p' will paste what you recently deleted or yanked.
* 'v' is for visual mode, it lets you highlight things for future actions.
* 'y' is for "yanking" (copying) text.
* 'yw' will yank one word.  likewise, you can combine it with counts and motions.
* 'yy' will copy the current line.
* 'dd' will delete the current line (you can paste it later with p, so this is effectively a 'cut' operation).
* '/' lets you search
* '?' lets you search in reverse.
* 'gg' moves to the start of the file.
* 'G' moves to the end of the file.
* 'dG' deletes everything from the cursor to the end of the file.
* 'dgg' deletes everything from the start of the file to the cursor.
* 'ctrl+g' will show you the current line number, and percent through the file.
* 'A' appends at the end of the line.  Useful since I frequently do this.
* 'I' inserts at the start of the line.

## Special Commands
* ":set nu" shows line numbers
* ":set nonu" hides line numbers
* ":split <filename>" splits the editor into two windows, and shows the other file in one of them.
* "ctrl-w ctrl-w" will switch between open windows. You can even have more than two.

## Motions
* "0" start of line
* "$" end of line
* "w" a 'word', moves to the start of a future word.
* "e" an 'end', moves to the end of a future word.
* "b" backwards one word, moves to the start of a word.
* "G" end of the file.
* "gg" start of the file.

# Less important shortcuts
* 'x' will delete the character under the cursor, without being in any mode. easily worked around by using the delete character instead.
* 'r' will let you replace the single character under the cursor, without being in any mode. easily worked around by using 'i' to insert text.
* 'c' let's you "change" things, it takes a motion.  It spares you from typing one extra key.

# Indent a bunch of lines

1. Go to the first line
2. ctrl+v
3. move up/down to the last/first line.
4. type shift + 'i'
5. type in your prefix
6. hit escape 

# Append to a bunch of lines

1. Go to the first line
2. ctrl + v --> it's important to use 'ctrl' here, not just 'v'!
3. move up/down to the last/first line
4. type End or $ to extend selection to the end of the lines
5. type shift + 'a' to append
6. type the suffix
7. hit escape

# Copy a Line

1. go to the line
2. type 'yy' (yank yank)
3. go where you want to insert it
4. type p

# Cut a Line

1. go to the line
2. type 'dd' (delete delete, it actually copies the line into memory)
3. go where you want to insert it
4. type p
