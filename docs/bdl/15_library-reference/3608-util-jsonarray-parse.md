---
title: "util.JSONArray.parse"
source: "fgl-topics/c_fgl_ext_util_JSONArray_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.parse"
type: "concept"
---

# util.JSONArray.parse

> Parses a JSON string and creates a JSON array object from it.

## Syntax

```
util.JSONArray.parse(
   s STRING )
  RETURNS util.JSONArray
```

1. s is a string value that contains JSON formatted data as
   a list of elements delimited by square brackets.

## Usage

The `util.JSONArray.parse()` method scans the JSON source string
passed as parameter and creates a new JSON array object from it.

The created object must be assigned to a program variable defined with the
`util.JSONArray` type.

The source string must follow the JSON format specification. Elements of the
list can contain multi-level structured data, but the string must follow the JSON
array string syntax '`[ element, ... ]`' with square brackets.

The method raises error [-8109](4483-genero-bdl-errors.md) if the JSON source string is not properly formatted. Consider enclosing the
`parse()` method call in a `TRY/CATCH` block, if the source string can
be malformed JSON.

## Example

```
IMPORT util
MAIN
    DEFINE da DYNAMIC ARRAY OF INTEGER
    DEFINE arr util.JSONArray
    LET arr = util.JSONArray.parse("[1,2,3,4,5]")
    DISPLAY arr.toString()
END MAIN
```

Output:

```
[1,2,3,4,5]
```

## Related links

**Related concepts**  

[util.JSONArray.fromFGL](3607-util-jsonarray-fromfgl.md "Creates a new JSON array object from a DYNAMIC ARRAY.")

[util.JSONArray.create](3606-util-jsonarray-create.md "Creates a new JSON array object.")
