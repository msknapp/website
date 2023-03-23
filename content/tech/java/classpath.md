---
title: "Classpath"
draft: false
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
prevlink: ..
---

Seeing the contents of a jar file:

jar tf <my.jar> | less

Extracting a jar file:

jar xf <my.jar>

I recommend copying/moving it to a tmp directory first so when you extract it, the files are not inter-mixed with others you wanted.

Finding what jars have a certain class:

$ cd ~/.m2/repository
$ find . -name '*.jar' -type f -exec jar tf {} | grep <name> \;

Files will have forward slashes in them like a directory, not like package names.  They will end with ".class" if they are class files.

Creating an index of classes and jars that contain them:

	function indexjars() {
	        for x in $(find ~/.m2/repository -name '*.jar' -type f); do
	                jar tf $x | sed -r 's|.*/(.*).class|\1 \0 '$x'|g' >> /tmp/classes_unordered.txt
	        done
	        sort < /tmp/classes_unordered.txt > /tmp/maven_index.txt
	}
	
Now you can use 'less' to search the index file for your class.  This can be used to answer the question "what jar has the class I'm seeing?"

Seeing dependencies with maven:

mvn dependency:tree | less
mvn dependency:resolve | less

You can add profiles using -P.  You can also set the scopes with --include-scope=compile|test|provided, etc.  Tree is preferred since it will tell you what dependency gave you a transitive one.  You can add an exclude statement to the dependency in the pom if necessary.

Viewing Source Files:
Unfortunately 'jar' and 'jad' are not very versatile.  "jar xf <my.jar> </path/to/My.class>" can be used to extract a single class, but it cannot go to standard output, it must be written to a file in the corresponding sub directory.  jad cannot read from standard input, only from a file.  jad can be piped to standard output though.  You will have to go to each file one at a time, extract what you want to see, decompile it, and then review its source code.
