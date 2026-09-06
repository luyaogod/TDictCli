---
title: "FALSE"
source: "fgl-topics/c_fgl_programs_FALSE.html"
breadcrumb: "Language basics > Predefined constants > FALSE"
type: "concept"
---

# FALSE

> FALSE is a predefined constant to be used in boolean expressions.

## Syntax

```
FALSE
```

## Usage

`FALSE` is a predefined constant that can be used as a boolean value in
[boolean expressions](0594-boolean-expressions.md "This section covers boolean expression evaluation rules.").

The `FALSE` constant is equal to 0 (zero), as an `INTEGER`
type.

For backward compatibility reasons, the `FALSE` constant is not of type
`BOOLEAN`. This can have impact on data conversions, for example when using
a JSON API where `true` / `false` are expected.

`TRUE` and `FALSE` are typically used as return values of
functions that give a boolean result.

## Example

```
MAIN
  DEFINE odd BOOLEAN
  LET odd = is_odd(125763)
  IF odd THEN
     DISPLAY "Number is odd."
  END IF
END MAIN

FUNCTION is_odd(value)
  DEFINE value INTEGER
  IF value MOD 2 = 1 THEN
     RETURN TRUE
  ELSE
     RETURN FALSE
  END IF
END FUNCTION
```

## Related links

**Related concepts**  

[BOOLEAN](0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.")
