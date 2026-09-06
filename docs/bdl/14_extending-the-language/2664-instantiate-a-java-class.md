---
title: "Instantiate a Java class"
source: "fgl-topics/c_fgl_JavaBridge_012.html"
breadcrumb: "Extending the language > The Java interface > Getting started with the Java interface > Instantiate a Java class"
type: "concept"
description: "In Java, objects are created with the new instruction. To create a new Java object in BDL, you must use the class-name .create() method, and assign the returned value to a program variable declared ..."
---

# Instantiate a Java class

In Java, objects are created with the `new` instruction. To create a new Java
object in BDL, you must use the `class-name.create()` method, and
assign the returned value to a program variable declared with the Java class name:

```
IMPORT JAVA java.lang.StringBuffer
MAIN
    DEFINE sb StringBuffer
    LET sb = StringBuffer.create()
END MAIN
```

If the Java class constructor uses parameters, pass the parameters to the `create()`
method:

```
IMPORT JAVA java.lang.StringBuffer
MAIN
    DEFINE sb1, sb2 StringBuffer
    -- Next code line uses StringBuffer(String str) constructor
    LET sb1 = StringBuffer.create("abcdef")
    -- Next code line uses StringBuffer(int capacity) constructor
    LET sb2 = StringBuffer.create(2048)
END MAIN
```

It is not possible to use Java generic types such as `java.util.Vector<E>`,
with a type parameter (for example, in pure Java: `Vector<MyClass> v = new
Vector<MyClass>()` ). However, it is possible to instantiate such classes without a
type parameter by using the `class-name.create()` method (in pure
Java: `Vector v = new Vector()`):

```
IMPORT JAVA java.util.Vector
MAIN
    DEFINE v java.util.Vector
    LET v = java.util.Vector.create()
    ...
END MAIN
```
