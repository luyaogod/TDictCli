---
title: "Calling a method of a class"
source: "fgl-topics/c_fgl_JavaBridge_013.html"
breadcrumb: "Extending the language > The Java interface > Getting started with the Java interface > Calling a method of a class"
type: "concept"
description: "Class methods (static method in Java) can be called without instantiating an object of the class. Static method invocation must be prefixed with the class name. In the following example, the compile() ..."
---

# Calling a method of a class

Class methods (static method in Java) can be called without instantiating an object of the class.
Static method invocation must be prefixed with the class name. In the following example, the
`compile()` class method of Pattern class returns a new instance of a Pattern
object:

```
IMPORT JAVA java.util.regex.Pattern
MAIN
    DEFINE p Pattern
    LET p = Pattern.compile("[,\\s]+")
END MAIN
```

If you define a variable with the same name as a Java class, you must
fully qualify the class when calling static methods, as shown in this
example:

```
IMPORT JAVA java.util.regex.Pattern
IMPORT JAVA java.util.regex.Matcher
MAIN
    DEFINE Pattern Pattern
    DEFINE Matcher Matcher
    -- static method, needs full qualifier
    LET Pattern = java.util.regex.Pattern.compile("[a-z]+")
    -- regular instance method, Pattern resolves to variable
    LET Matcher = Pattern.matcher("abcdef")
END MAIN
```

Note that in the Genero language, program variables are case-insensitive (Pattern = pattern).
