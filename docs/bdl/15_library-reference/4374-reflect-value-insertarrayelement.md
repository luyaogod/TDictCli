---
title: "reflect.Value.insertArrayElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_insertArrayElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.insertArrayElement"
type: "concept"
---

# reflect.Value.insertArrayElement

> Inserts a new element into an array.

## Syntax

```
insertArrayElement( index INTEGER )
```

1. index is the index of the new element.

## Usage

The `insertArrayElement()` method inserts a new element at the given index, into
the array represented by this `reflect.Value` object.

The `reflect.Value` object used to call this method must have been
created from a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md) variable, or
is a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dynamic array.

If the index is less than 1 or greater than the array length, the method returns silently without
error.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE val reflect.Value
    LET arr[1] = "aaaaa"
    LET arr[2] = "bbbbb"
    LET val = reflect.Value.valueOf( arr )
    CALL val.insertArrayElement(2)
    DISPLAY "len = ", arr.getLength()
    LET arr[2] = "xxxxx"
    DISPLAY "arr[1] = ", arr[1]
    DISPLAY "arr[2] = ", arr[2]
    DISPLAY "arr[3] = ", arr[3]
END MAIN
```
