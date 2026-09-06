---
title: "The sqlca diagnostic record"
source: "fgl-topics/c_fgl_sql_programming_sqlca.html"
breadcrumb: "SQL support > SQL programming > SQL basics > The sqlca diagnostic record"
type: "concept"
---

# The sqlca diagnostic record

> The sqlca variable is a predefined record containing SQL statement execution information.

## The sqlca record definition

The `sqlca` record is defined as follows:

```
DEFINE sqlca RECORD
  sqlcode INTEGER,
  sqlerrm VARCHAR(71),
  sqlerrp CHAR(7),
  sqlerrd ARRAY[6] OF BIGINT,
  sqlawarn CHAR(7)
END RECORD
```

1. `sqlcode`: contains the SQL execution code ( 0 = OK, 100 = not row found, <0 =
   error ).
2. `sqlerrm`: contains the error message parameter.
3. `sqlerrp` is not used at this time.
4. `sqlerrd[1]` (see note): contains the estimated number of rows affected after a
   successful `PREPARE` or `DECLARE`.
5. `sqlerrd[2]`: contains the last generated `SERIAL` or the native
   (non-Informix) SQL error code.
6. `sqlerrd[3]` (see note): contains the number of rows processed in the last statement , or the number of
   rows that were successfully processed before an error was detected.
7. `sqlerrd[4]` (see note): contains the estimated weighted sum of disk accesses and total rows
   processed.
8. `sqlerrd[5]` (see note): contains the offset in the statement text where the error was detected.
9. `sqlerrd[6]` (see note): contains the `ROWID` of the last row that was processed. Depends
   on how the database server processes a query, particularly for `SELECT`
   statements.
10. `sqlawarn` contains the ANSI warning represented
    by a W character at a given position in the string.
11. `sqlawarn[1]` is set to W when any of the other
    warning characters have been set to W.
12. `sqlawarn[2-7]` have specific meanings, see database
    server documentation for more details.

> **Note:**
>
> The content and meaning of some sqlca members registers depends on the target database engine
> type.

## Usage

The "sqlca" acronym stands for the SQL Communication Area variable.

The `sqlca` record can be used to get an SQL execution diagnostic. Error and
warning information can be found in this structure.

The `sqlca` record is filled after each SQL statement execution.

`sqlca` is not designed to be modified by user code. It must be used as a
read-only record.

## Portability

`sqlca.sqlcode` will be set to a specific IBM® Informix® SQL error code, provided that the
database driver can convert the native SQL error to an IBM Informix SQL error. In case of error,
`sqlca.sqlerrd[2]` will hold the native SQL error produced by the database
server.

Other `sqlca` record members are specific to IBM Informix. For example,
after inserting a row in a table with a `SERIAL` column,
`sqlca.sqlerrd[2]` will contain the new generated serial number. After an SQL error
occurred, `sqlca.sqlerrd[2]` will contain the native SQL error. Furthermore, the
`sqlca.sqlerrd[3]` member may be set with the number of processed rows, as long as
the database client supports this feature. Other `sqlca.sqlerrd[n]` members must be
considered as non portable.

## Example

```
MAIN
  WHENEVER ERROR CONTINUE
  DATABASE stores 
  SELECT COUNT(*) FROM foo   -- Table should not exist!
  DISPLAY sqlca.sqlcode, sqlca.sqlerrd[2]
END MAIN
```

## Related links

**Related concepts**  

[SQL error identification](0989-sql-error-identification.md "Identify SQL exceptions in your programs with sqlca.sqlcode.")

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[SQL errors on PREPARE](1403-sql-errors-on-prepare.md "SQL errors on PREPARE")

[The BIGSERIAL / SERIAL8 data types](1194-the-bigserial-serial8-data-types.md "The BIGSERIAL / SERIAL8 data types")
