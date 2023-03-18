---
title: "Bash Review"
draft: false
weight: 1
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
---
Review Questions:

What do these tell you:
	$$
	$?
	:
	$*
	$@
	$#
	((  ))
	$(  )
	${parameter}
	[  ]
	[[  ]]
	.
	source

What is the difference between [ ], [[ ]], ( ), and (( )) ?
How do you get the last column using gawk?
How do you make variables you set ALWAYS carry through to sub-shells? and undo that?
How do you make the '>' operator NEVER overwrite existing files? and undo that?
What keyboard shortcut clears the screen?
What keyboard shortcut exits a sub-shell?
How do you copy lines in vi?
How do you insert a prefix before many lines in vi?
How do you remove a prefix before many lines in vi?
In a bash script, how do you get input from a user?
How you you print a line without a new line?
How do you duplicate a line in vi?
In vi, how do you auto-indent? see line numbers?  and disable these features?
In vi, how do you set a bookmark and navigate between them?

==============================
Answers:

$$ gives you the process id of the current process

$? tells you the exit code of the last sub process
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ "false"
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $?
	1
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $?
	0
	

: is the no-op operator, it always exits with exit code of true

	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if ( : ); then echo "true" ; else echo "false" ; fi
	true
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if ( ! : ); then echo "true" ; else echo "false" ; fi
	false
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [ : ]; then echo "true" ; else echo "false" ; fi
	true
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [ ! : ]; then echo "true" ; else echo "false" ; fi
	false
	

$* means all arguments to the given script, not quoted, not including the script name
$@ is like $* except keeps text quoted.
$# tells you how many arguments there are to a bash script
. from command line means to execute the file.
source 
${parameter} gives you the value of the named variable parameter.  It does NOT evaluate the parameter as an expression.  For example:

	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo ${ls | head -1}
	bash: ${ls | head -1}: bad substitution
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $(ls | head -1)
	mesos.txt

$( ) will actually evaluate the expression inside it and return its output.  It is equivalent to using back-ticks, `expr`.  

'return' will exit a function with the provided exit code.  Zero is default.  If a function never has a return, then the exit code is that of its last operation.  Return statements do NOT set the value that is received by $(), you must use the standard output stream to accomplish that.

	#!/bin/bash
	
	function f1() {
	        return "foo"
	}
	
	function f2() {
	        echo "foo"
	        return 1
	}
	
	function f3() {
	        echo "foo"
	        return 0
	}
	
	
	echo "\${f1} value is: ${f1}, \${f2} value is ${f2}"
	# These both print nothing because ${} is used to get variable values, not to evaluate expressions
	
	echo "\$(f1) value is: $(f1)"
	# This prints nothing because return does not return a value, it sets the exit code.
	
	echo "\$(f2) value is: $(f2)"
	# prints foo because that was the output of running that expression.  f3 would have the same output.
	
	# All three of these claim to be successful, because they are wrapped in the 'test'.  
	if [ f1 ]; then echo "test f1 has an exit code of 0 (true/success)"; else echo "test f1 has an exit code of non-zero (false/failure)"; fi
	if [ f2 ]; then echo "test f2 has an exit code of 0 (true/success)"; else echo "test f2 has an exit code of non-zero (false/failure)"; fi
	if [ f3 ]; then echo "test f3 has an exit code of 0 (true/success)"; else echo "test f3 has an exit code of non-zero (false/failure)"; fi
	
	if f1; then echo "f1 has an exit code of 0 (true/success)"; else echo "f1 has an exit code of non-zero (false/failure)"; fi
	if f2; then echo "f2 has an exit code of 0 (true/success)"; else echo "f2 has an exit code of non-zero (false/failure)"; fi
	if f3; then echo "f3 has an exit code of 0 (true/success)"; else echo "f3 has an exit code of non-zero (false/failure)"; fi

   -------------------------  The output:  -----------------------------
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ ./example.sh 
	${f1} value is: , ${f2} value is 
	./example.sh: line 4: return: foo: numeric argument required
	$(f1) value is: 
	$(f2) value is: foo
	test f1 has an exit code of 0 (true/success)
	test f2 has an exit code of 0 (true/success)
	test f3 has an exit code of 0 (true/success)
	./example.sh: line 4: return: foo: numeric argument required
	f1 has an exit code of non-zero (false/failure)
	foo
	f2 has an exit code of non-zero (false/failure)
	foo
	f3 has an exit code of 0 (true/success)

   ------------ Lesson learned: -------------------
	It's probably a mistake to use [ ] around any function call, it will always return true.  Test is meant to evaluate special expressions, not function calls.

