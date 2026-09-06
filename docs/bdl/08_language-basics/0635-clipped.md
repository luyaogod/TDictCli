---
title: "CLIPPED"
source: "fgl-topics/c_fgl_operators_CLIPPED.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > CLIPPED"
type: "concept"
---

# CLIPPED

> The CLIPPED operator removes trailing blank spaces (ASCII 32) of a string expression.

## Syntax

```
expr CLIPPED
```

1. expr is a language expression.

## Usage

This operator removes all trailing blank space (ASCII 32) of a string expression.

The `CLIPPED` operator is typically used to remove the trailing blanks of a [`CHAR`](0557-char-size.md "The CHAR data type is a fixed-length character string data type.") value, which would be printed
otherwise.

## Example

```
MAIN
  DISPLAY "Some text   " CLIPPED
END MAIN
```

## Related links

**Related concepts**  

[String expressions](0597-string-expressions.md "This section covers string expression evaluation rules.")
