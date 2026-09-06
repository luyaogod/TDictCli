---
title: "The util.JSONArray class"
source: "fgl-topics/c_fgl_ext_util_JSONArray.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class"
type: "concept"
---

# The util.JSONArray class

> The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.

The `util.JSONArray` class is provided in the `util`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`util.JSONArray` class, you must import the `util` package in your
program:

```
IMPORT util
```

A `JSONArray` is an sequence of unnamed values. The format of a JSON array string
is a list of values wrapped in square brackets with commas between the
values:

```
[123,546,"abc","def","xyz"]
```

A `JSONArray` object must be created before usage with one of the class
methods like `util.JSONArray.create()`.

The `JSONArray` class provides methods for accessing, adding/replacing or deleting
the array values by index with the `get()`, `put()` and
`remove()` methods.

If the structure of the JSON array is not known at compile time, you can introspect the
elements of the array with the `getLength()` and `getType()`
methods.

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [util.JSONArray methods](3605-util-jsonarray-methods.md): Methods for the util.JSONArray class.
