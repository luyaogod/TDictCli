---
title: "DATABASE"
source: "fgl-topics/c_fgl_Connections_040.html"
breadcrumb: "SQL support > Database connections > Unique session mode connection instructions > DATABASE"
type: "concept"
---

# DATABASE

> Opens a new database connection in unique-session mode.

## Syntax

```
DATABASE { dbname[@dbserver] | variable | string } [EXCLUSIVE]
```

1. dbname identifies the database name.
2. dbserver identifies the IBM®
   Informix® database server (INFORMIXSERVER).
3. variable can be any character string defined variable containing
   the database specification.
4. string can be a string literal containing the database specification.

## Usage

The `DATABASE` instruction opens a connection to the database server,
like `CONNECT TO`, but without user and password specification.

```
MAIN
  DATABASE stores
  ...
END MAIN
```

It is possible to use a program variable containing the database specification.

```
MAIN
  DEFINE dbname VARCHAR(100)
  LET dbname = arg_val(1)
  DATABASE dbname
  ...
END MAIN
```

If a current connection exists, it is automatically closed before connecting to the new
database.

The connection is closed with the [`CLOSE
DATABASE`](1097-close-database.md "Closes the current database connection created by a DATABASE instruction.") instruction, or when the program ends.

When using the `DATABASE` instruction, the database user is typically
authenticated through the operating system user. If needed, the database user name and password can
be provided from a callback function, with the [`dbi.default.userauth.callback`](1093-user-authentication-callback-function.md) FGLPROFILE entry.

The `DATABASE` instruction raises an [exception](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.") if the connection can not be established,
for example, if you specify a database that the runtime system cannot locate, or cannot
open, or for which the user of your program does not have access privileges.

The `EXCLUSIVE` keyword can be used to open an IBM
Informix database in exclusive mode to
prevent access by anyone but the current user. This keyword is IBM
Informix specific and is not recommended
when writing a portable SQL application.

The [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instructions
allow better control over database connections; it is recommended that you use these instructions
instead of `DATABASE` and `CLOSE DATABASE`.

When used outside a program block, the `DATABASE` instruction defines the database
schema for compilation (for `DEFINE LIKE`). Furthermore, when
`DATABASE` is used before a `MAIN` block, it defines also an [implicit database connection](0994-implicit-database-connection.md "An implicit database connection is made with the DATABASE instruction used before MAIN; use SCHEMA to avoid the implicit connection."). To specify the database
schema for compilation, consider using [`SCHEMA`](../09_advanced-features/0793-schema.md "Defines the database schema files to be used for compilation.") instead.

## Related links

**Related concepts**  

[Opening a database connection](1061-opening-a-database-connection.md "A database connection identifies the SQL database server and the database entity the program connects to, in order to execute SQL statements.")
