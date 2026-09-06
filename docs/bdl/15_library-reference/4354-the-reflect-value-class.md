---
title: "The reflect.Value class"
source: "fgl-topics/c_fgl_ext_reflect_Value.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class"
type: "concept"
---

# The reflect.Value class

> The reflect.Value class is a generic API to inspect and modify variables.

This class is provided in the `reflect` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the `reflect` package
with:

```
IMPORT reflect
```

A `reflect.Value` object is the runtime representation of a [variable](../08_language-basics/0686-variables.md "Explains how to define program variables.").

A call to reflect.Value returns an object representing the type and data of a variable at
runtime.

Each variable has a Type. Several methods of this API are restricted to certain kinds of Types.
For example, the method [`insertArrayElement()`](4374-reflect-value-insertarrayelement.md "Inserts a new element into an array.") is restricted to the kind `"ARRAY"`.
The program should check the kind of a type before calling kind-specific functions.

> **Important:**
>
> The exceptions thrown by the
> `reflect.*` API can only be caught with a [`TRY/CATCH`](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") block: If [`WHENEVER
> ERROR CONTINUE`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") is active and the reflection API throws an exception, the program
> stops.

## Related links

**Related concepts**  

[IMPORT (C-Extension)](../09_advanced-features/0827-import-c-extension.md "The IMPORT instruction imports c extension module elements to be used by the current module.")

## Child topics

- [reflect.Value methods](4355-reflect-value-methods.md): Methods for the reflect.Value class.
