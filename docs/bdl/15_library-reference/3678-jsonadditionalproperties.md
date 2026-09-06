---
title: "JSONAdditionalProperties"
source: "fgl-topics/c_gws_JSON_attribute_JSONAdditionalProperties.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer attributes > JSONAdditionalProperties"
type: "concept"
---

# JSONAdditionalProperties

> Allows a record to accept JSON properties not explicitly listed in its schema definition.

## Syntax

```
JSONAdditionalProperties = { true | false }
```

`JSONAdditionalProperties` is an optional attribute. Default value is true.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

Set `JSONAdditionalProperties` on a record, dictionary type, or [user-defined type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") to control whether it accepts properties
not listed in its definition. The default is true: additional properties are accepted.

`JSONAdditionalProperties` represents the `additionalProperties`
keyword in the OpenAPI JSON schema specification.

## Example 1: using JSONAdditionalProperties

This example shows the three ways to set `JSONAdditionalProperties`:

- `MyRecord` sets it to `false`: the GWS engine rejects any
  JSON property not defined in the record.
- `MyArray` sets it explicitly to `true`: additional properties
  are accepted.
- `MyComplexType` omits the value, which defaults to `true`:
  additional properties are accepted.

```
TYPE MyRecord RECORD ATTRIBUTE(JSONAdditionalProperties="false")
  a,b INTEGER  # extra JSON properties rejected
END RECORD

TYPE MyArray DYNAMIC ARRAY OF RECORD ATTRIBUTE(JSONAdditionalProperties = "true")
   a,b INTEGER  # extra JSON properties accepted (explicit true)
END RECORD

TYPE MyComplexType RECORD ATTRIBUTE(JSONAdditionalProperties)
  a,b INTEGER  # extra JSON properties accepted (no value = default true)
END RECORD
```

The GWS engine exposes the `additionalProperties` property in the OpenAPI schema
for records defined with `JSONAdditionalProperties`. In the following schema,
`MyRecord` rejects additional properties and `MyComplexType` accepts
them. Open curly braces (`{}`) denote true.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the additionalProperties JSON schema property](../_images/rest_openapi_jsonadditionalproperties_attr_schema.png)

*Sample JSON schema with additionalProperties property*

## Example 2: capturing additional properties

With `JSONAdditionalProperties=true` on the record alone, you can accept unknown
JSON properties without error, but you won't be able to reference them in your code.

If you need to reference them in your code, add a `DICTIONARY` field also marked
with `JSONAdditionalProperties`. Any JSON property not matched to a named field
goes into that dictionary, with the property name as the key.

Declare the dictionary as `OF STRING` to capture additional properties as string
values:

```
TYPE MyStoredTypeString RECORD ATTRIBUTE(JSONAdditionalProperties)
  a,b INTEGER,
  p DICTIONARY ATTRIBUTE(JSONAdditionalProperties) OF STRING
END RECORD
```

Declare the dictionary as `OF RECORD` to capture structured additional
properties:

```
TYPE MyStoredTypeRecord RECORD ATTRIBUTE(JSONAdditionalProperties)
  a,b STRING,
  p DICTIONARY ATTRIBUTE(JSONAdditionalProperties) OF RECORD
    d,e INTEGER
  END RECORD
END RECORD
```

## Example 3: using JSONOneOf with JSONAdditionalProperties

This example shows how `JSONAdditionalProperties="false"` makes
`JSONOneOf` work when sub-records share common fields (`birthday`
and `gender`):

- **Without the attribute:** all sub-record schemas accept any request, so
  `oneOf` cannot resolve.
- **With the attribute set to false:** a sub-record is rejected if the request contains
  a property it does not define; only the variant whose fields match the request remains
  valid.

For the `JSONRequired` alternative, see [Example 2: using JSONOneOf with JSONRequired](3681-jsononeof.md).

```
# JSONAdditionalProperties=false is mandatory on each sub-record
# Otherwise all schemas are valid and oneOf cannot resolve
TYPE oneOfType RECORD ATTRIBUTE(JSONOneOf)
    _selector INTEGER ATTRIBUTE(JSONSelector),
    accountById RECORD ATTRIBUTE(JSONAdditionalProperties="false")    # variant 1: look up by id
        id INTEGER,
        birthday DATE,
        gender STRING
    END RECORD,
    accountByname RECORD ATTRIBUTE(JSONAdditionalProperties="false")  # variant 2: look up by name
        name STRING,
        birthday DATE,
        gender STRING
    END RECORD,
    accountBydate RECORD ATTRIBUTE(JSONAdditionalProperties="false")  # variant 3: look up by date
        date DATETIME YEAR TO SECOND,
        birthday DATE,
        gender STRING
    END RECORD
END RECORD

FUNCTION echoOneOfType(
    in oneOfType)
    ATTRIBUTE(WSPost, WSPath = "/echoOneOfType")
    RETURNS(oneOfType)
    DEFINE out oneOfType

    LET out = in

    DISPLAY out._selector

    CASE
        WHEN out._selector = 2
            DISPLAY "schema 1, id: ", out.accountById.id

        WHEN out._selector = 3
            DISPLAY "schema 2, name: ", out.accountByname.name

        WHEN out._selector = 4
            DISPLAY "schema 3, date: ", out.accountBydate.date
    END CASE

    RETURN out
END FUNCTION
```

The GWS engine generates an OpenAPI schema with `oneOf` and
`additionalProperties` set to false.

The output shown in this example is from the Firefox browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the oneOf and additionalProperties JSON schema properties](../_images/rest_openapi_jsonadditionalprops_jsononeof_schema.png)

*Sample JSON schema with oneOf and additionalProperties properties*

## Related links

**Related concepts**  

[JSONSelector](3684-jsonselector.md "Identifies the variant member of a JSONOneOf record to use during serialization or deserialization.")
