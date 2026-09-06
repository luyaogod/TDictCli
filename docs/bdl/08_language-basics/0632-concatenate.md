---
title: "Concatenate (||)"
source: "fgl-topics/c_fgl_operators_CONCATENATE.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > Concatenate (||)"
type: "concept"
---

# Concatenate (||)

> The || operator makes a string concatenation.

## Syntax

```
expr || expr
```

1. expr can be a character, numeric or date/time expression.

## Usage

The `||` operator is the concatenation operator that produces a string expression
from the expression elements on both sides of the operator.

If any of the members of a concatenation expression is `NULL`, the result string
will be `NULL`.

The `||` operator is typically used to concatenate strings to build a value passed
as parameters to a function.

The `||` concatenation operator has different formatting rules as when using the
the comma in instructions such as [`LET`](0701-let.md "The LET statement assigns values to variables.").
The comma implies fixed-size formatting for numeric and date/time values, while `||`
will compact and trim the values.

The `||` operator has a lower precedence than arithmetic operators; For example,
`a || b + c` is equivalent to `(a||(b+c))`.

However, the `||` operator has a higher precedence than [`LIKE`](0607-like.md "The LIKE operator returns TRUE if a string matches a given mask."), [`MATCHES`](0608-matches.md "The MATCHES operator returns TRUE if a string matches a given mask.") and [`USING`](0634-using.md "The USING operator converts date and numeric values to a string based on a formatting mask.") operators; For example, `v
USING "&&" || "&"` is equivalent to `v USING ("&&" ||
"&")` which is the same as `v USING "&&&"`.

## Example

Building strings to be passed as parameters to functions:

```
MAIN
    CALL print_message( 5, "Current date: " || TODAY )
    DISPLAY "Length: ", length( "ab" || "cdef" )
END MAIN

FUNCTION print_message(l TINYINT, s STRING)
    DISPLAY "Level: ", l, " ", s
END FUNCTION
```

Output:

```
Level:    5 Current date: 12/22/2018
Length:           6
```

Order of precedence
test:

```
MAIN
    DISPLAY 78 || 1 + 8                         -- 789
    DISPLAY 78 || 3 * 3                         -- 789
    DISPLAY 789 USING "&&" || "&"               -- 789
    DISPLAY (34 USING "&&") || "&"              -- 34&
    DISPLAY "abc" MATCHES "ab" || "c"           --      1 (TRUE)
    DISPLAY ("ab" MATCHES "ab") || "c"          -- 1c
    DISPLAY "a" MATCHES "a" || "&&"             --      0 (FALSE)
    DISPLAY 23 USING "a" MATCHES "a" || "&&"    -- *  (format = "0")
    DISPLAY 23 USING ("a" MATCHES "a") || "&&"  -- 123
END MAIN
```

Formatting differences between `||` and `,`
(comma):

```
MAIN
    DISPLAY "Date: ", TODAY, " day num: ", DAY(TODAY), " Pi=", 3.1415
    DISPLAY "Date: "||TODAY||" day num: "||DAY(TODAY)||" Pi="||3.1415
END MAIN
```

Output:

```
Date: 12/22/2018 day num:     22 Pi= 3.1415
Date: 12/22/2018 day num: 22 Pi=3.1415
```

## Related links

**Related concepts**  

[DISPLAY (to stdout)](../11_user-interface/1878-display-to-stdout.md "The DISPLAY instruction displays text in line mode to the standard output channel.")
