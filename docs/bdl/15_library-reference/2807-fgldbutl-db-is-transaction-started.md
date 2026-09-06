---
title: "fgldbutl.db_is_transaction_started()"
source: "fgl-topics/c_fgl_utility_functions_DB_IS_TRANSACTION_STARTED.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_is_transaction_started()"
type: "concept"
---

# fgldbutl.db_is_transaction_started()

> Indicates whether a nested transaction call is started.

## Syntax

```
FUNCTION db_is_transaction_started()
  RETURNS INTEGER
```

## Usage

The function returns `TRUE` if a transaction was started with [`db_start_transaction()`](2805-fgldbutl-db-start-transaction.md "Starts a nested transaction call."), and was not yet finished with a call to
the [`db_finish_transaction()`](2806-fgldbutl-db-finish-transaction.md "Terminates a nested transaction call.") function.

## Related links

**Related concepts**  

[Database transactions](../10_sql-support/1106-database-transactions.md "Database transaction concepts and handling.")
