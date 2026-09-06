---
title: "util.JSONArray.toString"
source: "fgl-topics/c_fgl_ext_util_JSONArray_toString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.toString"
type: "concept"
---

# util.JSONArray.toString

> Builds a JSON string from the elements contained in the JSON array object.

## Syntax

```
toString()
  RETURNS STRING
```

## Usage

The `toString()` method
produces a JSON formatted string from the elements
contained in the JSON array object.

## Example

```
IMPORT util
MAIN
    DEFINE ja util.JSONArray
    LET ja = util.JSONArray.create()
    CALL ja.put(1,"aa")
    CALL ja.put(2,"bb")
    CALL ja.put(3,"cc")
    DISPLAY ja.toString()
END MAIN
```

Output:

```
["aa","bb","cc"]
```

## Related links

**Related concepts**  

[util.JSONArray.toFGL](3614-util-jsonarray-tofgl.md "Fills a dynamic array variable with the elements contained in the JSON array object.")
