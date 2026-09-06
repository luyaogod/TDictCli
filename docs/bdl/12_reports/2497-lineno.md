---
title: "LINENO"
source: "fgl-topics/c_fgl_reports_LINENO.html"
breadcrumb: "Reports > Report operators > LINENO"
type: "concept"
---

# LINENO

> Contains the current line number in a report.

## Syntax

```
LINENO
```

## Usage

This operator takes no operand but returns the value of the line number of the report line
that is currently printing.

The report engine calculates the line number by calculating the number of lines from the
top of the current page, including the `TOP MARGIN`.

## Example

In this example, a `PRINT` statement instructs the report to calculate and
display the current line number, beginning in the tenth character position after the left
margin:

```
  ON EVERY ROW
    IF LINENO > 9 THEN
      PRINT COLUMN 10, "Line:", LINENO USING "<<<"
    END IF
```

## Related links

**Related concepts**  

[START REPORT](2470-start-report.md "The START REPORT instruction initializes a report execution.")
