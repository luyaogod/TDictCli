---
title: "reflect.Type.getKind"
source: "fgl-topics/c_fgl_ext_reflect_Type_getKind.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getKind"
type: "concept"
---

# reflect.Type.getKind

> Returns the kind of a type.

## Syntax

```
getKind()
  RETURNS STRING
```

1. Possible returned values: `PRIMITIVE`, `RECORD`,
   `ARRAY`, `DICTIONARY`, `FUNCTION`,
   `INTERFACE`.

## Usage

The `getKind()` method returns the kind of this `reflect.Type`
object.

The kind can be:

- `"PRIMITIVE"` for [primitive data
  types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") (`INTEGER`, `VARCHAR(50)`, etc)
- `"RECORD"`, when the object type represents a [structured record](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") (with or without associated [methods](../08_language-basics/0772-methods.md "A function declared with a receiver type defines a method for this type.")).
- `"ARRAY"` when the object type represents a [dynamic array](../08_language-basics/0734-dynamic-arrays.md) or [static array](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size.").
- `"DICTIONARY"` when the object type represents a [dictionary](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.").
- `"FUNCTION"` when the object type represents a [function type](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.").
- `"INTERFACE"` when the object type represents an [interface type](../08_language-basics/0778-interfaces.md "An interface groups a set of methods acting on a user-defined type.").

> **Tip:**
>
> Based on the kind of the `reflect.Type` object, use the methods
> corresponding to that kind: `getElementType()` for `"ARRAY"` or
> `"DICTIONARY"`, `getMethod()` for `"INTERFACE"`,
> etc.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE dic DICTIONARY OF STRING
    DEFINE rec RECORD pkey INT END RECORD
    DEFINE dtm DATETIME YEAR TO FRACTION(5)
    DISPLAY "arr kind = ", reflect.Type.typeOf(arr).getKind()
    DISPLAY "dic kind = ", reflect.Type.typeOf(dic).getKind()
    DISPLAY "rec kind = ", reflect.Type.typeOf(rec).getKind()
    DISPLAY "dtm kind = ", reflect.Type.typeOf(dtm).getKind()
END MAIN
```

Shows:

```
arr kind = ARRAY
dic kind = DICTIONARY
rec kind = RECORD
dtm kind = PRIMITIVE
```
