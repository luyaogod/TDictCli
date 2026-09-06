---
title: "Understanding exceptions"
source: "fgl-topics/c_fgl_Exceptions_002.html"
breadcrumb: "Advanced features > Exceptions > Understanding exceptions"
type: "concept"
---

# Understanding exceptions

> Exceptions are abnormal runtime events that can be trapped for control.

If an instruction executes abnormally, the runtime system throws exceptions (aka runtime error),
that can be trapped by the program to process the error and continue the execution.

Exceptions handling should be rare in a program. If something unexpected occurs such as a missing
database table, the program should stop. Use the default exception handler that stops the program in
case of error, and implement exception handlers in specific areas of the code, to catch runtime
errors that are possible, such as SQL lock timeout errors or invalid file names entered by the
user.

Exceptions can be trapped by a [`WHENEVER`](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") exception handler or by a [`TRY/CATCH`](0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") block. The `WHENEVER` exception is provided for
compatibility with Informix 4GL, consider using `TRY/CATCH` blocks instead.

If an exception is thrown by a method or function that can be used in an expression (an typically
returns an object), this exception cannot be handled with `WHENEVER ERROR CALL /
CONTINUE`. In such case, use a `TRY/CATCH` block can be used to handle
exceptions.

Some specific errors such as error [-1338](../15_library-reference/4483-genero-bdl-errors.md) cannot be trapped. In such case, the program will always stop and show the error, even
if a `WHENEVER` exception handler or `TRY/CATCH` block is used.

A Genero exception is identified by its number and has a description. For a complete list of BDL
errors, see [Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.").

Exception handlers are typically used to detect database errors when executing SQL statement. For
more details, see [SQL execution diagnostics](../10_sql-support/0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.").
