---
title: "reflect.Type.getElementType"
source: "fgl-topics/c_fgl_ext_reflect_Type_getElementType.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getElementType"
type: "concept"
---

# reflect.Type.getElementType

> Returns the type of the elements in an array or in a dictionary.

## Syntax

```
getElementType()
  RETURNS reflect.Type
```

## Usage

The `getElementType()` method returns a `reflect.Type` object that
represents the type used by the elements of the array or dictionary represented by this
`reflect.Type` object.

The `reflect.Type` object used to call this method must have been
created from a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md) or [`DICTIONARY`](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") variable, or is a
`reflect.Type` object returned from a method like [`getFieldType()`](4386-reflect-type-getfieldtype.md "Returns the type of a member of a record type."), and
references a dynamic array or dictionary.

> **Note:**
>
> The element type of a multi-dimensional array is an array-type. For example, when the
> `reflect.Type` object represents an `ARRAY[10,20] OF INTEGER`, the
> element-type is `ARRAY[20] OF INTEGER`.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE typ reflect.Type
    LET typ = reflect.Type.typeOf(arr)
    DISPLAY "element type = ", typ.getElementType().toString()
END MAIN
```

Shows:

```
element type = STRING
```
