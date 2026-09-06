---
title: "START REPORT"
source: "fgl-topics/c_fgl_reports_START_REPORT.html"
breadcrumb: "Reports > The report driver > START REPORT"
type: "concept"
---

# START REPORT

> The START REPORT instruction initializes a report execution.

## Syntax

```
START REPORT report-routine
  [ TO to-clause ]
  [ WITH dimension-option [,...]]
```

where to-clause is one of:

```
{ SCREEN
| PRINTER
| [FILE] filename
| PIPE program [ IN FORM MODE |  IN LINE MODE ]
| XML HANDLER sax-handler-object
| OUTPUT destination-expr [ DESTINATION { program  | filename} ]
}
```

where dimension-option is one of:

```
{ LEFT MARGIN = m-left
| RIGHT MARGIN = m-right
| TOP MARGIN = m-top
| BOTTOM MARGIN = m-bottom
| PAGE LENGTH = m-length
| TOP OF PAGE = c-top
}
```

1. report-routine is the name of the `REPORT` routine.
2. filename is a string expression specifying the file that receives
   report output.
3. program is a string expression specifying a program, a shell script,
   or a command line to receive report output.
4. destination-expr is a string expression that specifies one of:
   `SCREEN`, `PRINTER`, `FILE`, `PIPE`,
   `PIPE IN LINE MODE`, `PIPE IN FORM MODE`.
5. sax-handler-object is a variable referencing an `om.SaxDocumentHandler` instance.
6. m-left is the left margin in number of characters. The default is 5.
7. m-right is the right margin in number of characters. The default is 132.
8. m-top is the top margin in number of lines. The default is 3.
9. m-bottom is the bottom margin in number of lines. The default is 3.
10. m-length is the total number of lines on a report page. The default page
    length is 66 lines.
11. c-top is a string that defines the page-eject character sequence.

## Usage

The `START REPORT` statement initializes a report. The instruction allows
you to specify the report output destination, the page dimension, and margins.

`START REPORT` typically precedes a loop instruction such as
`FOR`, `FOREACH`, or `WHILE` in which
`OUTPUT TO REPORT` feeds the report routine with data. After the loop terminates,
`FINISH REPORT` completes the processing of the output.

```
DEFINE file_name VARCHAR(200), page_size INTEGER
...
START REPORT myrep
  TO FILE file_name 
  WITH PAGE LENGTH = page_size
```

If a `START REPORT` statement references a report that is already running,
the report is re-initialized; any output might be unpredictable.

## Output specification

The `TO` clause can be used to specify a destination for output. If you omit
the `TO` clause, the Genero runtime system sends report output to the destination
specified in the report routine definition. If the report routine does not define an
`OUTPUT` clause, the report output is sent by default to the report viewer
when in GUI mode, or to the screen when in TUI mode.

Report output can be specified dynamically as follows:

- The `TO FILE` option can specify the filename as
  a character variable that is assigned at runtime.
- The `TO PIPE` option can specify the program as
  a character variable that is assigned at runtime.
- The `TO OUTPUT` option can specify the report output with a string expression,
  described later in detail.

The `SCREEN` option specifies that output is to the report window. The way the
report is displayed to the end user depends on whether you are in TUI mode or GUI mode. In TUI mode,
the report output displays to the terminal screen. In GUI mode, the report output displays in a
dedicated pop-up window called the Report Viewer. When using `REPORT TO SCREEN` in
TUI mode, set a `PAGE LENGTH` no larger than the terminal can display, and include
`PAUSE` statements in the `FORMAT` section, to let the end user see
the output on the screen.

The `PRINTER` option instructs the runtime system to output the report to the
device or program defined by the [DBPRINT](../07_configuration/0514-dbprint.md "Defines the print device to be used by reports.")
environment variable.

When using the `TO PRINTER` clause, if the DBPRINT environment variable is set to
the value `"FGLSERVER"`, the document is sent to the printer configured in the Genero
Desktop Client (GDC).

When using the `FILE` option, you can specify a filename as the report destination.
Output will be sent to the specified file. If the file exists, its content will be overwritten by
the new report output. The `FILE` keyword is optional, but it's best to include it to
make your code more readable.

The `PIPE` option defines a program, shell script, or command line to which the
report output must be sent, using the standard input channel. When using the TUI mode, you can use
the `IN [LINE|FORM] MODE` option to specify whether the program is in line mode or in
formatted mode when report output is sent to a pipe. For more details about the `IN LINE
MODE` / `IN FORM MODE` clause, see [RUN](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.").

The `TO OUTPUT` option allows you to specify one of the output options dynamically
at runtime. The character string expression must be one of: `"SCREEN"`,
`"PRINTER"`, `"FILE"`, `"PIPE"`, `"PIPE IN LINE
MODE"`, `" PIPE IN FORM MODE"`. If the expression specifies
`"FILE"` or `"PIPE"`, you can also specify a
filename or program in a character variable following the
`DESTINATION` keyword.

The `XML HANDLER` option indicates that the report output will be generated as
XML and redirected to a SAX-document handler. When using XML output, the report result can be
shown in the Genero Report Engine installed on the front-end workstation. See XML output for more
details.

When `START REPORT` without a destination is used inside a `REPORT`
routine producing XML, the sub-report inherits the XML output target of the parent, and sub-report
nodes will be merged into the parent XML output.

## Page dimensions specification

The `WITH` clause defines the dimensions of each report page and the left, top,
right and bottom margins. The values corresponding to a margin and page length must be valid integer
expressions. The margins can be defined in any order, but a comma "," is required to separate two
page dimensions options.

- The `LEFT MARGIN` clause defines the number of blank spaces to include at the
  start of each new line of output. The default is 5.
- The `RIGHT MARGIN` clause defines the total number of characters in each line
  of output, including the left margin. If you omit this but specify `FORMAT EVERY ROW`,
  the default is 132.
- The `TOP MARGIN` clause specifies how many blank lines appear above the first
  line of text on each page of output. The default is 3.
- The `BOTTOM MARGIN` clause specifies how many blank lines follow the last line
  of output on each page. The default is 3.
- The `PAGE LENGTH` clause specifies the total number of lines on each page,
  including data, the margins, and any page headers or page trailers from the `FORMAT`
  section. The default page length is 66 lines.

In addition to the page dimension options, the `TOP OF PAGE` clause can specify
a page-eject sequence for a printer. On some systems, specifying this value can reduce the time
required for a large report to produce output, because [`SKIP TO TOP OF PAGE`](2495-skip.md "Skips a given number of lines in a report.") can substitute this
value for multiple line feeds.

> **Tip:**
>
> To produce reports with no page size limit, set margins to zero and use a page
> length of 1.

## Related links

**Related concepts**  

[The report routine](2474-the-report-routine.md "The report routine implements the body of a report, with formatting instructions.")

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")

[XML output for reports](2463-xml-output-for-reports.md "For better integration with external tools based on XML standards, reports can produce XML output.")

[The SaxDocumentHandler class](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events.")
