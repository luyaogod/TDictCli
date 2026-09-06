---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagdmg_013.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Database concepts > Transactions handling"
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

## Dameng®

Transactions in Dameng database engine:

- Beginning of transactions are implicit; two transactions are delimited by
  `COMMIT` or `ROLLBACK`.
- DDL statements can be executed in transactions, but when doing a `ROLLBACK`, only
  data changes will be reverted.

Savepoints in Dameng:

- `SAVEPOINT` and `ROLLBACK TO SAVEPOINT` instructions are
  supported.
- Dameng does not support `RELEASE SAVEPOINT`.

## Solution

The Informix behavior of `BEGIN WORK`,
`COMMIT WORK` and `ROLLBACK WORK` transaction instructions is emulated
with an autocommit mode in the Dameng database driver:

If no `BEGIN WORK` is performed, every SQL statement will be followed by an
implicit `COMMIT` executed by the Dameng ODI driver.

If the `BEGIN WORK` instruction is performed, the ODI drivers sets an internal
flag, and no implicit `COMMIT` is performed after SQL statements, until the
`COMMIT WORK` or `ROLLBACK WORK` instruction is used, resulting
respectively in a native `COMMIT` or `ROLLBACK`.

As result, the application code using the Informix-style transatoin control commands does not
need to be modified in order to work with Dameng.

When using savepoints in a transaction, always specify the savepoint name in `ROLLBACK TO
SAVEPOINT`. Do not use `RELEASE SAVEPOINT`.

See also [SELECT FOR UPDATE](1251-select-for-update.md)

## Related links

**Related concepts**  

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
