---
title: "Example 1: Primitive and complex types with null values"
source: "fgl-topics/c_gws_high_level_rest_bdl_null_value_conversion.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > OpenAPI types mapping: BDL > Null value conversions > Primitive and complex types with null values"
type: "concept"
---

# Example 1: Primitive and complex types with null values

> This example shows how null values are handled for primitive and complex data types during serialization.

In OpenAPI 3.0, null values for both primitive and complex data types are represented using the
`nullable: true` property.

To accept null values in Genero BDL, you must configure your variables in one of the following
ways:

- Define the variable with the [`json_null="null"`](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") attribute so it accepts NULL in addition to its base type
  (for example, `STRING`)
- Enable null handling globally by setting the [allowNullAsDefault](../15_library-reference/3687-allownullasdefault.md "Allow NULL values to be accepted during deserialization when the json_null=\"null\" attribute is not explicitly specified.") option in the
  [The json.Serializer class](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")

If null values are not handled using one of these approaches, serialization [error-15807](../15_library-reference/4483-genero-bdl-errors.md) is raised.

When working with user-defined record types, a special case applies. If a record is defined with
`json_null="null"`, its fields can also contain null values. The following example
shows a tAddress record type with nullable fields:

```
TYPE tAddress RECORD
    street STRING,
    city STRING,
    state STRING,
    zip STRING
END RECORD

TYPE customer RECORD
    lastname STRING ATTRIBUTES(json_null="null"),
    firstname STRING,
    address tAddress ATTRIBUTE(json_null="null")
END RECORD
```

In this example, fglrestful maps the Genero BDL types to the JSON object
schema as follows:

- The `nullable: true` property is set for fields that accept null values.
- The `anyOf` property is used to allow fields to accept either the expected type
  or NULL.
- The `tAddress` type is defined as a reference (`$ref`) in the
  `components/schemas` section, where its schema is declared.

When a property uses a `$ref`, OpenAPI requires it to be wrapped in an
`anyOf` clause if it appears alongside the nullable property. The following example
shows how this is represented in the generated OpenAPI document:

```
 "content": {
	"application/json": {
	    "schema": {
		  "type": "object",
		   "properties": {
			 "lastname": {
			     "type": "string",
			     "nullable": true
			 },
			 "firstname": {
			     "type": "string"
			 },
			 "address": {
			     "nullable": true,
			     "anyOf": [
			        {
			         "$ref": "#/components/schemas/tAddress"
			        }
			     ]
		         },
```

For more details, refer to the [OpenAPI](http://swagger.io/docs/specification/about/)
documentation.
