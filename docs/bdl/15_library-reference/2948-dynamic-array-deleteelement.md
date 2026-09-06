---
title: "DYNAMIC ARRAY.deleteElement"
source: "fgl-topics/c_fgl_Arrays_ARRAY_deleteElement.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.deleteElement"
type: "concept"
---

# DYNAMIC ARRAY.deleteElement

> Removes an element from the array.

## Syntax

```
deleteElement( index INTEGER )
```

1. index is the position of the element to be removed.

## Usage

This method removes the array element at the specified index.

No error is raised, if the index is out of bounds.

## Example

```
MAIN
  DEFINE a DYNAMIC ARRAY OF INTEGER
  LET a[10] = 9
  CALL a.deleteElement(5)
  DISPLAY a.getLength() -- shows 9
  DISPLAY a[9]  -- shows 9
END MAIN
```
