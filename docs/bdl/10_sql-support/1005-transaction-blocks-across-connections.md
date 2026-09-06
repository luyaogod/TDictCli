---
title: "Transaction blocks across connections"
source: "fgl-topics/c_fgl_sql_programming_013.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Transaction blocks across connections"
type: "concept"
---

# Transaction blocks across connections

> Transaction blocks manage transactions when connected to several database servers.

In some cases, you need to copy data from a database to another. Database vendor export /
import tools exist for this task and their use is preferred when large data transfer is
needed. However, it is also possible to achieve this with a BDL program connected to
both databases, reading data from the source database and inserting rows into the target
database.

If the rows created in the target database need to be committed as a whole, you must open a
transaction with the `BEGIN WORK` instruction, use `SET
CONNECTION` to switch between the connections to read/write rows, and terminate
the transaction with a `COMMIT WORK`.

In order to keep a transaction open when switching to another database connection, the
connection must be initiated with the `WITH CONCURRENT TRANSACTION` clause.
If this option is not used, databases servers might raise an error when changing the
connection context. For example IBM®
Informix® will return the SQL error -1801:
Multiple-server transaction not supported.

The example opens two database connections, reads rows from a table of the first database,
and uses a transaction to insert rows in a table of the second database:

```
MAIN
   DEFINE rec RECORD
               pk INTEGER,
               name VARCHAR(50)
          END RECORD

   CONNECT TO "test1+driver='dbmifx'" AS "s1"
           USER "ifxuser" USING "fourjs"
           WITH CONCURRENT TRANSACTION
   CREATE TEMP TABLE tt1 ( pk INT, name VARCHAR(50) )
   INSERT INTO tt1 VALUES ( 1, "Item 1" )
   INSERT INTO tt1 VALUES ( 2, "Item 2" )

   CONNECT TO "test1+driver='dbmmys'" AS "s2"
           USER "mysuser" USING "fourjs"
           WITH CONCURRENT TRANSACTION
   CREATE TEMP TABLE tt2 ( pk INT, name VARCHAR(50) )

   SET CONNECTION "s1"
   DECLARE c1 CURSOR FOR SELECT * FROM tt1

   SET CONNECTION "s2"
   BEGIN WORK

   SET CONNECTION "s1"
   FOREACH c1 INTO rec.*
       SET CONNECTION "s2"
       INSERT INTO tt2 VALUES ( rec.* )
       SET CONNECTION "s1"
   END FOREACH

   SET CONNECTION "s2"
   COMMIT WORK

END MAIN
```

## Related links

**Related concepts**  

[Multi-session mode connection instructions](1098-multi-session-mode-connection-instructions.md "Opening and closing a database for a unique session.")

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
