---
title: "json.Serializer.getOption"
source: "fgl-topics/c_gws_serializer_getOption.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods > json.Serializer.getOption"
type: "concept"
---

# json.Serializer.getOption

> Returns the current value of a JSON serializer option.

## Syntax

```
getOption(
   name STRING) RETURNS INTEGER
```

1. name is the name of an option of the JSON serializer of type
   `STRING`.

## Usage

Use this method to return the current value of a JSON serializer option. The returned
`INTEGER` value depends on the option; for valid values per option, see
[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example: return option value

Return the value of the [allowNullAsDefault](3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.") or
[serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.") options. In this example, the
`allowNullAsDefault` option is returned.

```
IMPORT json
DEFINE isAllowNullAsDefault INTEGER
MAIN
    DISPLAY json.Serializer.getOption("allowNullAsDefault")
    VAR isAllowNullAsDefault = json.Serializer.getOption("allowNullAsDefault")
    #...
END MAIN
```

## Related links

**Related concepts**  

[json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
