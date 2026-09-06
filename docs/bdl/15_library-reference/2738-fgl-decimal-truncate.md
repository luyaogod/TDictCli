---
title: "fgl_decimal_truncate()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DECIMAL_TRUNCATE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_decimal_truncate()"
type: "concept"
---

# fgl_decimal_truncate()

> Returns a decimal truncated to the precision passed as parameter.

## Syntax

```
FUNCTION fgl_decimal_truncate(
   x DECIMAL,
   scale INTEGER )
  RETURNS DECIMAL
```

1. x is the decimal to be converted.
2. scale defines the number of digits after the decimal point.

## Usage

This function truncates the decimal to the number of decimal digits specified.

The value is not rounded, it is just truncated. For example, when truncating 12.345
to 2 decimal digits, the result will be 12.34, not 12.35.
