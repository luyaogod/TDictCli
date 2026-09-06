---
title: "Lower or equal (<=)"
source: "fgl-topics/c_fgl_operators_LOWER_OR_EQUAL.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > Lower or equal (<=)"
type: "concept"
---

# Lower or equal (<=)

> The <= operator is provided to test whether a value or expression is lower than or equal to another.

## Syntax

```
expr <= expr
```

## Usage

Use the lower-or-equal operator to test if a value of the left operand is lower than or equal to
the value of the right operand.

This operator applies to expressions that evaluate to primitive data types such
as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers."), [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation."). It does not apply to the [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") and [`TEXT`](0569-text.md "The TEXT data type stores large text data.") types.

If one of the operands is `NULL`, the comparison expression evaluates to
`NULL`.

## Example

```
MAIN
  DEFINE n INT
  LET n = 100
  IF n <= 100 THEN
     DISPLAY "The variable is lower than or equal to 100."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")
