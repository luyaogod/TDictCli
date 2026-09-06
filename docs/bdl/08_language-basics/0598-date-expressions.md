---
title: "Date expressions"
source: "fgl-topics/c_fgl_Expressions_Date.html"
breadcrumb: "Language basics > Expressions > Date expressions"
type: "concept"
---

# Date expressions

> This section covers date expression evaluation rules.

A date expression evaluates to a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") data type value.

```
MAIN
  DEFINE r, c DATE
  LET c = TODAY + 4
  LET r = ( c - 2 )
END MAIN
```

The operands of a date expression can be one of:

- A [character string literal](0588-text-literals.md "Text literals define a character string in an expression.") that can be evaluated
  to a date based on the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment
  variable.
- A [variable](0686-variables.md "Explains how to define program variables.") or [constant](0710-constants.md "The definition of constants allows to centralize common static values.") of type [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.").
- A [function](0761-functions.md "Describes user defined functions.") returning a single date value.
- A unary `+` or `-` sign associated with an [integer expression](0595-integer-expressions.md "This section covers integer expression evaluation rules.") representing a number of days.
- The [`TODAY`](0663-today-function.md "The TODAY operator returns the current calendar date.") constant.
- A [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") expression with
  `YEAR TO DAY` qualifiers.
- An [`EXTEND`](0660-extend-function.md "The EXTEND() operator adjusts a date time value depending on the qualifier.") expression with
  `YEAR TO DAY` qualifiers.

If a date expression includes an operand whose value is not a date data type,
the runtime system attempts to convert the value to a date value following the
data conversion rules.

If an element of an date expression is [`NULL`](0572-null.md "The NULL constant defines a non-value."), the expression is evaluated to `NULL`.

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
