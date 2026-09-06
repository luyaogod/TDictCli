---
title: "FORMAT section in REPORT"
source: "fgl-topics/c_fgl_reports_The_FORMAT_Section.html"
breadcrumb: "Reports > The report routine > FORMAT section in REPORT"
type: "concept"
---

# FORMAT section in REPORT

> Defines the formatting directives inside a report routine.

## Syntax

Default format:

```
 FORMAT EVERY ROW
```

Custom format:

```
 FORMAT
  control-block
     [ report-statement
     | report-only-fgl-statement
     | sql-statement
     ]
   [...]
  [...]
```

where control-block can be one of:

```
{
[ FIRST ] PAGE HEADER
|  ON EVERY ROW
|  BEFORE GROUP OF report-variable
|  AFTER GROUP OF report-variable
|  PAGE TRAILER 
|  ON LAST ROW
}
```

1. report-statement is any report-specific instruction.
2. report-only-fgl-statement is any language instruction
   supported in the report routine.
3. sql-statement is any SQL statement supported by the
   language.
4. report-variable is the name of a variable declared
   in the DEFINE section.

## Usage

A report definition must contain a `FORMAT` section.

The `FORMAT` section determines how the output from the report will
look. It works with the values that are passed to the `REPORT` program
block through the argument list or with global or module variables in each record of
the report. In a source file, the `FORMAT` section begins with the
`FORMAT` keyword and ends with the `END REPORT` keywords.

The `FORMAT` section is made up of the following control blocks:

- `FIRST PAGE HEADER`
- `PAGE HEADER`
- `PAGE TRAILER`
- `BEFORE GROUP OF`
- `AFTER GROUP OF`
- `ON EVERY ROW`
- `ON LAST ROW`

If you use the `FORMAT EVERY ROW`, no other statements or control
blocks are valid. The `EVERY ROW` keywords specify a default output format,
including every input record that is passed to the report.

Control blocks define the structure of a report by specifying one or more statements
to be executed when specific parts of the report are processed.

If a report driver includes `START REPORT` and `FINISH REPORT`
statements, but no data records are passed to the report, no control blocks are executed.
That is, unless the report executes an `OUTPUT TO REPORT` statement that
passes at least one input record to the report; then neither the `FIRST PAGE
HEADER` control block nor any other control block is executed

Apart from `BEFORE GROUP OF` and `AFTER GROUP OF`, each
control block must appear only once.

More complex `FORMAT` sections
can contain control blocks like `ON EVERY ROW` or `BEFORE
GROUP OF`, which contain statements to execute while the report is being processed.
Control blocks can contain report execution statements and other executable statements.

A control block may invoke most language statements, except those listed in prohibited
statements.

The `BEFORE/AFTER GROUP OF` control blocks can include aggregate functions
to instruct the report engine to automatically compute such values.

A report-statement is a statement specially designed for the report
format section. It cannot be used in any other part of the program.

The sequence in which the `BEFORE GROUP OF` and `AFTER GROUP
OF` control blocks are executed depends on the sort list in the `ORDER
BY` section, regardless of the physical sequence in which these control blocks
appear within the `FORMAT` section.

## Related links

**Related concepts**  

[Prohibited report routine statements](2487-prohibited-report-routine-statements.md "Prohibited report routine statements")

## Child topics

- [FORMAT EVERY ROW](2480-format-every-row.md): Default format specification of a report.
- [FIRST PAGE HEADER](2481-first-page-header.md): Defines the printing commands for the first page of a report.
- [PAGE HEADER](2482-page-header.md): Defines the printing commands for the top of all pages of a report.
- [PAGE TRAILER](2483-page-trailer.md): Defines the printing commands for the tail of all pages of a report.
- [BEFORE/AFTER GROUP OF](2484-before-after-group-of.md): Defines printing commands of row grouping sections within a report.
- [ON EVERY ROW](2485-on-every-row.md): Defines printing commands for each row in a report.
- [ON LAST ROW](2486-on-last-row.md): Defines the printing commands of the last row in a report.
