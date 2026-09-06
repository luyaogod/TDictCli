---
title: "ON LAST ROW"
source: "fgl-topics/c_fgl_reports_ON_LAST_ROW.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > ON LAST ROW"
type: "concept"
---

# ON LAST ROW

> Defines the printing commands of the last row in a report.

The `ON LAST ROW` control block specifies the action that the runtime
system is to take after it processes the last input record that was passed to the report
definition and encounters the `FINISH REPORT` statement.

The statements in the `ON LAST ROW` control block are executed after
the statements in the `ON EVERY ROW` and `AFTER GROUP OF`
control blocks if these blocks are present.

When the runtime system processes the statements in an `ON LAST ROW`
control block, the variables that the report is processing still have the values from the
final record that the report processed. The `ON LAST ROW` control block
can use aggregate functions to display report totals.

## Related links

**Related concepts**  

[PRINT](2491-print.md "Formats and prints a row of data in a report routine.")

[PRINTX](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.")
