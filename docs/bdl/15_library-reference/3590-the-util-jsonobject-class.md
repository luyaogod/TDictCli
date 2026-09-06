---
title: "The util.JSONObject class"
source: "fgl-topics/c_fgl_ext_util_JSONObject.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class"
type: "concept"
---

# The util.JSONObject class

> The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.

The `util.JSONObject` class is provided in the `util`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`util.JSONObject` class, you must import the `util` package in your
program:

```
IMPORT util
```

A `JSONObject` is an unordered collection of name/value pairs. The format of a
JSON object string is a comma-separated "name":value pairs, wrapped in curly brackets. The value can
be simple numeric or string value, but it can also be an array of values enclosed in square brackets,
or a sub-element enclosed in curly
brackets:

```
{ 
  "cust_num":2735,
  "cust_name":"McCarlson",
  "order_ids":[234,3456,24656,34561],
  "address": {
          "street":"34, Sunset Bld",
          "city":"Los Angeles",
          "state":"CA"
        }
}
```

A `JSONObject` object must be created before usage with one of the class
methods like `util.JSONObject.create()`.

The `JSONObject` class provides methods for accessing, adding/replacing or
deleting the values by name with the `get()`, `put()` and
`remove()` methods.

The `get()` method can return a simple value, a
`util.JSONObject` or a `util.JSONArray` object reference.

The `put()` method can take a simple value, a `RECORD`, or an
`ARRAY` as parameter.

If the structure of the JSON object is not known at compile time, you can introspect the
elements of the object with the `getLength()`, `getType()` and
`name()` methods.

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [util.JSONObject methods](3591-util-jsonobject-methods.md): Methods for the util.JSONObject class.
