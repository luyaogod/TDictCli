---
title: "Solution 2: Generate serial numbers from your own sequence table"
source: "fgl-topics/c_fgl_sql_programming_075.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Auto-incremented columns (serials) > Solution 2: Generate serial numbers from your own sequence table"
type: "concept"
description: "Purpose The goal is to generate unique INTEGER or BIGINT numbers. These numbers will usually be used for primary keys. Prerequisites The database must use transactions. This is usually the case with ..."
---

# Solution 2: Generate serial numbers from your own sequence table

## Purpose

The goal is to generate unique `INTEGER` or `BIGINT` numbers.
These numbers will usually be used for primary keys.

## Prerequisites

1. The database must use transactions. This is usually the case with non-Informix databases, but
   IBM®
   Informix® databases default to auto commit mode. Make sure
   your IBM
   Informix database allows transactions.
2. The sequence generation must be called inside a transaction (`BEGIN
   WORK` / `COMMIT
   WORK`).
3. The transaction isolation level must guarantee that a row UPDATED in a transaction cannot be
   read or written by other db sessions until the transaction has ended (typically, `COMMITTED
   READ` is ok, but some db servers require a higher isolation level)
4. The lock wait mode must be `WAIT`. This is usually the case in non-Informix
   databases, but Informix defaults to `NOT WAIT`. You must change the lock wait mode
   with "`SET LOCK MODE TO WAIT`" or "`WAIT seconds`" when using
   IBM
   Informix.
5. Other applications or stored procedures must implement the same technique when inserting records
   in tables with auto-incremented columns.

## Principle

A dedicated table named "`SEQREG`"
is used to register sequence numbers. The key
is the name of the sequence. This name will usually be the table
name the sequence is generated for. In short,
this table contains a primary key that identifies
the sequence and a column containing the last generated
number.

The uniqueness is granted by the concurrency
management of the database server. The first
executed instruction is an `UPDATE` that sets an
exclusive lock on the `SEQREG` record.
When two processes try to get a sequence at
the same time, one will wait for the other until its transaction
is finished.

## Implementation

The "`fgldbutl.4gl`"
utility library implements a function called
"`db_get_sequence()`" which generates a new
sequence. You must create the `SEQREG` table
as described in the `fgldbutl.4gl` source
found in `FGLDIR/src`,
and make sure that every user has the privileges
to access and modify this table.

In order to guarantee the uniqueness
of the generated number, the call to
`db_get_sequence()` must be done inside a transaction
block that includes the `INSERT` statement.
Concurrent db sessions must wait for each
other in case of conflict and the transaction
isolation level must be high enough to make sure that the row
of the sequence table will not be read or
written by other db sessions until the transaction
end.

## Example

```
IMPORT FGL fgldbutl
DEFINE rec RECORD
           id    INTEGER,
           name  CHAR(100)
     END RECORD
...
BEGIN WORK
LET rec.id = db_get_sequence( "CUSTID" )
INSERT INTO CUSTOMER ( CUSTID, CUSTNAME ) VALUES ( rec.* )
COMMIT WORK
```

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")

[SET LOCK MODE](1114-set-lock-mode.md "Defines the behavior of the program that tries to access a locked row or table.")

[SET ISOLATION](1113-set-isolation.md "Defines the transaction isolation level for the current connection.")

**Related reference**  

[fgldbutl: Database utility module](../15_library-reference/2801-fgldbutl-database-utility-module.md "fgldbutl: Database utility module")
