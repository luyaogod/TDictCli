---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagsqt_012_2.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data manipulation > Transactions handling"
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

## SQLite

With SQLite:

- Individual SQL statements are auto-committed.
- Transactions start with `BEGIN TRANSACTION` and end with `COMMIT
  TRANSACTION` or `ROLLBACK TRANSACTION`.
- DDL statements can be executed (and canceled) in transaction blocks.

SQLite supports savepoints with some differences compared to Informix:

- `SAVEPOINT` can be used instead of `BEGIN TRANSACTION`. In this
  case, `RELEASE` is like a `COMMIT`.
- The syntax of a rollback to the savepoint is `ROLLBACK [TRANSACTION] TO [SAVEPOINT]
  name`.
- The syntax of a release of the savepoint is `RELEASE [SAVEPOINT]
  name`.
- Rollback must always specify the savepoint name.
- You cannot rollback to a savepoint if cursors are opened.
- In SQLite versions prior to 3.7, you cannot rollback are transaction if a cursor is open.

## Solution

Regarding transaction control instructions, BDL applications do not have to be modified in order
to work with SQLite. The `BEGIN WORK`, `COMMIT WORK` and
`ROLLBACK WORK` commands are translated the native commands of SQLite.

> **Note:**
>
> If you want to use savepoints, always specify the savepoint name in
> `ROLLBACK TO SAVEPOINT` and do not open cursors during transactions using savepoints.
> If you are using an SQLite versions prior to 3.7, it is not possible to perform a `ROLLBACK
> WORK` if a cursor (with hold) is currently open.

See also [SELECT
FOR UPDATE](1498-select-for-update.md)

## Related links

**Related concepts**  

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
