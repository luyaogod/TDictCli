---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagmsv_015.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Database concepts > Transactions handling"
type: "concept"
description: "Informix® With the Informix native mode (non ANSI): Transactions blocks start with BEGIN WORK and terminate with COMMIT WORK or ROLLBACK WORK . Statements executed outside a transaction are ..."
---

# Transactions handling

## Informix®

With the Informix native mode (non ANSI):

- Transactions blocks start with `BEGIN WORK` and terminate with `COMMIT
  WORK` or `ROLLBACK WORK`.
- Statements executed outside a transaction are automatically committed.
- DDL statements can be executed (and canceled) in transactions.

```
UPDATE tab1 SET ...   -- auto-committed
BEGIN WORK            -- start of TX block
UPDATE tab1 SET ...
UPDATE tab2 SET ...
...
COMMIT WORK           -- end of TX block
```

Informix version 11.50 introduces
savepoints:

```
SAVEPOINT name [UNIQUE]
ROLLBACK [WORK] TO SAVEPOINT [name] ]
RELEASE SAVEPOINT name
```

## Microsoft™ SQL Server

Microsoft SQL Server supports named and nested
transactions:

- Transactions are started with `BEGIN TRANSACTION [name]`.
- Transactions are validated with `COMMIT TRANSACTION
  [name]`.
- Transactions are canceled with `ROLLBACK TRANSACTION
  [name]`.
- Savepoints can be placed with `SAVE TRANSACTION name`.
- Transactions can be rolled back to a savepoint with `ROLLBACK TRANSACTION TO
  name`.
- Savepoints can not be released.
- Statements executed outside of a transaction are automatically committed (autocommit mode). This
  behavior can be changed with "`SET IMPLICIT_TRANSACTION ON`".
- DDL statements have specific behavior in transactions blocks: If a DDL statement fails, the
  whole transaction is canceled.

## Solution

Informix transaction
handling commands are automatically converted to Microsoft SQL Server instructions to start,
validate or cancel transactions.

Regarding the transaction control instructions, the BDL applications do not have to be modified
to work with Microsoft SQL Server.

> **Important:**
>
> If you want to use savepoints, do not use the `UNIQUE`
> keyword in the savepoint declaration, always specify the savepoint name in `ROLLBACK TO
> SAVEPOINT`, and do not drop savepoints with `RELEASE SAVEPOINT`.

## Related links

**Related concepts**  

[Concurrency management](1266-concurrency-management.md "Concurrency management")

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
