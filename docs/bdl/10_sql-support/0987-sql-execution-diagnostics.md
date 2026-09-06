---
title: "SQL execution diagnostics"
source: "fgl-topics/c_fgl_sql_programming_diag.html"
breadcrumb: "SQL support > SQL programming > SQL basics > SQL execution diagnostics"
type: "concept"
---

# SQL execution diagnostics

> If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.

## Trapping SQL errors

By default, SQL errors stop program execution and display the error message to the standard
output. Most SQL statements executed by a program should not return an error and thus do not require
error trapping. However, in some cases, a program must keep the control when an SQL error occurs.
For example, when connecting to the database, the user might enter an invalid password that will
raise a login denied error. The program must trap such SQL connection error in order to return to
the login dialog and let the user enter a new login and password.

To trap potential SQL errors, surround the SQL statements to be checked either with a [`WHENEVER ERROR`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") exception handler or with a
[`TRY` / `CATCH`](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")
block:

```
-- WHENEVER ERROR handler
WHENEVER ERROR CONTINUE
  INSERT INTO orders VALUES ( rec_ord. * )
  IF sqlca.sqlcode = -75623 THEN
     ...
  END IF
WHENEVER ERROR STOP -- restore the default

-- TRY/CATCH block
TRY
  INSERT INTO orders VALUES ( rec_ord. * )
CATCH
  IF sqlca.sqlcode = -75623 THEN
     ...
  END IF
END TRY
```

## Using sqlca.sqlcode

SQL error codes are provided in the [`sqlca.sqlcode`](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.") register. This register always contains an IBM® Informix® error code, even when
connected to a database that is different from IBM Informix.

`status` is the global error code register that can be set for any kind of error
(even non-SQL). When an SQL error occurs, both `sqlca.sqlcode` and
`status` hold the SQL error code.

Use `sqlca.sqlcode` for SQL error management, and use `status` to
detect errors with other language instructions.

When connecting to a database that is different from IBM
Informix, the database driver tries to convert the native
SQL error to an IBM Informix error which will be copied into the `sqlca.sqlcode` and
`status` registers. If the native SQL error cannot be converted,
`sqlca.sqlcode` and `status` will be set to [**-6372**](../15_library-reference/4483-genero-bdl-errors.md) (a general SQL error).
You can then check the native SQL error in `sqlca.sqlerrd[2]`. The native SQL error
code is always available in `sqlca.sqlerrd[2]`, even if it cannot be converted to an
IBM Informix
error.

## Using SQLSTATE

`SQLSTATE` is a register that contains an error code following ISO
standards. However, not all database servers support this standard.

Preferably use [`SQLSTATE`](../08_language-basics/0645-sqlstate-variable.md "The SQLSTATE predefined variable returns the code corresponding to the last SQL error.") for
SQL error checking, as long as the target databases support this feature.

> **Important:**
>
> The `SQLSTATE` codes are defined by the ISO standards. However, not all database
> types support this standard.

| Database Server Type | Supports SQLSTATE errors |
| --- | --- |
| IBM Informix | Yes, since IDS 10 |
| Microsoft™ SQL Server | Yes, since version 8 (2000) |
| Oracle® MySQL | Yes |
| Oracle Database Server | Not in version 10.2 |
| PostgreSQL | Yes, since version 7.4 |
| Dameng® | No |

## Centralize SQL error checking

SQL error identification sometimes requires complex code to check several RDBMS-specific
error numbers. Therefore, it is strongly recommended that you centralize SQL error identification in
a function. When needed, this allows you to write the RDBMS-specific code only once.

For maximum SQL portability, centralize SQL error checking in functions, to test either
`sqlca.sqlcode` or `SQLSTATE`, depending on the database
server type. Furthermore, consider defining error identifiers with constants:

```
CONSTANT SQLERR_INVALID_DATABASE = -1001,
         SQLERR_INVALID_USER = -1002,
         ...

FUNCTION do_connect()
  DEFINE uname, upswd VARCHAR(100)
  WHILE TRUE
    CALL login() RETURNING uname, upswd
    TRY
       CONNECT TO "stores" USER uname USING upswd
    CATCH
       CASE check_sql_error()
       WHEN SQLERR_INVALID_DATABASE
          DISPLAY SQLERRMESSAGE
          EXIT PROGRAM 1 -- Fatal error: Stop!
       WHEN SQLERR_INVALID_USER
          ERROR "Invalid login, try again"
          CONTINUE WHILE
       END CASE
    END TRY
    EXIT WHILE
  END WHILE
END FUNCTION
```

## SQL error messages

`SQLERRMESSAGE` is a register that contains the database-specific error
message. These messages are different for every database type.

Only use [`SQLERRMESSAGE`](../08_language-basics/0646-sqlerrmessage-variable.md "The SQLERRMESSAGE predefined variable holds the error message corresponding to the last SQL error.") to
print or log SQL execution diagnostics.

## SQL warnings

Some SQL instructions can produce SQL warnings. Unlike SQL errors, SQL warnings indicate a minor
issue that can often be ignored. For example, when connecting to an IBM
Informix database, a warning is returned to indicate that
a database connection was opened. Another warning may also be returned if the database supports
transactions. None of these facts are critical, but that information may help with program
execution.

If an SQL warning is raised, `sqlca.sqlcode` / `status` remain
zero, and the program flow continues. To detect if an SQL warning has occurred, the [sqlca.sqlawarn](1059-database-connections.md "Explains how to manage database connections in a program.") register must be
checked. `sqlca.sqlawarn` is defined as a  `CHAR(7)` variable. If
`sqlca.sqlawarn[1]` contains the `W` letter, it means that the last
SQL instruction returned a warning. The other character positions
(`sqlca.sqlawarn[2-8]`) may contain `W` letters too, depending on the
database server type and the type of SQL instruction that was executed.

If `sqlca.sqlawarn` is set, you can also check the `SQLSTATE` and
`sqlca.sqlerrd[2]` registers to get more details about the warning. The
`SQLERRMESSAGE` register might also contain the warning description.

In this example, the program connects to a database and displays the content of the
`sqlca.sqlawarn` register:

```
MAIN
  DATABASE stores 
  DISPLAY "[", sqlca.sqlawarn, "]"
END MAIN
```

When connecting to an IBM
Informix database with transactions, the program will
display the following:

```
[WW W ]
```

By default, SQL warnings do not stop the program execution. To trap SQL warnings with an
exception handler, use the `WHENEVER WARNING` instruction, as shown in this
example:

```
MAIN
  DEFINE cust_name VARCHAR(50)
  DATABASE stores 
  WHENEVER WARNING STOP
  SELECT cust_lname, cust_address INTO cust_name 
       FROM customer WHERE cust_id = 101
  WHENEVER WARNING CONTINUE
END MAIN
```

The `SELECT` statement in the above example uses two columns in the select list,
while only one `INTO` variable is provided. This is legal and does not raise an SQL
error. However, `sqlca.sqlawarn` is set to indicate that the number of target
variables does not match the select-list items.

See also [WHENEVER WARNING](../09_advanced-features/0849-understanding-exceptions.md "Exceptions are abnormal runtime events that can be trapped for control.")
exception.

## Display detailed debug information in case of internal driver error

If an unexpected problem happens within the database driver, it will return the error [**-6319**](../15_library-reference/4483-genero-bdl-errors.md) (indicating an internal
error in the database driver). When this SQL error occurs, set the [FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.") environment variable to get more
details about the internal error.

## Determine the number of processed rows

SQL statements such as `UPDATE` and `DELETE` can affect zero, one
or several rows.

If you need to check how many rows are processed by an SQL statement, check the
`sqlca.sqlerrd[3]` register.

In the next code example, the `sqlca.sqlcode` and `sqlca.sqlerrd[3]`
registers are used to diagnose the execution of an `UPDATE` statement:

```
TYPE t_country RECORD
  c_code CHAR(5),
  c_desc VARCHAR(50)
END RECORD

FUNCTION update_country_description(rec t_country) RETURNS INTEGER
  UPDATE country SET desc = rec.c_desc WHERE code = rec.c_code
  CASE
    WHEN sqlca.sqlcode < 0
      MESSAGE SFMT("SQL error:%1 [%2]", sqlca.sqlcode, SQLERRMESSAGE)
      RETURN -2
    WHEN sqlca.sqlcode == 0
      IF sqlca.sqlerrd[3] == 0 THEN
        MESSAGE SFMT("No country was found with code: %1", rec.c_code)
        RETURN -1
      ELSE
        MESSAGE SFMT("Updated country with code: %1", rec.c_code)
      END IF
  END CASE
  RETURN 0
END FUNCTION
```

For more details see [The sqlca diagnostic record](0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.").

## Related links

**Related concepts**  

[SQL error identification](0989-sql-error-identification.md "Identify SQL exceptions in your programs with sqlca.sqlcode.")
