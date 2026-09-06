---
title: "The util.JSON class"
source: "fgl-topics/c_fgl_ext_util_JSON.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class"
type: "concept"
---

# The util.JSON class

> The util.JSON class provides a basic interface to convert program variable values to/from JSON data.

The `util.JSON` class is provided in the `util`
[C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library. To use the
`util.JSON` class, you must import the `util` package in your
program:

```
IMPORT util
```

This class does not have to be instantiated; it provides class methods for
the current program.

The purpose of the `util.JSON` class is to convert a JSON
string from/to a BDL variable, to interface with other software based on the
JSON format.

The BDL variable can be:

- a simple variable (defined with a primitive type such as `DATE`,
  `INTEGER`, `VARCHAR(20)`)
- a structured variable (`RECORD ... END RECORD`)
- a dynamic array (`DYNAMIC ARRAY OF ...`)
- a dictionary (`DICTIONARY OF ...`)

It is not possible to modify JSON elements with this class. In order to
manipulate JSON objects, use the `util.JSONObject` and
`util.JSONArray` classes.

## Related links

**Related concepts**  

[JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.")

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [util.JSON methods](3580-util-json-methods.md): Methods for the util.JSON class.
- [Examples](3588-examples.md): util.JSON usage examples.
