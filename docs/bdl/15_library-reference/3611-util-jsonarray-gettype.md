---
title: "util.JSONArray.getType"
source: "fgl-topics/c_fgl_ext_util_JSONArray_getType.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.getType"
type: "concept"
---

# util.JSONArray.getType

> Returns the type of a JSON array element.

## Syntax

```
getType(
   index INTEGER )
  RETURNS STRING
```

1. index is the ordinal position of the element.

## Usage

The `getType()` method returns the data type name corresponding
to the JSON array element at the given position.

The index corresponding to the first element is 1.

This method can be used along with the `getLength()` method
to read the entries of a JSON array object.

Possible values returned by this method are:

- `NUMBER`: A numeric value.
- `STRING`: A string value delimited by double quotes.
- `BOOLEAN`: A boolean value (true/false)
- `NULL`: A non-existing element.
- `OBJECT`: A structured object.
- `ARRAY`: An ordered list of elements.

## Example

```
IMPORT util
MAIN
    DEFINE arr util.JSONArray
    LET arr = util.JSONArray.parse('[123,"abc",null]')
    DISPLAY arr.getType(1) -- NUMBER
    DISPLAY arr.getType(2) -- STRING
    DISPLAY arr.getType(3) -- NULL
END MAIN
```

Output:

```
NUMBER
STRING
NULL
```

## Related links

**Related concepts**  

[util.JSONArray.getLength](3610-util-jsonarray-getlength.md "Returns the number of elements in the JSON array object.")
