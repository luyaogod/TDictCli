---
title: "PAGE TRAILER"
source: "fgl-topics/c_fgl_reports_PAGE_TRAILER.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > PAGE TRAILER"
type: "concept"
---

# PAGE TRAILER

> Defines the printing commands for the tail of all pages of a report.

The `PAGE TRAILER` control block specifies what information, if any, appears at
the bottom of each page of output from the report.

The runtime system executes the statements in the `PAGE TRAILER` control block
before the `PAGE HEADER` control block when a new page is needed. New pages can be
initiated by any of the following conditions:

- [`PRINT`](2491-print.md "Formats and prints a row of data in a report routine.") attempts to print on a page
  that is already full.
- [`SKIP TO TOP OF PAGE`](2495-skip.md "Skips a given number of lines in a report.") is
  executed.
- [`SKIP n LINES`](2495-skip.md "Skips a given number of lines in a report.")
  specifies more lines than are available on the current page.
- [`NEED`](2493-need.md "Specifies the number of rows needed in a report section.") specifies more lines than are
  available on the current page.

You can use the [`PAGENO`](2498-pageno.md "Contains the current page number in a report.") operator
in a `PRINT` statement within a `PAGE TRAILER` control block to
automatically display the page number at the bottom of every page, as in this example:

```
PAGE TRAILER
   PRINT COLUMN 28, PAGENO USING "page <<<<"
```

The `BOTTOM MARGIN` specification (in the `OUTPUT` section ) affects how close
to the bottom of the page the output displays the page trailer.

> **Important:**
>
> The restrictions that apply to [`FIRST PAGE HEADER`](2481-first-page-header.md "Defines the printing commands for the first page of a report.")
> also apply to `PAGE TRAILER`.

## Related links

**Related concepts**  

[PAGE HEADER](2482-page-header.md "Defines the printing commands for the top of all pages of a report.")

[Report operators](2496-report-operators.md "Report operators can be used to print dynamic report information.")
