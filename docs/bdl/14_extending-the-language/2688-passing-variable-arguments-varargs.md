---
title: "Passing variable arguments (varargs)"
source: "fgl-topics/c_fgl_JavaBridge_varargs.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Passing variable arguments (varargs)"
type: "concept"
description: "Java supports variable arguments in method definitions with the ellipsis notation, allowing callers to pass a different number of arguments depending on the need. A typical example is a message print ..."
---

# Passing variable arguments (varargs)

Java supports variable arguments in method definitions with the ellipsis notation, allowing
callers to pass a different number of arguments depending on the need. A typical example is a
message print
method:

```
import java.lang.String;

public class MyClass {
    public static void ShowStrings( String... sl ) {
        for ( String s : sl )
            System.out.println(s);
    }
}
```

In order to call such a method from the Genero program, create a [Java array](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.") of the type of the variable argument, fill the array with objects and call the
method with that array:

```
IMPORT JAVA java.lang.String
IMPORT JAVA MyClass

MAIN
    TYPE sl_t ARRAY[] OF java.lang.String
    DEFINE sl ARRAY[] OF java.lang.String
    LET sl = sl_t.create(2)
    LET sl[1] = "Value 1"
    LET sl[2] = "Value 2"
    CALL MyClass.ShowStrings(sl)
END MAIN
```

Since Java arrays have a static size, you must create the Java array with the exact
number of variable arguments to be passed to the method.

If the Java class cannot be modified, consider implementing a function to wrap calls to
the Java method, with a varying number of arguments. It can for example take a BDL dynamic
array as parameter, to simplify the callers code:

```
IMPORT JAVA java.lang.String
IMPORT JAVA MyClass

MAIN
    DEFINE a DYNAMIC ARRAY OF STRING
    LET a[1] = "Value 1"
    LET a[2] = "Value 2"
    LET a[3] = "Value 3"
    CALL my_show_strings(a)
    LET a[4] = "Value 1"
    LET a[5] = "Value 2"
    CALL my_show_strings(a)
END MAIN

FUNCTION my_show_strings(sa)
    TYPE sl_t ARRAY[] OF java.lang.String
    DEFINE sa DYNAMIC ARRAY OF STRING
    DEFINE sl ARRAY[] OF java.lang.String
    DEFINE i INTEGER
    LET sl = sl_t.create(sa.getLength())
    FOR i=1 TO sa.getLength()
        LET sl[i] = sa[i]
    END FOR
    CALL MyClass.ShowStrings(sl)
END FUNCTION
```

If the Java class can be modified, a good practice is to write overloaded methods, using
a static number of arguments:

```
public class MyClass {
    private static void _ShowStrings( String... sl ) {
        for ( String s : sl )
            System.out.println(s);
    }
    public static void ShowStrings(String s1) {
        _ShowStrings(s1);
    }
    public static void ShowStrings(String s1, String s2) {
        _ShowStrings(s1, s2);
    }
    public static void ShowStrings(String s1, String s2, String s3) {
        _ShowStrings(s1, s2, s3);
    }
}
```

## Related links

**Related concepts**  

[Using Java arrays](2687-using-java-arrays.md "Using Java arrays")
