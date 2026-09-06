---
title: "Example 2: SQL error handling with WHENEVER"
source: "fgl-topics/c_fgl_Exceptions_016.html"
breadcrumb: "Advanced features > Exceptions > Examples > Example 2: SQL error handling with WHENEVER"
type: "concept"
description: "This code shows a typical SQL error handling block. It uses WHENEVER ERROR CONTINUE before executing SQL statements, tests the sqlca.sqlcode register for error after the INSERT SQL instruction, and ..."
---

# Example 2: SQL error handling with WHENEVER

This code shows a typical SQL error handling block. It uses `WHENEVER ERROR
CONTINUE` before executing SQL statements, tests the `sqlca.sqlcode` register
for error after the `INSERT` SQL instruction, and resets the default exception
handler with `WHENEVER ERROR STOP`:

```
MAIN
  DEFINE a_name VARCHAR(50)
  DEFINE cnt INTEGER

  # In the DATABASE statement, no error should occur...
  DATABASE stores 

  # Next INSERT may fail, if the user enters empty name and
  # the table has a NOT NULL constraint on cust_name column:
  WHENEVER ERROR CONTINUE
  PROMPT "Enter a customer name:" FOR a_name
  INSERT INTO customer (cust_name) VALUES (a_name)
  IF sqlca.sqlcode THEN
     DISPLAY "SQL Error occurred:", sqlca.sqlcode
     EXIT PROGRAM 1
  END IF
  WHENEVER ERROR STOP

  # Other instructions would stop program in case of error
  SELECT COUNT(*) INTO cnt FROM customer

END MAIN
```

Program output in case of invalid table name:

```
SQL Error occurred:      -217
```

Note that the above code can also be written
with a [`TRY/CATCH` block](0860-example-3-typical-try-catch-block.md).
