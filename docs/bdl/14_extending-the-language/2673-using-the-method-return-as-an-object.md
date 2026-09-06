---
title: "Using the method return as an object"
source: "fgl-topics/c_fgl_JavaBridge_020.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Using the method return as an object"
type: "concept"
description: "If a Java method returns an object, you can use the method call directly as an object reference to call another method: IMPORT JAVA java.util.regex.Pattern MAIN DEFINE p Pattern LET p = ..."
---

# Using the method return as an object

If a Java method returns
an object, you can use the method call directly as an object
reference to call another method:

```
IMPORT JAVA java.util.regex.Pattern
MAIN
  DEFINE p Pattern
  LET p = Pattern.compile("a*b")
  IF p.matcher("aaaab").matches() THEN
    DISPLAY "It matches..."
  END IF
END MAIN
```

In this code example, the `matcher()` method of
object `p` is invoked and returns an object of
type `java.util.regex.Matcher`. The object reference
returned by the `matcher()` method can be directly
used to invoke the `matches()` method of the `Matcher` class.
