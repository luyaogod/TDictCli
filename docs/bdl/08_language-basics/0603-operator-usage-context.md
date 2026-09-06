---
title: "Operator usage context"
source: "fgl-topics/c_fgl_operators_015.html"
breadcrumb: "Language basics > Operators > Operator usage context"
type: "concept"
---

# Operator usage context

> Some operators are specific to a context.

There are operators only available in the SQL syntax, for example:

- `expr BETWEEN expr AND
  expr`
- `expr comparison-operator SOME (
  sub-query )`

Other operators are only available in BDL code, such as:

- [expr :=
  expr](0651-assignment.md "The := operator assigns a variable with an expression and returns the result.")
- [`SFMT( expr [,...]
  )`](0639-sfmt-function.md "The SFMT() operator replaces place holders in a string with values.")

Some operators are common to the BDL and SQL language:

- [`IIF( bool-expr,
  true-expr, false-expr )`](0616-iif-function.md "The IIF() returns the second or third parameter depending on the boolean expression given as first argument.")
- `expr LIKE
  mask`

When using an operator such as `IIF()` or `LIKE` in an SQL
statement, it is processed by your database engine. In SQL, the behavior and semantics may be
slightly different from the BDL operator.

The following operators are only allowed in the `FORMAT` section of report
routines:

- [`PAGENO`](../12_reports/2498-pageno.md "Contains the current page number in a report.")
- [`WORDWRAP`](../12_reports/2500-wordwrap.md "Splits a character string to match a given margin limit.")
