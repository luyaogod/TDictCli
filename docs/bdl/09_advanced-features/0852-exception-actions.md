---
title: "Exception actions"
source: "fgl-topics/c_fgl_Exceptions_003.html"
breadcrumb: "Advanced features > Exceptions > WHENEVER directive > Exception actions"
type: "concept"
---

# Exception actions

> Exception actions define the type of action to be taken when an exception occurs.

There are five exception actions that can be executed if an exception is raised:

**`STOP`**

The program is immediately terminated. A message is displayed to the standard error with the
location of the related statement, the error number, and the details of the exception.

**`CONTINUE`**

The program continues normally. The exception is ignored, but can be checked by testing the
[`status`](0939-status.md "status is a predefined variable that contains the execution status of the last instruction.") register, or the [sqlca.sqlcode](../10_sql-support/0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.") register for SQL errors.

**`CALL`
exception-function**

The function exception-function is called by the runtime system. The [function](../08_language-basics/0761-functions.md "Describes user defined functions.") can be defined in any module, and must have zero
parameters and zero return values. The `status` variable will be set to the
corresponding error number.

**`GOTO`
exception-label**

The program execution continues at the label identified by exception-label,
as if a [`GOTO`](../08_language-basics/0681-goto.md "The GOTO instruction transfers program control to a labeled line within the same program block.") instruction was issued
after trapping the exception.

**`RAISE`**

This statement instructs the runtime system that the exception must propagated to the calling
function.
> **Important:**
>
> The `WHENEVER [ANY] ERROR
> RAISE` is not supported in a `REPORT` routine.

In the following code example, the `WHENEVER` instruction specifies that the
program flow must continue on [`ANY
ERROR`](0851-exception-classes.md "Exception classes define the kind of issues that can occur at runtime.") exception class:

```
WHENEVER ANY ERROR CONTINUE
```

## Related links

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
