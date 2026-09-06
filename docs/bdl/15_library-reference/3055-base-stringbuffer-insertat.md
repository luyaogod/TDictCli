---
title: "base.StringBuffer.insertAt"
source: "fgl-topics/c_fgl_ClassStringBuffer_insertAt.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.insertAt"
type: "concept"
---

# base.StringBuffer.insertAt

> Insert a string at a given position.

## Syntax

```
insertAt(
   index INTEGER,
   str STRING )
```

1. index is the position where the string must be
   inserted.
2. str is the string to be inserted.

## Usage

The `insertAt()` method inserts a string before the specified position in the
string buffer.

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
   CALL buf.insertAt(3, "xx")
   DISPLAY buf.toString()
END MAIN
```

Output:

```
abxxcdef
```
