---
title: "Example 1: Writing to a JSON file"
source: "fgl-topics/c_gws_jsonJSONWriter_example_1.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > Examples > Example 1: Writing to a JSON file"
type: "concept"
---

# Example 1: Writing to a JSON file

> This example shows two ways to write a JSON file.

## Example: Writing data to a JSON file

This program writes JSON data to the customers.json file. In this example, a
JSONWriter object, `writer`, is created. A call is made to the [setOutputCharset](3658-json-jsonreader-setinputcharset.md "Defines the charset used on the input stream.") method to specify the
charset needed and then the file is opened for writing. The data is written in various calls to the
`writer` object methods to write the JSON objects and data.

A `TRY/CATCH` block is used to detect potential JSON format
errors.

```
IMPORT json

DEFINE writer json.JSONWriter

MAIN
    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeTo("customer.json")
        CALL writer.startJSON()
        CALL writer.startArray()
        CALL writer.startObject()
        CALL writer.setProperty("num")
        CALL writer.setValue(823)
        CALL writer.setProperty("name")
        CALL writer.setValue("Mark Renbing")
        CALL writer.endObject()
        CALL writer.startObject()
        CALL writer.setProperty("num")
        CALL writer.setValue(234)
        CALL writer.setProperty("name")
        CALL writer.setValue("Clark Gambler")
        CALL writer.endObject()
        CALL writer.endArray()
        CALL writer.endJSON()
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
            status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
END MAIN
```

The file customers.json:

```
[
   {
        "num": 823,
        "name": "Mark Renbing"
    }, {
        "num": 234,
        "name": "Clark Gambler"
    }
]
```

## Example: Writing data to a JSON file via JSON serializer

This program writes data to the customers2.json file using a BDL variable
and the [variableToJSON](3673-json-serializer-variabletojson.md "Serializes a Genero BDL variable into JSON using a JSONWriter object.") serializer method.
In this example, a `writer` object is created. A call is made to the [setOutputCharset](3658-json-jsonreader-setinputcharset.md "Defines the charset used on the input stream.") method to specify the
charset needed and then the file is opened for writing. In this example, the data is serialized via
a Genero BDL variable into JSON using the [variableToJSON](3673-json-serializer-variabletojson.md "Serializes a Genero BDL variable into JSON using a JSONWriter object.") method.

A `TRY/CATCH` block is used to detect potential JSON format
errors.

```
IMPORT json

DEFINE writer json.JSONWriter
DEFINE custlist DYNAMIC ARRAY OF RECORD
    num INTEGER,
    name VARCHAR(40)
END RECORD

MAIN
    LET custlist[1].num = 823
    LET custlist[1].name = "Mark Renbing"
    LET custlist[2].num = 234
    LET custlist[2].name = "Clark Gambler"
    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeTo("customer2.json")
        CALL writer.startJSON()
        CALL json.Serializer.variableToJSON(custlist, writer)
        CALL writer.endJSON()
    CATCH
        DISPLAY "JSON Serializer ERROR: shouldn't raise error :",
            status || " " || sqlca.sqlerrm
        EXIT PROGRAM -1
    END TRY
END MAIN
```

The file customers2.json:

```
[
   {
        "num": 823,
        "name": "Mark Renbing"
    }, {
        "num": 234,
        "name": "Clark Gambler"
    }
]
```

## Related links

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
