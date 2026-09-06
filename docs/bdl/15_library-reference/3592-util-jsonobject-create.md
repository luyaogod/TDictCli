---
title: "util.JSONObject.create"
source: "fgl-topics/c_fgl_ext_util_JSONObject_create.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.create"
type: "concept"
---

# util.JSONObject.create

> Creates a new JSON object.

## Syntax

```
util.JSONObject.create()
  RETURNS util.JSONObject
```

## Usage

The `util.JSONObject.create()` method creates a new JSON object.

The created object must be assigned to a program variable defined with the
`util.JSONObject` type.

## Example

```
IMPORT util
MAIN
    DEFINE obj util.JSONObject
    LET obj = util.JSONObject.create()
    ...
END MAIN
```

## Related links

**Related concepts**  

[util.JSONObject.fromFGL](3593-util-jsonobject-fromfgl.md "Creates a new JSON object from a RECORD.")

[util.JSONObject.parse](3594-util-jsonobject-parse.md "Parses a JSON string and creates a JSON object from it.")
