---
title: "util.JSONArray.getLength"
source: "fgl-topics/c_fgl_ext_util_JSONArray_getLength.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.getLength"
type: "concept"
---

# util.JSONArray.getLength

> Returns the number of elements in the JSON array object.

## Syntax

```
getLength()
  RETURNS INTEGER
```

## Usage

The `getLength()` method returns the number of elements in the
JSON array object.

This method can be used along with the `get()` and
`getType()` method to read elements of a JSON array object.

## Example

```
IMPORT util
MAIN
    DEFINE arr util.JSONArray
    DEFINE i INTEGER
    LET arr = util.JSONArray.parse('[123,8723,9232]')
    FOR i=1 TO arr.getLength()
        DISPLAY i, ": ", arr.get(i)
    END FOR
END MAIN
```

Output:

```
          1:                   123.0
          2:                  8723.0
          3:                  9232.0
```

## Related links

**Related concepts**  

[util.JSONArray.getType](3611-util-jsonarray-gettype.md "Returns the type of a JSON array element.")
