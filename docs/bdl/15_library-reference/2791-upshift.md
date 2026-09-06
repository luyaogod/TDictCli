---
title: "upshift()"
source: "fgl-topics/c_fgl_BuiltInFunctions_UPSHIFT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > upshift()"
type: "concept"
---

# upshift()

> Converts a string to uppercase.

## Syntax

```
FUNCTION upshift(
   s STRING )
  RETURNS STRING
```

1. s is the character string to convert to uppercase letters.

## Usage

The `upshift()` function returns a string value in which all
lowercase characters in its argument are converted to uppercase.

The character conversion depends on
[locale settings](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.") (the LC\_CTYPE
environment variable). Non-alphabetic or uppercase characters are not altered.

## Related links

**Related concepts**  

[downshift()](2729-downshift.md "Converts a string to lowercase.")
