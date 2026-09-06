---
title: "Using Java arrays"
source: "fgl-topics/c_fgl_JavaBridge_034.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Using Java arrays"
type: "concept"
description: "Java arrays and Genero arrays are different. In order to interface with Java arrays, the Genero language has been extended with a new kind of array, called \"Java arrays\" . Java arrays have to be ..."
---

# Using Java arrays

Java arrays and [Genero arrays](../08_language-basics/0730-understanding-arrays.md "This is an introduction to arrays.") are
different. In order to interface with Java arrays, the Genero language has been extended with a new
kind of array, called ["Java arrays"](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.").

Java arrays have to be created with a given length. Like native Java
arrays, the length cannot be changed after the array is created.

In BDL Java arrays, the first element of is at index position 1, while in Java code, the first
element of an array starts at index 0 (zero).

To create a Java array in Genero, you must define a [`TYPE`](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") in order to call the [`create()`](../15_library-reference/2964-java-array-type-methods.md) type method of Java arrays. The type of the elements in a Java
array must be one of the language types that have a corresponding primitive type in Java (such as
INTEGER (int), FLOAT (double)), or it must be a Java class such as
`java.lang.String`.

The Java arrays are passed to Java methods by reference, so the elements
of the array can be manipulated in Java. Furthermore, Java arrays can be
created in Java code and returned to the Genero program.

This example shows how to create a Java array in Genero, to instantiate
a Java Array of [`INTEGER`](../08_language-basics/0562-integer.md "The INTEGER data type is used for storing large whole numbers.") elements:

```
MAIN
  TYPE int_array_type ARRAY[] OF INTEGER
  DEFINE ja int_array_type
  LET ja = int_array_type.create(100)
  LET ja[10] = 123
  DISPLAY ja[10], ja[20]
  DISPLAY ja.getLength()
END MAIN
```

This example shows a program creating a Java array of Java
strings:

```
IMPORT JAVA java.lang.String
MAIN
    TYPE string_array_type ARRAY[] OF java.lang.String
    DEFINE names string_array_type
    LET names = string_array_type.create(100)
    LET names[1] = "aaaaaaa"
    DISPLAY names[1]
END MAIN
```

To create a Java array of
structured [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") elements,
use the [`com.fourjs.fgl.lang.FglRecord`](2684-using-genero-records.md)
class:

```
IMPORT JAVA com.fourjs.fgl.lang.FglRecord
MAIN
  TYPE record_array_type ARRAY[]
       OF com.fourjs.fgl.lang.FglRecord
  DEFINE ra record_array_type
  TYPE r_t RECORD
       id INTEGER,
       name VARCHAR(100)
      END RECORD
  DEFINE r r_t
  LET ra = record_array_type.create(100)
  LET r.id = 123
  LET r.name = "McFly"
  LET ra[10] = r
  INITIALIZE r TO NULL
  LET r = CAST (ra[10] AS r_t)
  DISPLAY r.*
END MAIN
```

Java arrays of Java classes can be defined. This example introspects the
`java.lang.String` class by using Java array of
`java.lang.reflect.Method` to query the list of methods from the
`java.lang.String`
class:

```
IMPORT JAVA java.lang.Class
IMPORT JAVA java.lang.reflect.Method
MAIN
  DEFINE c java.lang.Class
  DEFINE ma ARRAY[] OF java.lang.reflect.Method
  DEFINE i INTEGER
  LET c = Class.forName("java.lang.String")
  LET ma = c.getMethods()
  FOR i = 1 TO ma.getLength()
     DISPLAY ma[i].toString()
  END FOR
END MAIN
```

Java arrays can be created in the Java code, to be returned from a method and
assigned to a program variable:

```
public static int [] createIntegerArray(int size) {
    return new int[size];
}
```
