---
title: "FOREACH (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_FOREACH.html"
breadcrumb: "SQL support > Result set processing > FOREACH (result set cursor)"
type: "concept"
---

# FOREACH (result set cursor)

> Processes a series of data rows returned from a database cursor.

## Syntax

```
FOREACH cid
   [ USING pvar {IN|OUT|INOUT} [,...] ]
   [ WITH REOPTIMIZATION ]
   [ INTO fvar [,...] ]
       {
         statement
       | CONTINUE FOREACH
       | EXIT FOREACH
       }
       [...]
END FOREACH
```

1. cid is the identifier of the database cursor.
2. pvar is a variable containing an input value for an SQL parameter.
3. fvar is a variable used as fetch buffer.

## Usage

Use the `FOREACH` instruction to retrieve and process database rows that were
selected by a query. This instruction is equivalent to using the [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor"), [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.") and [`CLOSE`](1153-close-result-set-cursor.md "Closes a database cursor and frees resources allocated on the database server for the result set.") cursor instructions:

1. Open the specified cursor
2. Fetch the rows selected
3. Close the cursor (after the last row has been fetched)

You must declare the cursor (by using the `DECLARE` instruction) before the
`FOREACH` instruction can retrieve the rows. A compile-time error occurs unless
the cursor was declared prior to this point in the source module. You can reference a sequential
cursor, a scroll cursor, a hold cursor, or an update cursor, but `FOREACH` only
processes rows in sequential order.

> **Important:**
>
> When fetching rows into a [dynamic
> array](../08_language-basics/0734-dynamic-arrays.md) element with `FOREACH cursor INTO
> dyn-arr[index].*`, the array element will be
> implicitly allocated in order to fetch the row into the array variable. This occurs also when
> reaching the end of the result set. Therefore, you always need to delete the last element of the
> dynamic array filled by a `FOREACH` loop.

The `FOREACH` statement performs successive fetches until all rows specified by
the `SELECT` statement are retrieved. Then the cursor is automatically closed. It is
also closed if a [`WHENEVER NOT FOUND`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")
exception handler within the `FOREACH` loop detects a `NOTFOUND`
condition.

After a `FOREACH` loop, `status` and [`sqlca.sqlcode`](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.") will not be set to
[`NOTFOUND/100`](../08_language-basics/0575-notfound.md "NOTFOUND is a predefined constant used to check if an SQL statement returns rows.") if no rows are
returned by the query: If no error occurred, these registers will hold the value zero.

The `USING` clause is required to provide the SQL parameter buffers, if the cursor
was declared with a prepared statement that includes (`?`) question mark
placeholders.
> **Important:**
>
> Older versions of Informix 4GL did not support the `USING` clause in the
> `FOREACH` instruction. The solution was to issue an [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") instruction with a
> `USING` clause, before the `FOREACH` (without `USING`
> clause). This practice is allowed by the Genero compiler for backward compatibility, but can lead to
> unexpected behavior, especially when enabling Informix optimization options such as OPTOFC=1. When
> the SQL statement has parameters, specify the parameter variables in a `USING` clause
> of the `FOREACH` instruction, and remove the unnecessary `OPEN`
> instruction before the `FOREACH`.

> **Tip:**
>
> Consider using the [`fglcomp -W
> fragile-cursor`](1117-using-program-variables-in-static-sql.md) option to detect risk SQL cursor usage.

The `IN`, `OUT` or `INOUT` options can be used to
call stored procedures having input / output parameters and generating a result set. Use the
`IN`, `OUT`, or `INOUT` options to indicate if a parameter
is respectively for input, output, or both.

The `INTO` clause can be used to provide the fetch buffers that receive the row
values.

Use the `WITH REOPTIMIZATION` clause to indicate that the query execution plan has
to be re-optimized.

The [`CONTINUE FOREACH`](../08_language-basics/0678-continue-block-name.md "The CONTINUE block-name instruction resumes execution of a loop or dialog statement.")
instruction interrupts processing of the current row and starts processing the next row. The runtime
system fetches the next row and resumes processing at the first statement in the block.

The [`EXIT FOREACH`](../08_language-basics/0679-exit-block-name.md "The EXIT block instruction transfers control out of the current program block.") instruction
interrupts processing and ignores the remaining rows of the result set.

The `IN`, `OUT`, or `INOUT` options can only be used
for simple variables; you cannot specify those options for a complete record with the record.\*
notation.

## Example

Fetching rows into a dynamic array:

```
MAIN
   DEFINE clist DYNAMIC ARRAY OF RECORD
           cnum INTEGER,
           cname CHAR(50)
         END RECORD
   DEFINE i INTEGER
   DEFINE mpk INTEGER
   DATABASE stores
   DECLARE c1 CURSOR FOR
     SELECT customer_num, cust_name FROM customer
       WHERE customer_num > ?
   LET i=1
   LET mpk = 100
   FOREACH c1 USING mpk INTO clist[i].*
      DISPLAY clist[i].*
      LET i=i+1
   END FOREACH
   CALL clist.deleteElement(i)
   DISPLAY "Number of rows found: ", clist.getLength()
END MAIN
```

Fetching rows into a static array:

```
MAIN
   DEFINE clist ARRAY[200] OF RECORD
           cnum INTEGER,
           cname CHAR(50)
         END RECORD
   DEFINE i INTEGER
   DEFINE mpk INTEGER
   DATABASE stores
   DECLARE c1 CURSOR FOR
      SELECT customer_num, cust_name FROM customer
       WHERE customer_num > ?
   LET i=0 -- Assign zero!
   LET mpk = 100
   FOREACH c1 USING mpk INTO clist[i+1].*
      LET i=i+1
      DISPLAY clist[i].*
   END FOREACH
   DISPLAY "Number of rows found: ", i
END MAIN
```

## Related links

**Related concepts**  

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
