---
title: "Method overloading in Java"
source: "fgl-topics/c_fgl_JavaBridge_018.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Method overloading in Java"
type: "concept"
description: "The Java language allows method overloading ; the parameter count and the parameter data types of a method are part of the method identification. Thus, the same method name can be used to implement ..."
---

# Method overloading in Java

The Java language allows method
overloading; the parameter count and the parameter data
types of a method are part of the method identification. Thus,
the same method name can be used to implement different versions of
the Java method, taking
different parameters:

```
DEFINE int2 SMALLINT, int4 INTEGER, flt FLOAT

-- Next call invokes method display(short) of the Java class
CALL myobj.display(int2)

-- Next call invokes method display(int) of the Java class
CALL myobj.display(int4)

-- Next call invokes method display(double) of the Java class
CALL myobj.display(flt)

-- Next call invokes method display(short,int) of the Java class
CALL myobj.display(int2,int4)
```
