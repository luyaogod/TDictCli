---
title: "DYNAMIC ARRAY.appendElement"
source: "fgl-topics/c_fgl_Arrays_ARRAY_appendElement.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.appendElement"
type: "concept"
---

# DYNAMIC ARRAY.appendElement

> Adds a new element to the end of the array.

## Syntax

```
appendElement( )
```

## Usage

This method creates a new element at the end of the array.

The element is initialized to `NULL`.

## Example

```
MAIN
  DEFINE a DYNAMIC ARRAY OF INTEGER
  DEFINE x INT
  FOR x=1 TO 5
      CALL a.appendElement()
      LET a[x] = 100+x
  END FOR
  DISPLAY a.getLength() -- shows 5
  DISPLAY a[3] -- shows 103
END MAIN
```

Since element allocation occurs automatically for dynamic arrays, you can omit the
call to the `appendElement()` method and directly assign the new element:

```
MAIN
  DEFINE a DYNAMIC ARRAY OF INTEGER
  LET a[100] = 87234 -- Array gets a length of 100 automatically
  LET a[101] = 98562 -- New element at position 101
END MAIN
```

However, for better code readability, you might want to use the `appendElement()`
method in some cases.
