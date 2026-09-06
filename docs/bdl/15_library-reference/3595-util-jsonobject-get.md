---
title: "util.JSONObject.get"
source: "fgl-topics/c_fgl_ext_util_JSONObject_get.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.get"
type: "concept"
---

# util.JSONObject.get

> Returns the value corresponding to the specified entry name.

## Syntax

```
get(
   name STRING )
  RETURNS result-type
```

1. name is the string identifying the JSON object property.
2. result-type can be a primitive type, a `util.JSONObject` or a
   `util.JSONArray` object reference.

## Usage

The `get()` method returns the value or JSON object corresponding to the element
name passed as parameter.

If the element identified by the name is a simple value, the method returns a
string. If the element is structured, the method returns a `util.JSONObject` instance
and the returned object must be assigned to a program variable defined with the
`util.JSONObject` type. If the element is a list of values, the method
`util.JSONArray` instance and the returned object must be assigned to a program
variable defined with the `util.JSONArray` type.
> **Note:**
>
> If the
> element is a simple value, pay attention to the data format when directly assigning the result of
> the `get()` method to a program variable. For example, if the JSON string value
> represents a date, it will use the ISO format (YYYY-MM-DD). Assigning this ISO-formatted date string
> to a `DATE` variable will fail, if DBDATE format does not match.

A name/value pair can be set with the `put()` method.

## Example

```
IMPORT util
MAIN
    DEFINE obj, sub util.JSONObject
    DEFINE jarr util.JSONArray
    DEFINE rec RECORD
               id INTEGER,
               name STRING
           END RECORD
    DEFINE arr DYNAMIC ARRAY OF INTEGER
    DEFINE x INT
    LET obj = util.JSONObject.create()
    -- Simple value
    CALL obj.put("simple", 234)
    LET x = obj.get("simple")
    DISPLAY "x     :", x
    -- Sub-element
    LET rec.id = 234
    LET rec.name = "Barton"
    CALL obj.put("record", rec)
    LET sub = obj.get("record")
    DISPLAY "sub   :", sub.toString()
    -- Array
    LET arr[1] = 234
    LET arr[2] = 2837
    CALL obj.put("array", arr)
    LET jarr = obj.get("array")
    DISPLAY "jarr  :", jarr.toString()
END MAIN
```

Output:

```
x     :        234
sub   :{"id":234,"name":"Barton"}
jarr  :[234,2837]
```

## Related links

**Related concepts**  

[util.JSONObject.put](3600-util-jsonobject-put.md "Sets a name-value pair in the JSON object.")
