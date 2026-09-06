---
title: "base.StringBuffer.toLowerCase"
source: "fgl-topics/c_fgl_ClassStringBuffer_toLowerCase.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.toLowerCase"
type: "concept"
---

# base.StringBuffer.toLowerCase

> Converts the string in the buffer to lower case.

## Syntax

```
toLowerCase()
```

## Usage

The `toLowerCase()` method
converts the current string to lower case.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("AbC")
   CALL buf.toLowerCase()
   DISPLAY buf.toString()
END MAIN
```

Output:

```
abc
```
