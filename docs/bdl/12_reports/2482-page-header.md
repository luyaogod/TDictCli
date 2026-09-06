---
title: "PAGE HEADER"
source: "fgl-topics/c_fgl_reports_PAGE_HEADER.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > PAGE HEADER"
type: "concept"
---

# PAGE HEADER

> Defines the printing commands for the top of all pages of a report.

This control block is executed whenever a new page is added to the report. The `PAGE
HEADER` control block specifies the action that the runtime takes before it begins
processing each page of the report. It can specify what information, if any, appears at the
top of each new page of output from the report.

The `TOP MARGIN` specification (in the [`OUTPUT`](2477-output-section-in-report.md "Specifies report destination and page format options.") section) affects how
many blank lines appear above the output produced by statements in the `PAGE HEADER`
control block.

You can use the [`PAGENO`](2498-pageno.md "Contains the current page number in a report.") operator
in a [`PRINT`](2491-print.md "Formats and prints a row of data in a report routine.") statement within a
`PAGE HEADER` control block to automatically display the current page number at
the top of every page.

The `FIRST PAGE HEADER`
control block overrides this control block on the first page of a report.

New group values can appear in the `PAGE HEADER` control block when this control
block is executed after a simultaneous end-of-group and end-of-page situation.

The runtime system delays the processing of the `PAGE HEADER` control block until
it encounters the first `PRINT`, [`SKIP`](2495-skip.md "Skips a given number of lines in a report."), or [`NEED`](2493-need.md "Specifies the number of rows needed in a report section.")
statement in the `ON EVERY ROW`, `BEFORE GROUP OF`, or `AFTER
GROUP OF` control block. This order guarantees that any group columns printed in the
`PAGE HEADER` control block have the same values as the columns printed in the
`ON EVERY ROW` control block.

> **Important:**
>
> The restrictions that apply to `FIRST PAGE HEADER` also
> apply to `PAGE HEADER`.

## Related links

**Related concepts**  

[PAGE TRAILER](2483-page-trailer.md "Defines the printing commands for the tail of all pages of a report.")

[Report operators](2496-report-operators.md "Report operators can be used to print dynamic report information.")
