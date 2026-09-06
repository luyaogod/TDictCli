---
title: "Avoiding long transactions"
source: "fgl-topics/c_fgl_sql_programming_052.html"
breadcrumb: "SQL support > SQL programming > SQL performance > Avoiding long transactions"
type: "concept"
---

# Avoiding long transactions

> Long transactions consume resources and decrease concurrent data access.

Old applications based on IBM®
Informix® database without
transaction logging might perform long running SQL modifications.

With recent database engines, using huge transactions can lead
to errors because of transaction log buffer overflow. For example,
if a table holds many rows, a "`DELETE FROM table`"
might produce a "snapshot too old" error in Oracle, if the rollback
segments are too small.

Therefore, you must avoid long transactions when connected to a
database using transactions:

- keep transactions as short as possible.
- access the least amount of data possible while in a transaction.
- split a long transaction into many short transactions. Use a loop
  to handle each block.
- to delete all rows from a table use the "`TRUNCATE TABLE`"
  instruction instead of "`DELETE FROM`" (Not for
  all vendors).
- In the end, increase the size of the transaction log to avoid
  filling it up.

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[Performance with transactions](1054-performance-with-transactions.md "Commit database changes by blocks of transaction speeds performance with some database servers.")
