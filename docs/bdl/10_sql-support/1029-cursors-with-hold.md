---
title: "Cursors WITH HOLD"
source: "fgl-topics/c_fgl_sql_programming_078.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Cursors WITH HOLD"
type: "concept"
---

# Cursors WITH HOLD

> Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.

## Type of WITH HOLD cursors

Make sure to distinguish cursors declared `WITH HOLD` for regular
`SELECT` statements, and `WITH HOLD` cursors declared for a
`SELECT` statements using the `FOR UPDATE` clause.

## Regular WITH HOLD cursors (without FOR UPDATE clause in SELECT)

Most database brands allow cursors to remain open across transactions, as long as the
`SELECT` statement is not defined with a `FOR UPDATE`
clause.

```
DECLARE c1 CURSOR WITH HOLD FOR SELECT * FROM customers
```

Therefore, a cursor declared `WITH HOLD` with a `SELECT` statement
not using the `FOR UPDATE` clause can be used with most databases. However, some
databases close any cursor (even declared without `FOR UPDATE`), when a transaction
is rolled back, or when the cursor is opened inside the transaction block.

| Database Server Type | WITH HOLD support |
| --- | --- |
| IBM® Informix® | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft™ SQL Server | Yes, [see details](1302-cursors-with-hold.md) |
| Oracle® MySQL / MariadDB | Yes, [see details](1349-cursors-with-hold.md) |
| Oracle Database Server | Yes, [see details](1405-cursors-with-hold.md) |
| PostgreSQL | Partial, [see details](1456-cursors-with-hold.md) |
| SQLite | Yes, [see details](1499-cursors-with-hold.md) |
| Dameng® | Yes, [see details](1251-select-for-update.md) |

## Cursors WITH HOLD using SELECT FOR UPDATE

IBM Informix
supports `WITH HOLD` cursors using the `FOR UPDATE`
clause:

```
DECLARE c1 CURSOR WITH HOLD FOR SELECT * FROM customers FOR UDPATE
```

Such cursors can remain open across transactions (when using `FOR UPDATE`, locks
are released at the end of a transaction, but the `WITH HOLD` cursor is not
closed).

Not all database brands allow "for update" cursors to remain open across transactions: The SQL
standards recommend closing `FOR UPDATE` cursors and release locks at the end of a
transaction.

Most database servers close `FOR UPDATE` cursors when a `COMMIT
WORK` or `ROLLBACK WORK` is done.

All database servers release locks when a transaction ends.

| Database Server Type | WITH HOLD FOR UPDATE support |
| --- | --- |
| IBM Informix | Yes, [native SQL feature](1186-what-are-the-supported-ibm-informix-sql-features.md) |
| Microsoft SQL Server | No, [see details](1302-cursors-with-hold.md) |
| Oracle MySQL / MariadDB | No, [see details](1349-cursors-with-hold.md) |
| Oracle Database Server | No, [see details](1405-cursors-with-hold.md) |
| PostgreSQL | No, [see details](1456-cursors-with-hold.md) |
| SQLite | No, [see details](1499-cursors-with-hold.md) |

It is mandatory to review code using `WITH HOLD` cursors with a
`SELECT` statement having the `FOR UPDATE` clause.

The standard SQL solution is to declare a `FOR UPDATE` cursor (without
`WITH HOLD` clause) outside the transaction and open the cursor inside the
transaction:

```
DECLARE c1 CURSOR FOR SELECT ... FOR UPDATE
BEGIN WORK
  OPEN c1
  FETCH c1 INTO ...
  UPDATE ...
COMMIT WORK
```

If you need to process a large result set by doing updates of master and detail rows, first fetch
the primary keys of all master rows into a program array, declare a cursor with the `SELECT
FOR UPDATE`, for each row in the array, start a transaction and perform the `UPDATE
WHERE CURRENT OF` for the current master record and the `UPDATE` for detail
rows, then commit the transaction and continue with the next master
record:

```
DEFINE x, mkeys DYNAMIC ARRAY OF INTEGER
DECLARE c1 CURSOR FOR SELECT key FROM master ...
FOREACH c1 INTO x
    LET mkeys[mkeys.getLength()+1] = x
END FOREACH
DECLARE c2 CURSOR FOR SELECT * FROM master WHERE key=? FOR UPDATE
FOR x = 1 TO mkeys.getLength()
  BEGIN WORK
  OPEN c2 USING mkeys[x]
  FETCH c2 INTO mrec.*
  IF status==NOTFOUND THEN
     ROLLBACK WORK
     CONTINUE FOR
  END IF
  UPDATE master SET ... WHERE CURRENT OF c2
  UPDATE detail SET ... WHERE master_key=mkeys[x]
  COMMIT WORK
END FOR
```

When the underlying database table is defined with a primary key column, consider replacing the
`WHERE CURRENT OF` clause by `WHERE pkey-column =
?`, which is a more portable syntax. With some types of databases, the `WHERE CURRENT
OF` clause has to be emulated, or is simply unsupported. For more details, see [Positioned UPDATE/DELETE](1028-positioned-update-delete.md "Using positioned updates/deletes with named database cursors.").

## Related links

**Related concepts**  

[DECLARE (SELECT … FOR UPDATE)](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes")
