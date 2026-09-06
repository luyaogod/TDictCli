---
title: "Implicit database connection"
source: "fgl-topics/c_fgl_sql_programming_004.html"
breadcrumb: "SQL support > SQL programming > SQL basics > Implicit database connection"
type: "concept"
---

# Implicit database connection

> An implicit database connection is made with the DATABASE instruction used before MAIN; use SCHEMA to avoid the implicit connection.

The `DATABASE` instruction can be used in two distinct ways, depending on the
context of the statement within its source module:

- To specify a default database.

  Typically used in a `GLOBALS` module, to define
  variables with the `DEFINE ... LIKE`, but it is also used for the
  `INITIALIZE` and `VALIDATE` statements. Using the
  `DATABASE` statement in this way results in that database being opened automatically
  at run time.
- To specify a current database.

  In `MAIN` or in a `FUNCTION`,
  used to connect to a database. A variable can be used in this context ( `DATABASE
  varname` ).

A default database is almost always used, because many programs
contain `DEFINE ... LIKE` statements. A
problem occurs when the production database name differs
from the development database name, because the default database specification
will result in an automatic connection (just after
`MAIN`):

```
DATABASE stock_dev  -- Default database, used at compile time
DEFINE
   p_cust RECORD LIKE customer.*
MAIN -- Connection to default database occurs at MAIN
   DEFINE dbname CHAR(30)
   LET dbname = "stock1"
   DATABASE dbname -- Real database used in production
   ...
END MAIN
```

In order to avoid the implicit connection, you can use the [`SCHEMA`](../09_advanced-features/0793-schema.md "Defines the database schema files to be used for compilation.") instruction instead of
`DATABASE`:

```
SCHEMA stock_dev -- Schema specification only
DEFINE
   p_cust RECORD LIKE customer.*
MAIN -- No default connection occurs...
   DEFINE dbname CHAR(30)
   LET dbname = "stock1"
   DATABASE dbname
END MAIN
```

This instruction will define the database schema for compilation
only, and will not make an implicit connection at runtime.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
