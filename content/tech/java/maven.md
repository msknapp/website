---
title: "Maven"
draft: false
description: my description.
summary: a useful page.
lastmod: 2023-03-14
date: 2023-03-14
tags: []
categories: []
series: []
keywords: []
nextlink: /tech/bash/
---

A basic maven pom template:

<project xmlns="http://maven.apache.org/POM/4.0.0"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0
                      http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>
 
  <!-- The Basics -->
  <groupId>...</groupId>
  <artifactId>...</artifactId>
  <version>...</version>
  <packaging>...</packaging>
  <dependencies>...</dependencies>
  <parent>...</parent>
  <dependencyManagement>...</dependencyManagement>
  <modules>...</modules>
  <properties>...</properties>
 
  <!-- Build Settings -->
  <build>...</build>
  <reporting>...</reporting>
 
  <!-- More Project Information -->
  <name>...</name>
  <description>...</description>
  <url>...</url>
 
  <scm>...</scm>
  <repositories>...</repositories>
  <profiles>...</profiles>
</project>



Making the jar shaded:


	<plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-shade-plugin</artifactId>
        <version>2.4.3</version>
        <executions>
          <execution>
            <phase>package</phase>
            <goals>
              <goal>shade</goal>
            </goals>
            <configuration>

              <artifactSet>
			<!-- There is also an excludes -->
                    <includes>
			<!-- You can use an asterisk in place of the artifact id. -->
			<include>groupid:artifactId[[:type]:classifier]</include>
		    </includes>
                </artifactSet>
              <transformers>
		<!-- for an executable jar -->
                <transformer implementation="org.apache.maven.plugins.shade.resource.ManifestResourceTransformer">
                  <mainClass>com.foo.bar.Main</mainClass>
                </transformer>
              </transformers>
		<!-- To resolve the conflict where signed jars prevent this from running. -->
		<filters>
	                <filter>
	                  <artifact>*:*</artifact>
	                  <excludes>
	                    <exclude>META-INF/*.SF</exclude>
	                    <exclude>META-INF/*.DSA</exclude>
	                    <exclude>META-INF/*.RSA</exclude>
	                  </excludes>
	                </filter>
		</filters>
            </configuration>
          </execution>
        </executions>
      </plugin>

Copying dependencies to a lib:

	The 'copy' goal is used when you want to just copy a few specific artifacts, it does not just copy all dependencies.  You probably want to use 'copy-dependencies' instead.


      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-dependency-plugin</artifactId>
        <version>2.10</version>
        <executions>
          <execution>
            <id>copy-dependencies</id>
            <phase>package</phase>
            <goals>
              <goal>copy-dependencies</goal>
            </goals>
            <configuration>
              <outputDirectory>${project.build.directory}/lib</outputDirectory>
              <overWriteReleases>false</overWriteReleases>
              <overWriteSnapshots>true</overWriteSnapshots>
              <overWriteIfNewer>true</overWriteIfNewer>
		<classifier>Specify classifier to look for</classifier>
		<excludeArtifactIds></excludeArtifactIds>
		<excludeClassifiers></excludeClassifiers>
		<excludeGroupIds></excludeGroupIds>
		<excludeScope></excludeScope>
		<excludeTypes></excludeTypes>
		<includeArtifactIds></includeArtifactIds>
		<includeClassifiers></includeClassifiers>
		<includeGroupIds></includeGroupIds>
		<includeScope></includeScope>
		<includeTypes></includeTypes>
		<!-- The plural properties are comma delimited. -->
            </configuration>
          </execution>
        </executions>
      </plugin>

Making a test jar and using it:

	In the first jar/project:
	<plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-jar-plugin</artifactId>
        <version>2.6</version>
        <executions>
          <execution>
            <goals>
              <goal>test-jar</goal>
            </goals>
          </execution>
        </executions>
      </plugin>


	In the project that depends on it:
	<dependency>
      <groupId>groupId</groupId>
      <artifactId>artifactId</artifactId>
      <type>test-jar</type>
      <version>version</version>
      <scope>test</scope>
    </dependency>

Using the jar plugin:
	<plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-jar-plugin</artifactId>
        <version>2.6</version>
        <configuration>
		<!-- Including/excluding classes in a jar (optional) -->
          <includes>
		<!-- This seems to take a path, not a package. -->
            <include>**/service/*</include>
          </includes>
		<!-- Making it executable (assuming that you are not using the shade or assembly plugins.) -->
          <archive>
            <manifest>
              <addClasspath>true</addClasspath>
              <mainClass>fully.qualified.MainClass</mainClass>
            </manifest>
          </archive>
        </configuration>
      </plugin>

Using the surefire plugin:

	<plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-surefire-plugin</artifactId>
        <version>2.19.1</version>
        <configuration>
	  <!-- Including some classes in tests (optional) -->
          <includes>
            <include>Sample.java</include>
          </includes>

	  <!-- Excluding some classes in tests (optional).  These appear to take file paths as glob patterns, not package names. -->
          <excludes>
            <exclude>**/TestCircle.java</exclude>
            <exclude>**/TestSquare.java</exclude>
          </excludes>

	  <!-- Running with some system properties (optional) -->
          <systemPropertyVariables>
            <propertyName>propertyValue</propertyName>
            <buildDirectory>${project.build.directory}</buildDirectory>
            [...]
          </systemPropertyVariables>
        </configuration>
      </plugin>

Configuring source/target java version:

	<plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-compiler-plugin</artifactId>
        <version>3.5</version>
        <configuration>
          <source>1.8</source>
          <target>1.8</target>
        </configuration>
      </plugin>


Using the scala maven plugin:


            <plugin>
                <groupId>net.alchim31.maven</groupId>
                <artifactId>scala-maven-plugin</artifactId>
                <version>3.2.1</version>
                <executions>
                    <execution>
                        <phase>compile</phase>
                        <goals>
                            <goal>compile</goal>
                            <goal>testCompile</goal>
                        </goals>
                    </execution>
                </executions>
                <configuration>
                    <args>
                        <arg>-Yresolve-term-conflict:object</arg>
                    </args>
                    <!--<addScalacArgs>-Yresolve-term-conflict=strategy</addScalacArgs>-->
                </configuration>
            </plugin>


    If there is a package and object with the same name on the classpath, it will keep the object/class and the package will be inaccessable because -Yresolve-term-conflict is set to object, I think the other option is "package".
