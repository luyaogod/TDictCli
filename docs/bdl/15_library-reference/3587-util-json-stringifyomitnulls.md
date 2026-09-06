---
title: "util.JSON.stringifyOmitNulls"
source: "fgl-topics/c_fgl_ext_util_JSON_stringifyOmitNulls.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.stringifyOmitNulls"
type: "concept"
---

# util.JSON.stringifyOmitNulls

> Produces a JSON formatted string from the input, by excluding empty records and empty arrays.

## Syntax

```
util.JSON.stringifyOmitNulls(
     value any-type
  )
  RETURNS STRING
```

1. value is the program variable to be converted to a JSON string.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `util.JSON.stringifyOmitNulls()` class method takes a variable as parameter,
and generates the corresponding data string in JSON format, as defined in the [[RFC4627]](http://www.ietf.org/rfc/rfc4627.txt) specification.

The input parameter can be a `RECORD`, an `ARRAY`
or a `DICTIONARY`, as well as a simple type such as `INTEGER` or
`STRING`.

Unlike [`util.JSON.stringify()`](3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays."), if a record member is null, it will not be written
to the JSON output. If all members of a record are null, the record itself will be omitted in the
output. If an array is empty, it will be excluded from the output.

Java Script distinguishes null and undefined, while Genero BDL does not
have a concept of undefined variable. The `util.JSON.stringifyOmitNulls`
method considers that BDL `NULL` values are undefined and must not be
serialized. However, if an `util.JSONObject` property exists and has the value
`NULL`, or when an `util.JSONArray` item exists and is
`NULL`, then this value will be serialized.

For full control on null and empty variables when serializing to JSON elements, use the
`json_null` variable definition attribute and the `stringify()`
method. The method `stringifyOmitNulls()` behaves like `stringify()`,
when all variable elements are defined with the `json_null="undefined"`
attribute.

The method raises error [-8110](4483-genero-bdl-errors.md)
if the JSON string cannot be generated.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE rec RECORD
             field1 INTEGER,
             field2 CHAR(1),
             subrec1 RECORD
               field11 INTEGER,
               fiedl12 VARCHAR(30)
             END RECORD,
             subarr1 DYNAMIC ARRAY OF INTEGER,
             jo util.JSONObject,
             ja util.JSONArray
           END RECORD
    INITIALIZE rec TO NULL
    LET rec.field1 = 999
    LET rec.jo = util.JSONObject.create()
    CALL rec.jo.put("name","Tiger BOUDS")
    CALL rec.jo.put("address",NULL) -- property exists and is NULL
    LET rec.ja = util.JSONArray.create()
    CALL rec.ja.put(1,"item1")
    CALL rec.ja.put(2,NULL) -- element exists and is NULL
    DISPLAY "stringify()         : ", util.JSON.format( util.JSON.stringify(rec) )
    DISPLAY "stringifyOmitNulls(): ", util.JSON.format( util.JSON.stringifyOmitNulls(rec) )
END MAIN
```

Output:

```
stringify()         : {
    "field1": 999,
    "subrec1": {},
    "subarr1": [],
    "jo": {
        "name": "Tiger BOUDS",
        "address": null
    },
    "ja": [
        "item1",
        null
    ]
}
stringifyOmitNulls(): {
    "field1": 999,
    "jo": {
        "name": "Tiger BOUDS",
        "address": null
    },
    "ja": [
        "item1",
        null
    ]
}
```

## Related links

**Related concepts**  

[util.JSON.stringify](3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.")

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")
