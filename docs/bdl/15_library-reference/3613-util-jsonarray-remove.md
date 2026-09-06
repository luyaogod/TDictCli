---
title: "util.JSONArray.remove"
source: "fgl-topics/c_fgl_ext_util_JSONArray_remove.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods > util.JSONArray.remove"
type: "concept"
---

# util.JSONArray.remove

> Removes the specified entry in the JSON array object.

## Syntax

```
remove(
   index INTEGER )
```

1. index is the index of the element in the JSON array object.

## Usage

The `remove()` method deletes an element in the JSON array object at the given position.

The index corresponding to the first element is 1.

If no element exists at the specified index, the method returns silently.

If the element is referenced by another `JSONObject` or
`JSONArray`, the data is still available from the other object.

## Example

```
IMPORT util
MAIN
    DEFINE arr util.JSONArray
    LET arr = util.JSONArray.parse('["aa","bb","cc"]')
    CALL arr.remove(2)
    DISPLAY arr.get(2)
END MAIN
```

Output:

```
cc
```

## Related links

**Related concepts**  

[util.JSONArray.put](3612-util-jsonarray-put.md "Sets an element by position in the JSON array object.")
