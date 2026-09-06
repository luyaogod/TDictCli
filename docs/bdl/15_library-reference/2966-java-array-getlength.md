---
title: "java-array.getLength"
source: "fgl-topics/c_fgl_Arrays_JavaArray_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > Java Array type as class > Java Array type methods > java-array.getLength"
type: "concept"
---

# java-array.getLength

> Returns the length of the Java array.

## Syntax

```
getLength( )
   RETURNS INTEGER
```

## Usage

This method returns the number of elements in the Java array.

## Example

```
IMPORT JAVA java.lang.String
MAIN
    TYPE string_array_type ARRAY[] OF java.lang.String
    DEFINE names string_array_type
    LET names = string_array_type.create(100)
    LET names[1] = "aaaaaaa"
    DISPLAY names.getLength()  -- Shows 100
END MAIN
```
