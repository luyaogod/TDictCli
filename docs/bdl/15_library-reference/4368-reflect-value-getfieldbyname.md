---
title: "reflect.Value.getFieldByName"
source: "fgl-topics/c_fgl_ext_reflect_Value_getFieldByName.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getFieldByName"
type: "concept"
---

# reflect.Value.getFieldByName

> Returns a field of a record, from the member name.

## Syntax

```
getFieldByName(
     name STRING )
  RETURNS reflect.Value
```

1. name is the name of the field in the `RECORD` structure. The
   field name is case sensitive.

## Usage

The `getFieldByName()` method returns a `reflect.Value` object that
is a reference to the field, which corresponds to a member with the name passed as parameter, for
the record structure represented by this `reflect.Value` object.

The name of the field (record member) is case sensitive. The parameter
passed to this method must match the name used in the `RECORD` or
`TYPE` definition. If the parameter does not match a record field name, the method
returns `NULL`.

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
    DISPLAY "value  = ", val.getFieldByName("name").toString()
    DISPLAY "type   = ", val.getFieldByName("name").getType().toString()
END MAIN
```

Shows:

```
value  = Mike FITZPATRICK
type   = VARCHAR(30)
```
