---
title: "OUTPUT TO REPORT"
source: "fgl-topics/c_fgl_reports_OUTPUT_TO_REPORT.html"
breadcrumb: "Reports > The report driver > OUTPUT TO REPORT"
type: "concept"
---

# OUTPUT TO REPORT

> The OUTPUT TO REPORT instruction provides a data row to the report execution.

## Syntax

```
OUTPUT TO REPORT report-name ( parameters )
```

1. report-name is the name of the report to which the
   parameters are sent.
2. parameters is the data that needs to be sent to the report.

## Usage

The `OUTPUT TO REPORT` instruction feeds the report routine with a single
set of data values (called an input record), which corresponds usually to one
printed line in the report output.

An input record is the ordered set of values
returned by the expressions that you list between the parentheses following the report name
in the `OUTPUT TO REPORT` statement. The specified values are passed to the
report routine, as part of the input record. The input record typically corresponds to a
retrieved row from the database.

The set of values is usually grouped in a `RECORD` variable and best
practice is to define a user defined type ([`TYPE`](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")) in order to ease the variable
definitions required in the code implementing the report driver and the report routine
definition, for example:

```
SCHEMA stores 
TYPE t_cust RECORD LIKE customer.*
...
DEFINE r_cust t_cust 
...
  OUTPUT TO REPORT cust_report(r_cust.*)
...
REPORT cust_report(r)
  DEFINE r t_cust 
  ...
```

The `OUTPUT TO REPORT` statement is included within a `WHILE`, `FOR`, or `FOREACH` loop, so that the
program passes data to the report one input record at a time. The example uses a
`FOREACH` loop to fetch data from the database and pass it as input record to
a report:

```
SCHEMA stores 
DEFINE o LIKE orders.*
...
DECLARE order_c CURSOR FOR
  SELECT orders.*
    FROM orders ORDER BY ord_cust 
  START REPORT order_list 
  FOREACH order_c INTO o.*
    OUTPUT TO REPORT order_list(o.*)
  END FOREACH
  FINISH REPORT order_list 
...
```

Special consideration is needed when using `OUTPUT TO REPORT` with row
ordering. For example if the report groups rows with [`BEFORE GROUP OF` or
`AFTER GROUP OF`](2484-before-after-group-of.md "Defines printing commands of row grouping sections within a report.") sections, the rows must be ordered by the column
specified in these sections, preferably by the report driver to avoid two-pass reports.

If `OUTPUT TO REPORT` is not executed, none of the control blocks of the
report routine are executed, even if the program also includes the `START REPORT` and `FINISH REPORT` statements.

The members of the input record that you specify in the expression list of the
`OUTPUT TO REPORT` statement must correspond to elements of the formal
argument list in the report definition; in their number and their position. They must
be also of compatible data types. At compile time, the number of parameters passed
with the `OUTPUT TO REPORT` instruction is not checked against the
`DEFINE section in REPORT` of the report
routine. This is known behavior in the language.

Arguments of the `TEXT` and `BYTE` data types are passed
by reference rather than by value; arguments of other data types are passed by value.
A report can use the `WORDWRAP` operator
with the `PRINT` statement to display
`TEXT` values. A report cannot display `BYTE` values; the
character string `<byte
value>` in output from the report indicates a `BYTE` value.

## Related links

**Related concepts**  

[The report routine](2474-the-report-routine.md "The report routine implements the body of a report, with formatting instructions.")

[Two-pass reports](2488-two-pass-reports.md "The report engine supports two-pass reports, to order rows automatically.")

[PRINTX](2492-printx.md "Prints an XML formatted row of data in a report, with an additional identifier for XML outputs.")

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
