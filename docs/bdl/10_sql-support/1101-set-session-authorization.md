---
title: "SET SESSION AUTHORIZATION"
source: "fgl-topics/c_fgl_Connections_045.html"
breadcrumb: "SQL support > Database connections > Multi-session mode connection instructions > SET SESSION AUTHORIZATION"
type: "concept"
---

# SET SESSION AUTHORIZATION

> Select the user under which database operations are performed in the current connection.

## Syntax

```
SET SESSION AUTHORIZATION TO login [ USING auth ]
```

1. login is the name of the database user.
2. auth is the password to authenticate the database user.

## Usage

The `SET SESSION AUTHORIZATION` instruction switches to the specified database
user in the current database connection. It allows a user to assume the identity of another database
user.

The `SET SESSION AUTHORIZATION` instruction is specific to IBM® Informix®.

If the current database connection is trusted (it was established with a [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") using the
`TRUSTED` keyword), the `SET SESSION AUTHORIZATION` instruction can be
used to switch to a database user associated to the trusted context of the user specified in
`CONNECT TO`. The password is optional, if the trusted context defines the user
association is defined `WITHOUT AUTHENTICATION`.

If the current database connection is not established for a trusted context, the `SET
SESSION AUTHORIZATION` can be used to switch to another database user, with an optional
password.

In both cases, the user specified in `CONNECT TO` must have all required database
privileges to switch to the user specified by the `SET SESSION AUTHORIZATION`
statement.

> **Tip:**
>
> Informix database trusted contexts and trusted connections are typically used
> within a three-tier application model, to increase overall system performances by reusing the
> current database connection: After a trusted connection is established, the middle-tier application
> can switch to different database users without reconnecting. Any SQL operation is done on behalf of
> the current database user. The application can then benefit from database server features such as
> SQL auditing and user privileges, which are based on database users.

See Informix database documentation for more details about `SET SESSION
AUTHORIZATION` and trusted contexts (`CREATE TRUSTED CONTEXT`).

## Example

```
MAIN
  DEFINE db, un, pw STRING
  LET db = "stores"
  LET un = "mike"
  LET pw = "..."
  CONNECT TO db USER un USING pw TRUSTED
  LET un = "scott"
  SET SESSION AUTHORIZATION TO un
  LET un = "robert"
  SET SESSION AUTHORIZATION TO un
END MAIN
```

## Related links

**Related concepts**  

[CONNECT TO](1099-connect-to.md "Opens a new database session in multi-session mode.")
