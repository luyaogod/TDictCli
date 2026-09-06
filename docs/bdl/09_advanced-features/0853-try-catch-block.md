---
title: "TRY - CATCH block"
source: "fgl-topics/c_fgl_Exceptions_007.html"
breadcrumb: "Advanced features > Exceptions > TRY - CATCH block"
type: "concept"
---

# TRY - CATCH block

> Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.

## Syntax

```
TRY
   instruction
   [...]
CATCH
   instruction
   [...]
END TRY
```

## Usage

Any language instruction in the `TRY` block will be executed until an exception
is thrown. After an exception the program execution continues in the `CATCH` block.
If no `CATCH` block is provided, the execution continues after `END
TRY`.

If no exception is raised by the statements between the `TRY` and
`CATCH` keywords, the instructions in the `CATCH` section are
ignored and the program flow continues after `END TRY`.

This code example shows a `TRY` block executing an SQL
statement:

```
TRY
    SELECT COUNT(*) INTO num_cust FROM customers WHERE ord_date <= max_date 
CATCH
    ERROR "Error caught during SQL statement execution:", sqlca.sqlcode
END TRY
```

The `TRY - CATCH` block is a pseudo-statement: The compiler does not generate
p-code for the code lines containing these keywords. As result, it is not possible to set a [debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs.") break point directly on lines with
`TRY`, `CATCH` or `END TRY`. If you try to set a break
point on a `TRY`, `CATCH` or `END TRY`, the debugger
will set the break code on the next code line where a break point can be placed.

A `TRY` block can be compared with `WHENEVER ANY ERROR GOTO`.
Here is the equivalent of the previous code
example:

```
WHENEVER ANY ERROR GOTO catch_error
    SELECT COUNT(*) INTO num_cust FROM customers WHERE ord_date <= max_date 
    GOTO no_error 
LABEL catch_error:
WHENEVER ERROR STOP
    ERROR "Error caught during SQL statement execution:", sqlca.sqlcode
LABEL no_error
```

The `TRY` statement can be nested in other `TRY` statements.
In this example, the instruction in line #5 will be executed in case of
SQL error:

```
TRY
    TRY
        SELECT COUNT(*) INTO num_cust FROM customers 
    CATCH
        ERROR "Try block 2: ", sqlca.sqlcode
    END TRY
CATCH
    ERROR "Try block 1: ", sqlca.sqlcode
END TRY
```

Unlike the [`WHENEVER`](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") instruction, a
`TRY/CATCH` block exception handler stops the evaluation of an [expression](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.") at first error. For example, in the expression
`[(Pi() / 0) + Pi()]` (where `Pi()` is a user function returning the
number Pi), the function is called only once, because the expression evaluation stops at division by
zero error [-1202](../15_library-reference/4483-genero-bdl-errors.md).

The `TRY` statement takes control over the current [`WHENEVER [ANY] ERROR` handler](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module."), for the code
lines between the `TRY` and `CATCH` keywords. However, between the
`CATCH` and `END TRY` keywords, the current `WHENEVER`
handler gets the control back. For example, in the next code example, the program output will show
"In CATCH/END
TRY":

```
DEFINE where STRING
MAIN
    WHENEVER ERROR CALL error_handler
    TRY
        LET where = "In TRY/CATCH"
        CONNECT TO "dummy"
    CATCH
        LET where = "In CATCH/END TRY"
        CONNECT TO "dummy"
    END TRY
END MAIN

FUNCTION error_handler()
    DISPLAY where
END FUNCTION
```

The [WHENEVER ERROR RAISE](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") instruction can be used
module-wide to define the behavior when an exception occurs in a function that is called from a
`TRY` / `CATCH` block. If an exception occurs in a statement after the
`WHENEVER ERROR RAISE` instruction, the program flow returns from the function and
raises the exception as if it had occurred in the code of the caller. If the exception is thrown in
the `MAIN` block, the program stops because the exception cannot be processed by a
caller. In this example, the instruction in line #5 will be executed if an exception occurs in the
*cust\_report()*
function:

```
MAIN
    TRY
        CALL cust_report()
    CATCH
        ERROR "An error occurred during report execution: ", status
    END TRY
END MAIN

FUNCTION cust_report()
    WHENEVER ERROR RAISE
    START REPORT cust_rep ...
    ...
END FUNCTION
```

## Related links

**Related concepts**  

[SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")

[The sqlca diagnostic record](../10_sql-support/0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")
