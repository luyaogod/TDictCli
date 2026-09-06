---
title: "Static fields of Java classes"
source: "fgl-topics/c_fgl_JavaBridge_022.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Static fields of Java classes"
type: "concept"
description: "Java classes can have object and class (\"static\") fields. Java static class fields can be declared as \"final\" (read-only). It is not possible to change the object or class fields in programs, even if ..."
---

# Static fields of Java classes

Java classes can have object
and class ("static") fields. Java static
class fields can be declared as "final" (read-only). It is not
possible to change the object or class fields in programs, even if
the field is not declared as "static final"; you can however
read from it:

```
IMPORT JAVA java.lang.Integer
MAIN
  DISPLAY Integer.MAX_VALUE
END MAIN
```
