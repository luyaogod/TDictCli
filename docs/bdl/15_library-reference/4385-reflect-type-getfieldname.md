---
title: "reflect.Type.getFieldName"
source: "fgl-topics/c_fgl_ext_reflect_Type_getFieldName.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getFieldName"
type: "concept"
---

# reflect.Type.getFieldName

> Returns the name of a member of a record type.

## Syntax

```
getFieldName(
     index INTEGER )
  RETURNS STRING
```

1. index is the ordinal position of the field in the `RECORD`
   structure.

## Usage

The `getFieldName()` method returns the name of the member at the given index, for
the record type represented by this `reflect.Type` object.

The `reflect.Type` object used to call this method must have been
created with a [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") variable, or is a
`reflect.Type` object returned from a method like [`reflect.Value.getType()`](4371-reflect-value-gettype.md "Returns the reflect.Type object of a reflect.Value object.") or [`getElementType()`](4383-reflect-type-getelementtype.md "Returns the type of the elements in an array or in a dictionary."), and
references a record structure.

## Example

```
IMPORT reflect
MAIN
    DEFINE rec RECORD
                   pkey INTEGER,
                   name VARCHAR(30)
               END RECORD
    DEFINE typ reflect.Type
    LET typ = reflect.Type.typeOf( rec )
    DISPLAY "field 1 name = ", typ.getFieldName(1)
END MAIN
```

Shows:

```
field 1 name = pkey
```