You can set defaults for parameters with a hyphen and then a default, ${parameter-default}, or you can insert a colon, ${parameter:-default}, like so:

	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $FRUIT
	
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo ${FRUIT-apple}
	apple

() starts a sub-shell, you can do multiple commands in it.  It is a sub-shell, but somehow it still observes variables you set outside of it.  Anything you set inside it is not seen outside of it.  It's like a closure.

	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $FRUIT
	apple
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ (echo $FRUIT)
	apple
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ (FRUIT=banana;echo $FRUIT)
	banana
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ (FRUIT=banana); echo $FRUIT
	apple

(( )) expands and evaluates an arithmetic expression. If the expression evaluates as zero, it returns an exit status of 1, or "false". A non-zero expression returns an exit status of 0, or "true".
	Notice that '=' is assignment inside these, so use '==' to compare values.

	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if (( 0 == 1-1 )); then echo "true"; else echo "false"; fi
	true
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [[ 0 == 1-1 ]]; then echo "true"; else echo "false"; fi
	false
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if (( 0 )); then echo "true"; else echo "false"; fi
	false
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if (( 1 )); then echo "true"; else echo "false"; fi
	true
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [[ 0 ]]; then echo "true"; else echo "false"; fi
	true
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [[ 1 ]]; then echo "true"; else echo "false"; fi
	true

	That can also be used to evaluate simple arithmetic operations, but it won't work correctly with floating point operations.


[ ] specifically checks the exit code of what it runs.  An exit code of 0 is treated as success/true, otherwise it is false. [ is a bash built-in, and is a synonym for test, not /usr/bin/test.
	the &&, ||, <, and > operators work within a [[ ]] test, despite giving an error within a [ ] construct.

[[  ]]	the &&, ||, <, and > operators work within a [[ ]] test, despite giving an error within a [ ] construct.
	Following an if statement, you don't need to have [ or [[, it's implied.

What is the difference between [ ], [[ ]], ( ), and (( )) ?
   

How do you get the last column using gawk?
   gawk -e '{print $NF}'  --> You must use single quotes here so the shell does not attempt to replace $NF.

How do you make variables you set ALWAYS carry through to sub-shells? and undo that?
   set -o allexport   
   set +o allexport

How do you make the '>' operator NEVER overwrite existing files? and undo that?
   set -o noclobber
   set +o noclobber

What keyboard shortcut clears the screen?
   ctrl+l

What keyboard shortcut exits a sub-shell?
   ctrl+d
   you can use "set -O ignoreeof" to disable that.

How do you copy lines in vi?
   in visual mode, use 'y' to copy.

How do you insert a prefix before many lines in vi?
   ctrl+v
   move to cover all necessary lines
   shift+i
   type what you want to be the prefix
   hit escape, wait two seconds.  done.

How do you remove a prefix before many lines in vi?
   ctrl+v
   cover the prefix you want to remove
   type 'd'.  done

In a bash script, how do you get input from a user?  'read'
How you you print a line without a new line? echo -n "my prompt: "

How do you duplicate a line in vi?
   Go to the line,
   type 'yy' or 'Y'
   Go to where you want the line cloned to
   type 'p'
 ~ or ~
   type 'YP' on the line you want duplicated.

In vi, how do you auto-indent? see line numbers?  and disable these features?
   :set autoindent
   :set noautoindent
   :set number
   :set nonumber

In vi, how do you set a bookmark and navigate between them?
   go to the line
   type 'm' then another lower-case letter from a to z
   Go about other editing
   type a back-tick `, and then the letter to go back to it.

==============================
Other tips:

from command line, do "bash -x <file>" to make it run in debug mode.
from within script, make the first line "#!/bin/bash -x" to have it always run in debug mode.
use "set -x" to put a script in debug mode, "set +x" to exit debug mode.


=====================================
Unexplained Bash behaviour:

if "false" returns an exit code of 1, why is it a 0 when run as part of a test?
	
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ "true"
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $?
	0
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ "false"
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ echo $?
	1
	michael@michael-GA-890GPA-UD3H ~/information_technology_notes $ if [ "false" ]; then echo "true"; else echo "false"; fi
	true

