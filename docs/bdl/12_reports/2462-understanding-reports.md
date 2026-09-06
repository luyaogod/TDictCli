---
title: "Understanding reports"
source: "fgl-topics/c_fgl_reports_002.html"
breadcrumb: "Reports > Understanding reports"
type: "concept"
---

# Understanding reports

> A report can arrange and format the data depending on the instructions used and display the output on the screen, send it to a printer, or store it as a file for future use.

To implement a report, a program must include two distinct components:

- The report driver specifies what data the report includes.
- The report routine formats the data for output.

The report driver retrieves the specified rows from a database, stores their values in program
variables, and sends these - one input record at a time - to the report routine. After the last
input record is received and formatted, the runtime system calculates any aggregate values based
on all the data and sends the entire report to some output device.

![Report driver and database cursor diagram](../_images/REPFig01.jpg)

*Report driver and database cursor*

By separating the two tasks of data retrieval and data formatting, the runtime system simplifies
the production of recurrent reports and makes it easy to apply the same report format to different
data sets.

The report engine supports the following features:

- The option to display report output to the screen, to a printer, to a file, or to a SAX handler
  to transform the output following XML standards. When using the GDC, it is possible to send to
  report to the local printer configured in GDC, by setting the DBPRINT environment variable to the
  value `"FGLSERVER"`.
- Full control of page layout, including first page header and generic page headers, page
  trailers, columnar presentation, and row grouping.
- Creating the report either from the rows returned by a cursor or from input records assembled
  from any other source, such as output from several different `SELECT` statements
  through the report driver.
- Control blocks to manipulate data from a database cursor on a row-by-row basis, either before
  or after the row is formatted by the report.
- Aggregate functions that can calculate frequencies, percentages, sums, averages, minimum, and
  maximum values.
- The `USING` operator and other built-in functions and operators for formatting
  and displaying information in output from the report.
- The `WORDWRAP` operator to format long character strings that occupy multiple
  lines of output from the report.
- The option to execute other language statements while generating a report.
- Stop a report in the report definition with `EXIT REPORT` or `TERMINATE
  REPORT`.

The report engine supports one-pass reports and two-pass reports:

- The one-pass mode requires sorted data to be produced by the report driver in order to handle
  row grouping with the `BEFORE GROUP` / `AFTER GROUP` blocks.
- The two-pass record handles sort automatically and does not need sorted data from the report
  driver. During the first pass, the report engine sorts the data and stores the sorted values in a
  temporary table in the current database. During the second pass, it calculates aggregate values
  and produces output from data stored in the temporary table.

## Related links

**Related concepts**  

[The report driver](2469-the-report-driver.md "The report driver retrieves data, starts the report engine and sends the data (as input records) to be formatted by the REPORT routine.")

[The report routine](2474-the-report-routine.md "The report routine implements the body of a report, with formatting instructions.")

[XML output for reports](2463-xml-output-for-reports.md "For better integration with external tools based on XML standards, reports can produce XML output.")
