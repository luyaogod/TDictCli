---
title: "length()"
source: "fgl-topics/c_fgl_BuiltInFunctions_LENGTH.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > length()"
type: "concept"
---

# length()

> Returns the number of characters in a string passed as parameter.

## Syntax

```
FUNCTION length(
   str STRING )
  RETURNS INTEGER
```

1. str is any valid character string expression supported by the language.

## Usage

The `length()` function counts the length of a character string.

> **Note:**
>
> Unlike the [`STRING.getLength()`](2923-string-getlength.md "Returns the length of the current string.") method, the `LENGTH()` function does not
> count the trailing blanks.

If the parameter is [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value."),
the function returns zero.

> **Important:**
>
> When using byte length semantics, the length is expressed in bytes. When
> using [char length semantics](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."), the
> unit is characters. This matters when using a multibyte locale such as UTF-8.

## Related links

**Related concepts**  

[fgl\_width()](2785-fgl-width.md "Returns the number of columns needed to represent the printed version of the expression.")

[fgl\_mblen()](2773-fgl-mblen.md "Returns the number of bytes of the first character in a string.")
