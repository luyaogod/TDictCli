---
title: "Understanding type definition"
source: "fgl-topics/c_fgl_user_types_002.html"
breadcrumb: "Language basics > Types > Understanding type definition"
type: "concept"
---

# Understanding type definition

> This is an introduction to types.

The `TYPE` instruction declares a user-defined type, which can be based on:

- [Primitive data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") (`INTEGER`,
  `VARCHAR(n)`)
- [Records](0715-records.md "Records allow structured program variables definitions.") (`RECORD ... END RECORD`,
  `RECORD ... LIKE tabname.*`)
- [Arrays](0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") (`DYNAMIC ARRAY OF ...`)
- [Dictionaries](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") (`DICTIONARY OF ...`)
- [Function references](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") (`FUNCTION
  name(type,...) RETURNS ...`)
- [Interfaces](0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.") (`INTERFACE ... END
  INTERFACE`)

Once declared, a type can be referenced in the declaration of program variables, or in other
types.

Types are typically defined to avoid the repetition of complex structured types.

> **Tip:**
>
> User-defined types improve code readability by centralizing data structure
> definitions. Consider using `PUBLIC TYPE` definitions to share types across modules
> with [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").

Types can also be completed with [methods](0772-methods.md "A function declared with a receiver type defines a method for this type.") to
encapsulate the data, and use the concept of polymorphism with [interfaces](0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.").

Anonymous types are defined automatically when defining variables and no explicit TYPE definition
is used. For more details, see [Anonymous types](0755-anonymous-types.md "Anonymous types are created from DEFINE instructions.").

## Related links

**Related concepts**  

[Variables](0686-variables.md "Explains how to define program variables.")
