---
title: "Example 2: Handling arrays with null values"
source: "fgl-topics/c_gws_high_level_rest_bdl_array_null_value_conversion.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > OpenAPI types mapping: BDL > Null value conversions > Arrays with null values"
type: "concept"
---

# Example 2: Handling arrays with null values

> This example shows how null values in arrays are handled during serialization.

The following program reads data with null values into the array arr, converts the array to JSON,
and writes the result to standard output. The conversion calls are enclosed in a
`TRY/CATCH` block to detect JSON serialization errors.

```
 # ... 
 DEFINE t TEXT
 DEFINE writer json.JSONWriter
 DEFINE arr DYNAMIC ARRAY OF STRING
    LOCATE t IN MEMORY
    INITIALIZE t to NULL

    LET arr[1] = "foo"
    LET arr[2] = NULL
    LET arr[3] = NULL

    TRY 
        LET writer = json.JSONWriter.Create()
        CALL writer.setOutputCharset("UTF-8")
        CALL writer.writeToText(t)
        CALL writer.startJSON()
        CALL json.Serializer.variableToJSON(arr, writer)
        CALL writer.endJSON()
        CALL writer.close()
    CATCH
        DISPLAY "\narr with NULL value:" || "ERROR " || status || " " || sqlca.sqlerrm
    END TRY
    # ...
```

When this array is serialized, [error-15807](../15_library-reference/4483-genero-bdl-errors.md) is raised:

```
arr with NULL value:ERROR -15807 Array cannot serialize 'null' elements. It requires json_null="null"
```

An array element can be serialized as `NULL` only if it is explicitly enabled on
the element
type:

```
DEFINE arr DYNAMIC ARRAY OF STRING ATTRIBUTE(json_null="null")
```

Alternatively,
you can allow null values globally by setting the [allowNullAsDefault](../15_library-reference/3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.")
option:

```
CALL json.Serializer.setOption( "allowNullAsDefault", 1 )
```

If the entire array is `NULL`, no error is raised. For
example:

```
INITIALIZE arr TO NULL
```

In this case, GWS serializes the array as an
empty object (`{}`), or as null if the [`json_null="null"`](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") attribute is applied to the
array:

```
DEFINE arr DYNAMIC ARRAY ATTRIBUTE(json_null="null") OF STRING ATTRIBUTE(json_null="null")
```

For more details, refer to the [OpenAPI](http://swagger.io/docs/specification/about/) documentation.
