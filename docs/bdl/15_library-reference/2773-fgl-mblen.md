---
title: "fgl_mblen()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_MBLEN.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_mblen()"
type: "concept"
---

# fgl_mblen()

> Returns the number of bytes of the first character in a string.

## Syntax

```
FUNCTION fgl_mblen(
   str STRING )
  RETURNS INTEGER
```

1. str is any valid string expression.

## Usage

The `fgl_mblen()` function returns the number of bytes used
to encode the first character of the specified string, in the current application
locale.

When using a multibyte character set like UTF-8, you can use this function
to compute the size in bytes of the first character in the string. Then, you
can use this size to identify the position of the next character in the
string. This function is mandatory to scan multibyte encoded strings.

If the parameter is NULL, the function returns zero.

If the parameter starts with an invalid multibyte character, the function
returns -1.

## Related links

**Related concepts**  

[length()](2787-length.md "Returns the number of characters in a string passed as parameter.")

[fgl\_width()](2785-fgl-width.md "Returns the number of columns needed to represent the printed version of the expression.")
