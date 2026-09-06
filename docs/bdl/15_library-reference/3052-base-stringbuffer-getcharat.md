---
title: "base.StringBuffer.getCharAt"
source: "fgl-topics/c_fgl_ClassStringBuffer_getCharAt.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.getCharAt"
type: "concept"
---

# base.StringBuffer.getCharAt

> Return the character at a specified position.

## Syntax

```
getCharAt(
   index INTEGER )
  RETURNS STRING
```

1. index is the character position in the string.

## Usage

The `getCharAt()` method returns the character from the string buffer at the position
that you specify.

The first character position is 1.

The method returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") if the position
is lower than 1 or greater than the length of the string.

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes. When
> using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is characters.
> This matters when using a multibyte locale such as UTF-8.

## Example

```
MAIN
   DEFINE buf base.StringBuffer 
   LET buf = base.StringBuffer.create()
   CALL buf.append("abcdef")
   DISPLAY buf.getCharAt(3)
END MAIN
```

Output:

```
c
```
