---
title: "Software Languages"
weight: 7
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---

This attempts to document languages not as how they work individually, but instead as how they 
deviate from a standard core programming language.  It starts with the most common language
features and functionality, and then shows how certain languages differ.

We should consider the following languagese: C, C++, java, javascript, python, typescript, golang, bash, Ruby, Rust

# Operators
* almost all languages have very similar operators.
* comparison: <, >, <=, >=, ==, !=, etc.
  * in C, C++, these return ints, not a boolean.
  * In Python, it also supports the "is" and "not" keywords.
  * In javascript, it also support ===, and !==, which is checking the 
    type of object as well.
* Assignment:
  * most languages support +=, -=, *=, /=, %=, |=, ^=, ~=, etc.
* Increment, decrement:
  * almost all languages support ++ and -- for incrementing/decrementing integers.
  * python is the exception, it does not support ++ or --, it supports += and -= though.
  * all support += and -=

# Types
* the way types are handled is split along patterns:
* statically typed languages:
  * C, C++, java, golang, Rust
  * optionally typed: Scala, Kotlin, Python, Typescript.
  * Split by two patterns, types before or after variable names:
    * types before variables and functions are used in C, C++, java
    * types after variables are used in Python (unenforced), Typescript, Golang, Rust, and Scala.
      * delimited with a colon: Scala, Rust, Python, Typescript
      * delimited with a space: Golang
* dynamically typed languages:
  * javascript (but Typescript adds types)
  * python (python supports type declarations but does not enforce them).
  * these do not declare the types of variables, and usually they can be re-assigned
    to completely different types.

# Generics
* Using <>: java, Rust,
* Using []: Scala, Golang
* Unsupported: older Golang, C, C++

# Loops
* all languages support keywords "for".
  * standard for loop: for (exp0; exp1; exp2) {...}
    is supported by most languages, with some deviations:
    * golang does not require the parentheses.
    * python does not support that format.
  * python standard for loop: for x in range(n): ...
* almost all languages support a lowercase "break" or "continue".
* only golang does not support "while", it uses for as its while loop.
* for with collections is supported in most languages:
  * java: for (int x : coll) {...;}
  * Python: for x in arr:...
  * golang: for index, x := range coll {...}
  * C, C++: unsure.
  * javascript: for (let x in coll) {...;}
    for (let x of coll) {...;}
    * "in" iterates over indexes.
    * "of" iterates over values.

# Code Structure
* Most languages use squirrel brackets to establish blocks.  Python and F# deviate 
  from this, using indentation instead.
* Semicolons are:
  * required in C, C++, java, Rust
  * recommended in javascript and typescript
  * allowed in Golang
  * not used in Python, Scala, F#, Ruby

# Lambda Functions
* Almost all languages allow functions to be first class citizens, objects, so they can 
  be assigned to variables and passed around.  Older versions of java (< 1.8) did not.
* Some support a short-hand form of functions, lambda expressions:
  * java: () -> {...;}
  * Python:
  * javascript, scala: () => {...}
  * 

# Switch Statements
* most languages support the standard:
  switch (expr) { case 0: ...; break; case 1: ...; break; default: ...; }
* some minor deviations:
  * golang doesn't require the parentheses, nor semicolons.
  * some languages don't require the "break" statements, but most do.
* Ruby uses "case", "when", and "end".  Break is not needed.
* Base uses "case" and "esac", conditions are string patterns followed by ")", and need ";;" to close them.
* Python does not have a switch statement, you can use if, elif, else instead.

# Constants
* In C, C++, javascript, golang: const keyword prefix
* In java, scala, groovy, jython, jruby, etc.: final keyword prefix.  Allowed in arguments too.
* In python, bash: not supported.

# Objects, Classes
* not supported in C, but you can use structs as args in functions.

# Variable Protections
* Most languages support them in some manner.
* Python does not enforce them, they are accomplished by the convention of prefixing
  variables with underscores.
* Javascript also has weak enforcement, using underscores to indicate a private field.
* Rust supports "pub" instead of public.  It tends to be very terse.

# Comments
* fall into two categories:
  * Using "//" and "/*  */": C, C++, java, javascript, typescript, Golang, F#, most squirrel bracket languages.
  * Using "#": Python, bash.

# Compilation
* System Languages: Compiles to OS bytecode: C, C++, Golang, Rust, 
* Compiles to a virtual runtime bytecode: java, scala, groovy, jython, jruby, 
  * most system languages can also be compiled into the WASM format.
* Interpreted or Scripted Languages: Python, Ruby, javascript, Lua, bash

# Garbage Collection
* Falls into two categories:
  * manually done: C, C++
  * Garbage collected: everything else.

# Use of Pointers, Passing by Value or Reference
* C, C++, Golang, and Rust support pointers.  So you have explicit control over 
  passing things by value or by reference.
  * Use "&" to get the memory address of a variable, which is a pointer.
  * C, C++, pointer types: <type>* <variablename>
    or: <type> *<variablename>
  * Golang pointer types: *<type>
  * de-reference a pointer variable, in all languages: *<variablename>
* In java, primitives are always passed by value, and objects are always passed by reference.
* In most scripting languages, passing by reference is normal.

# Pre-processing
* C, C++ support pre-processors, other languages do not typically use that.

