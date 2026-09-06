---
title: "FIRST PAGE HEADER"
source: "fgl-topics/c_fgl_reports_FIRST_PAGE_HEADER.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT > FIRST PAGE HEADER"
type: "concept"
---

# FIRST PAGE HEADER

> Defines the printing commands for the first page of a report.

This control block specifies the action that the runtime system takes before it begins processing
the first input record, and print the lines at the top of the first page in a report.

As the runtime system executes the `FIRST PAGE HEADER` control block before
generating any output, you can use this control block to initialize variables that you use in
the `FORMAT` section.

If a report driver includes `START REPORT` and `FINISH REPORT`
statements, but no data records are passed to the report, this control block is not executed.
That is, unless the report executes an `OUTPUT TO REPORT` statement that passes
at least one input record to the report, neither the `FIRST PAGE HEADER` control
block nor any other control block is executed.

As its name implies, you can also use a `FIRST PAGE HEADER` control block to
produce a title page as well as column headings. On the first page of a report, this control
block overrides any `PAGE HEADER` control block. That is, if both a `FIRST
PAGE HEADER` and a `PAGE HEADER` control block exist, output from the
first appears at the beginning of the first page, and output from the second begins all
subsequent pages.

The `TOP MARGIN` (set in the `OUTPUT` section) determines how
close the header appears to the top of the page.

Consider the following notes when programming the `FIRST PAGE HEADER` control
block:

1. You cannot include a `SKIP` integer `LINES` statement inside
   a loop within this control block.
2. The `NEED` statement is not valid within this control block.
3. If you use an `IF`…`THEN`…`ELSE` statement
   within this control block, the number of lines displayed by any `PRINT` statements
   following the `THEN` keyword must be equal to the number of lines displayed by any
   `PRINT` statements following the `ELSE` keyword.
4. If you use a `CASE`, `FOR`, or `WHILE` statement
   that contains a `PRINT` statement within this control block, you must terminate
   the `PRINT` statement with a semicolon ( ; ). The semicolon suppresses any
   `LINEFEED` characters in the loop, keeping the number of lines in the header
   constant from page to page.
5. You cannot use a `PRINT` filename statement to read and display text from a
   file within this control block

Corresponding restrictions also apply to `CASE`, `FOR`,
`IF`, `NEED`, `SKIP`, `PRINT`,
and `WHILE` statements in `PAGE HEADER` and `PAGE
TRAILER` control blocks.

## Related links

**Related concepts**  

[Report instructions](2489-report-instructions.md "The report instruction listed in this section can appear only in control blocks of the FORMAT section of a report routine.")
