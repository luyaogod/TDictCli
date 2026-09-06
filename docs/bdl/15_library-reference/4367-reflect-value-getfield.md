---
title: "reflect.Value.getField"
source: "fgl-topics/c_fgl_ext_reflect_Value_getField.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getField"
type: "concept"
---

# reflect.Value.getField

> Returns a field by index for a record.

## Syntax

```
getField(
     index INTEGER )
  RETURNS reflect.Value
```

1. index is the ordinal position of the field in the `RECORD`
   structure.

## Usage

The `getField()` method returns a `reflect.Value` object that is a
reference to the field, of the record structure represented by this `reflect.Value`
object, at the member position specified as parameter. First member starts at 1.

The `reflect.Value` object used to call this method must have been
created with a [`RECORD`](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") variable, or is a
`reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), [`getFieldByName()`](4368-reflect-value-getfieldbyname.md "Returns a field of a record, from the member name."), [`getArrayElement()`](4363-reflect-value-getarrayelement.md "Returns an element of an array."), or
[`getDictionaryElement()`](4365-reflect-value-getdictionaryelement.md "Returns an element of a dictionary."), and references a record structure.

The new `reflect.Value` object references the original
variable: Any call to a reflection manipulation method modifies the underlying variable
directly.

## Example

```
IMPORT reflect
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(30)
           END RECORD
    DEFINE val reflect.Value
    LET rec.pkey = 101
    LET rec.name = "Mike FITZPATRICK"
    LET val = reflect.Value.valueOf( rec )
    DISPLAY "value  = ", val.getField(1).toString()
    DISPLAY "type   = ", val.getField(1).getType().toString()
END MAIN
```

Shows:

```
value  = 101
type   = INTEGER
```
