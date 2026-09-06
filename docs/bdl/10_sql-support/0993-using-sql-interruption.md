---
title: "Using SQL interruption"
source: "fgl-topics/c_fgl_sql_programming_094.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Using SQL interruption"
type: "concept"
---

# Using SQL interruption

> Interrupt long running SQL queries, or interrupt queries waiting for locked data.

If the database server supports SQL interruption, a program can interrupt a long running SQL
statement.

> **Important:**
>
> Not all database servers support SQL interruption.

SQL interruption is not enabled by default. Use the [`OPTIONS SQL INTERRUPT ON`](../09_advanced-features/0934-enabling-disabling-sql-interruption.md "The OPTIONS SQL INTERRUPT instruction enables or disables SQL statement interruption.") instruction to enable SQL interruption.

The `OPTIONS SQL INTERRUPT ON` instruction must be used together with signal
handling instructions [`DEFER INTERRUPT` and
`DEFER QUIT`](../09_advanced-features/0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received."), otherwise the program will stop immediately in case of an
interruption event.

Depending on the type of database server and client, SQL interruption handling may require
additional overhead. Consider enabling SQL interruption only for SQL queries that can take a while
to execute or can be blocked for a long time because of concurrent access (locks).

When the program receives an interruption event (either a SIGINT signal from the system, or an
[interrupt event](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure.") from the front-end) and SQL
interrupt is enabled with `OPTIONS SQL INTERRUPT ON`, the following happens:

- The running SQL statement is stopped,
- The `int_flag` global variable is set to `TRUE`,
- The `sqlca.sqlcode` is set with error [-213](../15_library-reference/4483-genero-bdl-errors.md).

SQL interruption results in abnormal SQL statement execution and raises a runtime error.
Therefore, the SQL statement that can be interrupted must be protected by a `WHENEVER
ERROR` exception handler or `TRY/CATCH` block.

```
MAIN
  DEFINE cnt INTEGER
  DEFER INTERRUPT
  DATABASE test1
  WHENEVER ERROR CONTINUE
  OPTIONS SQL INTERRUPT ON
  -- Start long query (self join takes time)
  -- From now on, user can hit CTRL-C in TUI mode to stop the query,
  -- or use the special "interrupt" action (button) in GUI mode. 
  SELECT COUNT(*) INTO cnt FROM customers a, customers b 
       WHERE a.cust_id <> b.cust_id 
  OPTIONS SQL INTERRUPT OFF
  IF sqlca.sqlcode == -213 THEN
     DISPLAY "Statement was interrupted by user..."
     EXIT PROGRAM 1
  END IF
  WHENEVER ERROR STOP
END MAIN
```

When SQL interruption is supported by a database server type other than IBM® Informix®, the database drivers
will return error [-213](../15_library-reference/4483-genero-bdl-errors.md) in case
of interruption, to behave as in IBM Informix.

When [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") is set, `OPTIONS
SQL INTERRUPT ON/OFF` instructions are traced.

| Database Server Type | SQL Interruption API | SQL error code for interrupted query |
| --- | --- | --- |
| IBM Informix | [sqlbreak()](1186-what-are-the-supported-ibm-informix-sql-features.md) | Native error -213 |
| Microsoft™ SQL Server (Only 2005+ with SNC driver) | [SQLCancel()](1306-sql-interruption.md) | SQLSTATE HY008 |
| Oracle® MySQL | [KILL QUERY](1353-sql-interruption.md) | Native error -1317 |
| Oracle Database Server | [OCIBreak()](1409-sql-interruption.md) | Native error -1013 |
| PostgreSQL | [PQCancel()](1460-sql-interruption.md) | SQLSTATE 57014 |
| SQLite | [sqlite3\_interrupt()](1502-sql-interruption.md) | Native error SQLITE\_ABORT |
| Dameng® | [`cpi_cancel()`](1254-sql-interruption.md) | Native error -70019 |

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[TRY - CATCH block](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")

[WHENEVER directive](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")
