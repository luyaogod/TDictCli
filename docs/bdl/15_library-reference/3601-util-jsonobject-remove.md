---
title: "util.JSONObject.remove"
source: "fgl-topics/c_fgl_ext_util_JSONObject_remove.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods > util.JSONObject.remove"
type: "concept"
---

# util.JSONObject.remove

> Removes the specified element in the JSON object.

## Syntax

```
remove(
   name STRING )
```

1. name is the string identifying the JSON object
   property.

## Usage

The `remove()` method deletes a name-value pair identified by the name passed
as parameter.

If the property is referenced by another `JSONObject` or
`JSONArray`, the data is still available from the other object.

## Example

```
IMPORT util
MAIN
    DEFINE obj, subobj util.JSONObject
    LET subobj = util.JSONObject.create()
    CALL subobj.put("num", 5)
    CALL subobj.put("street", "Sunset Blvd")
    LET obj = util.JSONObject.create()
    CALL obj.put("address", subobj)
    DISPLAY '\nInitial content:'
    DISPLAY "obj    1:", obj.toString()
    DISPLAY "subobj 1:", subobj.toString()
    CALL subobj.remove("street")
    DISPLAY '\nAfter subobj.remove("street"):'
    DISPLAY "obj    2:", obj.toString()
    DISPLAY "subobj 2:", subobj.toString()
    DISPLAY '\nAfter obj.remove("address"):'
    CALL obj.remove("address")
    DISPLAY "obj    3:", obj.toString()
    DISPLAY "subobj 3:", subobj.toString()
END MAIN
```

Output:

```
Initial content:
obj    1:{"address":{"num":5,"street":"Sunset Blvd"}}
subobj 1:{"num":5,"street":"Sunset Blvd"}

After subobj.remove("street"):
obj    2:{"address":{"num":5}}   <-- affects obj because "address" references subobj data
subobj 2:{"num":5}

After obj.remove("address"):
obj    3:{}
subobj 3:{"num":5}               <-- subobj data still referenced and available
```

## Related links

**Related concepts**  

[util.JSONObject.put](3600-util-jsonobject-put.md "Sets a name-value pair in the JSON object.")

[util.JSONObject.get](3595-util-jsonobject-get.md "Returns the value corresponding to the specified entry name.")
