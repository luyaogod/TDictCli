---
title: "Performance with transactions"
source: "fgl-topics/c_fgl_sql_programming_051.html"
breadcrumb: "SQL support > SQL programming > SQL performance > Performance with transactions"
type: "concept"
---

# Performance with transactions

> Commit database changes by blocks of transaction speeds performance with some database servers.

To mimic the IBM® Informix® auto-commit behavior with an ANSI compliant RDBMS
like Oracle®, the database
driver must perform an implicit commit after each statement execution, if the SQL statement
is not inside a transaction block. This generates unnecessary database operations and can
slow down big loops. To avoid this implicit commit, you can control the transaction with
`BEGIN WORK` / `COMMIT WORK` around the code containing a lot
of SQL statement execution.

This technique is especially recommended with SQLite, because the SQLite database library
performs a lot of operations during a commit.

For example, the following loop will generate 2000 basic SQL operations ( 1000 inserts
plus 1000 commits ):

```
PREPARE s FROM "INSERT INTO tab VALUES ( ?, ? )"
FOR n=1 TO 100
   EXECUTE s USING n, c   -- Generates implicit COMMIT
END FOR
```

You can improve performance if you put a transaction block around
the loop:

```
PREPARE s FROM "INSERT INTO tab VALUES ( ?, ? )"
BEGIN WORK
FOR n=1 TO 100
   EXECUTE s USING n, c   -- In transaction -> no implicit COMMIT
END FOR
COMMIT WORK
```

With this code, only 1001 basic SQL operations will be executed ( 1000 inserts plus 1 commit ).

However, you must take care when generating large transactions because all modifications
are registered in transaction logs. This can result in a lack of database server resources
("transaction too long" errors, for example) when the number of operations is very big. If the
SQL operation does not require a unique transaction for database consistency reasons, you can
split the operation into several transactions, as in this
example:

```
PREPARE s FROM "INSERT INTO tab VALUES ( ?, ? )"
BEGIN WORK
FOR n=1 TO 100
   IF n MOD 10 == 0 THEN
      COMMIT WORK
      BEGIN WORK
   END IF
   EXECUTE s USING n, c   -- In transaction -> no implicit COMMIT
END FOR
COMMIT WORK
```

Note that the `LOAD` instruction automatically starts a transaction, if not
yet initiated. Therefore there is no need to enclose the `LOAD` statement
within a `BEGIN WORK` / `COMMIT WORK`, except if other SQL
statements are part of the transaction and need to be processed as a single atomic database
change.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
