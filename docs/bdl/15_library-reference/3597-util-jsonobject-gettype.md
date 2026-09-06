---
title: "util.JSONObject.getType"
source: "fgl-topics/c_fgl_ext_util_JSONObject_getType.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.getType"
type: "concept"
---

# util.JSONObject.getType

> Returns the type of a JSON object element.

## Syntax

```
getType(
   name STRING )
  RETURNS STRING
```

1. name is the name of the element.

## Usage

The `getType()` method returns the JSON data type name
corresponding to the element name passed as parameter.

This method can be used along with the `name()` and
`getLength()` methods, to read the entries of a JSON object.

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
    DEFINE obj util.JSONObject
    LET obj = util.JSONObject.create()
    CALL obj.put("id", 8723)
    DISPLAY obj.getType("id") -- NUMBER
    CALL obj.put("name", "Brando")
    DISPLAY obj.getType("name") -- STRING
    DISPLAY obj.getType("undef") -- NULL
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

[util.JSONObject.getLength](3596-util-jsonobject-getlength.md "Returns the number of name-value pairs in the JSON object.")

[util.JSONObject.name](3599-util-jsonobject-name.md "Returns the name of a JSON object entry by position.")
