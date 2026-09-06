---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagora_013_2.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Transactions handling"
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

## ORACLE

With ORACLE transactions:

- A transaction starts on first DML statement executed since last `COMMIT` or
  `ROLLBACK`.
- A transaction ends with a `COMMIT` or `ROLLBACK` statement.
- The current transaction is automatically committed when a DDL statement is executed.

ORACLE supports savepoints too. However, there are differences:

- Savepoints cannot be declared as `UNIQUE`
- Rollback must always specify the savepoint name
- You cannot release savepoints (`RELEASE SAVEPOINT`)

## Solution

Regarding transaction control instructions, BDL applications do not have to be modified in order
to work with ORACLE. The Informix behavior is simulated
with an autocommit mode in the ORACLE interface. A switch to the explicit commit mode is done when a
`BEGIN WORK` is performed by the BDL program.

When executing a DDL statement inside a
transaction, ORACLE automatically commits the transaction. Therefore,
you must extract the DDL statements from transaction blocks.

If you want to use savepoints, do not use the `UNIQUE` keyword in the savepoint
declaration, always specify the savepoint name in `ROLLBACK TO SAVEPOINT`, and do not
drop savepoints with `RELEASE SAVEPOINT`.

See also [SELECT FOR UPDATE](1406-select-for-update.md)

## Related links

**Related concepts**  

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
