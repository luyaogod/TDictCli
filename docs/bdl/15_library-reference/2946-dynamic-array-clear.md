---
title: "DYNAMIC ARRAY.clear"
source: "fgl-topics/c_fgl_Arrays_ARRAY_clear.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.clear"
type: "concept"
---

# DYNAMIC ARRAY.clear

> Removes all elements of the array.

## Syntax

```
clear( )
```

## Usage

This method clears the array, by removing all its elements.

For example, if the array is not empty, use the `clear()` method just before
filling the array with a new set of elements.

## Example

```
MAIN
    DEFINE a1 DYNAMIC ARRAY OF STRING
    CALL fill_array(a1)
    DISPLAY a1.getLength()
    DISPLAY a1[3]
END MAIN

FUNCTION fill_array(arr)
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE i INTEGER
    CALL arr.clear()
    FOR i=1 TO 10
       LET arr[i] = "Item #"||i
    END FOR
END FUNCTION
```
