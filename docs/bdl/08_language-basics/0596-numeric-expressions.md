---
title: "Numeric expressions"
source: "fgl-topics/c_fgl_Expressions_Number.html"
breadcrumb: "Language basics > Expressions > Numeric expressions"
type: "concept"
---

# Numeric expressions

> This section covers numeric expression evaluation rules.

A numeric expression evaluates to a decimal value of type [`DECIMAL`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") (or [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."))

```
MAIN
  DEFINE r, c DECIMAL(10,2)
  LET c = 456.22
  LET r = c * 2 + ( c / 4.55 )
END MAIN
```

The operands of a numeric expression can be one of:

- An [integer literal](0586-integer-literals.md "Integer literals define a whole number in an expression.").
- A [decimal literal](0587-numeric-literals.md "Numeric literals define values with a decimal part in an expression.").
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of numeric data type.
- A [function](0761-functions.md "Describes user defined functions.") returning a single numeric value.
- A [boolean expression](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.").
- The result of a [DATE](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") subtraction, as a number of days.

If a number expression includes an operand whose value is not a numeric data type,
the runtime system attempts to convert the value to a number following the data
conversion rules.

If an element of a number expression is `NULL`,
the expression is evaluated to `NULL`.

Depending on the operators and the resulting values, the data type of a numeric expression will
be `DECIMAL` or `INTEGER`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
