---
title: "Java objects as functions parameters and returns"
source: "fgl-topics/c_fgl_JavaBridge_019.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Java objects as functions parameters and returns"
type: "concept"
description: "Java objects must be instantiated and referenced by a program variable. The object reference is stored in the variable and can be passed as a parameter, or returned from a function. The Java objects ..."
---

# Java objects as functions parameters and returns

Java objects must be instantiated and referenced by a program variable. The object reference is
stored in the variable and can be passed as a parameter, or returned from a function.

The Java objects are passed by reference to functions. This means that the called function does
not get a clone of the object, but rather a handle to the original object. The function can then
manipulate and modify the original object provided by the
caller:

```
IMPORT JAVA java.lang.StringBuffer

MAIN
  DEFINE x java.lang.StringBuffer
  LET x = StringBuffer.create()
  CALL change(x)
  DISPLAY x.toString()
END MAIN

FUNCTION change(sb java.lang.StringBuffer)
  CALL sb.append("abc")
END FUNCTION
```

Similarly, Java object references can be returned from functions. Use the
`RETURNS` clause in the function definition, to specify the type of Java object
returned by the function:

```
IMPORT JAVA java.lang.StringBuffer

MAIN
  DEFINE x java.lang.StringBuffer
  LET x = build()
  DISPLAY x.toString()
END MAIN

FUNCTION build() RETURNS java.lang.StringBuffer
  DEFINE sb java.lang.StringBuffer
  LET sb = StringBuffer.create()   -- Creates a new object.
  CALL sb.append("abc")
  RETURN sb  -- Returns the reference to the object, not a copy/clone.
END FUNCTION
```
