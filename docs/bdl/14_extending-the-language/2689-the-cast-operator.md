---
title: "The CAST operator"
source: "fgl-topics/c_fgl_JavaBridge_035.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > The CAST operator"
type: "concept"
description: "Important consideration has to be taken when assigning object references to different target types or classes. A Widening Reference Conversion occurs when an object reference is converted to a ..."
---

# The CAST operator

Important consideration has to be taken when assigning object references to different target
types or classes.

A Widening Reference Conversion occurs when an object reference is converted to a
superclass that can accommodate any possible reference of the original type or class.

A Narrowing Reference Conversion occurs when an object reference of a superclass is
converted to a subtype or subclass of the original object reference.

For example, in a vehicle class hierarchy with `Vehicle` and `Car`
classes, `Car` is a subclass that inherits from the `Vehicle`
superclass. When assigning a `Car` object reference to a `Vehicle`
variable, Widening Reference Conversion takes place. When assigning a `Vehicle`
object reference to a `Car` variable, Narrowing Reference Conversion occurs.

While widening conversion does not require casts and will not produce
compilation or runtime errors, narrowing conversion needs the [`CAST` operator](../08_language-basics/0648-cast-function.md "The CAST operator converts a Java object to the user-defined type or Java class specified.") to
convert to the target type or class:

```
CAST( object_reference AS type_or_class )
```

The following example creates a `java.lang.StringBuffer` object, and assigns the
reference to a `java.lang.Object` variable (implying Widening Reference Conversion);
then the object reference is assigned back to the `java.lang.StringBuffer` variable
(implying Narrowing Reference Conversion and `CAST` operator
usage):

```
IMPORT JAVA java.lang.Object
IMPORT JAVA java.lang.StringBuffer
MAIN
  DEFINE o java.lang.Object
  DEFINE sb java.lang.StringBuffer
  LET sb = StringBuffer.create()
  -- Widening Reference Conversion
  LET o = sb
  -- Narrowing Reference Conversion needs CAST()
  LET sb = CAST( o AS StringBuffer )
END MAIN
```
