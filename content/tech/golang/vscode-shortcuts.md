---
title: VSCode Shortcuts
weight: 3
draft: false
description: my description.
summary: Keyboard hotkeys for VSCode that are useful for golang programming.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---


# Fatal Hotkeys
Do not press:
* `ctrl+alt+f10` - caused the screen to go blank and was unrecoverable.
* `super+l` - locks the screen.
* `ctrl+q` - closes vscode
* `ctrl+w` - closes the current editor tab
* `alt f4` - closes the current vscode window.


* command pallette: `f1, or ctrl+shift+p`
* view shortcut keys: `ctrl + k then s`
* delete line: `ctrl+shift+k`
* suggestion: `ctrl+space`
* quick fix: `ctrl + .`
* find all references: `ctrl+alt+f12`
* refactor: `ctrl+shift+r`
* paremeter hints: `ctrl+shift+space`
* format code: `shift+alt+f`
* move lines up/down: alt+up/down arrows.
* peek at implementations: `ctrl+shift+f12`
* open definition to side: `ctrl + k then f12`
* insert new line and move cursor:
  * above line: `ctrl+shift+enter`
  * below line: `ctrl+enter`

# Most wanted
* view docs on function: `ctrl + k then i`
* view parameter types, or return type, of function: `ctrl + k then i`
* Give me help information about a command, the 
  return types it has, etc.:
  * `ctrl+shift+f10` (may be the best option)
  * called "editor.action.showHover": `ctrl + k then i`
  * it is the same as hovering your mouse over the function.
* Show me the source of a function.
  * linux: `F12`
  * mac: `ctrl+F12`
  * alternatively, peek: `ctrl+shift+f10`
    * warning: do not use `ctrl+alt+f10`, it breaks everything.
  * go back afterwards: `ctrl+alt+-`
* Add an import.
  * easy way: `ctrl+.` (quick fix), then add import.
  * I added/customized: `ctrl+alt+o`
* show me the file in the explorer
  * called "explorer.autoreveal" or "Reveal in explorer view"
  * Ubuntu has no shortcut, but you can map one.
  * right click on the tab, then choose "Reveal in explorer view"
  * `ctrl+alt+r` does that for java, I configured the keybinding for files
    instead of just java.

Adding a golang import doesn't have a shortcut by default, but you can 
easily add one:
* `ctrl+s` to open shortcuts
* search for "import"
* for the golang import, set it to: `ctrl+alt+o`
* you can also search to see what hotkeys are already in use.
