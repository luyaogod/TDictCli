---
title: "util.JSONObject.toFGL"
source: "fgl-topics/c_fgl_ext_util_JSONObject_toFGL.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.toFGL"
type: "concept"
---

# util.JSONObject.toFGL

> Fills a record variable with the entries contained in the JSON object.

## Syntax

```
toFGL(
  recordRef record-type )
```

1. recordRef is the variable to be set with values of the JSON string.
   > **Important:**
   >
   > The recordRef is a `RECORD` variable
   > is passed by reference to the method.
2. record-type is a `RECORD ... END RECORD` type.

## Usage

The `toFGL()` method fills the `RECORD` variable
passed as parameter with the corresponding values defined in the JSON object.

The destination record must have the same structure as the JSON source data.
For more details see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE cust_rec RECORD
               cust_num INTEGER,
               cust_name VARCHAR(30),
               order_ids DYNAMIC ARRAY OF INTEGER
           END RECORD
    DEFINE js STRING
    DEFINE obj util.JSONObject
    LET js='{ "cust_num":2735, "cust_name":"McCarlson",
              "order_ids":[234,3456,24656,34561] }'
    LET obj = util.JSONObject.parse( js )
    CALL obj.toFGL( cust_rec )
    DISPLAY cust_rec.cust_name
    DISPLAY cust_rec.order_ids[4]
END MAIN
```

Output:

```
McCarlson
      34561
```

## Related links

**Related concepts**  

[util.JSONObject.toString](3603-util-jsonobject-tostring.md "Builds a JSON string from the values contained in the JSON object.")
