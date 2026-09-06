---
title: "util.JSONArray.create"
source: "fgl-topics/c_fgl_ext_util_JSONArray_create.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.create"
type: "concept"
---

# util.JSONArray.create

> Creates a new JSON array object.

## Syntax

```
util.JSONArray.create()
  RETURNS util.JSONArray
```

## Usage

The `util.JSONArray.create()` method create a new JSON array object.

The created object must be assigned to a program variable defined with the
`util.JSONArray` type.

## Example

```
IMPORT util
MAIN
    DEFINE arr util.JSONArray
    LET arr = util.JSONArray.create()
    ...
END MAIN
```

## Related links

**Related concepts**  

[util.JSONArray.fromFGL](3607-util-jsonarray-fromfgl.md "Creates a new JSON array object from a DYNAMIC ARRAY.")

[util.JSONArray.parse](3608-util-jsonarray-parse.md "Parses a JSON string and creates a JSON array object from it.")
