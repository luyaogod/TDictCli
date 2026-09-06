---
title: "SPACES"
source: "fgl-topics/c_fgl_operators_SPACES.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > SPACES"
type: "concept"
---

# SPACES

> The SPACES operator returns a character string with blanks.

## Syntax

```
int-expr SPACES
```

1. int-expr is an integer expression.
2. `SPACE` (without S) is an alias for this operator.

## Usage

The `SPACE` operator
is typically used in reports to print spaces to align data
in the report output.

## Example

```
MAIN
  DISPLAY 20 SPACES || "xxx"
END MAIN
```

## Related links

**Related concepts**  

[Reports](../12_reports/2461-reports.md "Reports")

[Integer expressions](0595-integer-expressions.md "This section covers integer expression evaluation rules.")
