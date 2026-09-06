---
title: "Integer expressions"
source: "fgl-topics/c_fgl_Expressions_Integer.html"
breadcrumb: "Language basics > Expressions > Integer expressions"
type: "concept"
---

# Integer expressions

> This section covers integer expression evaluation rules.

An integer expression evaluates to a whole number of type [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers.") (or [`DECIMAL`](0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage."))

```
MAIN
  DEFINE r, c INTEGER
  LET c = 4
  LET r = c * ( 2 + c MOD 4 ) / getRowCount("customers")
END MAIN
```

The operands of an integer expression can be:

- An [integer literal](0586-integer-literals.md "Integer literals define a whole number in an expression.").
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of type [`TINYINT`](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers."),
  [`SMALLINT`](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers."), [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers.") or [`BIGINT`](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.").
- A [function](0761-functions.md "Describes user defined functions.") returning a single integer value.
- A [boolean expression](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.").
- The result of a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") subtraction,
  as a number of days.

If an integer expression includes an operand whose value is not an integer data type,
the runtime system attempts to convert the value to an integer following the data
conversion rules.

If an element of an integer expression is `NULL`,
the expression is evaluated to `NULL`.

The data type of an integer expression is typically `INTEGER`. However, if the
expression contains operators that can produce a decimal number like the [division operator](0626-division.md "The / operator divides a number by another."), the expresion type can be
`DECIMAL`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
