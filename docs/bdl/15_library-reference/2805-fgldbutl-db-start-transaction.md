---
title: "fgldbutl.db_start_transaction()"
source: "fgl-topics/c_fgl_utility_functions_DB_START_TRANSACTION.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_start_transaction()"
type: "concept"
---

# fgldbutl.db_start_transaction()

> Starts a nested transaction call.

## Syntax

```
FUNCTION db_start_transaction()
  RETURNS INTEGER
```

## Usage

On most database engines, you can only have a unique transaction, that is started with
`BEGIN WORK` and ended with `COMMIT WORK` or
`ROLLBACK WORK`. But in a some cases, you may need to do complex nested
function calls, executing several SQL instruction that must all be grouped in a single
transaction. The nested transaction utility functions help you to implement this.

With this nested transaction technique, you encapsulate transaction start and end withing
the utility function. Custom functions doing SQL operations can then be reused in different parts of
your application. If the caller does not start the transaction, the called function will
automatically start and end the transaction.

The `db_start_transaction()` function encapsulates the `BEGIN WORK`
instruction to start a transaction, in order to implement nested transactions.

These transaction encapsulation functions are provided for special cases, where the function call
graph is complex. In general, it is recommended that you simply use the standard `BEGIN
WORK` / `COMMIT WORK` / `ROLLBACK WORK` instructions to
implement transaction blocks.

These transaction management functions execute a real transaction instruction at the
boundaries of the subsequent start/finish calls.

## Example

```
IMPORT FGL fgldbutl

MAIN
  DEFINE s INTEGER
  DATABASE mydb
  LET s = db_start_transaction() -- real BEGIN WORK
  IF s != 0 THEN DISPLAY "error 1" END IF
  WHENEVER ERROR CONTINUE
  UPDATE customer SET cust_name = 'Undef'
  WHENEVER ERROR STOP
  LET s = sqlca.sqlcode
  IF s != 0 THEN
     DISPLAY "error 2"
  ELSE
     LET s = do_update()
     IF s != 0 THEN DISPLAY "error 3" END IF
  END IF
  LET s = db_finish_transaction(s==0) -- real COMMIT or ROLLBACK WORK
  IF s != 0 THEN DISPLAY "error 4" END IF
END MAIN

FUNCTION do_update()
  DEFINE s INTEGER
  LET s = db_start_transaction() -- no SQL command (nested)
  IF s != 0 THEN
     DISPLAY "error 1.1"
  ELSE
     WHENEVER ERROR CONTINUE
     UPDATE customer SET cust_status = 'X'
     WHENEVER ERROR STOP
     LET s = sqlca.sqlcode
     IF s != 0 THEN
        DISPLAY "error 1.2"
     END IF
  END IF
  LET s = db_finish_transaction(s==0) -- no SQL command (nested)
  IF s != 0 THEN DISPLAY "error 1.3" END IF
  RETURN s
END FUNCTION
```

## Related links

**Related concepts**  

[Database transactions](../10_sql-support/1106-database-transactions.md "Database transaction concepts and handling.")
