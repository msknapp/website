---
title: "Javascript"
draft: false
weight: 4
description: javascript programming
summary: Notes about the javascript programming language.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

* output with `console.log()`
* semicolons are recommended but optional.
* comment syntax is normal `//` or `/*  */`
* Uses `{ }` for code blocks.
* has "number" types, does not separate integers from floats.  
  So "//" is not integer division (Python), it is a comment.
* has "boolean" and "string".  No byte or rune type exists.
* has arrays.  "object" is their catch-all dictionary and class type.
* double or single quotes can be used to define strings.
* uses "=" for definition, never ":=" (golang, c?)
* uses standard `if (...) {...} else if (...) {...} else {...}` format.
* Has a switch statement, cases need a break statement or consecutive blocks run, 
  format: switch (...) { case 0: ...; break; case 1: ...; break; default: ...; }
  switch checks types too, like ===, so the type must also match.
* for loops should include "let" in first part.  for (let x=0;x < 3; x++) {...;}
* for loops support "in" syntax, and can work on arrays or objects (dicts),
  but the order is not preserved.  for (let x in arg) {...; }
  arr.forEach, Array.forEach, preserve order.
* for of loops are similar.  for (let x of arg) {...;}
* supports while and do while loops
* you can treat defined or undefined vars as truthy or falsy.
  the number 0, and empty strings are considered falsy.  "undefined" is falsy.
  null is falsy.  NaN is falsy.  The string "false" is truthy because it is not empty.
  What about empty arrays and objects?
* use ?? for defaulting vars.  Example, if x is undefined, y = x ?? "z"; will make y = "z"
  it's called nullish coalescing.  Works for null and undefined.  Empty strings, and zeros 
  might still get used.
* You can use "?." to avoid NPEs.  `a?.b?.c`
* javascript is flexible with args to functions.  You could call a function with 
  more or less args than expected.
* variable names can have "$" and "_".  Can't start with digit.
  conventionally uses lower camelcase.  Hyphens can't be used in var names.
  Note: upper camelcase is called Pascal case.
  Capitalization does not change visibility like golang.
* Use "const" for variables that should not change.  Use "var" when it can
  change and must run in older browsers.  Use "let" when it can change.
  * "var" can be re-declared later, "let" cannot.
  * "var" can even be re-declared with a different type.
  * "var" can't have block scope, "let" and "const" are block scoped.
  * "let" variables can be re-declared in sub-blocks.
* can use let to declare multiple variables, separated by commas.
* don't need to initialize the variables.  "let x;" is fine.
* private variables/functions don't exist but _ prefixes are used as a convention.
* $ is often used to indicate the main or primary function.
* "hoisting" means that variable declarations automatically move to the top of code,
  or top of their scopes/blocks.  "let" variables can't be initialized and hoisted.
* const means the variable cannot be re-assigned, but its contents can change.
* supports "**" exponentiation, "++" and "--", and "+=", "*=" etc.
* supports ternary operator ?
* "===" and "!==" check the types of variables too.
* check types with "typeof(x)" or "x.instanceof(object)" or similar.
* supports bitwise operators & | ~ ^ << etc.
* does implicit casting, like adding numbers and strings converts all numbers to strings.
* arrays use normal bracket notation: [1, "2", false, etc.] (json)
* objects use {} notation, properties are separated with colons (json)
* "undefined" is a type.  new Object() is equivalent to {}
  object key names don't need to be quoted when declared in the literal syntax.
* objects can have new properties, and even methods, added at any time.
* You can reference object properties with either the dot notation or index style, 
  obj.prop or obj["prop"] work.
* objects get methods as if they are variables with function values.
* supports the "this" keyword.
* "this" from code not owned by an object refers to the "global" object.  Set "use strict" to 
  make it be undefined instead.
* strings length is a property not a method, and not a function you call on it.
  myString.length is right.
* === will always be false for objects.
* "string" is a built-in primitive, "String" is an object declared with "new", and is not
  recommended for use, it's not the same type as string.
* supports string interpolation with back-ticks, it can refer to variables and other expressions 
  within back-ticks.
* 100 / "10" produces 10.  It automatically casts strings to numbers if needed.
* NaN and Infinity are special numbers supported.
* arrays can be declared with [] or new Array().  You don't need to declare their size, even
  when assigning items to an index.  typeof for arrays returns object.  Arrays have a 
  forEach method/function.  Use Array.isArray to check its type.  or x instanceof Array.
* arrays support "push".  You can assign to an index beyond its length, the gap is filled with undefined.
* arrays support "pop" or "shift", removing the first item.  Using "delete" results in undefined holes.
* arrays support functional programming, or streaming methods.
* supports break and continue with labels.  labels are followed with colons.  Can escape blocks
  that are not even loops.
* typeof new Date() - returns "object"
* typeof null - returns "object"
* null == undefined - is true.  They differ by their type only.  null is an object, undefined is not.
* numbers to strings: (3).toString() or String(3)
* strings to numbers: Number("3") works.  Number("") returns 0.  Number("x") returns NaN.
* supports try {...;} catch (err) {...;} finally {...;}
* you can use throw, its arg can be a string, number, boolean, or object.
* lambda methods are supported, called "arrow functions", () => { ...; return ...; },
  shorter syntaxes are allowed too: () => 3

Classes:
* use "class" keyword, capitalize them, have "constructor" method.  Don't need to 
  use object syntax where a var name refers to a function.  Can declare methods directly.
* not an object, is a template for one.
* use the new keyword to create them, calling its constructor automatically.
* supports inheritance with the extends keyword, and can refer to super constructor.
* supports special "get" and "set" methods, so properties can actually call methods.  
* class declarations are not hoisted, you cannot refer to them before they are defined
  in the code.
* get and set can also be used in object definitions.
* supports the static keyword so your class can have static functions.

Functions:
* can define with function keyword, or as an expression.  
  * when defined as an expression, they will never be hoisted.  Also they must be 
    followed with a semi-colon.
* functions can have properties and methods, so they are actually like objects.
* typeof will return "function".
* supports arrow functions: const x = (arg) => "foo"
  curly braces, and return statements, are only needed if there are multiple statements.
  * arrow functions are never hoisted, and they cannot use "this"
* you can call functions with fewer arguments than parameters, the missing ones are undefined.
  then the code could check for this and set a default value.
* you can establish defaults for parameters using "=".
* You can use "..." before an arg, as the last parameter.  It is passed as an array.
* you can call functions with more arguments than parameters.  The extra args don't get assigned
  to any local variables, but they are reachable with a special "arguments" object that all
  functions have.  No error is thrown if a function is called with excess arguments.
  "arguments" is an array, it has a length property.
* arguments that are not objects are always passed by value, changing them will not change higher scope values.
  objects passed to functions can be mutated.
* functions declared at the code root become methods of the "window" object.
* if you preceed a function call with "new", then it is treated as a constructor and actually 
  returns an object.  The return statement is implicit and not needed.  Use "this" to assign properties.
* for events, "this" refers to the html element that had the event.
* functions can be nested within functions, this enables having closures.  functions can refer
  to variables in their parent scopes.  Closures effectively let us have private variables.


Strict Mode:
* "use strict" goes at top of code, or at top of function block.
* variables must be declared before used, or else an error is thrown.
* enforces a lot of other things too.

Special Objects:
* new Set([...])
* new Map([ ["key", value], ... ]) - retains order of keys.


Web Objects:
* console - a log
* document - the web page
* window

Events: always lowercase and starting with "on", are bound to html elements.

Special Property Attributes:
* enumerable
* configurable
* writable
