---
title: "Implicit and explicit JSON conversion for primitive types"
source: "fgl-topics/c_gws_JSONSerializer_primitive_types.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > Implicit and explicit conversion > Primitive type conversion"
type: "concept"
---

# Implicit and explicit JSON conversion for primitive types

> This topic describes how primitive 4GL values are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.

For detailed comparisons and examples, refer to the tables in the sections dedicated to each data
type.

## BOOLEAN

| 4GL value | JSON output | json.Serializer | util.json |
| --- | --- | --- | --- |
| `TRUE` | `true` | OK | OK |
| `FALSE` | `false` | OK | OK |
| `"true"` (`STRING` to `FGL-Boolean` assignment) | `null` | ERROR (unless authorized by [serializeNullAsDefault](3690-serializenullasdefault.md "Allow NULL values during serialization even when constraints are set, such as when JSONRequired is defined or when json_null=\"null\" is not defined.")) | OK |

| JSON input | Result | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `true` | `1` | OK | OK | Direct mapping |
| `false` | `0` | OK | OK | Direct mapping |
| `"true"` | `1` | ERROR | OK | Implicit `JSON-String` to `FGL-Boolean` 1 |
| `"false"` | `0` | ERROR | OK | Implicit cast `JSON-String` to `FGL-Boolean` 1 |
| `1` | `1` | ERROR | OK | Implicit cast: `JSON-Integer` to `FGL-Boolean` 1 |
| `0` | `0` | ERROR | OK | Implicit cast 1 |
| `"1"` | `1` | ERROR | OK | Implicit `JSON-String` to `FGL-Boolean` 1 |
| `"0"` | `0` | ERROR | OK | Implicit cast `JSON-String` to `FGL-Boolean` 1 |
| `"toto"` | - | ERROR | OKsilent cast to `NULL` | `util.json` casts to `NULL` without raising an error. |
| `"toto"` and with `allowImplicitConversion = 1` | - | ERROR | OKsilent cast to `NULL` | [allowImplicitConversion](3686-allowimplicitconversion.md "Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another.") does not help. |
| `{}` object | - | ERROR | ERROR |  |
| `[]` array | - | ERROR | ERROR |  |

