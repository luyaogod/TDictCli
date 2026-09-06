---
title: "The reflect.Type class"
source: "fgl-topics/c_fgl_ext_reflect_Type.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class"
type: "concept"
---

# The reflect.Type class

> The reflect.Type class is a generic API to inspect types.

This class is provided in the `reflect` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the `reflect` package
with:

```
IMPORT reflect
```

A `reflect.Type` object is the runtime representation of a type. A type can be
[anonymous](../08_language-basics/0755-anonymous-types.md "Anonymous types are created from DEFINE instructions.") (when defining a variable directly with a
type specification), or it can be a user-defined and named type, created from a [`TYPE`](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") instruction.

Several methods of this API are restricted to special kinds of types. For example, only the
`reflect.Type` objects of kind `"ARRAY"` and
`"DICTIONARY"` can return an element-type with [`getElementType()`](4383-reflect-type-getelementtype.md "Returns the type of the elements in an array or in a dictionary."). The
program should check the kind of a type before calling kind-specific methods. The kind of a type an
be checked with the [`getKind()`](4388-reflect-type-getkind.md "Returns the kind of a type.") method.

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

- [reflect.Type methods](4380-reflect-type-methods.md): Methods for the reflect.Type class.
