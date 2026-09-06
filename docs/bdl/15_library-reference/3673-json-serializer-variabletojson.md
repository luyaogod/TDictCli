---
title: "json.Serializer.VariableToJSON"
source: "fgl-topics/c_gws_serializer_VariableToJSON.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods > json.Serializer.VariableToJSON"
type: "concept"
---

# json.Serializer.VariableToJSON

> Serializes a Genero BDL variable into JSON using a JSONWriter object.

## Syntax

```
VariableToJSON(
   var AnyType,
   json json.JSONWriter )
```

1. var can be of [various
   kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").
2. json is a `JSONWriter` object.

## Usage

This method serializes a Genero BDL variable into JSON using a JSONWriter object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
IMPORT json
MAIN
    DEFINE writer json.JSONWriter
    DEFINE var RECORD
        cust_num INTEGER,
        cust_name VARCHAR(30),
        order_ids DYNAMIC ARRAY OF INTEGER
    END RECORD
    LET var.cust_num = "2735"
    LET var.cust_name = "McCarlson"
    LET var.order_ids[1] = 234
    LET var.order_ids[2] = 3456
    LET var.order_ids[3] = 24656
    LET var.order_ids[4] = 34561
    TRY
        DISPLAY "START test"
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeTo("toto1.json")
        CALL writer.startJson()
        CALL json.Serializer.VariableToJSON(var, writer) # serialize var into the file
        CALL writer.endJson()
        CALL writer.close()
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
        status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
    DISPLAY "END test ok"
END MAIN
```

Output:

```
{"cust_num": 2735,"cust_name": "McCarlson","order_ids": [234,3456,24656,34561]}
```

## Related links

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
