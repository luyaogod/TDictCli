---
title: "Serializing JSON nulls with serializeNullAsDefault"
source: "fgl-topics/c_gws_json_serialization_options.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > Examples > serializeNullAsDefault example"
type: "concept"
---

# Serializing JSON nulls with serializeNullAsDefault

> Use the serializeNullAsDefault option to accept BDL NULL values during JSON serialization.

By default, Genero Web Services BDL to JSON serialization will raise errors if
`NULL` values are sent.

Null values can be allowed in Genero BDL in two ways:

- By defining BDL variables with the attributes value [json\_null=null](3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.").
- By setting the [serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.")
  option.

In this example, the `str` variable is not set with a value and is therefore
`NULL` when passed in the call to `json.Serializer.variableToJSON(str,
writer)` and an error is raised.

```
IMPORT JSON

CONSTANT TNAME = "one_value_withoutNull()"
DEFINE 
    writer json.JSONWriter,
    t TEXT,
    str STRING

MAIN
    LOCATE t in MEMORY
    DISPLAY "\nSTART ", TNAME

    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeToText(t)
        CALL writer.startJSON()
        CALL json.Serializer.variableToJSON(str, writer)
        CALL writer.endJSON()
        CALL writer.close()
        DISPLAY "Should raise a serialization error"
        DISPLAY "json: ", t
        EXIT PROGRAM 1
    CATCH
        DISPLAY status || ": " || sqlca.sqlerrm
        DISPLAY "END " || TNAME || " ok"
    END TRY
END MAIN
```

[error-15807](4483-genero-bdl-errors.md) is raised and
the following output is returned because "`str`" is
`NULL`:

```
START one_value_withoutNull()
-15807: Primitive value cannot be serialized to 'null'
END one_value_withoutNull() ok
```

To avoid serialization errors, call `json.Serializer.setOption("serializeNullAsDefault",
1)` to allow null values, even when a [`JSONRequired`](3683-jsonrequired.md "Specify properties that are required in a JSON schema.") constraint is set on the record. This is demonstrated in the
following example. Note that you must reset the [serializeNullAsDefault](3685-json-serializer-options.md) option to its default setting at the end of the
program.

```
IMPORT JSON

CONSTANT TNAME = "object_withOption()"

DEFINE rec RECORD
    a STRING ATTRIBUTE(JSONRequired),
    b INTEGER
END RECORD
DEFINE 
    writer json.JSONWriter,
    t TEXT
MAIN
    LOCATE t in MEMORY
    DISPLAY "\nSTART ", TNAME

    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeToText(t)
        CALL writer.startJSON()
        CALL json.Serializer.setOption("serializeNullAsDefault",1)
        CALL json.Serializer.variableToJSON(rec, writer)
        CALL writer.endJSON()
        CALL writer.close()
        DISPLAY "json: ", t
    CATCH
        DISPLAY status || ": " || sqlca.sqlerrm
        EXIT PROGRAM 1
    END TRY

    DISPLAY "END " || TNAME || " ok"

    # reset serializeNullAsDefault to default
    CALL json.Serializer.setOption("serializeNullAsDefault",0)

END MAIN
```

No serialization error is raised and the following output is produced with null values:

```
START object_withOption()
json: {"a":null,"b":0}
END object_withOption() ok
```

For more details about `serializeNullAsDefault`, go to [json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.").

## Serialization error with nulls in an array

By default, Genero Web Services BDL to JSON serialization will raise errors if
`NULL` values are sent in an array as in this example:

```
IMPORT JSON

CONSTANT TNAME = "array_withoutNull()"
DEFINE arr DYNAMIC ARRAY OF STRING
DEFINE 
    writer json.JSONWriter,
    t TEXT

MAIN
    LOCATE t in MEMORY
    DISPLAY "\nSTART ", TNAME

    LET arr[1] = "foo"
    LET arr[2] = NULL

    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeToText(t)
        CALL writer.startJSON()
        CALL json.Serializer.variableToJSON(arr, writer)
        CALL writer.endJSON()
        CALL writer.close()
        DISPLAY "Should generate serialization error"
        DISPLAY "json: ", t
        EXIT PROGRAM 1
    CATCH
        DISPLAY status || ": " || sqlca.sqlerrm
        DISPLAY "END " || TNAME || " ok"
    END TRY
END MAIN
```

[error-15807](4483-genero-bdl-errors.md) is raised and the following output is returned because "`arr[2]`"
is `NULL`:

```
START array_withoutNull()
-15807: Array cannot serialize 'null' elements. It requires json_null="null"
END array_withoutNull() ok
```

To prevent serialization errors in arrays, you should call
`json.Serializer.setOption("serializeNullAsDefault", 1)` before writing to the array.
This will allow null values. Note that you must reset the [serializeNullAsDefault](3685-json-serializer-options.md) option to its default setting at the end of the
program.

```
IMPORT JSON

CONSTANT TNAME = "array_withOption()"
DEFINE arr DYNAMIC ARRAY OF STRING
DEFINE 
    writer json.JSONWriter,
    t TEXT

MAIN
    LOCATE t in MEMORY
    DISPLAY "\nSTART ", TNAME

    LET arr[1] = "foo"
    LET arr[2] = NULL

    TRY
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeToText(t)
        CALL writer.startJSON()
        CALL json.Serializer.setOption("serializeNullAsDefault",1)
        CALL json.Serializer.variableToJSON(arr, writer)
        CALL writer.endJSON()
        CALL writer.close()
        DISPLAY "json: ", t
    CATCH
        DISPLAY status || ": " || sqlca.sqlerrm
        EXIT PROGRAM 1
    END TRY

    DISPLAY "END " || TNAME || " ok"
    # reset serializeNullAsDefault to default
    CALL json.Serializer.setOption("serializeNullAsDefault",0)
END MAIN
```

No serialization error is raised and the following output is produced with null values:

```
START array_withOption()
json: ["foo",null]
END array_withOption() ok
```

For more details about `serializeNullAsDefault`, go to [json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.").

## Using `json_null="null"` to manage JSON serialization of null values

Alternatively, it is recommended that you define BDL variables with the attributes value
`json_null="null"` to allow JSON serialization of null values.

```
DEFINE str STRING ATTRIBUTE(json_null="null")
```

You can define an array with the
`json_null="null"` attribute to handle a null index or an attempt to serialize it
without elements. For examples, go to [Using the json\_null="null" attribute](../09_advanced-features/0959-nulls-and-empty-structures.md).

## Related links

**Related concepts**  

[json.Serializer.getOption](3674-json-serializer-getoption.md "Returns the current value of a JSON serializer option.")

[json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.")

[json.JSONWriter methods](3619-json-jsonwriter-methods.md "Methods for the json.JSONWriter class.")
