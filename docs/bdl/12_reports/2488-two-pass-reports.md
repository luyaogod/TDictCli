---
title: "Two-pass reports"
source: "fgl-topics/c_fgl_reports_Two_Pass_Reports.html"
breadcrumb: "Reports > Two-pass reports"
type: "concept"
---

# Two-pass reports

> The report engine supports two-pass reports, to order rows automatically.

The one-pass report requires sorted data to be produced by the report driver in order to handle
`BEFORE/AFTER GROUP OF` blocks properly. The two-pass report handles sorts internally
and does not need sorted data from the report driver. During the first pass, the report engine sorts
the data and stores the sorted values in a temporary file in the database. During the second pass,
it calculates any aggregate values and produces output from data in the temporary files.

A report is defined as a two-pass report if it includes any of the following items:

- An `ORDER BY` section without the
  `EXTERNAL` keyword.
- The [`GROUP PERCENT(*)`](2503-percent.md "Calculates the percentage of rows matching a condition.") aggregate
  function anywhere in the report.
- Any [aggregate](2501-report-aggregate-functions.md "Report aggregate functions can be used to compute data.") function that has no
  `GROUP` keyword in any control block other than `ON LAST ROW`.

Two-pass reports create temporary tables. The `FINISH REPORT` statement uses values from
these tables to calculate any global aggregates, and then deletes the tables. Since two-pass
reports create temporary tables, the report engine requires a database connection, and the
database server must support temporary tables with indexes.

Consider avoiding two-pass reports when a regular report is possible.

## Related links

**Related concepts**  

[The report routine](2474-the-report-routine.md "The report routine implements the body of a report, with formatting instructions.")
