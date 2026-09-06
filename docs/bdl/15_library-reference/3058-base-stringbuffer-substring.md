---
title: "base.StringBuffer.subString"
source: "fgl-topics/c_fgl_ClassStringBuffer_subString.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.subString"
type: "concept"
---

# base.StringBuffer.subString

> Return the substring at the specified position.

## Syntax

```
subString(
   startIndex INTEGER,
   endIndex INTEGER )
  RETURNS STRING
```

1. startIndex is the substring to be found.
2. endIndex is the ending position.

## Usage

The `subString()` method returns the substring defined by the start and end
positions passed as parameter.

The first character is at position 1.

> **Important:**
>
> When using byte length semantics, the positions are expressed in bytes. When
> using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is characters.
> This matters when using a multibyte locale such as UTF-8.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("abcdefg")
   DISPLAY buf.subString(2,5)
END MAIN
```

Output:

```
bcde
```
