---
title: "base.StringBuffer.toUpperCase"
source: "fgl-topics/c_fgl_ClassStringBuffer_toUpperCase.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.toUpperCase"
type: "concept"
---

# base.StringBuffer.toUpperCase

> Converts the string in the buffer to upper case.

## Syntax

```
toUpperCase()
```

## Usage

The `toUpperCase()` method
converts the current string to upper case.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("AbC")
   CALL buf.toUpperCase()
   DISPLAY buf.toString()
END MAIN
```

Output:

```
ABC
```
