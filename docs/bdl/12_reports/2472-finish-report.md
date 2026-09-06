---
title: "FINISH REPORT"
source: "fgl-topics/c_fgl_reports_FINISH_REPORT.html"
breadcrumb: "Reports > The report driver > FINISH REPORT"
type: "concept"
---

# FINISH REPORT

> The FINISH REPORT instruction finalizes a report execution.

## Syntax

```
FINISH REPORT report-name
```

1. report-name is the name of the report to be ended.

## Usage

`FINISH REPORT` closes the report driver. Therefore, it must be the last statement
in the report driver and must follow a `START REPORT` statement that specifies the
name of the same report.

`FINISH REPORT` must be the last statement in the report driver.

`FINISH REPORT` does the following:

1. Completes the second pass, if report is a two-pass report. These 'second pass' activities handle
   the calculation and output of any aggregate values that are based on all the input records in
   the report, such as `COUNT(*)` or `PERCENT(*)` with no
   `GROUP` qualifier.
2. Executes any [`AFTER GROUP
   OF`](2484-before-after-group-of.md "Defines printing commands of row grouping sections within a report.") control blocks.
3. Executes any [`PAGE HEADER`](2482-page-header.md "Defines the printing commands for the top of all pages of a report."),
   [`ON LAST ROW`](2486-on-last-row.md "Defines the printing commands of the last row in a report."), and [`PAGE TRAILER`](2483-page-trailer.md "Defines the printing commands for the tail of all pages of a report.") control blocks to
   complete the report.
4. Copies data from the output buffers of the report to the destination.
5. Closes the Select cursor on any temporary table that was created to order the input records or
   to perform aggregate calculations.

If the `FINISH REPORT` instruction fails, the runtime system will raise the error
[-8140](../15_library-reference/4483-genero-bdl-errors.md), with the reason for the
failure.

## Related links

**Related concepts**  

[TERMINATE REPORT](2473-terminate-report.md "The TERMINATE REPORT instruction cancels a report execution.")

[EXIT REPORT](2490-exit-report.md "Cancels the report processing.")

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")
