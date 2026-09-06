---
title: "Troubleshooting Java interface issues"
source: "fgl-topics/c_fgl_JavaBridge_038.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Troubleshooting Java interface issues"
type: "concept"
---

# Troubleshooting Java interface issues

> This section describes common issues related to the Java interface.

## FGL Error -6622: Missing Java classes

When compiling a .4gl source code using Java classes, the
fglcomp compiler will report the error [-6622](../15_library-reference/4483-genero-bdl-errors.md), if the imported Java classes references classes that cannot be found at compile
time.

For example, with the following code:

```
$ cat main.4gl
IMPORT JAVA org.apache.commons.vfs2.Selectors
MAIN
END MAIN
```

The fglcomp compiler may report the following error, if one
of the referenced Java classes is not installed on your system:

```
$ fglcomp -M main.4gl
main.4gl:2:13:2:45:error:(-6222) org.apache.commons.vfs2.Selectors: class not found.
```

This kind of error will not occur when compiling pure Java code using the same class,
because the javac compiler and the fglcomp compiler behave
differently.

When using `IMPORT JAVA java-class`, fglcomp
calls the Java class-loader by using the JNI method
`FindClass(java-class)`, which is the equivalent of
`java.lang.Class.forName(java-class)`. This can fail, if other
classes referenced by the imported Java classes are missing.

To find what referenced classes are missing in your environment, try to load the classes imported
in your FGL code, by using a pure a Java sample:

```
$ cat Test.java
public class Test {
  public static void main(String[] args) throws Exception {
    java.lang.Class.forName("org.apache.commons.vfs2.Selectors");
  }
}

$  javac Test.java && java Test
Exception in thread "main" java.lang.NoClassDefFoundError: org/apache/commons/lang3/Range
        at org.apache.commons.vfs2.FileDepthSelector.<init>(FileDepthSelector.java:35)
        at org.apache.commons.vfs2.FileDepthSelector.<init>(FileDepthSelector.java:54)
        at org.apache.commons.vfs2.Selectors.<clinit>(Selectors.java:27)
        at java.base/java.lang.Class.forName0(Native Method)
        at java.base/java.lang.Class.forName(Class.java:315)
        at Test.main(Test.java:3)
Caused by: java.lang.ClassNotFoundException: org.apache.commons.lang3.Range
        at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:583)
        at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)
        at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:521)
        ...
```

Here we can see that the Java class `org.apache.commons.lang3.Range` is
missing.

The problem can be solved by installing the `org.apache.commons.lang3`
package.
