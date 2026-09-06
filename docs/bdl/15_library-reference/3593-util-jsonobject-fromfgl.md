---
title: "util.JSONObject.fromFGL"
source: "fgl-topics/c_fgl_ext_util_JSONObject_fromFGL.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.fromFGL"
type: "concept"
---

# util.JSONObject.fromFGL

> Creates a new JSON object from a RECORD.

## Syntax

```
util.JSONObject.fromFGL(
   record record-type )
  RETURNS util.JSONObject
```

1. record is the record variable used to create the JSON object.
2. record-type is a `RECORD ... END RECORD` type.

## Usage

The `util.JSONObject.fromFGL()` method creates a new JSON object
from the `RECORD` variable passed as parameter.

The created object must be assigned to a program variable defined with the
`util.JSONObject` type.

The members
of the `RECORD` are converted to name/value pairs in the JSON object.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE cust_rec RECORD
               cust_num INTEGER,
               cust_name VARCHAR(30),
               order_ids DYNAMIC ARRAY OF INTEGER
           END RECORD
    DEFINE obj util.JSONObject
    LET cust_rec.cust_num = 345
    LET cust_rec.cust_name = "McMaclum"
    LET cust_rec.order_ids[1] = 4732
    LET cust_rec.order_ids[2] = 9834
    LET cust_rec.order_ids[3] = 2194
    LET obj = util.JSONObject.fromFGL(cust_rec)
    DISPLAY obj.toString()
END MAIN
```

Output:

```
{"cust_num":345,"cust_name":"McMaclum","order_ids":[4732,9834,2194]}
```

## Related links

**Related concepts**  

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")

[util.JSONObject.create](3592-util-jsonobject-create.md "Creates a new JSON object.")

[util.JSONObject.parse](3594-util-jsonobject-parse.md "Parses a JSON string and creates a JSON object from it.")
