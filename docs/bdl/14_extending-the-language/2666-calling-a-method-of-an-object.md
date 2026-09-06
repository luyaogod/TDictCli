---
title: "Calling a method of an object"
source: "fgl-topics/c_fgl_JavaBridge_014.html"
breadcrumb: "Extending the language > The Java interface > Getting started with the Java interface > Calling a method of an object"
type: "concept"
description: "Once the class has been instantiated as an object, and the object reference has been assigned to a variable, you can call a method of the Java object by using the variable as the prefix: IMPORT JAVA ..."
---

# Calling a method of an object

Once the class has been instantiated as an object, and the object
reference has been assigned to a variable, you can call a
method of the Java object
by using the variable as the prefix:

```
IMPORT JAVA java.util.regex.Pattern
IMPORT JAVA java.util.regex.Matcher
MAIN
    DEFINE p Pattern
    DEFINE m Matcher
    LET p = java.util.regex.Pattern.compile("[a-z]+")
    LET m = p.matcher("abcdef")
    DISPLAY m.matches()
END MAIN
```

In this example, the last line of the `MAIN` module
calls an object method that returns a boolean value that
is converted to an `INTEGER` and displayed.
