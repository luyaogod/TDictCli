---
title: "DYNAMIC ARRAY.insertElement"
source: "fgl-topics/c_fgl_Arrays_ARRAY_insertElement.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.insertElement"
type: "concept"
---

# DYNAMIC ARRAY.insertElement

> Inserts a new element at the given index.

## Syntax

```
insertElement( index INTEGER )
```

1. index is the position where a new element must be inserted.

## Usage

This method inserts a new element in the array, before the specified index.

If the specified index equals the array length plus one, a new element will be created at the
specified index. This includes the case when inserting an element at index 1 and the array is
empty.

Otherwise, when the index is out of bounds, no error is raised and the array is left
untouched.

## Example

```
MAIN
  DEFINE a DYNAMIC ARRAY OF INTEGER

  LET a[10] = 11
  CALL a.insertElement(10) -- insert at 10
  LET a[10] = 10
  DISPLAY a.getLength() -- shows 11
  DISPLAY a[10]  -- shows 10
  DISPLAY a[11]  -- shows 11

  CALL a.clear()
  CALL a.insertElement(10) -- nop
  DISPLAY a.getLength() -- shows 0
  CALL a.insertElement(1) -- insert at 1
  DISPLAY a.getLength() -- shows 1
  CALL a.insertElement(50) -- nop
  DISPLAY a.getLength() -- shows 1
  CALL a.insertElement(2) -- append
  DISPLAY a.getLength() -- shows 2

END MAIN
```
