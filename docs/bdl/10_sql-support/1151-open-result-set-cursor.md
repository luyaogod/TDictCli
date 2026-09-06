---
title: "OPEN (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_OPEN.html"
breadcrumb: "SQL support > Result set processing > OPEN (result set cursor)"
type: "concept"
---

# OPEN (result set cursor)

> Executes the SQL statement with result set associated with the specified database cursor

## Syntax

```
OPEN cid
   [ USING pvar {IN|OUT|INOUT} [,...] ]
   [ WITH REOPTIMIZATION ]
```

1. cid is the identifier of the database cursor.
2. pvar is a variable containing an input value
   for an SQL parameter.

## Usage

The `OPEN` instruction executes the SQL statement of a cursor created by the
[`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") instruction. The result
set is produced on the server side and rows can be fetched.

The `USING` clause is required to provide the SQL parameters as program variables,
if the cursor was declared with a prepared statement that includes (`?`) question
mark placeholders.

A subsequent `OPEN` statement closes the cursor and then reopens it. When the
database server reopens the cursor, it creates a new result set, based on the current values of the
variables in the `USING` clause. If the variables have changed since the previous
`OPEN` statement, reopening the cursor can generate an entirely different result
set.

The `IN`, `OUT` or `INOUT` options can be used to
call stored procedures having input / output parameters and generating a result set. Use the
`IN`, `OUT` or `INOUT` options to indicate if a
parameter is respectively for input, output or both.

Sometimes, query execution plans need to be re-optimized when SQL parameter values change. Use
the `WITH REOPTIMIZATION` clause to indicate that the query execution plan has to
be re-optimized on the database server (this operation is normally done during the
`DECLARE` instruction). If this option is not supported by the database server,
it is ignored.

In an IBM® Informix®
database that is ANSI-compliant, you receive an error code if you try to open a cursor that is
already open. **Informix only!**

A cursor is closed with the [`CLOSE`](1153-close-result-set-cursor.md "Closes a database cursor and frees resources allocated on the database server for the result set.")
instruction, or when the parent connection is terminated (typically, when the program ends). By
using the `CLOSE` instruction explicitly, you release resources allocated for the
result set in the db client library and on the database server.

The database server evaluates the values that are named in the `USING` clause of
the `OPEN` statement only when it opens the cursor. While the cursor is open,
subsequent changes to program variables in the `OPEN` clause do not change the result
set of the cursor; you must re-open the cursor to re-execute the statement.

If you release cursor resources with a [`FREE`](1154-free-result-set-cursor.md "Releases SQL cursor resources allocated by the DECLARE instruction.")
instruction, you cannot use the cursor unless you declare the cursor again.

The `IN`, `OUT` or `INOUT` options can only be
used for simple variables, you cannot specify those options for a complete record with the record.\*
notation.

## Example

```
MAIN
   DEFINE k INTEGER
   DEFINE n VARCHAR(50)
   DATABASE stores 
   DECLARE c1 CURSOR FROM "SELECT cust_name FROM customer WHERE cust_id > ?"
   LET k = 102
   OPEN c1 USING k 
   FETCH c1 INTO n 
   LET k = 103
   OPEN c1 USING k 
   FETCH c1 INTO n 
END MAIN
```
