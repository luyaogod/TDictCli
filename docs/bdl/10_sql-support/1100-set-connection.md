---
title: "SET CONNECTION"
source: "fgl-topics/c_fgl_Connections_043.html"
breadcrumb: "SQL support > Database connections > Multi-session mode connection instructions > SET CONNECTION"
type: "concept"
---

# SET CONNECTION

> Selects the current session when in multi-session mode.

## Syntax

```
SET CONNECTION {
    { session | DEFAULT } [DORMANT]
    | CURRENT DORMANT }
```

1. session is a string expression identifying
   the name of the database session to be set as current.

## Usage

The `SET CONNECTION` instruction
makes a given connection current.

The session name is
case-sensitive.

When using the `DEFAULT` keyword, it identifies the default database server
connection established with a `CONNECT TO DEFAULT` or a `DATABASE`
instruction.

The `DEFAULT` clause of `SET CONNECTION` is specific to IBM® Informix®.

To make the current connection dormant, use `CURRENT DORMANT` keyword.

The `DORMANT` clause of `SET CONNECTION` is specific to IBM Informix.

A `SET CONNECTION` statement cannot be executed with dynamic SQL
(`PREPARE` + `EXECUTE`).

## Example

```
MAIN
  DEFINE c1, c2, c3 INT
  CONNECT TO "stores1"
  CONNECT TO "stores2" AS "SA"
  CONNECT TO "stores3" AS "SB"
  SET CONNECTION "stores1"    -- Select first session
  SELECT COUNT(*) INTO c1 FROM customers
  SET CONNECTION "SA"         -- Select second session
  SELECT COUNT(*) INTO c2 FROM customers
  SET CONNECTION "SB"         -- Select third session
  SELECT COUNT(*) INTO c3 FROM customers
  SET CONNECTION "stores1"    -- Select first session again
END MAIN
```

## Related links

**Related concepts**  

[CONNECT TO](1099-connect-to.md "Opens a new database session in multi-session mode.")

[Dynamic SQL management](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
