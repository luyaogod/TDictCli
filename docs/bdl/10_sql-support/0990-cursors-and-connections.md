---
title: "Cursors and connections"
source: "fgl-topics/c_fgl_sql_programming_043.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Cursors and connections"
type: "concept"
---

# Cursors and connections

> How to use database cursors across connections?

Several database connections can be opened simultaneously with the [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction. Once connected,
you can [`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") cursors or [`PREPARE`](1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") statements to be used in
parallel within different connection contexts. This section describes how to use SQL cursors and SQL
statements in a multiple-connection program.

When you `DECLARE` a cursor or when you `PREPARE` a
statement, you actually create an SQL statement handle; the runtime system
allocates resources for that statement handle before sending the SQL text to the database
server via the database driver.

The SQL statement handle is created in the context of the current connection, and must be
used in that context, until it is freed or recreated with another `DECLARE` or
`PREPARE` statement. Using an SQL statement handle in a different connection
context than the one for which it was created will produce a runtime error.

The [`SET CONNECTION`](1100-set-connection.md "Selects the current session when in multi-session mode.") instruction
changes the connection context. Connections are identified by a name. The `AS`
clause of the `CONNECT TO` instruction allows you to specify a connection name.
If the `AS` clause is omitted, the connection gets a default name based on the
data source name.

This small program example illustrates the use of two cursors with two different
connections:

```
MAIN
   CONNECT TO "dbsource1" AS "s1"
   CONNECT TO "dbsource2" AS "s2"
   SET CONNECTION "s1"
   DECLARE c1 CURSOR FOR SELECT tab1.* FROM tab1
   SET CONNECTION "s2"
   DECLARE c2 CURSOR FOR SELECT tab1.* FROM tab1
   SET CONNECTION "s1"
   OPEN c1
   SET CONNECTION "s2"
   OPEN c2
   ...
END MAIN
```

The `DECLARE` and `PREPARE` instructions are a type of creator
instruction; if an SQL statement handle is recreated in a connection other than the original
connection for which it was created, old resources are freed and new resources are allocated in
the current connection context.

This allows you to re-execute the same cursor code in different connection contexts, as in
this example:

```
MAIN
   CONNECT TO "dbsource1" AS "s1"
   CONNECT TO "dbsource2" AS "s2"
   SET CONNECTION "s1"
   IF checkForOrders() > 0 ...
   SET CONNECTION "s2"
   IF checkForOrders() > 0 ...
   ...
END MAIN

FUNCTION checkForOrders(d)
   DEFINE d DATE, i INTEGER
   DECLARE c1 CURSOR FOR SELECT COUNT(*) FROM orders WHERE ord_date = d
   OPEN c1
   FETCH c1 INTO i
   CLOSE c1
   FREE c1
   RETURN i
END FUNCTION
```

If the SQL statement handle was created in a different connection, the resources used in
the old connection context are freed automatically, and new statement handle resources are
allocated in the current connection context.

## Related links

**Related concepts**  

[Database connections](1059-database-connections.md "Explains how to manage database connections in a program.")
