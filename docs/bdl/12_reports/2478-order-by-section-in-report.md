---
title: "ORDER BY section in REPORT"
source: "fgl-topics/c_fgl_reports_The_ORDER_BY_Section.html"
breadcrumb: "Reports > The report routine > ORDER BY section in REPORT"
type: "concept"
---

# ORDER BY section in REPORT

> Forces a sort order of unsorted data rows in reports.

## Syntax

```
ORDER [ EXTERNAL ] BY report-variable [ DESC | ASC ] [,...]
```

1. report-variable identifies one of the variables passed
   to the report routine to be used for sorting rows.

## Usage

When grouping rows in a report, values that the report definition receives from
the report driver are significant in determining how `BEFORE GROUP OF/AFTER GROUP
OF` control blocks will process the data in the formatted report
output.

The `ORDER BY` section defines how the variables of the input
records are to be sorted. It is required if the report driver does not send sorted
data to the report. The specified sort order determines the order in which the
runtime system processes any `GROUP OF` control blocks in the
`FORMAT`
section.

If you omit the `ORDER BY` section, the runtime system processes
input records in the order received from the report driver and processes any
`GROUP OF` control blocks in their order of appearance in the
`FORMAT` section. If records are not sorted in the report driver,
the `GROUP OF` control blocks might be executed at random intervals
(that is, after any input record) because unsorted values tend to change from record
to record.

If you specify only one variable in the `GROUP OF` control blocks,
and the input records are already sorted in sequence on that variable by the
`SELECT` statement, you do not need to include an `ORDER BY`
section in the report.

Specify `ORDER EXTERNAL BY` if the input records have already been
sorted by the `SELECT` statement used by the report driver. The list
of variables after the keywords `ORDER EXTERNAL BY` control the execution
order of `GROUP BY` control blocks.

Without the `EXTERNAL` keyword, the report becomes a two-pass report,
meaning that the report engine processes the set of input records twice. During the
first pass, the report engine sorts the data and stores the sorted values in a temporary
table in the database. During the second pass, it calculates any aggregate values and
produces output from data in the temporary files.

With the `EXTERNAL` keyword, the report engine only needs to make a
single pass through the data: it does not need to build the temporary table in the
database for sorting the data. However, If the report routine contains aggregations
functions such as [`GROUP PERCENT(*)`](2503-percent.md "Calculates the percentage of rows matching a condition."),
the report will become a two-pass report because such aggregation function needs all
rows to compute the value.

The `DESC` or `ASC` clause defines the sort order.

## Related links

**Related concepts**  

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")

[OUTPUT section in REPORT](2477-output-section-in-report.md "Specifies report destination and page format options.")

[The report driver](2469-the-report-driver.md "The report driver retrieves data, starts the report engine and sends the data (as input records) to be formatted by the REPORT routine.")
