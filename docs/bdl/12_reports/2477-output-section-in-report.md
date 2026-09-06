---
title: "OUTPUT section in REPORT"
source: "fgl-topics/c_fgl_reports_The_OUTPUT_Section.html"
breadcrumb: "Reports > The report routine > OUTPUT section in REPORT"
type: "concept"
---

# OUTPUT section in REPORT

> Specifies report destination and page format options.

## Syntax

```
 OUTPUT 
[
  REPORT TO
  {
      SCREEN
    | PRINTER
    | [ FILE ] filename 
    | PIPE [ IN FORM MODE | IN LINE MODE ] program  
   } 
]
[
  [ LEFT MARGIN m-left ] 
  [ RIGHT MARGIN m-right ] 
  [ TOP MARGIN m-top ] 
  [ BOTTOM MARGIN m-bottom ] 
  [ PAGE LENGTH m-length ] 
  [ TOP OF PAGE c-top] 
]
```

1. program defines the name of a program, shell script, command receiving the
   output.
2. filename defines the file which receives the output of the report.
3. m-left is the left margin in number of characters. The default is 5.
4. m-right is the right margin in number of characters. The default is 132.
5. m-top is the top margin in number of lines. The default is 3.
6. m-bottom is the bottom margin in number of lines. The default is 3.
7. m-length is the total number of lines on a report page. The default page
   length is 66 lines.
8. c-top is a string that defines the page-eject character sequence.

## Usage

The `OUTPUT` section can specify the destination and dimensions for output from
the report and the page-eject sequence for the printer. If you omit the `OUTPUT`
section, the report uses default values to format each page. This section is superseded by any
corresponding `START REPORT` specifications.

The `OUTPUT` section can direct the output from the report to a printer, file, or
pipe, and can initialize the page dimensions and margins of report output. If
`PRINTER` is specified, the [DBPRINT](../07_configuration/0514-dbprint.md "Defines the print device to be used by reports.") environment variable specifies which printer.

The `START REPORT` statement of the report driver can override all of these
specifications by assigning another destination in its `TO` clause or by assigning
other dimensions, margins, or another page-eject sequence in the `WITH` clause.

As the size specifications for the dimensions and margins of a page of report output that the
`OUTPUT` section can specify must be literal integers, consider defining page
dimensions in the `START REPORT` statement, where you can use variables to assign
these values dynamically at runtime.

> **Tip:**
>
> To produce reports with no page size limit, set margins to zero and use a page
> length of 1.

## Related links

**Related concepts**  

[START REPORT](2470-start-report.md "The START REPORT instruction initializes a report execution.")
