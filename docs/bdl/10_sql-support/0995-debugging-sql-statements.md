---
title: "Debugging SQL statements"
source: "fgl-topics/c_fgl_sql_programming_debug.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Debugging SQL statements"
type: "concept"
---

# Debugging SQL statements

> The runtime system can display debug information for SQL statements executed by the program.

SQL debug information is printed by the runtime system, when using a value different from zero in
the [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") environment variable or with
the [`fgl_sqldebug()`](../15_library-reference/2783-fgl-sqldebug.md "Sets the SQL debug level from program code.")
built-in function.

The debug messages are sent to the standard error stream (stderr). Consider
to redirect the standard error output into a file, to inspect the debug log after program
execution.

UNIX™ (shell) example:

```
FGLSQLDEBUG=3
export FGLSQLDEBUG
fglrun myprog 2>sqldbg.txt
```

When the debug level is set with a positive value, the runtime system prints an SQL debug message
for each SQL statement executed by the program.

When the debug level is set to -1, the runtime system only prints debug information for SQL
statements that produce an error.

SQL debug messages show detailed information about the SQL statement execution, for
example:

```
SQL: DATABASE
 | 4gl source      : c.4gl line=2
 | loading driver  : [/opt/fgl/dbdrivers/dbmdefault]
 | db driver type  : ifx
 | sqlcode         : 0
 | curr driver     : ident='dbmdefault'
 | curr connection : ident='_1' (dbspec=[test1])
 | Execution time  :   0 00:00:00.02689
SQL: DELETE FROM mytable WHERE pkey IS NULL
 | 4gl source      : c.4gl line=4
 | sqlcode         : -206
 |   sqlstate      : 42000
 |   sqlerrd2      : -111
 |   sql message   : The specified table (mytable) is not in the database.
 |   sql msg param : mytable
 | curr driver     : ident='dbmdefault'
 | curr connection : ident='_1' (dbspec=[test1])
 | Execution time  :   0 00:00:00.00035
```

The most important information is the SQL error code and the source code line where the SQL
statement failed. For performance tuning, the execution time can be used to identify slow SQL
statements.

If the SQL debug log generates too many messages, define the debug leve to `-1`,
in order to produce an SQL debug message only for the SQL statements that produce an SQL error. Once
the SQL statement producing the error is identified, you can find the source code line where the
error occurred, to reproduce the problem. Before reproducing the error, set the debug level to a
high positive level, to get detailed debug information.

When the debug level is set to a positive value, an SQL debug header is printed before
executing the underlying database driver code. If the driver code crashes or raises an assertion,
you can easily find the last SQL instruction that was executed by the program, and report to your
support center. When the debug level is set to -1, the SQL debug header is not printed before
executing the database driver code, because the SQL execution status
(`sqlca.sqlcode`) is not known before executing the statement. If the driver code
crashes, no error message will be printed. If you experience database driver crashes, use a positive
debug level in order to identify the problem.

It is also possible to use the `fgl_sqldebug()` built-in function to set the SQL
debug level by program:

```
MAIN
    DATABASE stores
    DELETE FROM mytable   -- SQL debug output if FGLSQLDEBUG is set
    CALL fgl_sqldebug(3)
    DELETE FROM mytable   -- SQL debug output is enabled by program
    CALL fgl_sqldebug(0)
    DELETE FROM mytable   -- SQL debug output if FGLSQLDEBUG is set
    CALL fgl_sqldebug(-1)
    DELETE FROM undefined -- SQL debug output only in case of error
END MAIN
```

## Related links

**Related concepts**  

[The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
