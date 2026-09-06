---
title: "downshift()"
source: "fgl-topics/c_fgl_BuiltInFunctions_DOWNSHIFT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > downshift()"
type: "concept"
---

# downshift()

> Converts a string to lowercase.

## Syntax

```
FUNCTION downshift(
   s STRING )
  RETURNS STRING
```

1. s is the character string to convert to lowercase letters.

## Usage

The `downshift()` function returns a string value in which all
uppercase characters in its argument are converted to lowercase.

The character conversion depends on [locale
settings](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.") (the LC\_CTYPE environment variable). Non-alphabetic or lowercase
characters are not altered.

## Related links

**Related concepts**  

[upshift()](2791-upshift.md "Converts a string to uppercase.")
