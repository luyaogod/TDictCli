---
title: "JSONRequired"
source: "fgl-topics/c_gws_JSON_attribute_JSONRequired_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > json.Serializer attributes > JSONRequired"
type: "concept"
---

# JSONRequired

> Specify properties that are required in a JSON schema.

## Syntax

```
JSONRequired
```

`JSONRequired` is an optional attribute.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

You can set `JSONRequired` on elements of a record or a user-defined [type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.") to specify properties that are required in the JSON
Schema. You can also set `JSONRequired` directly at the root of the type, and in this
case all members are required.

If the `JSONRequired` attribute is placed at the root of a record or on one of its
primitive members, ​​it may be also necessary to set the [json\_null="null"](4847-openapi-types-mapping-bdl.md) attribute on the relevant field, or set the [serializeNullAsDefault](../15_library-reference/3685-json-serializer-options.md) option for the [The json.Serializer class](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa."), to allow
nulls; otherwise, [error-15807](../15_library-reference/4483-genero-bdl-errors.md)
will be raised if a `null` value is sent. Any attempt to send a `null`
value to a field marked with `JSONRequired` will result in this error. For examples
allowing null values with `JSONRequired`, go to [Serializing JSON nulls with serializeNullAsDefault](../15_library-reference/3693-serializenullasdefault-example.md "Use the serializeNullAsDefault option to accept BDL NULL values during JSON serialization.").

The `JSONRequired` attribute supports the `required` keyword in the
JSON schema property of the Swagger and OpenAPI specification.

## Example 1 using JSONRequired at the root of a type definition

In this example, the `JSONRequired` attribute is set directly at the root of the
type definition, and in this case all members are `required` properties in the JSON
schema.

```
TYPE profileType RECORD ATTRIBUTE(JSONRequired)
    id INTEGER ATTRIBUTES(WSDescription = "Internal Identifier"),
    name VARCHAR(100) ATTRIBUTES(WSDescription = "Lastname"),
    email VARCHAR(255),
    category VARCHAR(10) ATTRIBUTES(WSDescription = "User Demographic"),
    status INTEGER,
    ccode VARCHAR(3) ATTRIBUTES(WSDescription = "Country Code")
END RECORD
```

## Example 2 using JSONRequired in members of a type definition

In this example, the `JSONRequired` attribute specifies the `id`,
`name`, and `email` members of the `profileType` record
as `required` properties in the JSON schema.

```
TYPE profileType RECORD ATTRIBUTE(WSTypeDescription = "profile of user")
    id INTEGER ATTRIBUTES (WSDescription="Internal Identifier", JSONRequired), 
    name VARCHAR(100) ATTRIBUTES (WSDescription="Lastname", JSONRequired),
    email VARCHAR(255) ATTRIBUTE(JSONRequired),
    category VARCHAR(10) ATTRIBUTES (WSDescription="User Demographic"),
    status INTEGER,
    ccode VARCHAR(3) ATTRIBUTES (WSDescription="Country Code")
    # ...
    END RECORD
```

In the OpenAPI documentation, the GWS engine exposes the `required` property in
the `schema` for the record elements defined with `JSONRequired`.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

In this sample schema defining a record of `profileType`, a request to the service
will require that all members are required; otherwise, the request will fail.

![Image from the OpenAPI document showing the profileType JSON schema with all required properties](../_images/rest_openapi_api_jsonrequired_json_schema_all.png)

*Sample JSON schema with all required properties*

In this sample schema defining a record of `profileType`, a request to the service
will require that each user has an id, name and email address; otherwise, the request will fail.
Providing the properties not required is optional.

![Image from the OpenAPI document showing the profileType JSON schema with some required properties](../_images/rest_openapi_api_jsonrequired_json_schema.png)

*Sample JSON schema with some required properties*

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

**Related reference**  

[json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
