---
title: "Platform-specific notes for the JVM"
source: "fgl-topics/c_fgl_JavaBridge_008.html"
breadcrumb: "Extending the language > The Java interface > Prerequisites and installation > Platform-specific notes for the JVM"
type: "concept"
description: "Unix-based platforms The JAVA_HOME environment variable must be set to use the Java interface on Unix-based platforms platforms. If needed, the LD_LIBRARY_PATH environment variable can contain the ..."
---

# Platform-specific notes for the JVM

## Unix-based platforms

The JAVA\_HOME environment variable must be set to use the Java interface on Unix-based platforms
platforms.

If needed, the LD\_LIBRARY\_PATH environment variable can contain the path to the JVM library. This
is however not required if JAVA\_HOME is properly set.

## Microsoft™ Windows®

The JAVA\_HOME environment variable must be set to use the Java interface on Microsoft Windows
platforms.

If needed, the PATH environment variable can contain the path to the JVM library. This is however
not required if JAVA\_HOME is properly set.

> **Important:**
>
> If PATH is defined to find the `javac` compiler or JVM
> library, make sure that it does not contain double quotes around the path to the
> JVM.DLL dynamic library, otherwise the DLL loader will fail to load the
> JVM.

See also [Microsoft Windows platform notes](../04_installation/0043-microsoft-windows-platform-notes.md).

## Apple® macOS®

The JAVA\_HOME environment variable must be set to use the Java interface on macOS.

To find the JAVA\_HOME path on macOS, use the /usr/libexec/java\_home
tool:

```
export JAVA_HOME=`/usr/libexec/java_home`
```

> **Important:**
>
> On macOS™, the usage of DYLD\_LIBRARY\_PATH
> is strongly discouraged, especially since Mac OS X 10.11 this environment variable is no longer
> exported in sub processes.

The Genero runtime system uses the Java Native Interface (JNI) to interact with the JVM and
execute Java code. Make sure that JNI is available, refer to JDK documentation for more details.

Read also [Apple macOS platform notes](../04_installation/0042-apple-macos-platform-notes.md).

## Android™

The Java Interface can be used in Genero apps built for the Android / GMA platform.

JDK 17 is required to build Android apps. For the latest information regarding system requirements and
Java support, please refer to the Supported platforms and databases document, available
on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).

On Android
devices, some system functions can only be accessed in the context of a JVM. Use the [Java Interface](2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") with the
`com.fourjs.gma.vm.FglRun` class to access such system specifics. Custom Java classes
need to be part of the .apk package and can be used without any further
configuration.

For more details, see [Executing Java code with GMA](2693-executing-java-code-with-gma.md).

## iOS

> **Important:**
>
> The Java interface cannot be used in apps running on iOS devices: There is
> no standard free JVM available.

## IBM® AIX®

Consider the following notes when using the Java Interface with the IBM Java VM on AIX:

If you get `java.lang.UnsatisfiedLinkError` exceptions, set the path to native
shared libraries in the LIBPATH environment variable:

```
$ LIBPATH=$JAVA_HOME/jre/bin:$JAVA_HOME/jre/bin/j9vm:$JAVA_HOME/jre/lib/ppc64:$LIBPATH
$ export LIBPATH
```

This is required when using Java code that needs to access native code supplied as part of the
JRE. For example, without setting LIBPATH to the appropriate path, the JVM cannot find the shared
library libnet.so.

Using the `-Djava.library.path=path-to-native-library` java VM
option does not seem to help.

See also [IBM AIX platform notes](../04_installation/0041-ibm-aix-platform-notes.md).
