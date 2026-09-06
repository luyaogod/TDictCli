---
title: "DEFINE section in REPORT"
source: "fgl-topics/c_fgl_reports_The_DEFINE_Section.html"
breadcrumb: "Reports > The report routine > DEFINE section in REPORT"
type: "concept"
---

# DEFINE section in REPORT

> Defines report parameters and local variables.

## Syntax

```
DEFINE variable-name type-specification [,...]
```

1. variable-name is an identifier.
2. type-specification can be one of:
   - A [primitive type](../08_language-basics/0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](../08_language-basics/0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](../08_language-basics/0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")

## Usage

The `DEFINE` section declares a data type for each formal argument in the
`REPORT` prototype and for any additional local variables that can be referenced only
within the `REPORT` program block.

The scope of report variables is the `REPORT` routine. However, report variables
are not allocated on the FGL stack like regular `FUNCTION` variables: The lifetime of
report variables is the duration of the program, like module or global variables. As result, they
persist across `OUTPUT TO REPORT` and `FINISH REPORT` calls. When
restarting a report, consider initializing the local report variables in [`FIRST PAGE HEADER`](2481-first-page-header.md "Defines the printing commands for the first page of a report.").

Include arguments in the report prototype and declare them in the `DEFINE`
section, if any of the following conditions is true:

- If you specify [`FORMAT EVERY
  ROW`](2480-format-every-row.md "Default format specification of a report.") to create a default report, you must pass all the values for each record of the
  report.
- If an `ORDER BY`
  section is included, you must pass all the values that `ORDER BY` references for each
  input record of the report.
- If you use the [`AFTER GROUP
  OF`](2484-before-after-group-of.md "Defines printing commands of row grouping sections within a report.") control block, you must pass at least the arguments that are named in that
  control block.
- If an aggregate that depends on all records of the report appears anywhere except in the [`ON LAST ROW`](2486-on-last-row.md "Defines the printing commands of the last row in a report.") control block, you must
  pass each of the records of the report through the argument list.

Aggregates dependent on all records include:

- [`GROUP PERCENT(*)`](2503-percent.md "Calculates the percentage of rows matching a condition.") (anywhere in a
  report).
- Any [aggregate](2501-report-aggregate-functions.md "Report aggregate functions can be used to compute data.") without the
  `GROUP` keyword (anywhere outside the `ON LAST ROW` control
  block).

When a report calls an aggregate function, an error might result if any argument of an aggregate
function is not also a format argument of the report. You can, however, use global or module
variables as arguments of aggregates, if the value of the variable does not change while the report
is executing.

## Related links

**Related concepts**  

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")
