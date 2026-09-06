---
title: "ON EVERY ROW"
source: "fgl-topics/c_fgl_reports_ON_EVERY_ROW.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > ON EVERY ROW"
type: "concept"
---

# ON EVERY ROW

> Defines printing commands for each row in a report.

The `ON EVERY ROW` control block specifies the action to be taken by the runtime
system for every input record that is passed to the report definition.

The runtime system executes the statements within the `ON EVERY ROW` control
block for each new input record that is passed to the report. The following example is from a
report that lists all the customers, their addresses, and their telephone numbers across the
page:

```
  ON EVERY ROW
    PRINT r.fname, " ", r.lname, " ",
         r.address1, " ", r.cust_phone
```

The runtime system delays processing the `PAGE HEADER` control block (or the
`FIRST PAGE HEADER` control block, if it exists) until it encounters the first
`PRINT`, `SKIP`, or `NEED` statement in the
`ON EVERY ROW` control block.

If a `BEFORE GROUP OF` control block is triggered by a change in the value of a
variable, the runtime system executes all appropriate `BEFORE GROUP OF` control
blocks (in the order of their priority) before it executes the `ON EVERY ROW`
control block. Similarly, if execution of an `AFTER GROUP OF` control block is
triggered by a change in the value of a variable, the runtime system executes all appropriate
`AFTER GROUP OF` control blocks (in the reverse order of their priority) before it
executes the `ON EVERY ROW` control block.

## Related links

**Related concepts**  

[PRINT](2491-print.md "Formats and prints a row of data in a report routine.")

[PRINTX](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.")
