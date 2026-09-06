---
title: "util.JSONObject.toString"
source: "fgl-topics/c_fgl_ext_util_JSONObject_toString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.toString"
type: "concept"
---

# util.JSONObject.toString

> Builds a JSON string from the values contained in the JSON object.

## Syntax

```
toString()
  RETURNS STRING
```

## Usage

The `toString()` method produces a JSON
formatted string from the name-value pairs contained in the JSON object.

## Example

```
IMPORT util
MAIN
    DEFINE obj util.JSONObject
    LET obj = util.JSONObject.create()
    CALL obj.put("num", "75263")
    CALL obj.put("name", "Ferguson")
    CALL obj.put("address", "12 Marylon Street")
    DISPLAY obj.toString()
END MAIN
```

Output:

```
{"num":"75263","name":"Ferguson","address":"12 Marylon Street"}
```

## Related links

**Related concepts**  

[util.JSONObject.toFGL](3602-util-jsonobject-tofgl.md "Fills a record variable with the entries contained in the JSON object.")
