---
title: "Implicit and explicit JSON conversion for complex types"
source: "fgl-topics/c_gws_JSONSerializer_complex_types_2.html"
breadcrumb: "Web services > Reference > JSON serialization rules and customization > Implicit and explicit conversion > Complex type conversion"
type: "concept"
---

# Implicit and explicit JSON conversion for complex types

> This topic describes how complex 4GL structures such as RECORD, DICTIONARY, arrays, and util.JSONObject are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.

> **Important:**
>
> **Type checking array/dictionary types**
>
> When using
> `json.Serializer`, **type checking is strictly enforced** only for the following structures:
>
> - `STATIC ARRAY`
> - `DYNAMIC ARRAY`
> - `DICTIONARY`

JSON objects in Genero BDL can be generated in three ways:

1. `RECORD`: a statically typed structure; each member is typed and must conform to
   [primitive type conversion rules](../15_library-reference/3664-primitive-type-conversion.md "This topic describes how primitive 4GL values are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.").
2. `DICTIONARY`: a dynamic object with string keys; each value must conform to
   conversion rules similar to array elements. (For details, go to [Table 2](../15_library-reference/3665-complex-type-conversion.md))
3. [`util.JSONObject`](../15_library-reference/3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax."): a dynamic,
   open object; no strict type enforcement is applied during serialization or deserialization.

In the tables below, the term FGL-Object refers to 4GL structured types such as
`RECORD` and `DICTIONARY`. These types are schema-aware and their
members are subject to strict validation when using `json.Serializer`.
FGL-Object does not include `util.JSONObject`, which behaves as an open
container with no enforced type structure.

The following tables compare these different data types along with their structures, typing,
conversion rules, and best use cases. Refer to [Detailed
Compatibility Matrix for JSON Arrays](../15_library-reference/3665-complex-type-conversion.md) for a detailed comparisons between
`util.json` and `json.Serializer`.

| Type | Structure | Typing | Conversion rules | Best use case |
| --- | --- | --- | --- | --- |
| `RECORD` | Fixed fields | Strong typing | Strict (per field) | Known schema, validated structures |
| `DICTIONARY` | Dynamic fields | Dynamic, typed values | Like `ARRAY` (see [Table 2](../15_library-reference/3665-complex-type-conversion.md)) | JSON with dynamic or unknown keys |
| `util.JSONObject` | Fully dynamic | No enforced typing | None | Raw JSON, flexible usage |

JSON arrays can be generated in Genero BDL in two ways:

1. **FGL-Array** (`DYNAMIC` or `STATIC ARRAY`):

   This is a typed
   array structure where each element must conform to strict conversion rules based on the array's
   declared base type (for example, `DYNAMIC ARRAY OF INTEGER`, `DYNAMIC ARRAY OF
   RECORD`). Type mismatches lead to deserialization errors.

   The term
   FGL-Array refers to both `DYNAMIC ARRAY OF
   element-type` and `ARRAY[n] OF
   element-type`.
2. **util.JSONArray**:

   [`util.JSONArray`](../15_library-reference/3604-the-util-jsonarray-class.md "The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.") is a dynamic, open array. No strict type enforcement is
   applied during serialization or deserialization; elements of any type (primitive or complex) can
   coexist.

On the other hand, the elements of a dynamic array must comply with the conversion rules
described in [Table 2](../15_library-reference/3665-complex-type-conversion.md).

For conversion behavior when using the `util.JSON` class, including arrays,
objects, and dictionaries, refer to [JSON to BDL
type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.") and [BDL to JSON
type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

| Target type | JSON token type accepted (default) | With `allowImplicitConversion = true` |
| --- | --- | --- |
| `BOOLEAN` | `true`, `false`, `0`, `1` | `"true"`, `"false"`, `"0"`, `"1"` |
| `FGL-INTEGER` | integer | string representing integer |
| `FGL-NUMBER` | number | string representing integer/float |
| `FGL-STRING` | string | number, boolean |
| `DATE` | string (ISO 8601) | - |
| `DATETIME` | string | - |
| `FGL-ARRAY` `util.JSONArray` | array | - |
| `FGL-OBJECT` `util.JSONObject` | object | - |
| `with attribute json_null attribute` | `null` | - |

| Target type | Accepted JSON inputs | json.Serializer | util.json | Notes |
| --- | --- | --- | --- | --- |
| `BOOLEAN` |  |  |  |  |
|  | `true`, `false` | OK | OK | Direct mapping |
|  | `1`, `0`, `"true"`, `"false"`, `"1"`, `"0"` | ERROR (unless [allowImplicitConversion](../15_library-reference/3686-allowimplicitconversion.md "Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another.") is used) | OK | Implicit cast: JSON-String/Number to `FGL-BOOLEAN` |
|  | `"toto"` | ERROR | OKsilent cast to `NULL` |  |
| `FGL-INTEGER` |  |  |  |  |
|  | `123` | OK | OK | Direct mapping |
|  | `"123"` | ERROR (unless [allowImplicitConversion](../15_library-reference/3686-allowimplicitconversion.md "Allow implicit type conversions during JSON deserialization when the JSON property specifies one type, but the values suggest another.") is used) | OK | `JSON-String` to `FGL-INTEGER` |
|  | `"abc"` | ERROR | OKsilent cast to `NULL` |  |
| `FGL-NUMBER` | `123.45`, `"123.45"` | OK | OK | Same behavior as `FGL-INTEGER` |
| `FGL-STRING` | `"foo"` | OK | OK |  |
| `DATE` |  |  |  |  |
|  | `"YYYY-MM-DD"`, timestamp | OK | OK | JSON-String must match format or be numeric |
|  | `"not-a-date"` | ERROR | OKsilent cast to `NULL` |  |
| `FGL-Object``util.JSONObject` | `{...}` | OK | OK |  |
| `FGL-ARRAY``util.JSONArray` |  |  |  |  |
|  | `[valid_elements]` | OK | OK | Recursively deserialized |
|  | `[]` | OK | OK | Accepted even if empty |
|  | `[invalid_element]` | ERROR | OK | `util.JSON`: Invalid when assigning a primitive to a complex type, or vice versa. |
| `with attribute json_null attribute` | `null` | ERROR | OK | `json.Serializer:` `NULL` handling configurable |

For conversion behavior when using the `util.JSON` class, including arrays,
objects, and dictionaries, refer to
[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.")
and
[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

## Related links

**Related concepts**  

[Implicit and explicit JSON conversion for primitive types](../15_library-reference/3664-primitive-type-conversion.md "This topic describes how primitive 4GL values are converted during JSON serialization and deserialization, and compares the behavior of json.Serializer and util.JSON.")

[Best practices](../15_library-reference/3669-best-practices.md "Recommended practices for using json.Serializer to enhance performance and avoid common pitfalls.")

[Null and empty value handling in json.Serializer and util.JSON](../15_library-reference/3666-null-and-empty-value-handling-in-json-serializer-and-util-js.md "This topic compares how json.Serializer and util.JSON handle null and empty values during JSON serialization and deserialization.")

[JSON to BDL type conversion rules](../09_advanced-features/0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")

**Related reference**  

[json.Serializer options](../15_library-reference/3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
