---
title: "Stored procedure calls"
source: "fgl-topics/c_fgl_sql_programming_031.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Stored procedure calls"
type: "concept"
description: "PostgreSQL supports stored procedures and stored functions. Important: Stored procedures are supported since PostgreSQL version 11: Only stored functions are supported in releases prior to PostgreSQL ..."
---

# Stored procedure calls

PostgreSQL supports stored procedures and stored functions.

> **Important:**
>
> Stored procedures are supported since PostgreSQL version 11: Only stored
> functions are supported in releases prior to PostgreSQL 11.

To create a stored procedure in a PortgreSQL database, use the `CREATE PROCEDURE`
statement. Stored procudes can take `IN`, `OUT` and
`INOUT` parameters that will be returned as a result set row.
> **Note:**
>
> Until PostgreSQL
> version 13, only `IN` and `INOUT` parameter markers are supported with
> stored procedures, requiring to use the `INOUT` marker for output-only parameters.
> Starting with PostgreSQL version 14, you can use the `OUT` parameter marker, for
> output-only parameters.

To create a stored function in a PortgreSQL database, use the `CREATE FUNCTION`
statement. Stored functions can take `IN` parameters and can also return more that
one value when specify the returning values as function parameters with the `OUT`
keyword. However, to be a scalar valued function that can be used in SQL expressions, a stored
function must return a single value with the `RETURNS` clause. To return a resut set
with multiple rows, define a stored function with the `RETURNS SETOF` clause.

Pay attention to the procedure/function signature; PostgreSQL allows overloading. For example,
`func(int)` and `func(char)` are two different functions. To drop a
procedure or function, you must specify the parameter type to identify signature properly.

## Child topics

- [Stored procedure with output parameters](1463-stored-procedure-with-output-parameters.md)
- [Stored functions with result set](1464-stored-functions-with-result-set.md)
- [Stored function with output parameters](1465-stored-function-with-output-parameters.md)