For behavior when using the permissive `util.JSON` parser, refer to
[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.").

**1.** Implicit conversion in [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")
was removed in version 6.00.00.

## FGL-Integer / FGL-Number

The term `FGL-Integer` refers to the 4GL types: `TINYINT`,
`SMALLINT`, `INTEGER`, and `BIGINT`. A JSON number with
a decimal part of zero (for example, `12.0000`) is **considered valid for
FGL-INTEGER** if no fractional digits remain after parsing.

The term `FGL-Number` in this document refers to all 4GL types:
`DECIMAL`, `MONEY`, `FLOAT`, and
`SMALLFLOAT`. When serialized to JSON, these are represented as JSON
`number`.

| 4GL value | JSON output | json.Serializer | util.json |
| --- | --- | --- | --- |
| `123` | `123` | OK | OK |
| `123.45` | `123.45` | OK | OK |

| JSON input | Target type | Result | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- | --- |
| `12` | `FGL-Integer` | 12 | OK | OK |  |
| `12.000` | `FGL-Integer` | 12 | OK | OK |  |
| `12.35` | `FGL-Integer` | 12 | ERROR | OK | Truncation, not rounding 2 |
| `"123"` | `FGL-Integer` | 123 | ERROR | OK | Implicit `JSON-String` to `FGL-Integer` 2 |
| `"123.5"` | `FGL-Integer` | 123 | ERROR | OK | Implicit `JSON-String` to `FGL-Integer` 2 |
| `12` | `FGL-Number` | 12 | OK | OK |  |
| `12.35` | `FGL-Number` | 12.35 | OK | OK |  |
| `"12.35"` | `FGL-Number` | 12.35 | ERROR | OK | Implicit `JSON-String` to `FGL-Integer` 2 |
| `"toto"` | `FGL-Integer``FGL-Number` | - | ERROR | OKsilent cast to `NULL` | Invalid conversion |
| `true/false` | `FGL-Integer``FGL-Number` | 1 / 01.0 / 0.0 | ERROR | OKsilent cast to NULL | Type mismatch `json.Serializer:` [allowImplicitConversion](3686-allowimplicitconversion.md "Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another.") enables JSON-Boolean =>JSON-Integer/Number |
| `{}` | `FGL-Integer``FGL-Number` | - | ERROR | ERROR |  |
| `[]` | `FGL-Integer``FGL-Number` | - | ERROR | ERROR |  |

For permissive parsing rules and implicit casting behavior provided by
`util.JSON`, refer to [JSON to BDL
type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.").

**2.** Implicit conversion in [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")
was removed in version 6.00.00.

## FGL-String

The term **FGL-String** refers to 4GL string types: `STRING`,
`CHAR`, `VARCHAR`, `TEXT`, `BYTE`. When
serialized to JSON, these are represented as JSON `string` values.

| 4GL value | JSON output | json.Serializer | util.json |
| --- | --- | --- | --- |
| `"abc"` | `"abc"` | OK | OK |
| `""` (empty) | `""` | OK | OK |

| JSON input | Result | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `""` | `""` | OK | OK | Empty string preserved |
| `"hello"` | `"hello"` | OK | OK |  |
| `123` | `"123"` | ERROR | OK | Implicit cast JSON-Integer to FGL-String 3 |
| `true/false` | `"true"/"false"` | ERROR | OK | Implicit cast JSON-Boolean to FGL-String 3 |
| `{}` | - | ERROR | ERROR |  |
| `[]` | - | ERROR | ERROR |  |

For permissive JSON parsing behavior provided by `util.JSON`, refer to
[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.").

**3.** Implicit conversion in [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.")
was removed in version 6.00.00.

## DATE

| 4GL value | JSON output | json.Serializer | util.json | Format |
| --- | --- | --- | --- | --- |
| `TODAY` | `"2025-12-11"` | OK | OK | OK – ISO-8601 like (`YYYY-MM-DD`) |

| JSON input | Result | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `""` | NULL | OK | OK |  |
| `"2025-05-16"` | `"05/16/2025"` | OK | OK | Valid ISO-like format |
| `"16/05/2025"` | - | ERROR | OK silent cast to `NULL` | Not supported |
| `"05/16/2025"` | - | ERROR | OK silent cast to `NULL` | US format not accepted |
| `"not-a-date"` | - | ERROR | OK silent cast to `NULL` |  |
| `{}` | - | ERROR | ERROR |  |
| `[]` | - | ERROR | ERROR |  |

For JSON parsing behavior using the `util.JSON` class, refer to
[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.").

## DATETIME

| 4GL value | JSON output | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `CURRENT YEAR TO SECOND` | `"2025-12-11 14:23:01"` | OK | OK | Not ISO-8601 compliant yet. 4 |

**4.** We do a regex check for the format: `\d{4}-\d{2}-\d{2}\d{2}:\d{2}:\d{2}`
> **Note:**
>
> ISO 8601 format with `'T'` separator is not yet supported in
> `json.Serializer` but will be in the future.

| JSON input | Result | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `"2025-05-16 14:23:01"` |  | OK | OK |  |
| `"2025-05-16T14:23:01"` | `2025-05-16 14:23:01` | OK | OK | Convert to FGL format |
| `"not-a-date"` | - | ERROR | OKsilent cast to `NULL` |  |
| `{}` | - | ERROR | ERROR |  |
| `[]` | - | ERROR | ERROR |  |

For JSON parsing behavior using the `util.JSON` class, refer to
[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.").

## Related links

**Related concepts**  

[Implicit and explicit JSON conversion for complex types](3665-complex-type-conversion.md "This topic describes how complex 4GL structures such as RECORD, DICTIONARY, arrays, and util.JSONObject are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.")

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")
