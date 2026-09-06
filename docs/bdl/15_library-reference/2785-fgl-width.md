---
title: "fgl_width()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_WIDTH.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_width()"
type: "concept"
---

# fgl_width()

> Returns the number of columns needed to represent the printed version of the expression.

## Syntax

```
FUNCTION fgl_width(
   str STRING )
  RETURNS INTEGER
```

1. str is any valid string expression.

## Usage

The `fgl_width()` function returns the number of columns that
will be used if you display str on a text terminal.

If the parameter is [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value."),
the function returns zero.

The number of columns used by a character depends on the glyph (that is the graphical symbol used
to draw the character on the screen). For example, an ASCII character like "A" uses one column,
while one Chinese character uses 2 columns (this means that on a text terminal, the size of one
Chinese character takes the same size as "AB").

Trailing blanks are counted in the length of the string.

## Related links

**Related concepts**  

[length()](2787-length.md "Returns the number of characters in a string passed as parameter.")

[fgl\_mblen()](2773-fgl-mblen.md "Returns the number of bytes of the first character in a string.")
