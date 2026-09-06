---
title: "Various type specification"
source: "fgl-topics/c_fgl_runtime_stack_various_types.html"
breadcrumb: "Advanced features > Runtime stack > Various type specification"
type: "concept"
---

# Various type specification

> Some Genero APIs use variant types for parameters or returns.

In Genero BDL, we distinguish different kind of types:

- [Primitive data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), such as
  `INTEGER`, `VARCHAR(50)`.
- [Structured types](../08_language-basics/0715-records.md "Records allow structured program variables definitions."), defined with
  `RECORD` (with or without associated [methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type.")).
- [Array types](../08_language-basics/0734-dynamic-arrays.md) (`DYNAMIC ARRAY OF
  ...`).
- [Dictionary types](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") (`DICTIONARY OF
  ...`).
- [Function reference types](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression."), defined with
  `TYPE + FUNCTION`.
- [Interface types](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type."), defined with `TYPE +
  INTERFACE`.

Some [built-in function](../15_library-reference/2724-built-in-functions.md "A built-in function is a predefined function that is part of the runtime system, or provided as a library function automatically loaded when a program starts. The built-in functions are part of the language.") or methods of [built-in classes](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.") or [extension classes](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.") can take various sort of types
as parameter, or return various types of values.

For example, the [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") method can take a simple primitive data type value, a
structured record, or an array as parameter, and convert that to its corresponding JSON string.

When a Genero API accepts such various type specification as parameter or return value, the
syntax diagram will mention this as
"`any-type`":

```
util.JSON.stringify(
     value any-type
   )
  RETURNS STRING
```
