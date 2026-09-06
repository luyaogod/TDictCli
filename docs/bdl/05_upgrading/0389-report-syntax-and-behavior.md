---
title: "REPORT syntax and behavior"
source: "fgl-topics/c_fgl_MigI4GL_046.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > REPORT syntax and behavior"
type: "concept"
---

# REPORT syntax and behavior

> This topic describes differences between I4GL and Genero BDL in the report engine.

## FINISH REPORT is mandatory to get output

With IBM® Informix® 4GL, even when the `FINISH REPORT` instruction is not used, the report
engine produces some output.

To handle report output correctly, Genero BDL needs to know if the report is finished or if more
rows will be printed. Therefore, the `FINISH REPORT` instruction is mandatory with
FGL, in order to get a report output.

Fix the code and properly end the report with `FINISH REPORT`.

## BEFORE/AFTER GROUP OF ignores trailing blanks

With IBM Informix 4GL RDS (fglpc/fglgo), trailing blanks are
significant when present in a `VARCHAR` variable used by a `BEFORE GROUP
OF` or `AFTER GROUP OF` clause in a `REPORT` routine.

For example, the value `"ABC__"` (here the spaces are represented with an
underscore) and `"ABC_"` or `"ABC"` are considered different with I4GL
`BEFORE/AFTER GROUP OF`, while Genero BDL considers these values as equal and group
the report rows consequently.

The Genero BDL behavior is consistent with other character string comparison areas, like when
using a `VARCHAR` variable in an `IF` statement.

Note that `CHAR` variables are always blank-padded. Consequently, it's not
possible to have a different number of trailing spaces.

## Report SUM/AVG when using MONEY

With IBM Informix 4GL, in a `REPORT` routine, when using the `SUM` or
`AVG` aggregate functions with a variable of type `MONEY`, the
resulting type remains a `MONEY`, and the currency symbols defined by DBMONEY is used
in the output.

With Genero BDL, the result type for `SUM` or `AVG` is always a
`DECIMAL`, and the currency symbol is not used to format the result. The precision of
the result type depends on how often the report routine has been called. To get a predictable
result, format the result of those aggregate functions with a `USING` clause.

## Using string expression in COLUMN operator

With IBM Informix 4GL, when using `WHENEVER ANY ERROR`, if a PRINT instruction uses the
`COLUMN` operator to shift forward the output position, if the expression passed as
argument is not a valid integer number, I4GL produces a runtime error -1213 "A character to numeric
conversion process failed".

Genero BDL does not produce this error when using `WHENEVER ANY ERROR`. The
expression evaluates to `NULL` and no space characters are produced by the
`COLUMN` operator.

Note that in the context of `WHENEVER ERROR`, the behavior of I4GL and Genero BDL
is identical in this case.

Fix the code in order to use a valid `INTEGER` variable for
`COLUMN` arguments.

## The FGLSKIPNXTPG environment variable

With IBM Informix 4GL, the FGLSKIPNXTPG environment variable can be used to control the behavior of the
report engine. Since I4GL 4.1, the `SKIP TO TOP OF PAGE` command has no effect when a
report is already at the top of the page (that is, when no data has yet been sent to a new page). To
get the older behavior of `SKIP TO TOP OF PAGE` can override this effect using the
environment variable FGLSKIPNXTPG.

Genero BDL implements the new I4GL 4.1 behavior, but there is not way to get the former behavior
as when setting the FGLSKIPNXTPG environment variable with I4GL.

## Related links

**Related concepts**  

[Reports](../12_reports/2461-reports.md "Reports")
