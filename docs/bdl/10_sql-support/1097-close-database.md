---
title: "CLOSE DATABASE"
source: "fgl-topics/c_fgl_Connections_041.html"
breadcrumb: "SQL support > Database connections > Unique session mode connection instructions > CLOSE DATABASE"
type: "concept"
---

# CLOSE DATABASE

> Closes the current database connection created by a DATABASE instruction.

## Syntax

```
CLOSE DATABASE
```

## Usage

The `CLOSE DATABASE` instruction closes the current database connection opened by
a [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") instruction.

The
current connection is automatically closed when the program ends.

## Example

```
MAIN
  DATABASE stores1
  CLOSE DATABASE
  DATABASE stores2
  CLOSE DATABASE
END MAINs
```
