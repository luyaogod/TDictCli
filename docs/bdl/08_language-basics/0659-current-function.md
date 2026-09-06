---
title: "CURRENT [function]"
source: "fgl-topics/c_fgl_operators_CURRENT.html"
breadcrumb: "Language basics > Operators > List of expression elements > Date and time operators > CURRENT [function]"
type: "concept"
---

# CURRENT [function]

> The CURRENT function operator returns the current system date and time.

## Syntax

```
CURRENT [ qual1 TO qual2 [(scale)] ]
```

1. qual1 and qual2 define
   the date time qualifiers. The qualifiers can be a combination of `YEAR`,
   `MONTH`, `DAY`, `HOUR`, `MINUTE`,
   `SECOND` and `FRACTION[(scale)]` as in the
   definition of a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
   type.
2. scale defines the precision of the fractional part of a second, from 1 to 5 digits:
   - `FRACTION(1)` : `10-1 s` (deciseconds)
   - `FRACTION(2)` : `10-2 s` (centiseconds)
   - `FRACTION(3)` : `10-3 s` (milliseconds)
   - `FRACTION(4)` : `10-4 s` (100 microseconds)
   - `FRACTION(5)` : `10-5 s` (10 microseconds)

## Usage

The `CURRENT` operator returns the system date/time in the current local timezone.

This operator can be used to assign the current system date and time to a [`DATETIME`](0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") variable.

Use optional datetime qualifiers to specify the precision of the returned value. The possible
qualifiers are the same as in a `DATETIME` data type definition.

If the datetime qualifiers are not specified after the `CURRENT` keyword, the
precision defaults to `YEAR TO FRACTION(3)`.

## Example

```
MAIN
    DISPLAY CURRENT -- defaults to YEAR TO FRACTION(3)
    DISPLAY CURRENT YEAR TO DAY
    DISPLAY CURRENT HOUR TO SECOND
    DISPLAY CURRENT HOUR TO FRACTION -- defaults to FRACTION(3)
    DISPLAY CURRENT HOUR TO FRACTION(5)
    DISPLAY CURRENT SECOND TO FRACTION(5)
    DISPLAY CURRENT FRACTION TO FRACTION(5)
END MAIN
```

Output:

```
2025-03-15 09:49:03.041
2025-03-15
09:49:03
09:49:03.042
09:49:03.04201
03.04201
.04202
```

## Related links

**Related concepts**  

[EXTEND() [function]](0660-extend-function.md "The EXTEND() operator adjusts a date time value depending on the qualifier.")