# Hoisting, Code order requirements, finding code
* In C and C++, the compiler is not intelligent enough to use code before it has been declared.
  You must declare a definition of variables and functions before they are referred to.
  This is why you need to #include the header files from imports.
* In almost everything else, the compiler is able to find and use code as it is called, even if
  its definition comes later in the file or in a completely different file.

# Declaring Variables
* Without initialization:
  * C, C++, java: <type> <name>;
  * Go: <var|const> <name> <type>;
  * Rust:
  * 
# Naming Conventions
* Variables:
  * C, C++, java, Golang: lowerCamelCase
  * Python: snake_case
* Functions:
  * C, C++, java: lowerCamelCase
  * Go: 
    * public: UpperCamelCase
    * private: lowerCamelCase
  * Python: snake_case
* Packages, modules, directories...

# Strings
* Single quotes are characters: C, C++, java, Golang
* double quotes are strings in every language.
* single quotes are also strings: Python, Ruby, bash, javascript, typescript
* interpolated strings:
  * java: more recently supported with triple back-ticks
  * scala: supported with pipes and triple quotes.
* are primitives: in golang
* are objects: in java and all higher level (scripted) languages.
* C, C++ don't have a string type, you must use arrays of characters: char s[] = "foo";
  can still use double-quotes though, and don't need to declare the size.
  hence it is mutable.  They end with \0, the null terminating character.  That counts
  toward its size.
* in languages where strings are objects, they tend to have a lot of methods built in.
  In other languages, string functions tend to be in a separate package, like "strings" (golang),
  or in #include <string.h> in C, C++

# Associative Arrays
* aka dict, maps, hashes.
* Golang: x := map[string]string{"foo": "bar"}
* C, C++:
* java: Map<String, String> x = new HashMap<>();
* ruby: x = {'a' => 'x', 'b' => 'y'}

# Regular Expressions
* Using the /.../ syntax: javascript, ruby, 
* Other languages have a class for it.

# Keywords not, and, or
* supported in Python, Ruby, 

# Functional Programming Patterns
* supported in java > 8
* Python has some support with lambdas.
* Golang has limited support.
* Scala, Haskell, Kotlin, Ruby all have very strong support for functional programming techniques.
* Javascript has decent support, has the forEach method.

# Null Type
* In Golang, Ruby: nil
* In java, C, C++, javascript: null
* Javascript also supports "undefined"

# What is considered falsy
* Javascript is flexible: undefined, null, empty strings, 0, false
* Python is the only language that uses capitals for their boolean types: True, False
* C, C++, the expression must produce 0 or an integer that is not 0
* Java: the expression must produce a boolean

# Boolean Types
* Java: boolean
* Golang: bool
* C, C++: no built-in type, <stdbool.h> adds the type but it's actually an int 0 or 1.
* Python: bool, True or False

# Declaring Functions
* javascript: function name() {...;}
* java: [private|public|protected] [static] [synchronized] <type|void> name() {...;}
* C, C++: <type|void> name() {...;}
* golang: func name() [type] {...}
* Rust: fn name() {...}
* Python: def name():
* Ruby: def name
  or: def name( args )
* bash: function name {...}
  or name() {...}

# Arrays
* C, C++: int x[] = {1, 2, 3};
* java: int[] x = {1, 2, 3};
* golang: 
  * array: var [size]int x
  * slice: []int x = {1, 2, 3}
* python: x = [1, 2, 3]
* rust: let mut arr: [i32, 3] = [0; 3];
* javascript: const x = [1, 2, 3];
* in practically every language, array members are accessed like: arrName[i]
* Array Literals:
  * declared with {}: C, C++, java
  * declared with []: Python, javascript
* Constant Size:
  * C, C++, Golang, Rust, java, most system languages.
* Dynamic size:
  * Golang slices, Python lists, java collections like Lists.
  * In javascript you can assign to indexes outside its current range and it just extends
    to meet it.

# Code Structuring
* There are three patterns:
  * Squirrel bracket blocks: most languages
  * Indentation: Python, F#
    * Tends to use four spaces as indentation.
  * Token based: Ruby, Bash
    * blocks end when a token is reached like "end", "fi", "done", "esac", etc.

# Calling Package Functions
* Ruby, C, C++: <package>::<function>
* Python: <package>.<function>  (depends on the import statement used)
* Java: not possible, must refer to the class's static functions.

# Instantiation
* Using "new" keyword: java, C, C++, Go (optional, type is in parens)
* Using static type's new function: Ruby
* Using type's name: Python, javascript
  * implies there can only be one constructor.  These languages tend to support
    optional arguments with defaults for functions/methods.

# Support for Default Arguments, or function overloading
* java: functions can be declared multiple times with different parameters.
* Python, javascript: functions can be declared with default values for parameters.

# Exceptions, Errors, try catch
* Golang methods return errors, which can be ignored or re-thrown.
  It also supports panic, and try catch
* java, supports checked or unchecked exceptions.  functions can declare they throw 
  some exceptions.  Supports try catch finally.  Supports try with resources.
* Python: errors are always unchecked.  Supports try and catch.

# Scope Rules
* most languages say a variable exists within its code block, class, or globally.
* javascript allows "var" variables to exist after their block finishes.
* Python and bash also allow their variables to exist after their block finishes.

