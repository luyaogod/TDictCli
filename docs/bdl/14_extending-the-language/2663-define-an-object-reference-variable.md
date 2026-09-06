---
title: "Define an object reference variable"
source: "fgl-topics/c_fgl_JavaBridge_011.html"
breadcrumb: "Extending the language > The Java interface > Getting started with the Java interface > Define an object reference variable"
type: "concept"
description: "Before creating a Java object in your program, you must declare a program variable to reference the object. The type of the variable must be the name of the Java class, and can be fully qualified if ..."
---

# Define an object reference variable

Before creating a Java object in your program, you must declare a [program variable](../08_language-basics/0686-variables.md "Explains how to define program variables.") to reference the
object. The type of the variable must be the name of the Java class, and
can be fully qualified if needed:

```
IMPORT JAVA java.util.regex.Pattern
MAIN
  DEFINE p1 Pattern
  DEFINE p2 java.util.regex.Pattern
END MAIN
```

The variables declared with a class are only the handles to reference an object (meaning the
object is not yet [created](2664-instantiate-a-java-class.md)).
