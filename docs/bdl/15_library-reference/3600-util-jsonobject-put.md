---
title: "util.JSONObject.put"
source: "fgl-topics/c_fgl_ext_util_JSONObject_put.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.put"
type: "concept"
---

# util.JSONObject.put

> Sets a name-value pair in the JSON object.

## Syntax

```
put(
   name STRING,
   value value-type )
```

1. name is a string defining the entry name.
2. value is the value to be associated to the name.
3. value-type can be a primitive type, a
   `RECORD`, `DYNAMIC ARRAY`, a `util.JSONObject` or
   `util.JSONArray`.

## Usage

The `put()` method adds a name-value pair to the JSON object.

The first parameter is the name of the element.

The second parameter can be a primitive typed value such as a
`STRING` or `DECIMAL`, a complex variable defined as
`RECORD` or `DYNAMIC ARRAY`, a `util.JSONObject` or a
`util.JSONArray`.

If the element exists, the existing value is replaced.

For backward compatibility, the [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") and [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") constants are of the `INTEGER` type (1 and 0). If you
use these constants directly in an API writing JSON data, they will result in numerical values
`1` or `0`. To obtain JSON's `true` and
`false` boolean values, use a `BOOLEAN` variable instead.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE obj util.JSONObject
    DEFINE rec RECORD
               id INTEGER,
               name STRING
           END RECORD
    DEFINE arr DYNAMIC ARRAY OF INTEGER
    LET obj = util.JSONObject.create()
    CALL obj.put("simple", 234)
    LET rec.id = 234
    LET rec.name = "Barton"
    CALL obj.put("record", rec)
    LET arr[1] = 234
    LET arr[2] = 2837
    CALL obj.put("array", arr)
    DISPLAY obj.toString()
END MAIN
```

Output:

```
{"simple":234,"record":{"id":234,"name":"Barton"},"array":[234,2837]}
```

## Related links

**Related concepts**  

[util.JSONObject.get](3595-util-jsonobject-get.md "Returns the value corresponding to the specified entry name.")
