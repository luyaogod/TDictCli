---
title: "util.JSONObject.getLength"
source: "fgl-topics/c_fgl_ext_util_JSONObject_getLength.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.getLength"
type: "concept"
---

# util.JSONObject.getLength

> Returns the number of name-value pairs in the JSON object.

## Syntax

```
getLength()
  RETURNS INTEGER
```

## Usage

The `getLength()` method returns the number of name-value
pairs in the JSON object.

This method can be used along with the `name()` and
`getType()` methods to read the entries of a JSON object.

## Example

```
IMPORT util
MAIN
    DEFINE obj util.JSONObject
    DEFINE i INTEGER
    LET obj = util.JSONObject.parse('{"id":123,"name":"Scott"}')
    FOR i=1 TO obj.getLength()
        DISPLAY i, ": ", obj.name(i), "=", obj.get(obj.name(i))
    END FOR
END MAIN
```

Output:

```
          1: id=                  123.0
          2: name=Scott
```

## Related links

**Related concepts**  

[util.JSONObject.name](3599-util-jsonobject-name.md "Returns the name of a JSON object entry by position.")

[util.JSONObject.getType](3597-util-jsonobject-gettype.md "Returns the type of a JSON object element.")
