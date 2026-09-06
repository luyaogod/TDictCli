---
title: "Garbage collection of unused objects"
source: "fgl-topics/c_fgl_JavaBridge_007.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Garbage collection of unused objects"
type: "concept"
description: "Java objects do not need to be explicitly destroyed; as long as an object is referenced by a variable, on the stack or in an expression, it will remain. When the last reference to an object is ..."
---

# Garbage collection of unused objects

Java objects do not need
to be explicitly destroyed; as long as an object is referenced
by a variable, on the stack or in an expression, it will remain.
When the last reference to an object is removed, the object is destroyed
automatically.

This example shows how an unique object can be referenced twice, using two variables:

```
IMPORT JAVA java.lang.StringBuffer

MAIN
    -- Declare 2 variables to reference a StringBuffer object
    DEFINE sb1, sb2 java.lang.StringBuffer
    -- Create object and assign reference to variable
    LET sb1 = StringBuffer.create()
    -- Same object is now referenced by 2 variables
    LET sb2 = sb1
    -- Object is modified through first variable
    CALL sb1.append("abc")
    -- Object is modified through second variable
    CALL sb2.append("def")
    -- Shows content of StringBuffer object
    DISPLAY sb1.toString()
    -- Same output as previous line
    DISPLAY sb2.toString()
    -- Object is only referenced by second variable
    LET sb1 = NULL
    -- sb2 removed from stack, object is no longer referenced and is destroyed.
END MAIN
```
