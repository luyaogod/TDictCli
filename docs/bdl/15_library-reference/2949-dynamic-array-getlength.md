---
title: "DYNAMIC ARRAY.getLength"
source: "fgl-topics/c_fgl_Arrays_ARRAY_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.getLength"
type: "concept"
---

# DYNAMIC ARRAY.getLength

> Returns the length of the array.

## Syntax

```
getLength( )
   RETURNS INTEGER
```

## Usage

This method returns the number of elements in the array.

> **Tip:**
>
> To improve performances in `FOR … END FOR` loops, assign a variable
> with the array length and use this variable as argument for the `TO` clause.

## Example

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING,
           i, m INTEGER
    LET arr[1] = "aaa"
    LET arr[2] = "bbbb"
    LET arr[3] = "ccccc"
    LET m = arr.getLength()
    FOR i=1 TO m
        DISPLAY arr[i]
    END FOR
END MAIN
```
