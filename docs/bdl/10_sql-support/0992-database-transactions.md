---
title: "Database transactions"
source: "fgl-topics/c_fgl_sql_programming_014.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Database transactions"
type: "concept"
---

# Database transactions

> Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.

The BDL instructions to define a transaction block are:

- `BEGIN WORK`
- `COMMIT WORK`
- `ROLLBACK WORK`

Additional BDL instructions related to transaction management is available, such as `SET
LOCK MODE ...`, `SET ISOLATION ...`, `SAVEPOINT`, etc. For a
complete description, see [Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

When performing a transaction instruction, the database drivers execute the corresponding native
SQL instruction (or database client API call) to begin, commit or rollback a transaction.

In this example a basic transaction block executes inside a [`TRY/CATCH` block](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") to rollback the transaction
in case of SQL error:

```
MAIN
   CONNECT TO ...
   TRY
     BEGIN WORK
     UPDATE tab1 SET col1 = 'aaa' WHERE pkey = 123
     UPDATE tab2 SET col2 = 'bbb' WHERE pkey = 456
     ...
     COMMIT WORK
   CATCH
     ROLLBACK WORK
   END TRY
END MAIN
```

## Related links

**Related concepts**  

[Concurrent data access](0991-concurrent-data-access.md "Understanding concurrent data access and data consistency.")

[Transaction blocks across connections](1005-transaction-blocks-across-connections.md "Transaction blocks manage transactions when connected to several database servers.")

[Handling nested transactions](1004-handling-nested-transactions.md "You can manage nested transactions in different parts of a program.")

[Avoiding long transactions](1055-avoiding-long-transactions.md "Long transactions consume resources and decrease concurrent data access.")

[Performance with transactions](1054-performance-with-transactions.md "Commit database changes by blocks of transaction speeds performance with some database servers.")
