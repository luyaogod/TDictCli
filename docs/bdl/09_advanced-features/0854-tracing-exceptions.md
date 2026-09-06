---
title: "Tracing exceptions"
source: "fgl-topics/c_fgl_Exceptions_010.html"
breadcrumb: "Advanced features > Exceptions > Tracing exceptions"
type: "concept"
---

# Tracing exceptions

> Exception can be logged in a file when using the STARTLOG() function.

Exceptions are automatically logged in a file, if all the following conditions are true:

- The [STARTLOG](../15_library-reference/2790-startlog.md "Initializes error logging and opens the error log file passed as the parameter.") function has been
  previously called to specify the name of the exception logging file.
- The [exception action](0852-exception-actions.md "Exception actions define the type of action to be taken when an exception occurs.") is set to
  `CALL`, `GOTO` or `STOP`. Exceptions are not logged
  when the action is `CONTINUE` or `RAISE`.
- The [exception class](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.") is an `ERROR`,
  `ANY ERROR` or `WARNING`. `NOT FOUND` exceptions cannot
  be logged.

In other words, errors will not be logged in the case of `WHENEVER { [ANY]
ERROR | WARNING } CONTINUE`, or when controlled by a
`TRY`/`CATCH` block.

Each log entry contains:

- The system-time
- The location of the related instruction (source-file, line)
- The error-number
- The text of the error message, giving human-readable details for the exception

It is good practice to generate the call stack trace with the [base.Application.getStackTrace()](../15_library-reference/2976-base-application-getstacktrace.md "Returns the function call stack trace.") function,
typically in the handler used `WHENEVER ERROR CALL my_handler()`.

## Related links

**Related concepts**  

[SQL error identification](../10_sql-support/0989-sql-error-identification.md "Identify SQL exceptions in your programs with sqlca.sqlcode.")

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
