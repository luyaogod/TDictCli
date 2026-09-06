---
title: "OpenAPI types mapping: BDL"
source: "fgl-topics/r_gws_high_level_rest_bdl_json_type_mapping.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > OpenAPI types mapping: BDL"
type: "reference"
---

# OpenAPI types mapping: BDL

> Conversion mapping for Genero BDL data types in OpenAPI documentation.

| Genero BDL data types | JSON data types | Format | Additional format keywords |
| --- | --- | --- | --- |
| STRING | STRING |  |  |
| VARCHAR(n) | STRING |  | `maxLength: n`For an example, go to [VARCHAR(N)](4848-char-varchar-and-decimal-conversions.md) |
| CHAR(n) | STRING |  | `maxLength: n`For an example, go to [CHAR(n)](4848-char-varchar-and-decimal-conversions.md) |
| DATE | STRING | `format: date` | For examples of date conversion, go to [DATE](4852-date-datetime-and-interval-conversions.md) |
| DATETIME YEAR TO DAY | STRING | `format: date` | For examples of datetime conversion, go to [DATETIME](4852-date-datetime-and-interval-conversions.md) |
| DATETIME YEAR TO MINUTE, DATETIME YEAR TO SECOND, DATETIME YEAR TO FRACTION(n) | STRING | `format: date-time` | For examples of datetime conversion, go to [DATETIME](4852-date-datetime-and-interval-conversions.md) |
| DATETIME YEAR TO YEAR, DATETIME YEAR TO MONTH, DATETIME YEAR TO HOUR, or any qualifier that does not start at YEAR (such as HOUR TO MINUTE or MONTH TO DAY) | STRING |  | For examples of datetime conversion, go to [DATETIME](4852-date-datetime-and-interval-conversions.md) |
| INTERVAL (DAY TO SECOND and YEAR TO MONTH classes) | STRING | `format: duration` |  |
| TEXT | STRING |  |  |
| BYTE | STRING | `format: binary` |  |
| BOOLEAN | BOOLEAN |  |  |
| INTEGER | INTEGER | `format: int32` |  |
| TINYINT | INTEGER |  |  |
| SMALLINT | INTEGER |  |  |
| BIGINT | INTEGER | `format: int64` |  |
| DECIMAL | NUMBER |  | For examples of decimal conversion, go to [DECIMAL(p,s)](4848-char-varchar-and-decimal-conversions.md) |
| SMALLFLOAT | NUMBER | `format: float` |  |
| FLOAT | NUMBER | `format: double` |  |
| MONEY | NUMBER | `format: double` | `multipleOf: 0.01``minimum:-100000000000000``exclusiveMinimum: true``maximum: 100000000000000``exclusiveMaximum: true` |
| RECORD | OBJECT |  |  |
| DYNAMIC ARRAY OF | ARRAY |  | For examples handling nulls in arrays, go to [Example 2: Handling arrays with null values](4851-arrays-with-null-values.md "This example shows how null values in arrays are handled during serialization.") |
| ARRAY [ ] OF | N/A |  | N/A |
| DICTIONARY OF | OBJECT |  | `additionalProperties:"type":"object"` |
| [util.JSONObject](../15_library-reference/3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.") | OBJECT |  | `additionalProperties: {}` |
| [util.JSONArray](../15_library-reference/3604-the-util-jsonarray-class.md "The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.") | OBJECT |  | `additionalProperties: {}` |
| [`json_null="null"`](../15_library-reference/3668-json-null-attribute.md "This attribute controls the representation of null or empty values. json.Serializer uses it during both serialization and deserialization, but util.JSON applies it only when serializing.") attribute set to allow null values to deserialize JSON null values ​​into BDL variables, or to enable serialization of null values from BDL to JSON. | All primitive and complex data types | `nullable: true` | For examples, go to [Example 1: Primitive and complex types with null values](4850-primitive-and-complex-types-with-null-values.md "This example shows how null values are handled for primitive and complex data types during serialization.") |

## Converting OpenAPI types to BDL in client stubs

When fglrestful generates a client stub, it uses the OpenAPI format keywords
in the Additional format keywords column above to convert OpenAPI property
types to precise BDL types. For example, a JSON `STRING` property with
`maxLength: 50` becomes `VARCHAR(50)` in the stub, and a
`NUMBER` property with `multipleOf`, `minimum`, and
`maximum` keywords becomes `DECIMAL(P,S)`.

Use the `--ignore-restrictions` option with fglrestful to
bypass this conversion: string properties fall back to `STRING` and numeric
properties to unparameterized `DECIMAL` (maximum-precision). For details, see [--ignore-restrictions](../13_programming-tools/2524-fglrestful.md).

## Child topics

- [CHAR, VARCHAR, and DECIMAL conversions](4848-char-varchar-and-decimal-conversions.md): GWS refines CHAR, VARCHAR, and DECIMAL data types in high-level RESTful web services by applying JSON Schema restriction keywords.
- [Null value conversions](4849-null-value-conversions.md): These examples show how null values are handled during data type conversion.
- [DATE, DATETIME, and INTERVAL conversions](4852-date-datetime-and-interval-conversions.md): GWS serializes DATE, DATETIME, and INTERVAL values as JSON strings using either the Genero BDL native format or standard RFC 3339 / ISO 8601 formats.
