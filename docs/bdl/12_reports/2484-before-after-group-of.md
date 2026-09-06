---
title: "BEFORE/AFTER GROUP OF"
source: "fgl-topics/c_fgl_reports_BEFORE_AFTER_GROUP_OF.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > BEFORE/AFTER GROUP OF"
type: "concept"
---

# BEFORE/AFTER GROUP OF

> Defines printing commands of row grouping sections within a report.

The `BEFORE/AFTER GROUP OF` control blocks specify what action the runtime
system takes respectively before or after it processes a group of input records. Group hierarchy
is determined by the `ORDER BY` specification in the `SELECT`
statement or in the report definition.

A group of records is all of the input records that contain the same value for the variable
whose name follows the `AFTER GROUP OF` keywords. This group variable must be
passed through the report arguments. A report can include no more than one `AFTER
GROUP OF` control block for any group variable.

When the runtime system executes the statements in a `BEFORE/AFTER GROUP OF`
control block, the report variables have the values from the first / last record of the new
group. From this perspective, the `BEFORE/AFTER GROUP OF` control block can be
thought of as the "on first / last record of group" control block.

Each `BEFORE GROUP OF` block is executed in order, from highest to lowest
priority, at the start of a report (after any `FIRST PAGE HEADER` or `PAGE HEADER` control blocks, but before
processing the first record) and on these occasions:

- Whenever the value of the group variable changes (after any `AFTER GROUP OF`
  block for the old value completes execution)
- Whenever the value of a higher-priority variable in the sort list changes (after any
  `AFTER GROUP OF` block for the old value completes execution)

The runtime system executes the `AFTER GROUP OF` control block on these
occasions:

- Whenever the value of the group variable changes.
- Whenever the value of a higher-priority variable in the sort list changes.
- At the end of the report (after processing the last input record but before the runtime
  system executes any `ON LAST ROW` or `PAGE TRAILER` control
  blocks). In this case, each `AFTER GROUP OF` control block is executed in
  ascending priority.

How often the value of the group variable changes depends in part on whether the input
records have been sorted by the `SELECT` statement:

- If records are already sorted, the `BEFORE/AFTER GROUP OF` block executes
  before the runtime system processes the first record of the group.
- If records are not sorted, the `BEFORE GROUP OF` block might be executed
  after any record because the value of the group variable can change with each record. If no
  `ORDER BY` section is specified, all `BEFORE/AFTER GROUP OF`
  control blocks are executed in the same order in which they appear in the `FORMAT`
  section. The `BEFORE/AFTER GROUP OF` control blocks are designed to work with
  sorted data.

You can sort the records by specifying a sort list in either of the following areas:

- An `ORDER BY` section in the report definition
- The `ORDER BY` clause of the `SELECT` statement in the
  report driver

To sort data in the report definition (with an `ORDER BY` section), make
sure that the name of the group variable appears in both the `ORDER BY` section
and in the `BEFORE GROUP OF` control block.

To sort data in the `ORDER BY` clause of a `SELECT` statement,
perform the following tasks:

- Use the column name in the `ORDER BY` clause of the `SELECT`
  statement as the group variable in the `BEFORE GROUP OF` control block.
- If the report contains `BEFORE` or `AFTER GROUP OF` control
  blocks, make sure that you include an `ORDER EXTERNAL BY` section in the report
  to specify the precedence of variables in the sort list.

If you specify sort lists in both the report driver and the report definition, the sort list
in the `ORDER BY` section of the `REPORT` takes precedence. When
the runtime system starts to generate a report, it first executes the `BEFORE GROUP OF`
control blocks in descending order of priority before it executes the `ON EVERY ROW`
control block. If the report is not already at the top of the page, the `SKIP TO TOP OF
PAGE` statement in a `BEFORE GROUP OF` control block causes the output for
each group to start at the top of a page.

If the sort list includes more than one variable, the runtime system sorts the records by
values in the first variable (highest priority). Records that have the same value for the first
variable are then ordered by the second variable and so on until records that have the same values
for all other variables are ordered by the last variable (lowest priority) in the sort list.

The `ORDER BY` section determines the order in which the runtime system processes
`BEFORE GROUP OF` and `AFTER GROUP OF` control blocks. If you omit
the `ORDER BY` section, the runtime system processes any `GROUP OF`
control blocks in the lexical order of their appearance within the `FORMAT` section.

If you include an `ORDER BY` section, and the `FORMAT` section
contains more than one `BEFORE GROUP OF` or `AFTER GROUP OF` control
block, the order in which these control blocks are executed is determined by the sort list in the
`ORDER BY` section. In this case, their order within the `FORMAT`
section is not significant because the sort list overrides their lexical order.

The runtime system processes all the statements in a `BEFORE GROUP OF` or
`AFTER GROUP OF` control block on these occasions:

- Each time the value of the current group variable changes.
- Each time the value of a higher-priority variable changes.

How often the value of the group variable changes depends in part on whether the input records
have been sorted. If the records are sorted, `AFTER GROUP OF` executes after the
runtime system processes the last record of the group of records; `BEFORE GROUP OF`
executes before the runtime system processes the first records with the same value for the group
variable. If the records are not sorted, the `BEFORE GROUP OF` and `AFTER
GROUP OF` control blocks might be executed before and after each record because the value of
the group variable might change with each record. All the `AFTER GROUP OF` and
`BEFORE GROUP OF` control blocks are executed in the same lexical order in which they
appear in the `FORMAT` section.

In the `AFTER GROUP OF` control block, you can include the `GROUP`
keyword to qualify [aggregate report
functions](2501-report-aggregate-functions.md "Report aggregate functions can be used to compute data.") like `GROUP SUM()` in the following
example:

```
  AFTER GROUP OF r.order_num
    PRINT r.order_date, 7 SPACES,
       r.order_num USING"###&", 8 SPACES,
         r.ship_date, " ",
         GROUP SUM(r.total_price) USING"$$$$,$$$,$$$.&&"
  AFTER GROUP OF r.customer_num
    PRINT 42 SPACES, "-------------------"
    PRINT 42 SPACES, GROUP SUM(r.total_price) USING"$$$$,$$$,$$$.&&"
```

Using the `GROUP` keyword to qualify an aggregate function is only valid within
the `AFTER GROUP OF` control block. It is not valid, for example, in the
`BEFORE GROUP OF` control block.

After the last input record is processed, the runtime system executes the `AFTER GROUP
OF` control blocks before it executes the `ON LAST ROW` control block.

## Related links

**Related concepts**  

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")

[PRINT](2491-print.md "Formats and prints a row of data in a report routine.")

[PRINTX](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.")
