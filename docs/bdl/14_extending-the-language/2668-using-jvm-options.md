---
title: "Using JVM options"
source: "fgl-topics/c_fgl_JavaBridge_016.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Using JVM options"
type: "concept"
description: "When using the Java interface, you can instruct fglrun or fglcomp to pass Java VM specific options during JNI initialization, by using the --java-option command line argument. In the example, fglrun ..."
---

# Using JVM options

When using the Java interface, you can instruct [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") or [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") to pass Java VM specific options during JNI initialization, by
using the `--java-option` command line argument.

In the example, fglrun passes `-verbose:gc` to the Java Virtual
Machine:

```
$ fglrun --java-option=-verbose:gc myprog.42m
```

If you want to pass several options to the JVM, repeat the `--java-option`
argument as in this
example:

```
$ fglrun --java-option=-verbose:gc --java-option=-esa myprog.42m
```

You may want to pass the Java class path as a command line option to fglrun
with `-Djava.class.path` option as in this
example:

```
$ fglrun --java-option=-Djava.class.path=$FGLDIR/lib/fgl.jar:$MYCLASSPATH
 myprog.42m
```

Regarding class path specification, the `java` runtime or `javac`
compiler provides the `-cp` or `-classpath` options but when loading
the JVM library from fglrun or fglcomp, only
`-Djava.class.path` option is supported by the JNI interface.

Alternatively, you can define JVM options with the JAVA\_TOOL\_OPTIONS environment
variable:

```
$ export JAVA_TOOL_OPTIONS="-Xmx5G"
$ fglrun myprog.42m
```

However, this solution can have some limitations. For more details, see the [Java documentation](https://docs.oracle.com/en/java/javase/17/troubleshoot/environment-variables-and-system-properties.html#GUID-BE6E7B7F-A4BE-45C0-9078-AA8A66754B97).
