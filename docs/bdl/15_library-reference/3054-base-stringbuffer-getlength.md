---
title: "base.StringBuffer.getLength"
source: "fgl-topics/c_fgl_ClassStringBuffer_getLength.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.getLength"
type: "concept"
---

# base.StringBuffer.getLength

> Return the length of a string.

## Syntax

```
getLength()
  RETURNS INTEGER
```

## Usage

Use the `getLength()` method to return the number of characters in the
current string buffer, including trailing spaces.

The length of an empty string buffer is 0.

> **Important:**
>
> When using byte length semantics, the string length is expressed in bytes.
> When using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is
> characters. This matters when using a multibyte locale such as UTF-8.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("abc")
   DISPLAY buf.getLength()
   -- append three spaces to the end of the string 
   CALL buf.append("   ") 
   DISPLAY buf.getLength()
END MAIN
```

Output:

```
          3
          6
```
