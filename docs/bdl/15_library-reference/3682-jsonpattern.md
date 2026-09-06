---
title: "JSONPattern"
source: "fgl-topics/c_gws_JSON_attribute_JSONPattern.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer attributes > JSONPattern"
type: "concept"
---

# JSONPattern

> Specify a regular expression pattern for string values in a JSON schema.

## Syntax

```
JSONPattern = "pattern"
```

`JSONPattern` is an optional attribute that specifies a regular expression pattern for string values, following the JSON Schema `pattern` keyword.

> **Important:**
>
> This serializer-specific attribute is supported only by [json.Serializer](3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") and does not work with [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.").

## Usage

The `JSONPattern` attribute allows you to define a regular expression on
`json-string` types that string values must match to be considered valid. It is
designed to be used on fields (record members) or types, not on functions.

The `pattern` keyword in JSON Schema (and the `JSONPattern`
attribute in BDL) expects a regular expression that adheres to the regex syntax defined by the
JavaScript language standard [ECMA-262](https://262.ecma-international.org/5.1/#sec-15.10.1) (external link):

- You can utilize many of the same regex features found in JavaScript, including character classes
  (`[0-9]`), quantifiers (`*`, `+`,
  `{n,m}`), anchors (`^`, `$`), and groups
  (`(...)`).
- However, not all advanced features, like lookbehind assertions, are guaranteed to be supported.
  Some JSON Schema validators or the BDL compiler may only implement a subset of these features for
  compatibility and portability. It is advisable to adhere to the subset defined in the [JSON
  Schema.](https://262.ecma-international.org/5.1/#sec-15.10.1) (external link).

The regex pattern is validated at compilation time. During runtime, serialization and
deserialization will check the value against the pattern.

**Practical implications for JSONPattern:**

- You should use standard JavaScript regex syntax.
- It is recommended to escape backslashes in JSON and YAML strings, as required by those formats.
  In BDL, enclose literal strings (or other text containing special
  characters) in single quotes or back ticks to avoid escaping with backslashes and improve
  readability. Example: `` `your "string" with \special\ chars` ``.
- Stick to common regex constructs for maximum compatibility (avoid advanced or non-standard
  features).

## YAML Example

YAML schema using `pattern`:

This pattern matches a US Social Security Number (SSN) in the format **123-45-6789**:

- `^` anchor matches the start of the string.
- `\\d{3}` matches exactly three digits.
- `-` matches a literal dash.
- `\\d{2}` matches exactly two digits.
- `\\d{4}` matches exactly four digits.
- `$` anchor matches the end of the string.

```
ssn:
type: string
pattern: '^\\d{3}-\\d{2}-\\d{4}$'
```

## JSON Example

JSON schema using `pattern` (note the escaped backslashes):

```
{
"ssn": {
  "type": "string",
  "pattern": "^\\d{3}-\\d{2}-\\d{4}$"
}
}
```

## BDL example type using `JSONPattern`

Escaped backslashes:

```
TYPE security RECORD
  ssn STRING ATTRIBUTE (JSONPattern="^\\d{3}-\\d{2}-\\d{4}$")
END RECORD
```

Using back ticks:

```
TYPE ssnSchemaRec RECORD
    ssn STRING ATTRIBUTE(JSONPattern =`^\d{3}-\d{2}-\d{4}$`)
END RECORD
```

## JSON Data Example

Valid JSON data for the above pattern:

```
{"ssn":"443-58-8799"}
```

## Compilation and validation rules

1. **String type**: Only `json-string` types are supported.
2. **Accepted BDL types**: The following types are accepted: `STRING`,
   `CHAR`, `VARCHAR`, `DATE`, `DATETIME`,
   `INTERVAL`, `TEXT`. Using `JSONPattern` with
   unsupported types or on functions will result in [error-9156](4483-genero-bdl-errors.md).
3. **Value requirement**: The value must be present and cannot be null.
4. **Validation**:

   - The regular expression is validated at compilation. An invalid regular expression will trigger
     [error-9157](4483-genero-bdl-errors.md).
   - Serialization and deserialization will check the value at runtime.

## GWS engine and OpenAPI

- **Serialization**: The GWS engine validates the value against the regular expression before
  sending JSON.
- **Deserialization**: The engine checks the received value to ensure it meets the expected
  criteria.
- **OpenAPI documentation**: The "`pattern`" keyword is displayed when declared
  in Genero BDL code.
- **JSONPattern attribute**: The `fglrestful` generates the
  `JSONPattern` attribute on the type, along with the corresponding regular
  expression.

## Related links

**Related concepts**  

[Using RESTful attributes in functions](../16_web-services/4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

[Regular expression patterns](3545-regular-expression-patterns.md "Summary of util.Regexp.compile() pattern syntax.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
