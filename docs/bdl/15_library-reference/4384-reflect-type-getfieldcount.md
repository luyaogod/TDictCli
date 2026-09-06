---
title: "reflect.Type.getFieldCount"
source: "fgl-topics/c_fgl_ext_reflect_Type_getFieldCount.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getFieldCount"
type: "concept"
---

# reflect.Type.getFieldCount

> Returns the number of fields of record type.

## Syntax

```
getFieldCount()
  RETURNS INTEGER
```

## Usage

The `getFieldCount()` method returns the number of members of the record type
represented by this `reflect.Type` object.

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
    DISPLAY "field count  = ", typ.getFieldCount()
END MAIN
```

Shows:

```
field count  =           2
```
