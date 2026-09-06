---
title: "The reflect.Method class"
source: "fgl-topics/c_fgl_ext_reflect_Method.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Method class"
type: "concept"
---

# The reflect.Method class

> The reflect.Method class is a generic API to inspect methods.

This class is provided in the `reflect` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the `reflect` package
with:

```
IMPORT reflect
```

A `reflect.Method` object is the runtime representation of a [method](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type."). It can be used to describe method
parameters and return values and identify the method by its name.

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

- [reflect.Method methods](4395-reflect-method-methods.md): Methods for the reflect.Method class.
