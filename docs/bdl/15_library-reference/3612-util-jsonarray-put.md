---
title: "util.JSONArray.put"
source: "fgl-topics/c_fgl_ext_util_JSONArray_put.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.put"
type: "concept"
---

# util.JSONArray.put

> Sets an element by position in the JSON array object.

## Syntax

```
put(
   index INTEGER,
   value value-type )
```

1. index is the index of the element in the JSON array object.
2. value is the value to be associated to the index.
3. value-type can be a primitive type, a
   `RECORD`, `DYNAMIC ARRAY`, a `util.JSONObject` or
   `util.JSONArray`.

## Usage

The `put()` method sets an element value by position in the JSON array object.

The first parameter is the index of the element.

The second parameter can be a primitive typed value such as a
`STRING` or `DECIMAL`, a complex variable defined as
`RECORD` or `DYNAMIC ARRAY`, a `util.JSONObject` or a
`util.JSONArray`.

The index corresponding to the first element is 1.

If the element exists, the existing value is replaced.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE ja util.JSONArray
    DEFINE rec RECORD
               id INTEGER,
               name STRING
           END RECORD
    DEFINE arr DYNAMIC ARRAY OF INTEGER
    LET ja = util.JSONArray.create()
    CALL ja.put(1, 234)
    LET rec.id = 234
    LET rec.name = "Barton"
    CALL ja.put(2, rec)
    LET arr[1] = 234
    LET arr[2] = 2837
    CALL ja.put(3, arr)
    DISPLAY ja.toString()
END MAIN
```

Output:

```
[234,{"id":234,"name":"Barton"},[234,2837]]
```

## Related links

**Related concepts**  

[util.JSONArray.get](3609-util-jsonarray-get.md "Returns the value of a JSON array element.")
