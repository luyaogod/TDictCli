---
title: "Deserializing JSON nulls with allowNullAsDefault"
source: "fgl-topics/c_gws_json_deserialization_null_options.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > Examples > allowNullAsDefault example"
type: "concept"
---

# Deserializing JSON nulls with allowNullAsDefault

> Use the allowNullAsDefault option to accept JSON null values during JSON to BDL deserialization.

By default, Genero Web Services JSON to Genero BDL deserialization will raise errors if the JSON
content does not match the BDL variable receiving the data. For example, you have a Genero BDL
record defined as follows:

```
TYPE tAddress RECORD
    street STRING,
    city STRING,
    state STRING,
    zip STRING
END RECORD

TYPE tCustomer RECORD
    id STRING,
    name STRING,
    address tAddress
END RECORD
```

The following JSON document will by default raise conversion errors because the "address" element
is defined as "null", which does not correspond with the BDL variable:

```
{
  "id": "1",
  "name": "John Doe",
  "address": null
}
```

To avoid the conversion error, set the following option to allow nulls in your server
module:

```
CALL json.Serializer.setOption( "allowNullAsDefault", 1 )
```

For more details, go to [json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.").

Alternatively, it is recommended that you define BDL variables with the attributes value
`json_null="null"` to handle JSON serialization of null
values.

```
TYPE tAddress RECORD ATTRIBUTES(json_null="null")
    street STRING,
    city STRING,
    state STRING,
    zip STRING
END RECORD
```

For more information and examples using `json_null`, go to
[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.").
