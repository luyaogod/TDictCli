---
title: "util.JSONObject.name"
source: "fgl-topics/c_fgl_ext_util_JSONObject_name.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.name"
type: "concept"
---

# util.JSONObject.name

> Returns the name of a JSON object entry by position.

## Syntax

```
name(
   index INTEGER )
  RETURNS STRING
```

1. index is the index of the name-value pair in the JSON object.

## Usage

The `name()` method returns the entry name in the JSON object at the
given position.

The index corresponding to the first name-value pair is 1.

If no entry exists at the given index, the method returns `NULL`.

This method can be used along with the `getLength()` and
`getType()` methods to read the entries of a JSON object.

## Example

```
IMPORT util
MAIN
    DEFINE obj util.JSONObject
    DEFINE i INTEGER
    LET obj = util.JSONObject.parse('{"id":123,"name":"Scott"}')
    FOR i=1 TO obj.getLength()
        DISPLAY i, ": ", obj.name(i)
    END FOR
END MAIN
```

Output:

```
          1: id
          2: name
```

## Related links

**Related concepts**  

[util.JSONObject.getLength](3596-util-jsonobject-getlength.md "Returns the number of name-value pairs in the JSON object.")

[util.JSONObject.getType](3597-util-jsonobject-gettype.md "Returns the type of a JSON object element.")
