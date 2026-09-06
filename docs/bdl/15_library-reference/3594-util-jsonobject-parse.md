---
title: "util.JSONObject.parse"
source: "fgl-topics/c_fgl_ext_util_JSONObject_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.parse"
type: "concept"
---

# util.JSONObject.parse

> Parses a JSON string and creates a JSON object from it.

## Syntax

```
util.JSONObject.parse(
   s STRING )
  RETURNS util.JSONObject
```

1. s is a string value that contains JSON formatted data.

## Usage

The `util.JSONObject.parse()` method scans the JSON source string
passed as parameter and creates a JSON object from it.

The created object must be assigned to a program variable defined with the
`util.JSONObject` type.

The source string must follow the JSON
format specification. It can contain multi-level structured data, but it must start
with a curly brace.

The method raises error [-8109](4483-genero-bdl-errors.md) if the JSON source string is not properly formatted. Consider enclosing the
`parse()` method call in a `TRY/CATCH` block, if the source string can
be malformed JSON.

## Example

```
IMPORT util
MAIN
    DEFINE js STRING
    DEFINE obj util.JSONObject
    LET js='{ "cust_num":2735, "cust_name":"McCarlson",
              "orderids":[234,3456,24656,34561] }'
    LET obj = util.JSONObject.parse( js )
    DISPLAY obj.get("cust_name")
END MAIN
```

Output:

```
McCarlson
```

## Related links

**Related concepts**  

[util.JSONObject.create](3592-util-jsonobject-create.md "Creates a new JSON object.")

[util.JSONObject.fromFGL](3593-util-jsonobject-fromfgl.md "Creates a new JSON object from a RECORD.")
