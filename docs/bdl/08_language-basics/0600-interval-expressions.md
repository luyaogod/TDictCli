---
title: "Interval expressions"
source: "fgl-topics/c_fgl_Expressions_Interval.html"
breadcrumb: "Language basics > Expressions > Interval expressions"
type: "concept"
---

# Interval expressions

> This section covers interval expression evaluation rules.

An interval evaluates to an [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") data type.

```
MAIN
  DEFINE r, c INTERVAL HOUR TO MINUTE
  LET c = "12:45"
  LET r = c + ( DATETIME(14:02) HOUR TO MINUTE - DATETIME(10:43) HOUR TO MINUTE )
END MAIN
```

The operands of an interval expression must be one of:

- An [interval literal](0591-interval-literals.md "Interval literals define an interval value in an expression.").
- A [character string literal](0588-text-literals.md "Text literals define a character string in an expression.") representing an
  Interval with the format `YYYY-MM-DD hh:mm:ss.fffff`.
- An integer expression using the [`UNITS`](0669-units.md "The UNITS operator converts an integer to an interval.") operator.
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of [`INTERVAL`](0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.")
  type.
- A [function](0761-functions.md "Describes user defined functions.") returning a single interval value.
- The result of a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
  subtraction.

If an interval expression includes an operand whose value is not an interval
data type, the runtime system attempts to convert the value to an interval value
following the data conversion rules.

If an element of an interval expression is [`NULL`](0572-null.md "The NULL constant defines a non-value."), the expression is evaluated to `NULL`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
