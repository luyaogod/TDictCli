---
title: "Java software requirements for FGL"
source: "fgl-topics/c_fgl_JavaBridge_006.html"
breadcrumb: "Extending the language > The Java interface > Prerequisites and installation > Java software requirements for FGL"
type: "concept"
description: "In order to use the Java Interface in your application programs, you need the Java software installed and properly configured. Install a Java Development Kit (JDK) on development sites (if you need to ..."
---

# Java software requirements for FGL

In order to use the Java Interface in your application programs, you need the Java software
installed and properly configured.

- Install a Java Development Kit (JDK) on development sites (if you need to compile your own Java
  classes)
- Install a Java Runtime Environment (JRE) on production sites (on the server where your programs
  are running)

Any class and jar file shipped with the FGL package (\*.class or
\*.jar) is compiled with the command `javac -source
release -target release` where
release is the oldest release supported by javac version used
on the build system. For example, java 17 supports compatibility options down to version 7. This
allows the use of the FGL to java wrapper classes (`IMPORT JAVA
com.fourjs.fgl.lang.class-name`) on installation machines using older
java versions than the java version used on the porting machine.

As a general rule, always install the current Long-Term-Support (LTS) Java version available on
your platform, with Java Native Interface (JNI) support.

For a detailed list of supported JVMs, refer to the Supported platforms and
databases document, available on the Products download page of the [Four Js Web site](https://4js.com/download/products/).

The version of the installed Java software can be shown with the
command:

```
java --version
```

In order to execute Java byte code, the Genero runtime system uses the JNI interface. The JVM is
loaded as a shared library and its binary format must match the binary format of the Genero runtime
system. For example, a 64-bit Genero package requires a 64-bit JVM.

When implementing Java classes for Genero Mobile for Android™ (GMA), check the JDK version required by the Android SDK. For more information, see the [Android Studio web site](http://developer.android.com/sdk/installing/studio.html).
