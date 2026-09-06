---
title: "json.Serializer.setOption"
source: "fgl-topics/c_gws_serializer_setOption.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods > json.Serializer.setOption"
type: "concept"
---

# json.Serializer.setOption

> Sets an option on the JSON serializer.

## Syntax

```
setOption(
   name STRING,
   value INTEGER )
```

1. name is the name of a JSON serializer option of type
   `STRING`.
2. value is an `INTEGER`.

## Usage

Use this method to configure options for the JSON serializer. For a list of available
options and their values, see [json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example: Set `allowNullAsDefault`

In this example, the [allowNullAsDefault](3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.") option is
set to allow nulls. For more details, go to [Deserializing JSON nulls with allowNullAsDefault](3692-allownullasdefault-example.md "Use the allowNullAsDefault option to accept JSON null values during JSON to BDL deserialization.").

```
IMPORT json
MAIN
    CALL json.Serializer.setOption("allowNullAsDefault",1)
    #...
END MAIN
```

For examples setting the [serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.") option, go to [Serializing JSON nulls with serializeNullAsDefault](3693-serializenullasdefault-example.md "Use the serializeNullAsDefault option to accept BDL NULL values during JSON serialization.").

## Related links

**Related concepts**  

[json.Serializer.getOption](3674-json-serializer-getoption.md "Returns the current value of a JSON serializer option.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
