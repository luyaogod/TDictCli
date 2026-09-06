---
title: "fgl_decimal_power()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DECIMAL_POWER.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_decimal_power()"
type: "concept"
---

# fgl_decimal_power()

> Raises decimal to the power of the real exponent.

## Syntax

```
FUNCTION fgl_decimal_power(
      x DECIMAL,
      y DECIMAL ) 
    RETURNS DECIMAL
```

1. x is the decimal to be raised to the power of y.
2. y is the exponent.

## Usage

Unlike the `**` operator, the `fgl_decimal_power()`
function supports real numbers for the exponent.

## Related links

**Related concepts**  

[Exponentiation (\*\*)](../08_language-basics/0627-exponentiation.md "The ** operator calculates an exponentiation.")
