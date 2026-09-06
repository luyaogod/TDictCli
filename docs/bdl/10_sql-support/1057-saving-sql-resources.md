---
title: "Saving SQL resources"
source: "fgl-topics/c_fgl_sql_programming_055.html"
breadcrumb: "SQL support > SQL programming > SQL performance > Saving SQL resources"
type: "concept"
---

# Saving SQL resources

> SQL cursors and prepared statement consume resources that should be freed when useless.

To write efficient SQL in your programs, you can use prepared SQL.
However, when using prepared SQL, an SQL statement handle is allocated
on the client and server side, consuming resources. Depending to the
database type, this can be a few bytes or a significant amount of
memory. When executing several static SQL statements, the same SQL
statement handle is reused and thus less memory is needed.

The BDL language allows you to use either static SQL or dynamic SQL,
so it's in your hands to choose memory or performance. However, in some
cases the same code will be used by different kinds of programs,
needing either low resource usage or good performance. In many OLTP
applications you can actually distinguish two type of programs:

- Programs where memory usage is not a problem but good performance
  is needed (typically, batch programs executed as a unique instance
  during the night).
- Programs where performance is less important but memory usage
  must be limited (typically, interactive programs executed as multiple
  instances for each application user).

To reuse the same code for interactive programs and batch programs,
use the following pattern:

1. Define a local module variable as an indicator for the prepared
   statement.
2. Write a function returning the type of program (for example,
   'interactive' or 'batch' mode).
3. Then, in a reusable function using SQL statements, prepare and
   free the statement based on the type of program, as shown in the next
   example.

```
PRIVATE DEFINE up_prepared BOOLEAN

FUNCTION getUserPermissions( username )
  DEFINE username VARCHAR(20)
  DEFINE cre, upd, del CHAR(1)

  IF NOT up_prepared THEN
     PREPARE up_stmt FROM "SELECT can_create, can_update, can_delete"
                             || " FROM user_perms WHERE name = ?"
     LET up_prepared = TRUE
  END IF

  EXECUTE up_stmt USING username INTO cre, upd, del

  IF isInteractive() THEN
     FREE up_stmt
     LET up_prepared = FALSE
  END IF

  RETURN cre, upd, del

END FUNCTION
```

The first time this function is called, `up_prepared` is `FALSE`,
the statement is prepared and the flag is set to `TRUE`. The statement is then
executed, values are fetched. If the program is interactive, the statement is freed and the variable
is set back to `FALSE`, forcing statement preparation in the next call of this
function. If the program is not interactive, the statement handle is kept and the next call will not
require to re-prepare the statement.

Modern database client interfaces support "deferred prepare", to postpone the SQL statement
preparation to the first execution: This avoids a roundtrip with the DB server. In such case, the
benefit of using prepared statements makes only sense when the SQL is executed many times.

## Related links

**Related concepts**  

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
