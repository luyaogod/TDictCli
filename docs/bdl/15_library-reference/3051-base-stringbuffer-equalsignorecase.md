---
title: "base.StringBuffer.equalsIgnoreCase"
source: "fgl-topics/c_fgl_ClassStringBuffer_equalsIgnoreCase.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.equalsIgnoreCase"
type: "concept"
---

# base.StringBuffer.equalsIgnoreCase

> Compare strings (case insensitive)

## Syntax

```
equalsIgnoreCase(
   str STRING )
  RETURNS BOOLEAN
```

1. str is the string to compare with.

## Usage

The `equalsIgnoreCase()` method
compares the current string buffer with the passed string, ignoring
the character case.

Since the parameter for the method must
be a string, you can use the [`toString()`](3060-base-stringbuffer-tostring.md "Create a STRING from the string buffer.") method
to convert a `base.StringBuffer` object in order
to compare it.

The method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") if
the strings are identical, otherwise it returns [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.").

## Example

```
MAIN
   DEFINE buf3 base.StringBuffer
   LET buf3 = base.StringBuffer.create()
   CALL buf3.append("there")
   IF buf3.equalsIgnoreCase("There") THEN
      DISPLAY "buf matches There ignoring case"
   END IF
END MAIN
```

Output:

```
buf matches There ignoring case
```
