---
title: "CONNECT TO"
source: "fgl-topics/c_fgl_Connections_042.html"
breadcrumb: "SQL support > Database connections > Multi-session mode connection instructions > CONNECT TO"
type: "concept"
---

# CONNECT TO

> Opens a new database session in multi-session mode.

## Syntax

```
CONNECT TO { dbname | DEFAULT } [ AS session ]
    [ USER login USING auth [ TRUSTED ] ]
    [ WITH CONCURRENT TRANSACTION ]
```

1. dbname is the database specification.
2. session identifies the database session. By default, it is
   dbname.
3. login is the name of the database user.
4. auth is the password to authenticate the database user, or an external
   authentication token.

## Usage

The `CONNECT TO` instruction opens a database connection. If the instruction
successfully connects to the database environment, the connection becomes the current database
session for the program.

The `DEFAULT` clause of `CONNECT TO` is specific to IBM® Informix®.

The session name is case-sensitive.

A program can connect to several database environments at the same time (using different database
drivers), and it can establish multiple connections to the same database environment, provided each
connection has a unique connection name.

The connection is closed with the [`DISCONNECT`](1102-disconnect.md "Terminates database sessions when in multi-session mode.") instruction, or when the program ends.

When the `USER login USING auth` clause is
specified, the database user is identified by login and auth,
ignoring all other user settings defined in FGLPROFILE or as connection string parameters.

The auth parameter can be a simple password for internal database users, but
for some types of database engines, it can be used to specify an external authentication token, such
as a distinguished name (DN). For more details, see the SQL adaptation guide for your database
type.

When the `CONNECT TO` instruction does not specify the `USER/USING`
clause, the database user is typically authenticated through the operating system user. If needed,
the database user name and password can be provided from a callback function, with the [`dbi.default.userauth.callback`](1093-user-authentication-callback-function.md) FGLPROFILE
entry.

The `TRUSTED` keyword can be used after the `USER/USING` clause, to
indicate that the specified user is trusted by an IBM Informix trusted context. When a trusted connection is
established, a program can use the [`SET SESSION
AUTHORIZATION`](1101-set-session-authorization.md "Select the user under which database operations are performed in the current connection.") instruction, to switch between database users associated to the
connection user, as defined by the related trusted context. See Informix database server
documentation for more details.

The `TRUSTED` clause of `CONNECT TO` is specific to IBM Informix.

The `CONNECT TO` instruction raises an [exception](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.") if the connection can not be established, for example, if you specify a database
that the runtime system cannot locate, or cannot open, or for which the user of your program does
not have access privileges.

The `WITH CONCURRENT TRANSACTION` clause allows a program to open several
transactions concurrently in different database sessions: The transaction can be started with the
`BEGIN WORK` statement in a given connection context, then the program can switch to
another connection with [`SET
CONNECTION`](1100-set-connection.md "Selects the current session when in multi-session mode."), and when done, switch back to the first connection to issue a
`COMMIT WORK` or `ROLLBACK WORK`. This is supported for IBM Informix database
servers.

The `WITH CONCURRENT TRANSACTION` option is ignored with non-Informix database
server types, but it can be used in the `CONNECT TO` statement, for consistency with
Informix.

A `CONNECT TO` statement cannot be executed with dynamic SQL
(`PREPARE` + `EXECUTE`).

With IBM Informix database servers, when using the `CONNECT TO DEFAULT`, you connect to
the default IBM Informix database server, identified by the INFORMIXSERVER environment variable, without any
database selection.

When using IBM Informix on UNIX™, the only restriction on establishing
multiple connections to the same database environment is that an program can establish only one
connection to each local server that uses the shared-memory connection mechanism. To find out
whether a local server uses the shared-memory connection mechanism or the local-loopback connection
mechanism, examine the $INFORMIXDIR/etc/sqlhosts file.

## Example

```
MAIN
  DEFINE uname, upswd VARCHAR(50) 
  CONNECT TO "stores1" -- Session name is "stores1"
  CONNECT TO "stores1" AS "SA" -- Session name is "SA"
  CALL login_dialog() RETURNING uname, upswd
  CONNECT TO "stores2" AS "SB" USER uname USING upswd
END MAIN
```

## Related links

**Related concepts**  

[Connection parameters in database specification](1076-connection-parameters-in-database-specification.md "Connection parameters can be provided in the database specification string passed to the DATABASE and CONNECT TO instructions.")

[Opening a database connection](1061-opening-a-database-connection.md "A database connection identifies the SQL database server and the database entity the program connects to, in order to execute SQL statements.")
