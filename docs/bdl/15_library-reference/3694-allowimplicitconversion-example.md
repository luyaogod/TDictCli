---
title: "Deserializing JSON values with allowImplicitConversion"
source: "fgl-topics/c_gws_json_deserialization_allow_implicit_conversion.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > Examples > allowImplicitConversion example"
type: "concept"
---

# Deserializing JSON values with allowImplicitConversion

> Use the allowImplicitConversion option to accept JSON values that require implicit type conversion during deserialization.

By default, Genero Web Services deserialization will raise errors if the JSON content does not
match the BDL variable receiving the data. For more information on implicit and explicit conversion
of JSON to Genero BDL, go to [Implicit and explicit JSON conversion in json.Serializer and util.JSON](3663-implicit-and-explicit-conversion.md "This topic provides an overview of implicit and explicit JSON conversion behavior in json.Serializer and util.JSON, highlighting the differences between strict and permissive conversion models.").

In this example, you have a BDL variable that is defined as
follows:

```
PRIVATE TYPE TypedLists RECORD
    ints DYNAMIC ARRAY OF INT,
    strings DYNAMIC ARRAY OF STRING
END RECORD
```

The following JSON document will raise conversion errors by default because the "strings" array
does not contain valid string values. For example, it should be formatted as "`strings":
["1", "2"]`" to correspond with the expected data type. Here is the original
JSON:

```
{
    "ints": [1, 2],
    "strings": [1, 2]
}
```

To avoid the conversion error, you can set the following option in your server module to allow
implicit type
conversion:

```
CALL json.Serializer.setOption( "allowImplicitConversion", 1 )
```

For more details, go to [json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.").
