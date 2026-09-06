---
title: "How GWS maps BDL types into the OpenAPI description"
source: "fgl-topics/r_gws_openapi_bdl_mapping.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > How GWS maps BDL types into the OpenAPI description"
type: "reference"
---

# How GWS maps BDL types into the OpenAPI description

> GWS maps Genero BDL types to JSON schema objects in the OpenAPI description. The mapping depends on the content type and whether the data type is named or anonymous.

## Anonymous types

When an operation uses an anonymous type, or when the content type is set to
`application/xml`, the GWS engine exposes the object structure directly in the
OpenAPI description. The properties appear inline within the `schema` of the request
or response body.

This mapping is used when the type is not explicitly named in your BDL code or when XML content
is required.

![Anonymous type showing properties exposed inline under the schema element](../_images/rest_openapi_anonymous_type.png)

*Anonymous type exposed inline*

## Named types

When an operation uses a named type and the content type is
`application/json`, the GWS engine represents the object using a
`$ref`. The reference points to a definition in the
`components/schemas` section of the OpenAPI description.

In this documentation, a named type refers to a schema that is defined once and
reused through a `$ref` reference. This follows the JSON Schema pattern used by
OpenAPI. An anonymous type refers to an inline schema whose properties are written
directly inside the `schema` element. Anonymous types cannot be reused, while named
types appear under `components/schemas` and can be referenced by multiple operations.

This reuse model keeps the OpenAPI description concise and ensures that changes to a named type
are reflected consistently across all operations that reference it.

![Named type using a JSON schema reference under components/schemas](../_images/rest_openapi_named_type.png)

*Named type represented by a $ref*

## Parameters

Parameters defined with `WSHeader`, `WSQuery`, or
`WSCookie` attributes are mapped to the corresponding OpenAPI parameter entries.
These parameters are independent of the `$ref` mechanism and are always described
inline.

## Error handling

Error types defined with the `WSError` attribute follow the same mapping rules as
named or anonymous types, depending on their content type. The generated OpenAPI description shows
error responses in the `responses` section of each operation.

## Related links

**Related tasks**  

[Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")
