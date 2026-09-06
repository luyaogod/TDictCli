---
title: "Datetime expressions"
source: "fgl-topics/c_fgl_Expressions_Datetime.html"
breadcrumb: "Language basics > Expressions > Datetime expressions"
type: "concept"
---

# Datetime expressions

> This section covers date-time expression evaluation rules.

A datetime expression evaluates to a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") data type.

```
MAIN
  DEFINE r, c DATETIME YEAR TO SECOND
  LET c = CURRENT YEAR TO SECOND
  LET r = c + INTERVAL( 234-02 ) YEAR TO MONTH
END MAIN
```

The operands of a datetime expression can be one of:

- A [datetime literal](0590-datetime-literals.md "Datetime literals define date/time value in an expression.").
- A [character string literal](0588-text-literals.md "Text literals define a character string in an expression.") representing a
  datetime with the format `YYYY-MM-DD hh:mm:ss.fffff`.
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of `DATETIME` type.
- A [function](0761-functions.md "Describes user defined functions.") returning a single datetime value.
- A unary `+` or `-` sign associated with an [interval expression](0600-interval-expressions.md "This section covers interval expression evaluation rules.").
- A [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") expression.
- An [`EXTEND`](0660-extend-function.md "The EXTEND() operator adjusts a date time value depending on the qualifier.") expression.

If a datetime expression includes an operand whose value is not a datetime data type,
the runtime system attempts to convert the value to a datetime value following the data
conversion rules.

If an element of an datetime expression is [`NULL`](0572-null.md "The NULL constant defines a non-value."), the expression is evaluated to `NULL`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
