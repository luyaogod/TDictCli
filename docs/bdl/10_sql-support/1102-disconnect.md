---
title: "DISCONNECT"
source: "fgl-topics/c_fgl_Connections_044.html"
breadcrumb: "SQL support > Database connections > Multi-session mode connection instructions > DISCONNECT"
type: "concept"
---

# DISCONNECT

> Terminates database sessions when in multi-session mode.

## Syntax

```
DISCONNECT { ALL | DEFAULT | CURRENT | session }
```

1. session is a string expression identifying the name of the database session
   to be terminated.

## Usage

The `DISCONNECT` instruction closes a given database connection opened with a
[`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction.

The session name is case-sensitive.

When using the `DEFAULT` keyword, it identifies the default database server
connection established with a `CONNECT TO DEFAULT` or a `DATABASE`
instruction.

The `DEFAULT` clause of `DISCONNECT` is specific to IBM® Informix®.

Use the `ALL` keyword to terminate all opened connections. From that point, you must
establish a new connection to execute SQL statements.

Use the `CURRENT` keyword
to terminate the current connection only. From that point, in order to execute SQL statements, you must
select another connection with `SET CONNECTION`, or establish a new connection with
`CONNECT TO`.

A `DISCONNECT` statement cannot be executed with dynamic SQL (i.e. `PREPARE`
+ `EXECUTE`).

If a `DISCONNECT` statement is used while a database transaction is active, the transaction
is automatically rolled back.

## Example

```
MAIN
  CONNECT TO  "stores1"  -- Will be identified by "stores1"
  CONNECT TO  "stores1" AS "SA"
  CONNECT TO  "stores2" AS "SB" USER "scott" USING "tiger"
  DISCONNECT "stores1"
  DISCONNECT "SB"
  SET CONNECTION "SA"
END MAIN
```

## Related links

**Related concepts**  

[Database transactions](1106-database-transactions.md "Database transaction concepts and handling.")
