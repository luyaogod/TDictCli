---
title: "reflect.Value.appendArrayElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_appendArrayElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.appendArrayElement"
type: "concept"
---

# reflect.Value.appendArrayElement

> Appends a new element to an array.

## Syntax

```
appendArrayElement()
```

## Usage

The `appendArrayElement()` method adds a new element to the end of the array
represented by this `reflect.Value` object.

The `reflect.Value` object used to call this method must have been
created from a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md) variable, or
is a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dynamic array.

## Example

```
IMPORT reflect
MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING
    DEFINE val reflect.Value
    LET val = reflect.Value.valueOf( arr )
    CALL val.appendArrayElement()
    DISPLAY "len = ", arr.getLength()
END MAIN
```

Shows:

```
len =           1
```
