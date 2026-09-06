---
title: "base.StringBuffer.equals"
source: "fgl-topics/c_fgl_ClassStringBuffer_equals.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.equals"
type: "concept"
---

# base.StringBuffer.equals

> Compare strings (case sensitive).

## Syntax

```
equals(
   str STRING )
  RETURNS BOOLEAN
```

1. str is the string to compare with.

## Usage

Use the `equals()` method
to determine whether the value of a `base.StringBuffer` object
is identical to a specified string.

This method is case-sensitive.

Since
the parameter for the method must be a string, you can use the [`toString()`](3060-base-stringbuffer-tostring.md "Create a STRING from the string buffer.")
method to convert a `base.StringBuffer` object in
order to compare it.

The method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") if
the strings are identical, otherwise it returns [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.").

## Example

```
MAIN
   DEFINE buf, buf2 base.StringBuffer,
          mystring STRING

   LET buf = base.StringBuffer.create()
   CALL buf.append("there")

   -- compare to a STRING
   IF buf.equals("there") THEN
      DISPLAY "buf matches there"
   END IF

   -- compare to a STRING variable
   LET mystring = "there"
   IF buf.equals(mystring) THEN
      DISPLAY "buf matches mystring"
   END IF

   -- compare to another StringBuffer object
   LET buf2 = base.StringBuffer.create()
   CALL buf2.append("there")
   IF buf.equals(buf2.toString()) THEN
      DISPLAY "buf matches buf2"
   END IF
END MAIN
```

Output:

```
buf matches there
buf matches mystring
buf matches buf2
```
