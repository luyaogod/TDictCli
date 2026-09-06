---
title: "base.StringBuffer.getIndexOf"
source: "fgl-topics/c_fgl_ClassStringBuffer_getIndexOf.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.getIndexOf"
type: "concept"
---

# base.StringBuffer.getIndexOf

> Return the position of a substring.

## Syntax

```
getIndexOf(
   str STRING,
   startIndex INTEGER )
  RETURNS INTEGER
```

1. str is the substring to be found.
2. startIndex is the starting position.

## Usage

The `getIndexOf()` method returns the position of a substring
in the string buffer. Specify the substring and an integer specifying the
position at which the search should begin. Use 1 if you want to start at the
beginning of the string buffer.

The method returns zero if the substring is not found.

```
CALL buf.append("abcdef")
DISPLAY buf.getIndexOf("def",1) -- Shows 4
```

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes. When
> using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is characters.
> This matters when using a multibyte locale such as UTF-8.

## Example

This example iterates through the complete string to display the
position of multiple occurrences of the same substring.

```
MAIN
   DEFINE buf base.StringBuffer
   DEFINE pos INTEGER
   DEFINE s STRING
   LET buf = base.StringBuffer.create()
   CALL buf.append("---abc-----abc--abc----")
   LET pos = 1
   LET s = "abc"
   WHILE TRUE
      LET pos = buf.getIndexOf(s,pos)
      IF pos == 0 THEN
         EXIT WHILE
      END IF
      DISPLAY "Pos: ", pos
      LET pos = pos + length(s)
   END WHILE
END MAIN
```

Output:

```
Pos:           4
Pos:          12
Pos:          17
```
