---
title: "String expressions"
source: "fgl-topics/c_fgl_Expressions_String.html"
breadcrumb: "Language basics > Expressions > String expressions"
type: "concept"
---

# String expressions

> This section covers string expression evaluation rules.

A string expression includes at least one character string value and evaluates to a string data
type value.

```
MAIN
  DEFINE r, c VARCHAR(100)
  LET c = "abcdef" 
  LET r = c[1,3] || ": " || TODAY USING "YYYY-MM-DD" || " " || length(c)
END MAIN
```

The data type of string expression result is
`STRING`.

At least one of the operands in a string expression must be
one of:

- A [character string literal](0588-text-literals.md "Text literals define a character string in an expression.").
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of [`CHAR`](0557-char-size.md "The CHAR data type is a fixed-length character string data type."), [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") or [`TEXT`](0569-text.md "The TEXT data type stores large text data.") data type.
- A [function](0761-functions.md "Describes user defined functions.") returning a single character
  value.

Other operands whose values are not character string data types are converted to strings
by following the data conversion rules.

If an element of a string expression is `NULL`, the expression is evaluated to
`NULL`.

An empty string ("") is equivalent to `NULL`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
