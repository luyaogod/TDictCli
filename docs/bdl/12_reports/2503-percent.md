---
title: "PERCENT"
source: "fgl-topics/c_fgl_reports_PERCENT.html"
breadcrumb: "Reports > Report aggregate functions > PERCENT"
type: "concept"
---

# PERCENT

> Calculates the percentage of rows matching a condition.

## Syntax

```
[GROUP] PERCENT(*) [ WHERE condition ]
```

1. condition is a boolean expression evaluated to compute the aggregate value.

## Usage

This aggregate report instruction returns the percentage of the total number of records qualified by
the optional `WHERE` condition.

Using the `GROUP` keyword causes
the aggregate instructions to include only data of the current group of records that have the same
value for the variable that you specify in the `AFTER GROUP OF` control block.

This aggregate instruction makes a two-pass report when not using the `GROUP`
keyword and is used in any control block other than `ON LAST ROW`, or when using
the `GROUP PERCENT(*)` anywhere in the report.

## Related links

**Related concepts**  

[Report engine configuration](2508-report-engine-configuration.md "Report engine behavior can be controlled with FGLPROFILE settings.")

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")
