---
title: "json.Serializer.JSONToVariable"
source: "fgl-topics/c_gws_serializer_JSONToVariable.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods > json.Serializer.JSONToVariable"
type: "concept"
---

# json.Serializer.JSONToVariable

> Serializes entries contained in a JSONReader object into a Genero BDL variable.

## Syntax

```
JSONToVariable(
  json json.JSONReader, 
  var AnyType )
```

1. json is a `JSONReader` object.
2. var can be of [various
   kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `JSONToVariable()` method fills the variable passed as parameter with the
corresponding values defined in the JSON object.

The destination record must have the same structure as the JSON source data.
For more details see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
IMPORT json
MAIN
    DEFINE cust_rec RECORD
               cust_num INTEGER,
               cust_name VARCHAR(30),
               order_ids DYNAMIC ARRAY OF INTEGER
           END RECORD
    DEFINE i INTEGER
    DEFINE reader json.JSONReader
    TRY
        DISPLAY "START test"
        LET reader = json.JSONReader.Create()
        CALL reader.setInputCharset("UTF-8")
        CALL reader.readFrom("toto1.json")
        CALL reader.next()
        CALL json.Serializer.JSONToVariable(reader, cust_rec)
        CALL reader.close()
        DISPLAY "Cust name:", cust_rec.cust_name
        FOR i = 1 TO cust_rec.order_ids.getLength()
            DISPLAY "Order id:", cust_rec.order_ids[i]
        END FOR
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
            status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
    DISPLAY "END test ok"
END MAIN
```

The toto1.json file :

```
{
    "cust_num": 2735,
    "cust_name": "McCarlson",
    "order_ids": [234, 3456, 24656, 34561]
}
```

## Related links

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
