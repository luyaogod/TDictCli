---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagmys_014.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Database concepts > Transactions handling"
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

## Oracle® MySQL and MariaDB

- Transactions are started with `START TRANSACTION`.
- Transactions are validated with `COMMIT [WORK]`.
- Transactions are canceled with `ROLLBACK [WORK]`.
- Savepoints can be placed with `SAVEPOINT name`.
- Transactions can be rolled back to a savepoint with `ROLLBACK [WORK] TO [SAVEPOINT]
  name`.
- Savepoints can be released with `RELEASE SAVEPOINT name`.
- Statements executed outside of a transaction are automatically committed.
- DDL statements cannot be combined with other statements within a transaction block.

## Solution

Informix transaction handling commands are
automatically converted to MySQL instructions to start, validate or cancel transactions.

MySQL does not support transactions by default. You must set the server system parameter
`table_type=InnoDB`.

Regarding the transaction control instructions, the BDL applications do not have to be modified
in order to work with MySQL, as long as you have a transaction manager installed with MySQL.

If you want to use savepoints, do not use the `UNIQUE` keyword in the savepoint
declaration, always specify the savepoint name in `ROLLBACK TO SAVEPOINT`, and do not
drop savepoints with `RELEASE SAVEPOINT`.

## Related links

**Related concepts**  

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
