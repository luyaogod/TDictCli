---
title: "base.StringBuffer.toString"
source: "fgl-topics/c_fgl_ClassStringBuffer_toString.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.toString"
type: "concept"
---

# base.StringBuffer.toString

> Create a STRING from the string buffer.

## Syntax

```
toString()
  RETURNS STRING
```

## Usage

The `toString()` method
creates a [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.")
value from the current string buffer.

Use this method if
you need to pass the string to another method or instruction that
expects a `STRING` as parameter.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("abc")
   DISPLAY buf.toString()
END MAIN
```

Output:

```
abc
```
