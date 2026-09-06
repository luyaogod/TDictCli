---
title: "CHAR, VARCHAR, and DECIMAL conversions"
source: "fgl-topics/c_gws_high_level_rest_bdl_data_type_conversion.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > OpenAPI types mapping: BDL > CHAR, VARCHAR, and DECIMAL conversions"
type: "concept"
---

# CHAR, VARCHAR, and DECIMAL conversions

> GWS refines CHAR, VARCHAR, and DECIMAL data types in high-level RESTful web services by applying JSON Schema restriction keywords.

The OpenAPI specification defines basic data types such as string, number, and date. These types
generally map to Genero BDL primitive data types.

When you need to enforce additional constraints, for example to reflect database column
definitions, the GWS refines these data types by applying JSON Schema keywords to the generated
schema.

The following sections describe how the GWS refines `CHAR(n)`,
`VARCHAR(N)`, and `DECIMAL(P,S)` data types.

## CHAR(n)

OpenAPI represents fixed-length character data as a string. To define a `CHAR(n)`
with a maximum length of n, GWS applies the `maxLength`
keyword.

```
schema
type           "string"
maxLength      10
```

In this case, fglrestful generates the client stub by mapping the schema to a
Genero BDL `VARCHAR(10)` type. Pass `--ignore-restrictions` to
generate `STRING` instead.

This applies whether the field is declared directly as `CHAR(n)` or through a
[`LIKE`](../08_language-basics/0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.") clause on a `CHAR`
database column.

## VARCHAR(N)

OpenAPI represents variable-length character data as a string. To define a
`VARCHAR(N)` with a maximum length of N, GWS applies the `maxLength`
keyword.

```
schema
type           "string"
maxLength	10
```

In this case, fglrestful generates the client stub by mapping the schema to a
Genero BDL `VARCHAR(10)` type.

## DECIMAL(p,s)

OpenAPI represents [fixed point decimal](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") values as
a number. To model a Genero BDL `DECIMAL(p,s)`, GWS applies a combination of JSON
Schema restriction keywords:

```
schema
type	           "number"
multipleOf	      0.01
minimum	        -1000
exclusiveMinimum	true
maximum	         1000
exclusiveMaximum	true
```

From this schema, fglrestful derives the precision and scale as follows.

## Calculating precision (p)

Precision is determined from the effective minimum and maximum values.

- Maximum value:

  ```
  1000 - 0.01 = 999.99
  ```
- Minimum value:

  ```
  -1000 + 0.01 = -999.99
  ```

  When
  `exclusiveMinimum` and `exclusiveMaximum` are set to true, the
  boundary values themselves are excluded. In this example, -1000 and 1000 are not valid values.

## Calculating scale (s)

Scale is derived from the multipleOf value:

- `multipleOf` = 0.01 → `DECIMAL(p,2)` (two decimal places)
- `multipleOf` = 1 → `DECIMAL(p,0)` (no decimal places)

If `multipleOf` is not defined, the value is mapped to
`DECIMAL(p)`, meaning it can include or omit decimal places.

In this case, fglrestful generates the client stub with a Genero BDL
`DECIMAL(5,2)` type. The precision is 5 because the largest value (999.99) requires
five digits in total, including the fractional part.

When a number is fully unconstrained, with no `multipleOf`, `minimum`,
or `maximum` keyword at all, fglrestful maps the value to
`DECIMAL(32)` (the maximum precision), rather than to the FGL language default
`DECIMAL(16)` used for an unparameterized `DECIMAL` (without precision or
scale).

## Related links

**Related concepts**  

[Set data format with WSMedia](4741-set-data-format-with-wsmedia.md "It is important to set the correct MIME type for a Web service request or response. You can specify the data format via the WSMedia attribute.")

[Setting MIME type at runtime](4772-setting-mime-type-at-runtime.md "Override the GWS default media format for messages.")

[NULLs and empty structures](../09_advanced-features/0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.")

**Related reference**  

[OpenAPI types mapping: BDL](4847-openapi-types-mapping-bdl.md "Conversion mapping for Genero BDL data types in OpenAPI documentation.")

[How GWS maps BDL types into the OpenAPI description](4775-how-gws-maps-bdl-types-into-the-openapi-description.md "GWS maps Genero BDL types to JSON schema objects in the OpenAPI description. The mapping depends on the content type and whether the data type is named or anonymous.")
