---
title: "The report prototype"
source: "fgl-topics/c_fgl_reports_The_Report_Prototype.html"
breadcrumb: "Reports > The report routine > The report prototype"
type: "concept"
description: "When defining a report routine, the report name must immediately follow the REPORT keyword. The name must be unique among function and report names within the program. Its scope is the entire program. ..."
---

# The report prototype

When defining a report routine, the report name must immediately follow the
`REPORT` keyword. The name must be unique among function and
report names within the program. Its scope is the entire program.

The list of formal arguments of the report must be enclosed in parentheses and separated by
commas. These are local variables that store values that the calling routine passes to the
report.

When you call a report, the formal arguments are assigned values from the
argument list of the `OUTPUT TO REPORT` statement. These actual
arguments that you pass must match, in number and position, the formal arguments
of the `REPORT` routine. The data types must be compatible, but
they need not be identical. The runtime system can perform some conversions
between compatible data types.

The names of the actual arguments and the formal arguments do not have to
match.

You must include the following items in the list of formal arguments:

- All the values for each row sent to the report in the following
  cases:
  - If you include an `ORDER BY` section
    or `GROUP PERCENT(*)` function
  - If you use a global aggregate function (one over all rows of the report)
    anywhere in the report, except in the [`ON LAST ROW`](2486-on-last-row.md "Defines the printing commands of the last row in a report.")
    control block
  - If you specify the `FORMAT EVERY ROW`
    default format
- Any variables referenced in the following group control blocks:
  - [`AFTER GROUP
    OF`](2484-before-after-group-of.md "Defines printing commands of row grouping sections within a report.")
  - [`BEFORE GROUP
    OF`](2484-before-after-group-of.md "Defines printing commands of row grouping sections within a report.")

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
