---
title: "Handling nested transactions"
source: "fgl-topics/c_fgl_sql_programming_005.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Handling nested transactions"
type: "concept"
---

# Handling nested transactions

> You can manage nested transactions in different parts of a program.

A program can become very complex if it contains a lot of nested functions calls, doing SQL
processing within transactions. You may want to centralize transaction control commands in wrapper
functions. The fgldbutl.4gl library contains special functions to manage the
beginning and the end of a transaction with an internal counter, in order to implement nested
function calls inside a unique transaction.

```
MAIN
  IF a() <> 0 THEN
       ERROR "..."
  END IF
  IF b() <> 0 THEN
       ERROR "..."
  END IF
END MAIN

FUNCTION a()
  DEFINE s INTEGER
  LET s = db_start_transaction()
  UPDATE ...
  LET s = sqlca.sqlcode
  IF s = 0 THEN
     LET s = b()
  END IF
  LET s = db_finish_transaction((s==0))
  RETURN s
END FUNCTION

FUNCTION b()
   DEFINE s INTEGER
   LET s = db_start_transaction()
   UPDATE ...
   LET s = sqlca.sqlcode
   LET s = db_finish_transaction((s==0))
   RETURN s 
END FUNCTION
```

In this example, you see in the `MAIN` block that both functions
`a()` and `b()` can be called separately. However, the transaction SQL
commands will be used only if needed: When function `a()` is called, it starts the
transaction, then calls `b()`, which does not start the transaction since it was
already started by `a()`. When function `b()` is called directly, it
starts the transaction.

The function `db_finish_transaction()` is called with the expression
`(s==0)`, which is evaluated before the call. This allows you to write in one
line the equivalent of the following `IF` statement:

```
IF s==0 THEN
   LET s = db_finish_transaction(1)
ELSE
   LET s = db_finish_transaction(0)
END IF
```

## Related links

**Related concepts**  

[The MAIN block / function](../09_advanced-features/0789-the-main-block-function.md "The MAIN block is the starting point of the program.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
