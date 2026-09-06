---
title: "Exponentiation (**)"
source: "fgl-topics/c_fgl_operators_EXPONENTIATION.html"
breadcrumb: "Language basics > Operators > List of expression elements > Arithmetic operators > Exponentiation (**)"
type: "concept"
---

# Exponentiation (**)

> The ** operator calculates an exponentiation.

## Syntax

```
num-expr ** int-expr
```

1. num-expr is a numeric expression (real number).
2. int-expr is the exponent, as a whole number.

## Usage

The `**` operator returns a value calculated by raising the left-hand
operand to a power corresponding to the integer part of the right-hand
operand.

```
 n ** e   =   n * n * n ... (e times)
```

If the right operand (the exponent) is a number with a decimal part, it is rounded to a whole
integer before computing the exponentiation.

If one of the operands is `NULL`, the arithmetic expression evaluates to
`NULL`.

The `**` operator takes an integer as exponent. If you need to raise a number to
an exponent that is a real number, use the [`fgl_decimal_power()`](../15_library-reference/2742-fgl-decimal-power.md "Raises decimal to the power of the real exponent.")
utility function.

## Example

```
MAIN
  DISPLAY 2 ** 8
  DISPLAY 10 ** 4
END MAIN
```

## Related links

**Related concepts**  

[Numeric expressions](0596-numeric-expressions.md "This section covers numeric expression evaluation rules.")

[fgl\_decimal\_power()](../15_library-reference/2742-fgl-decimal-power.md "Raises decimal to the power of the real exponent.")
