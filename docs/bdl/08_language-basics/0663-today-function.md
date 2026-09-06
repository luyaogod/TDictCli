---
title: "TODAY [function]"
source: "fgl-topics/c_fgl_operators_TODAY.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > TODAY [function]"
type: "concept"
---

# TODAY [function]

> The TODAY operator returns the current calendar date.

## Syntax

```
TODAY
```

## Usage

`TODAY` returns the current system date as a [`DATE`](0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") value, in the current local
timezone.

This operator can be used to assign the current system date to a `DATE`
variable.

The `TODAY` operator is the `DATE` equivalent for the [`CURRENT`](0659-current-function.md "The CURRENT function operator returns the current system date and time.") operator used for
`DATETIME`.

## Example

```
MAIN
  DISPLAY TODAY
END MAIN
```

## Related links

**Related concepts**  

[Date expressions](0598-date-expressions.md "This section covers date expression evaluation rules.")
