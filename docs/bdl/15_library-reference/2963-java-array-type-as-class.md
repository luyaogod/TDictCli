---
title: "Java Array type as class"
source: "fgl-topics/c_fgl_Arrays_JavaArray_as_class.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > Java Array type as class"
type: "concept"
---

# Java Array type as class

> The Java Array type provides a set of utility methods to manipulate array elements.

Java array methods can be invoked with a type reference or the array variable,
for example:

```
IMPORT JAVA java.lang.String
MAIN
    TYPE string_array_type ARRAY[] OF java.lang.String
    DEFINE names string_array_type
    LET names = string_array_type.create(100)
    LET names[1] = "aaaaaaa"
    DISPLAY names.getLength()
END MAIN
```

## Related links

**Related concepts**  

[The Java interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")

[ARRAY](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.")

## Child topics

- [Java Array type methods](2964-java-array-type-methods.md)
