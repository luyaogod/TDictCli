---
title: "reflect.Type.getFieldTypeByName"
source: "fgl-topics/c_fgl_ext_reflect_Type_getFieldTypeByName.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getFieldTypeByName"
type: "concept"
---

# reflect.Type.getFieldTypeByName

> Returns the type of a member of a record type, from the member name.

## Syntax

```
getFieldTypeByName(
     name STRING )
  RETURNS reflect.Type
```

1. name is the name of the field in the `RECORD` structure. The
   field name is case sensitive.

## Usage

The `getFieldTypeByName()` method returns a `reflect.Type` object
that corresponds to a member with the name passed as parameter, for the record type represented by
this `reflect.Type` object.

The name of the field (record member) is case sensitive. The parameter
passed to this method must match the name used in the `RECORD` or
`TYPE` definition. If the parameter does not match a record field name, the method
returns `NULL`.

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
    DISPLAY "type = ", typ.getFieldTypeByName("name").toString()
END MAIN
```

Shows:

```
type = VARCHAR(30)
```
