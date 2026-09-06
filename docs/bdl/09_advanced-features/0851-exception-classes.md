---
title: "Exception classes"
source: "fgl-topics/c_fgl_Exceptions_005.html"
breadcrumb: "Advanced features > Exceptions > WHENEVER directive > Exception classes"
type: "concept"
---

# Exception classes

> Exception classes define the kind of issues that can occur at runtime.

During a program execution, the following kind of issues may happen:

1. Language errors : In case of invalid use of language instructions or API. For
   example, when specifying an unexisting form with the `OPEN FORM` instruction, or when
   opening an unexisting file with `base.Channel.openFile()`.
2. Expression errors : When an expression evaluation produces an error. For example,
   when a character string cannot be converted to a number type.
3. SQL errors : If an SQL statement execution or SQL instruction returns an error. For
   example, when an `INSERT` statement violates a `UNIQUE` table
   constraint, or when doing a `FETCH` on an unopened cursor.
4. SQL not found : When an SQL statement returns no data row(s). For example, when a
   `FETCH` statement reaches the end of the SQL cursor result set.
5. SQL Warnings : If an SQL statement produces warning(s). For example, when fetching a
   character string value that gets truncated because the target `VARCHAR` variable is
   not long enough.

The default action can be changed by specifying the exception class in the
`WHENEVER` instruction.

The following exception classes exist:

**`ERROR` or `SQLERROR`**

Triggered by a Language error or SQL error. Default exception action
is [`STOP`](0852-exception-actions.md). The `WHENEVER [SQL]ERROR` and `WHENEVER ANY
[SQL]ERROR` directives are mutually exclusive.

**`ANY ERROR` or `ANY SQLERROR`**

Triggered by a Language error, SQL error or Expression
error. The `WHENEVER ANY [SQL]ERROR` and `WHENEVER [SQL]ERROR`
directives are mutually exclusive. `WHENEVER ANY ERROR` has no default action,
because the default is `WHENEVER ERROR STOP`, where Expression errors do
not stop the program execution.

**`NOT FOUND`**

Triggered by SQL statements returning status `NOTFOUND`. Default exception action
is [`CONTINUE`](0852-exception-actions.md). The action for `WHENEVER NOT FOUND` directive can
be specified independently from other exception classes.

**`WARNING`**

SQL statements setting `sqlca.sqlawarn` flags. Default exception action is [`CONTINUE`](0852-exception-actions.md).
The action for `WHENEVER WARNING` directive can be specified independently from other
exception classes.

In the following code example, the `WHENEVER` instruction defines the behavior for
the `ANY ERROR` exception
class:

```
WHENEVER ANY ERROR CONTINUE
```

## Related links

**Related reference**  

[Genero BDL errors](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.")
