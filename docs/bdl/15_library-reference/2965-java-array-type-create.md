---
title: "java-array-type.create"
source: "fgl-topics/c_fgl_Arrays_JavaArray_create.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > Java Array type as class > Java Array type methods > java-array-type.create"
type: "concept"
---

# java-array-type.create

> Creates a new Java array of the given type.

## Syntax

```
java-array-type.create( size INTEGER )
   RETURNS java-array-type
```

1. size defines the actual number of elements of the array.
2. java-array-type is a user type defined with the `TYPE`
   instruction.

## Usage

This class method creates a new instance of the Java array specified by the type used, with the
size provided as parameter.

The type must be declared as a user defined type define with the `ARRAY [ ]
OF` notation reserved for Java arrays.

## Example

```
IMPORT JAVA java.lang.String
MAIN
    TYPE string_array_type ARRAY[] OF java.lang.String
    DEFINE names string_array_type
    LET names = string_array_type.create(100)
    LET names[1] = "aaaaaaa"
    DISPLAY names[1]
END MAIN
```

## Related links

**Related concepts**  

[Types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
