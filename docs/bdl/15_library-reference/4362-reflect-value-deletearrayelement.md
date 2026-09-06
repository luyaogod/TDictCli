---
title: "reflect.Value.deleteArrayElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_deleteArrayElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.deleteArrayElement"
type: "concept"
---

# reflect.Value.deleteArrayElement

> Deletes an element of an array.

## Syntax

```
deleteArrayElement( index INTEGER )
```

1. index is the index of the array element to be deleted.

## Usage

The `deleteArrayElement()` method removes an element at the given index, from the
array represented by this `reflect.Value` object.

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
    CALL val.deleteArrayElement(1)
    DISPLAY "len = ", arr.getLength()
    DISPLAY "arr[1] = ", arr[1]
END MAIN
```

Shows:

```
len =           1
arr[1] = bbbbb
```
