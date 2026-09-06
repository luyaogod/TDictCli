---
title: "JSONEnum"
source: "fgl-topics/c_gws_JSON_attribute_JSONEnum.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer attributes > JSONEnum"
type: "concept"
---

# JSONEnum

> Defines an explicit, typed list of acceptable values for a field and maps to the enum keyword in JSON Schema.

## Syntax

```
 JSONEnum  = '{ "enum-value" } [,...]'
```

Where:

- `JSONEnum` is a set of acceptable JSON literal values for the field (string
  (`"string"`), number (`42`, `3.14`), boolean
  (`true` or `false`), or `null`).
- `"enum-value"` is one or more values; values are separated by commas.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

The `JSONEnum` attribute supports the `enum` keyword in the JSON
schema property of the Swagger and OpenAPI specification.

It allows the definition of a typed list of acceptable values for a given field. Only values
strictly matching one of the entries in the list will be accepted at runtime.

All comparisons are **strictly typed and case-sensitive** (for example, `"42"`
is not equal to `42`). Avoid mixing JSON types in `JSONEnum`. While
you can list values of different types (for example, `"string"`, `42`,
`true`), runtime validation enforces the declared field type. If the field is
`STRING`, only string tokens will pass validation. Integer or Boolean tokens will
fail, even if they appear in the `JSONEnum`
list.

```
DEFINE mixedRec RECORD
    mixed STRING ATTRIBUTE(JSONEnum = "string", 42, true)
END RECORD
```

```
# JSON input:
{ "mixed": 42 }
```

```
# Result:
Runtime error – value 42 does not match declared type STRING.
```

## BDL Example

Use single quotes or back ticks for the whole attribute string to avoid having to escape double
quotes. Double quotes are fine for numeric enumerated values. Examples:

- Strings (outer single quotes, inner double quotes):

  ```
  TYPE response RECORD
    status STRING ATTRIBUTE (JSONEnum='"pending","approved","rejected"')
  END RECORD
  ```

  Or strings with backslashes (outer double quotes and escaped
  backslashes):

  ```
  TYPE response RECORD
    status STRING ATTRIBUTE (JSONEnum="\"pending\",\"approved\",\"rejected\"")
  END RECORD
  ```
- Numbers (double quotes for the whole attribute are fine):

  ```
  # Or with numeric types:
  DEFINE rating SMALLINT ATTRIBUTE (JSONEnum="1,2,3,4,5")
  ```

## YAML Schema

Defining a record type using `enum` :

```
status:
  type: string
  enum: [pending, approved, rejected]
```

## JSON Schema

Defining a record type using `enum` :

```
{
  "status": {
    "type": "string",
    "enum": ["pending", "approved", "rejected"]
  }
}
```

**JSON data:**

```
{"status": "approved"}
{"rating": 4}
```

## Compilation and validation rules

- **JSON types:** Accepted JSON value types are limited to JSON primitives only — for example:
  `"string"`, `42`, `3.14`, `true`,
  `false`, and `null`.
- **BDL types:** Accepted BDL types are
  - `STRING`, `CHAR`, `VARCHAR`, `TEXT`
  - `TINYINT`, `SMALLINT`, `INTEGER`,
    `BIGINT`
  - `DECIMAL`, `FLOAT`, `SMALLFLOAT`,
    `MONEY`
  - `DATE`, `DATETIME`, `INTERVAL`, and
    `BOOLEAN`.
- **Web Service:** The Web Service function attribute is **not allowed**.
- **String format:** Strings can be escaped using backslashes (`\"value\"`) if
  necessary, but prefer using outer single quotes or back ticks with inner double quotes for each
  enumerated value to avoid escaping.
- **Numeric equality:** For integer and numeric comparisons, values that are numerically equal
  but have different formatting are considered equal (for example, `3.00 = 3` and
  `12.4500 = 12.45` ).

## GWS engine

**Serialization**

- If the BDL value is not in the declared `JSONEnum`, serialization fails.
- If the value is `NULL`, it is only serialized as null in JSON if `json_null="null"` is declared.

**Deserialization**

- Incoming values must match one of the enumeration entries.
- `NULL` is accepted only if `json_null="null"` is set and the
  `null` token is declared in `JSONEnum`.

## OpenAPI documentation

The OpenAPI schema should display the `enum` keyword and the corresponding allowed
values when declared via BDL using `JSONEnum`.

**Handling null in OpenAPI enums:** If `json_null="null"` is set, and the `null` token is included in the
enumerated values, the OpenAPI generates the `enum` as
follows:

```
type: string
enum: ["yes", "no", null] 
nullable: true
```

Enabling `json_null="null"` and adding `null` to the enumerated
list is required to handle the null option.

## Summary

Below is a summary of how the `JSONEnum` attribute behaves across key aspects:

| Aspect | Behavior |
| --- | --- |
| Format | JSON literals separated by comma |
| Accepted types | Only primitive JSON types |
| Comparison | Strict (type and case sensitive) |
| `NULL` handling | Only allowed with `json_null="null"` |
| Mixed types | Not recommended |
| Duplicate values | Not verified |
| OpenAPI generation | `enum` generated with `nullable: true` if applicable |
| Runtime validation | Enforced during serialization/deserialization |
| Object/array values | Not supported |

## Related links

**Related concepts**  

[Using RESTful attributes in functions](../16_web-services/4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
