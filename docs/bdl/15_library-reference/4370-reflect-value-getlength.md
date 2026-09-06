---
title: "reflect.Value.getLength"
source: "fgl-topics/c_fgl_ext_reflect_Value_getLength.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getLength"
type: "concept"
---

# reflect.Value.getLength

> Returns the number of elements in an array.

## Syntax

```
getLength( )
  RETURNS INTEGER
```

## Usage

The `getLength()` method returns the number of elements in the array represented
by this `reflect.Value` object.

The `reflect.Value` object must have been created with a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md) or a [static `ARRAY[...]`](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size.") variable, or is a
`reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dynamic array or static array.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE val reflect.Value
    LET arr[1] = "aaaaa"
    LET arr[2] = "bbbbb"
    LET val = reflect.Value.valueOf( arr )
    DISPLAY "len = ", val.getLength()
END MAIN
```

Shows:

```
len =           2
```
