---
title: "WHENEVER directive"
source: "fgl-topics/c_fgl_Exceptions_006.html"
breadcrumb: "Advanced features > Exceptions > WHENEVER directive"
type: "concept"
---

# WHENEVER directive

> Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.

## Syntax

```
WHENEVER exception-class
 exception-action
```

where [exception-class](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.") is one
of:

```
{ {ERROR|SQLERROR}
| ANY {ERROR|SQLERROR}
| NOT FOUND
| WARNING
}
```

and [exception-action](0852-exception-actions.md "Exception actions define the type of action to be taken when an exception occurs.") is one
of:

```
{ CONTINUE
| STOP
| CALL [module.]function 
| RAISE
| GOTO label
}
```

1. `ERROR` and `SQLERROR` are synonyms and have the same effect.
2. function can be any function name defined in the program.
3. module is the name of a module imported with `IMPORT FGL`.
4. label must be a label defined in the current program block (main,
   function or report routine).

## Usage

The `WHENEVER` directive defines the exception handling by associating an [*exception class*](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.") with an [*exception action*](0852-exception-actions.md "Exception actions define the type of action to be taken when an exception occurs.").

The `WHENEVER` directive is similar to a C preprocessor macro: Its effect is local
to the module and defines the error handling for the rest of the module, unless a new
`WHENEVER` directive for the same exception class is encountered by the
compiler, or a `TRY/CATCH` block is used.

A `WHENEVER` directive defines only the exception handling for the current module
and therefore does not affect callers. Use the directive `WHENEVER ERROR RAISE` to
propagate exceptions occuring in the module to the callers.

A `WHENEVER [ANY] [SQL]ERROR` directive overwrites a previous `WHENEVER
[ANY] [SQL]ERROR`. In the next example, the last `WHENEVER` takes control:

```
WHENEVER ERROR STOP
WHENEVER ANY ERROR CONTINUE
WHENEVER ANY ERROR STOP -- overwrites previous directives
...
```

If no `WHENEVER` directive is used, the default action for `WHENEVER
ERROR` is `STOP`. Stopping the program in case of error is the recommended
default. However, this default does not catch expression errors like type conversion errors.
Consider using the `fglrun.mapAnyErrorToError` FGLPROFILE entry, to catch conversion
errors. For more details, see [Default exception handling](0855-default-exception-handling.md "Default exception handling must be adapted to your programming pattern.").

This code example shows a typical `WHENEVER` directive usage:

```
WHENEVER ERROR CONTINUE
DROP TABLE mytable -- SQL error will be ignored 
CREATE TABLE mytable ( k INT, c VARCHAR(20) )
WHENEVER ERROR STOP
IF sqlca.sqlcode != 0 THEN
   ERROR "Could not create the table..."
END IF
```

Unlike the [`TRY/CATCH`](0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") block, a `WHENEVER ERROR CONTINUE` exception
handler does not stop the evaluation of an [expression](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language."). For example, in the expression `[(Pi() / 0) + Pi()]` (where
`Pi()` is a user function returning the number Pi), the function is called twice,
even if the division by zero produces error [-1202](../15_library-reference/4483-genero-bdl-errors.md). This behavior exists to be
compatible with Informix 4GL legacy code.

If an exception is thrown by a method or function that can be used in an expression (an typically
returns an object), this exception cannot be handled with `WHENEVER ERROR CALL /
CONTINUE`. In such case, only a `TRY/CATCH` block can be used to handle
exceptions.

Exception classes `ERROR` and `SQLERROR` are synonyms: In the
previous example it is also possible to use `WHENEVER SQLERROR` instead of
`WHENEVER ERROR`.

Actions for classes `[SQL]ERROR`, `WARNING` and `NOT
FOUND` can be set independently:

```
WHENEVER ERROR STOP
WHENEVER WARNING CONTINUE
WHENEVER NOT FOUND GOTO not_found_handler 
...
```

The effect of the `WHENEVER` directive is local to the current module, and applies
to all source lines following that directive (cross-function definitions). In the next example,
module1.4gl uses a `WHENEVER` directive that takes effect in the
second function, but does not affect the calling module
main.4gl:

```
$ head module1.4gl main.4gl
==> module1.4gl <==
FUNCTION function1()
    WHENEVER ANY ERROR CONTINUE -- applies to subsequent lines
END FUNCTION

FUNCTION function2()
    DEFINE x INTEGER
    LET x = "aaa"
    DISPLAY "function2: x = ", x, " status = ", status
END FUNCTION

==> main.4gl <==
IMPORT FGL module1
MAIN
    DEFINE x INTEGER
    WHENEVER ANY ERROR STOP
    CALL module1.function2()
    LET x = "aaa"
END MAIN

$ fglcomp main.4gl && fglrun main.42m
function2: x =             status =       -1213
Program stopped at 'main.4gl', line number 6.
FORMS statement error number -1213.
A character to numeric conversion process failed.
```

In the above example, the error -1213 in line 6 of main.4gl is expected,
because the `WHENEVER` directive of the main.4gl module
applies.

When using the `WHENEVER ... CALL function` directive, the program flow will go to
the specified function and return to the code block where the exception
occurred:

```
MAIN
    DEFINE x INTEGER
    WHENEVER ANY ERROR CALL error_handler
    -- WHENEVER handler takes effect
    LET x = 1/0
    DISPLAY "Back in MAIN..."
END MAIN

FUNCTION error_handler()
    DISPLAY "error_handler: ", status
END FUNCTION
```

Output:

```
error_handler:       -1202
Back in MAIN...
```

In a `WHENEVER ... CALL` directive, do not specify parentheses after the function
name.

A [`TRY/CATCH`](0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")
block takes precedence over the last `WHENEVER` directive, see the following
example:

```
MAIN
    DEFINE x INTEGER
    WHENEVER ANY ERROR CONTINUE
    -- WHENEVER handler takes effect
    LET x = 1/0
    DISPLAY "WHENEVER: ", status
    -- WHENEVER handler is hidden by TRY/CATCH block
    TRY
        LET x = 1/0
    CATCH
        DISPLAY "CATCH   : ", status
    END TRY
    -- WHENEVER handler takes again effect
    CALL func()
END MAIN

FUNCTION func()
    DEFINE x INTEGER
    LET x = 1/0
    DISPLAY "WHENEVER: ", status
END FUNCTION
```

Output:

```
WHENEVER:       -1202
CATCH   :       -1202
WHENEVER:       -1202
```

The `RAISE` option can be used to propagate exceptions to the caller, which
typically traps the error in a [TRY/CATCH
block](0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.")

main.4gl:

```
IMPORT FGL myutils
MAIN
    TRY
       -- Pass a NULL form name to get error -1110
       CALL myutils.open_form(NULL)
    CATCH
       DISPLAY "Error: ", status
    END TRY
END MAIN
```

myutils.4gl:

```
FUNCTION open_form(fn)
    DEFINE fn STRING
    WHENEVER ERROR RAISE -- Propagate exceptions to caller
    OPEN FORM f1 FROM fn
END FUNCTION
```

Note that `WHENEVER [ANY] ERROR RAISE` is not supported in a
`REPORT` routine.

## Related links

**Related concepts**  

[The sqlca diagnostic record](../10_sql-support/0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.")

[status](0939-status.md "status is a predefined variable that contains the execution status of the last instruction.")

## Child topics

- [Exception classes](0851-exception-classes.md): Exception classes define the kind of issues that can occur at runtime.
- [Exception actions](0852-exception-actions.md): Exception actions define the type of action to be taken when an exception occurs.
