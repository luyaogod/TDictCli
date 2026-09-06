---
title: "reflect.Value.getArrayElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_getArrayElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getArrayElement"
type: "concept"
---

# reflect.Value.getArrayElement

> Returns an element of an array.

## Syntax

```
getArrayElement(
     index INTEGER )
  RETURNS reflect.Value
```

1. index is the index of the element in the array.

## Usage

The `getArrayElement()` method returns a `reflect.Value` object
that is a reference to the element of the array represented by this `reflect.Value`
object, at the index specified as parameter. First element starts at 1.

The `reflect.Value` object used to call this method must have been
created from a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md) variable, or
is a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dynamic array.

This method throws an error if the index is invalid. This is in contrast to
the ARRAY-subscript operator, which automatically creates a new element. This behaviour is by
design.

The new `reflect.Value` object references the original
variable: Any call to a reflection manipulation method modifies the underlying variable
directly.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE val reflect.Value
    DEFINE elem reflect.Value
    LET arr[1] = "aaaaa"
    LET arr[2] = "bbbbb"
    LET val = reflect.Value.valueOf( arr )
    LET elem = val.getArrayElement(2)
    DISPLAY "val elem 2 = ", elem.toString()
    CALL elem.set(reflect.Value.copyOf("xxxxx"))
    DISPLAY "arr[2]     = ", arr[2]
END MAIN
```

Shows:

```
val elem 2 = bbbbb
arr[2]     = xxxxx
```
