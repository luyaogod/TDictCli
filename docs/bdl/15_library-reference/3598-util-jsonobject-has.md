---
title: "util.JSONObject.has"
source: "fgl-topics/c_fgl_ext_util_JSONObject_has.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.has"
type: "concept"
---

# util.JSONObject.has

> Checks if the JSON object contains a specific entry name.

## Syntax

```
has(
    name STRING )
  RETURNS BOOLEAN
```

1. name is a string identifying a JSON object property.

## Usage

The `has()` method determines if the JSON object holds a property identified by
the element name passed as parameter.

The method returns `TRUE` if the name/value pair exists in the JSON object.

A name/value pair can be set with the `put()` method.

## Related links

**Related concepts**  

[util.JSONObject.put](3600-util-jsonobject-put.md "Sets a name-value pair in the JSON object.")
