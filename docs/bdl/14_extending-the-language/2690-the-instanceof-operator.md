---
title: "The INSTANCEOF operator"
source: "fgl-topics/c_fgl_JavaBridge_036.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > The INSTANCEOF operator"
type: "concept"
description: "When manipulating an object reference with a variable defined with a superclass of the real class used to instantiate the object, you sometimes need to identify the real class of the object. This is ..."
---

# The INSTANCEOF operator

When manipulating an object reference with a variable defined with
a superclass of the real class used to instantiate the object,
you sometimes need to identify the real class of the object.

This is possible with the [`INSTANCEOF` operator](../08_language-basics/0649-instanceof.md "The INSTANCEOF checks the class of an object.").

This operator checks whether the left operand is an instance of
the type or class specified by the right operand:

```
object_reference INSTANCEOF type_or_class
```

This example creates a `java.lang.StringBuffer` object,
assigns the reference to a `java.lang.Object` variable,
and tests whether the class type of the object reference is
a `java.lang.StringBuffer`:

```
IMPORT JAVA java.lang.Object
IMPORT JAVA java.lang.StringBuffer
MAIN
  DEFINE o java.lang.Object
  LET o = StringBuffer.create()
  DISPLAY o INSTANCEOF StringBuffer    -- Shows 1 (TRUE)
END MAIN
```
