---
title: "The report routine"
source: "fgl-topics/c_fgl_reports_Report_Definition.html"
breadcrumb: "Reports > The report routine"
type: "concept"
---

# The report routine

> The report routine implements the body of a report, with formatting instructions.

## Syntax 1 (legacy syntax):

```
[PUBLIC|PRIVATE] REPORT report-name ( parameter-name [,...] )
 [ define-section ]
 [ output-section ]
 [ sort-section ]
 [ format-section ] 
END REPORT
```

## Syntax 2 (full typed):

```
[PUBLIC|PRIVATE] REPORT report-name (
     parameter-name data-type
        [ {ATTRIBUTE|ATTRIBUTES} ( attribute [ = "value" ] [,...] ) ]
     [,...]
  )
 [ define-section ]
 [ output-section ]
 [ sort-section ]
 [ format-section ] 
END REPORT
```

where define-section
is:

```
DEFINE variable-definition [,...]
```

where output-section
is:

```
OUTPUT 
[
  REPORT TO
  {
   SCREEN
 | PRINTER
 | [ FILE ] filename 
 | PIPE program [ IN FORM MODE |  IN LINE MODE ]
   } 
]
[
  [ WITH ]
  [ LEFT MARGIN m-left ] 
  [ RIGHT MARGIN m-right] 
  [ TOP MARGIN m-top ] 
  [ BOTTOM MARGIN m-bottom ] 
  [ PAGE LENGTH m-length ] 
  [ TOP OF PAGE c-top]
 ]
```

where sort-section
is:

```
 ORDER [ EXTERNAL ] BY report-variable [,...]
```

where format-section
is:

```
FORMAT EVERY ROW
```

or:

```
 FORMAT
   control-block
   [ report-only-fgl-statement
   | sql-statement
   | report-statement ]
   [...]
 [...]
```

where control-block can be one
of:

```
{
[ FIRST ]  PAGE HEADER
|  ON EVERY ROW
|  BEFORE GROUP OF report-variable
|  AFTER GROUP OF report-variable
|  PAGE TRAILER
|  ON LAST ROW
}
```

1. variable-definition follows the `DEFINE` instruction syntax
   and declares report-variables.
2. report-variable is the name of a variable declared in the
   `DEFINE` section.
3. report-only-fgl-statement is a subset of all the regular language
   statements.
4. sql-statement is a valid static SQL statement.

## Usage

The report definition formats input records. Like the `FUNCTION` or
`MAIN` statement, it is a program block that can have the scope of local variables.
It is not, however, a function; it is not reentrant, and `CALL` cannot invoke it. The
report definition receives data from its driver in sets called input records. These records can
include program records, but other data types are also supported. Each input record is formatted and
printed as specified by control blocks and statements within the report definition. Most statements
and functions can be included in a report definition, and certain specialized statements and
operators for formatting output can appear only in a report definition.

Like `MAIN` or `FUNCTION`, the report definition must appear outside
any other program block. It must begin with the `REPORT` statement and must end with
the `END REPORT` keywords.

Some statements are prohibited in a
`REPORT` routine control block. For example, it is not possible to use
`CONSTRUCT`, `INPUT`, `DEFER`, `DEFINE`,
`REPORT`, `RETURN` instructions in a control block of a report.

By default, report routines are public; they can be called by any other module of the program. If
a report routine is only used by the current module, you may want to hide that routine from other
modules, to make sure that it will not be called by mistake. To keep a report routine local to the
module, add the `PRIVATE` keyword before the report header. Private report routines
are only hidden to external modules; all functions of the current module can still call local
private report routines.

The define section declares the data types of local variables used within the
report, and of any variables (the input records) that are passed as arguments to the report by the
calling statement. Reports without arguments or local variables do not require a `DEFINE` section.

The output-section can set margin and page size values, and can also specify
where to send the formatted output. Output from the report consists of successive pages, each
containing a fixed number of lines whose margins and maximum number of characters are fixed.

The sort-section specifies how the rows have to be sorted. The specified sort
order determines the order in which the runtime system processes any `GROUP OF`
control blocks in the `FORMAT` section.

The format-section is required. It specifies the appearance of the report,
including page headers, page trailers, and aggregate functions of the data. It can also contain
control blocks that specify actions to take before or after specific groups of rows are processed.
(Alternatively, it can produce a default report by only specifying `FORMAT EVERY
ROW`).

## Related links

**Related concepts**  

[OUTPUT section in REPORT](2477-output-section-in-report.md "Specifies report destination and page format options.")

[ORDER BY section in REPORT](2478-order-by-section-in-report.md "Forces a sort order of unsorted data rows in reports.")

[FORMAT section in REPORT](2479-format-section-in-report.md "Defines the formatting directives inside a report routine.")

## Child topics

- [The report prototype](2475-the-report-prototype.md)
- [DEFINE section in REPORT](2476-define-section-in-report.md): Defines report parameters and local variables.
- [OUTPUT section in REPORT](2477-output-section-in-report.md): Specifies report destination and page format options.
- [ORDER BY section in REPORT](2478-order-by-section-in-report.md): Forces a sort order of unsorted data rows in reports.
- [FORMAT section in REPORT](2479-format-section-in-report.md): Defines the formatting directives inside a report routine.
- [Prohibited report routine statements](2487-prohibited-report-routine-statements.md)
