---
title: "Performances with SQL interruption"
source: "fgl-topics/c_fgl_Migrate_to_320_sql_interruption.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Performances with SQL interruption"
type: "concept"
---

# Performances with SQL interruption

> With some database drivers, performances can be impacted when using OPTIONS SQL INTERRUPT ON.

Starting with Genero BDL 3.20, `dbmsnc` / MS ODBC SQL database drivers have been
reviewed to improve support for SQL interruption, by using asynchronous ODBC calls.

This solution requires additional overhead, that can slow down regular SQL statements which
do not need SQL interruption handling.
> **Note:**
>
> Most of SQL statements of a program do not need SQL
> interruption handling.

Instead of enabling SQL interruption with [`OPTIONS SQL INTERRUPT ON`](../10_sql-support/0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") at the beginning of the program, consider enabling
this option only for the execution of SQL queries that can be interrupted. This programming pattern
is anyway a good practice, since SQL interruption requires exception handling that needs to be
treated.

> **Note:**
>
> To easily identify when `OPTIONS SQL INTERRUPT ON/OFF` is used in your
> programs, the [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") output has been
> enhanced to trace this
> instruction:
>
> ```
> SQL: SQL INTERRUPT ON
>  | 4gl source      : sqlint.4gl line=5
>  | curr driver     : ident='dbmdefault'
>  | curr connection : ident='_1' (dbspec=[test1])
>  | Execution time  :   0 00:00:00.00003
> ...
> SQL: SQL INTERRUPT OFF
>  | 4gl source      : sqlint.4gl line=12
>  | curr driver     : ident='dbmdefault'
>  | curr connection : ident='_1' (dbspec=[test1])
>  | Execution time  :   0 00:00:00.00003
> ```

## Related links

**Related concepts**  

[SQL Server drivers performance](0167-sql-server-drivers-performance.md "SQL Server ODI drivers based on FreeTDS, Easysoft and MS ODBC have been reviewed to achieve better execution times.")
