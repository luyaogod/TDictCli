---
title: "fgldbutl.db_finish_transaction()"
source: "fgl-topics/c_fgl_utility_functions_DB_FINISH_TRANSACTION.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_finish_transaction()"
type: "concept"
---

# fgldbutl.db_finish_transaction()

> Terminates a nested transaction call.

## Syntax

```
FUNCTION db_finish_transaction(
   commit INTEGER )
  RETURNS INTEGER
```

1. commit is a boolean that indicates whether the transaction
   must be committed.

## Usage

This function encapsulates the `COMMIT WORK` or `ROLLBACK WORK`
instructions to end a transaction.

When the number of calls to
[`DB_START_TRANSACTION()`](2805-fgldbutl-db-start-transaction.md "Starts a nested transaction call.")
matches, this function executes a `COMMIT WORK` if the passed parameter is
`TRUE`; if the passed parameter is `FALSE`, it executes a
`ROLLBACK WORK`.

If the number of start/finish calls does not match, the function does nothing.

## Related links

**Related concepts**  

[Database transactions](../10_sql-support/1106-database-transactions.md "Database transaction concepts and handling.")
