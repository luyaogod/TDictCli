---
title: "base.StringBuffer.append"
source: "fgl-topics/c_fgl_ClassStringBuffer_append.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.append"
type: "concept"
---

# base.StringBuffer.append

> Append a string at the end of the current string.

## Syntax

```
append(
   str STRING )
```

1. str is the string to append to the string buffer.

## Usage

The `append()` method appends a string to the internal string buffer.

## Example

```
DEFINE buf base.StringBuffer
LET buf = base.StringBuffer.create()
CALL buf.append("abc")
...
```

For a complete example, see [Example 1: Add strings to a StringBuffer](3069-example-1-add-strings-to-a-stringbuffer.md).
