---
title: "fglhint_* SQL comments"
source: "fgl-topics/c_fgl_DynamicSql_fglhints.html"
breadcrumb: "SQL support > Dynamic SQL management > fglhint_* SQL comments"
type: "concept"
---

# fglhint_* SQL comments

> Using special SQL comment hints to control statement execution.

## Syntax

```
/*
{ fglhint_insert
| fglhint_update
| fglhint_delete
| fglhint_select
| fglhint_other
}
[ fglhint_no_ifxemul ]
*/
```

The `fglhint_*` indicators help the database driver to identify the SQL statement
and control Informix® emulation:

1. `fglhint_insert`: Regular `INSERT` (preparable, without result set).
2. `fglhint_update`: Regular `UPDATE` (preparable, without result set).
3. `fglhint_delete`: Regular `DELETE` (preparable, without result set).
4. `fglhint_select`: Regular `SELECT` (preparable, with result set).
5. `fglhint_other`: Any other non-preparable SQL statement (maybe with result set).
6. `fglhint_no_ifxemul`: Disable Informix emulation (like the FGLPROFILE entry, but for this statement only).

## Usage

When preparing an SQL statement, the ODI database drivers must identify the type of SQL statement,
in order to set database client API options to perform the statement properly.

For example, when executing an `SELECT` statement that returns a result set, ODBC
options for a server-side cursor must be set with Microsoft® SQL Server clients. If a server-side cursor is not used only
one active cursor is used in a Genero program.

Or, for example, when executing an `SELECT` statement that creates temporary tables
with SQL Server, it must be done with a direct ODBC execution (using ODBC SQLExecDirect instead of
SQLPrepare/SQLExecute), otherwise the temporary table is created in the scope of the sp\_prepare()
stored procedure, and is dropped immediately after returning from the sp\_prepare() call.

Specified inside `/* */` C-style comments, the `fglhint_*` keywords
can be used as SQL statement recognition directives:

- When the SQL statement looks like a regular `INSERT` statement, but uses an
  `OUTPUT INSERTED.*` clause, add the `/* fglhint_select */` SQL comment,
  to indicate that it must be treated as a regular, preparable `SELECT` statement
  returning a result set.
- When the SQL statement looks like a regular `SELECT` statement, but has an
  `INTO newtable` clause to create a new table, add the `/* fglhint_insert */`
  SQL comment, to indicate that it must be treated as a regular, preparable `INSERT`
  statement.
- When the SQL statement looks like a regular `SELECT` statement, but has an
  `INTO #newtemptable` clause to create a new temporary table, add the
  `/* fglhint_other */` SQL comment, to indicate that it must be treated as a
  non-preparable `INSERT` statement (to be executed directly).

`/* */` C-style comments are only allowed in dynamic SQL statements, the BDL
language does not allow such comments in static SQL statements, which is based on the Informix-SQL
syntax.

For example, this statement performs an `INSERT` with an
`OUTPUT` clause, that will produce a result set. Use the
`fglhint_select` hint, to indicate that the statement must be executed as a regular
`SELECT`:

```
DECLARE c1 CURSOR
   FROM "/* fglhint_select */ INSERT INTO table1 OUTPUT INSERTED.* SELECT * FROM customers"
```

This statement performs a `SELECT INTO newtable`, that will be executed as an
`INSERT` (without a result set), because we use the `fglhint_insert`
hint:

```
EXECUTE IMMEDIATE "/* fglhint_insert */ SELECT * INTO newtable FROM customers WHERE cust_valid='Y'"
```

In order to perform a `SELECT INTO #temptable`, the statement must be executed
directly without using the sp\_prepare/sp\_execute procedures of SQL Server (otherwise the temp table
will only exist in the context of the stored procedures).
Therefore, we use the `fglhint_other` hint to force a direct SQL execution:

```
EXECUTE IMMEDIATE "/* fglhint_other */ SELECT * INTO #tmptab1 SELECT * FROM customers"
DECLARE c2 CURSOR FROM "SELECT * FROM #tmptab1 ..."
```

Furthermore, Informix SQL emulation can
be disabled by using the `fglhint_no_ifxemul`
hint:

```
PREPARE s1 FROM "/* fglhint_no_ifxemul */ SELECT * FROM [dbo].[mytable]"
```

When using the `fglhint_no_ifxemul` hint, Informix emulation will only be disabled for this SQL statement. Other SQL
statements executed by your program can use the Informix emulations.

SQL statement type identification and Informix
emulation hints can be combined:

```
PREPARE s1 FROM "/* fglhint_no_ifxemul fglhint_select */ INSERT INTO [dbo].[mytable] OUTPUT ..."
```

## Related links

**Related concepts**  

[IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.")

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
