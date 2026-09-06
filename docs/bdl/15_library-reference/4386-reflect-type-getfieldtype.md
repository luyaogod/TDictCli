---
title: "reflect.Type.getFieldType"
source: "fgl-topics/c_fgl_ext_reflect_Type_getFieldType.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getFieldType"
type: "concept"
---

# reflect.Type.getFieldType

> Returns the type of a member of a record type.

## Syntax

```
getFieldType(
     index INTEGER )
  RETURNS reflect.Type
```

1. index is the ordinal position of the field in the `RECORD`
   structure.

## Usage

The `getFieldType()` method returns a `reflect.Type` object at the
given index, for the record type represented by this `reflect.Type` object.

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
    DISPLAY "type = ", typ.getFieldType(1).toString()
END MAIN
```

Shows:

```
type = INTEGER
```
