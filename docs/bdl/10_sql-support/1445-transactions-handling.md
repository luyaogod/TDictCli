---
title: "Transactions handling"
source: "fgl-topics/c_fgl_odiagpgs_013_2.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data manipulation > Transactions handling"
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

## PostgreSQL

PostgreSQL supports transaction with savepoints:

- Transactions are started with `BEGIN WORK`.
- Transactions are validated with `COMMIT WORK`.
- Transactions are canceled with `ROLLBACK WORK`.
- Savepoints can be placed with `SAVEPOINT name`.
- Transactions can be rolled back to a savepoint with `ROLLBACK TO SAVEPOINT
  name`.
- Savepoints can be released with `RELEASE SAVEPOINT name`.
- Statements executed outside of a transaction are automatically committed.
- DDL statements can be executed (and canceled) in transactions.
- If an SQL error occurs in a transaction, the whole transaction is aborted.

Transactions in stored procedures: avoid using transactions in stored procedures to allow the
client applications to handle transactions, depending on the transaction model.

The main difference between Informix and PostgreSQL
resides in the fact that PostgreSQL cancels the entire transaction if an SQL error occurs in one of
the statements executed inside the transaction.

In such case, PostgreSQL will not allow to execute any other SQL statement until you rollaback
the current transaction or rollback to a previously declarated savepoint.

The error reported by PostgreSQL in such case is SQLSTATE `25P02`, with the
message `"current transaction is aborted, commands ignored until end of transaction
block"`.

The following code example illustrates this difference:

```
CREATE TABLE tab1 ( k INT PRIMARY KEY, c CHAR(10) )
WHENEVER ERROR CONTINUE
BEGIN WORK
INSERT INTO tab1 ( 1, 'abc' )
INSERT INTO tab1 ( 1, 'abc' ) -- PKey violation => SQL Error => TX aborted
INSERT INTO tab1 ( 2, 'def' ) -- PostgreSQL error SQLSTATE 25P02
COMMIT WORK
```

With Informix, this code will continue after the
second invalid `INSERT`, and finish the transaction, by creating two new rows.

With PostgreSQL, the table will remain empty after executing this piece of code, because the
server will rollback the whole transaction.

To workaround this problem in PostgreSQL you can use `SAVEPOINT` as described in
[Solution](1423-transactions-handling.md).

## Solution

Informix transaction handling commands are automatically
converted to PostgreSQL instructions to start, validate or cancel transactions.

Regarding the `BEGIN WORK`, `COMMIT WORK`, `ROLLBACK
WORK` and `SAVEPOINT` transaction control instructions, the BDL applications
do not have to be modified in order to work with PostgreSQL.

Check the code between `BEGIN WORK`/`COMMIT WORK` commands, to find
SQL statements that can raise an SQL error, that is handled to continue with the transaction. For
example, next (dirty code) works with Informix, but needs to be reviewed with
PostgreSQL:

```
MAIN
  DATABASE test1
  CREATE TABLE tab1 ( k INT PRIMARY KEY, c CHAR(10) )
  INSERT INTO tab1 VALUES ( 1, 'abc' ) -- have some data before TX
  WHENEVER ERROR CONTINUE -- want to trap errors and continue TX
  BEGIN WORK
  INSERT INTO tab1 VALUES ( 1, 'abc' ) -- dup pkey error
  IF sqlca.sqlcode<0 THEN -- let's try another pkey ...
     INSERT INTO tab1 VALUES ( 2, 'abc' )
  END IF
  COMMIT WORK
END MAIN
```

In the next code example, nothing needs to be modified, because any potential SQL error will make
the program flow jump to the `CATCH` clause, and rollback the entire transaction.
Using such `TRY`/`CATCH` block is a proper coding pattern with any
kind of database engine:

```
MAIN
  DATABASE test1
  CREATE TABLE tab1 ( k INT PRIMARY KEY, c CHAR(10) )
  BEGIN WORK  -- Can be outside TRY (should never fail if coded properly)
  TRY
     INSERT INTO tab1 VALUES ( 1, 'abc' )
     INSERT INTO tab1 VALUES ( 1, 'abc' ) -- dup pkey
     INSERT INTO tab1 VALUES ( 2, 'def' )
     COMMIT WORK
  CATCH
     ROLLBACK WORK
  END TRY
END MAIN
```

However, if the transaction must continue on SQL error, the SQL statements that can potentially
raise an SQL error must be protected with a `SAVEPOINT`. If an error occurs, just
rollback to the savepoint:

```
MAIN
  DATABASE test1
  CREATE TABLE tab1 ( k INT PRIMARY KEY, c CHAR(10) )
  INSERT INTO tab1 VALUES ( 1, 'abc' )
  WHENEVER ERROR CONTINUE
  BEGIN WORK
  CALL sql_protect()
  INSERT INTO tab1 VALUES ( 1, 'abc' ) -- dup key
  CALL sql_unprotect()
  IF sqlca.sqlcode<0 THEN
     INSERT INTO tab1 VALUES ( 2, 'abc' )
  END IF
  COMMIT WORK
END MAIN

FUNCTION sql_protect()
  IF fgl_db_driver_type()!="pgs" THEN
     RETURN
  END IF
  SAVEPOINT _sql_protect_
END FUNCTION

FUNCTION sql_unprotect()
  IF fgl_db_driver_type()!="pgs" THEN
     RETURN
  END IF
  IF sqlca.sqlcode < 0 THEN
     ROLLBACK WORK TO SAVEPOINT _sql_protect_
  ELSE
     RELEASE SAVEPOINT _sql_protect_
  END IF
END FUNCTION
```

> **Note:**
>
> If you want to use savepoints, do not use the `UNIQUE` keyword
> in the savepoint declaration, always specify the savepoint name in `ROLLBACK TO
> SAVEPOINT`, and do not drop savepoints with `RELEASE SAVEPOINT`.

## Related links

**Related concepts**  

[Database transactions](0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
