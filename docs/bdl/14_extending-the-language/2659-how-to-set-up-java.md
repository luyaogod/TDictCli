---
title: "How to set up Java"
source: "fgl-topics/t_fgl_JavaBridge_setup.html"
breadcrumb: "Extending the language > The Java interface > Prerequisites and installation > How to set up Java"
type: "task"
description: "This short procedure describes how to set up a Java environment to be used with Genero. Download the latest JDK from your preferred Java provider. On production sites, you only need a Java Runtime ..."
---

# How to set up Java

This short procedure describes how to set up a Java environment to be used with Genero.

1. Download the latest JDK from your preferred Java provider.
   On production sites, you only need a Java Runtime Environment (JRE).
2. Install the package by following the vendor installation notes.
3. Configure your environment for Java Native Interface (JNI) support:

   - In a development or runtime environment, on Unix-based platforms (including macOS®) and on Microsoft™
     Windows®, define the JAVA\_HOME environment variable, to allow
     fglrun and fglcomp to find the Java Runtime library
     (libjvm.so or JVM.DLL). Defining LD\_LIBRARY\_PATH on Unix
     or PATH on Microsoft Windows to find the Java runtime library is not required.
   - In a development environment, set the PATH environment variable to find the Java tools
     (javac, java).
   - See also [Platform-specific notes for the JVM](2660-platform-specific-notes-for-the-jvm.md)
4. Set the CLASSPATH or pass the `--java-option=-Djava.class.path=<pathlist>`
   [option](2668-using-jvm-options.md) to fglrun with the
   directories of the Java packages you want to use. You must add
   $FGLDIR/lib/fgl.jar to the class path in order to compile Java code with
   language specific classes such as [`com.fourjs.fgl.lang.FglDecimal`](2679-using-the-decimal-type.md) or [`com.fourjs.fgl.lang.FglRecord`](2684-using-genero-records.md).
5. Try your JDK by compiling a small Java sample and executing it.
