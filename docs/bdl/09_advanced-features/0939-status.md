---
title: "status"
source: "fgl-topics/c_fgl_programs_STATUS.html"
breadcrumb: "Advanced features > Program registers > status"
type: "concept"
---

# status

> status is a predefined variable that contains the execution status of the last instruction.

## Syntax

```
status
```

## Usage

`status` is a predefined variable that contains the execution status of the
last program instruction.

`status` allows diagnostic information about procedural, interactive, and SQL
instructions to be obtained.

The data type of `status` is `INTEGER`.

While `status` can be modified by hand, it is not recommended except in specific
situations as shown in the status
example.

`status` is typically used with `WHENEVER ERROR CONTINUE` or
`WHENEVER ERROR CALL`, or `TRY/CATCH` blocks, to identify the type of
error that occurred.

`status` will be set for expression evaluation errors only when
`WHENEVER ANY ERROR` is used.

After an SQL statement execution, `status` contains the value of [`sqlca.sqlcode`](../10_sql-support/0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.").

`status` is set to an error code when an instruction produces an error, or it is
reset to zero when non-assignment instructions succeed.

While SQL or screen interaction statements set or reset `status` variable. When
using `WHENEVER ANY ERROR` or in the scope of a `TRY/CATCH` block,
expressions will set the `status` on error, but will not reset
`status` to zero on success: The `status` variable keeps the last
value until a new expression error in thrown.

A typical mistake is to test `status` after a `DISPLAY status`
instruction, written after an SQL statement:

```
WHENEVER ERROR CONTINUE
DELETE FROM _invalid_table_name_ where col = 1
WHENEVER ERROR STOP
DISPLAY "status:", status   -- this DISPLAY instruction reset status to zero
IF status<0 THEN            -- Will never be the case, since status==0
   DISPLAY "SQL Error!"
   EXIT PROGRAM 1
END IF
```

> **Tip:**
>
> Use `sqlca.sqlcode` for SQL error detection, and use `status` for
> other language instructions.

## Example

```
MAIN
  DISPLAY is_number(NULL)
  DISPLAY is_number("abc")
  DISPLAY is_number("-12.45")
END MAIN

FUNCTION is_number(s)
  DEFINE s STRING
  DEFINE f FLOAT, l_status INTEGER
  IF length(s)==0 THEN
     RETURN FALSE
  END IF
  WHENEVER ANY ERROR CONTINUE
  LET status=0 # Needed, as status won't be set if succeeds
  LET f = s
  LET l_status = status
  WHENEVER ANY ERROR CONTINUE
  IF l_status == 0 THEN
     RETURN TRUE
  ELSE
     RETURN FALSE
  END IF
END FUNCTION
```

## Related links

**Related concepts**  

[Example 2: SQL error handling with WHENEVER](0859-example-2-sql-error-handling-with-whenever.md "Example 2: SQL error handling with WHENEVER")

[WHENEVER directive](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.")

[Status variable handling](../05_upgrading/0391-status-variable-handling.md "I4GL and Genero BDL support the status variable differently.")
