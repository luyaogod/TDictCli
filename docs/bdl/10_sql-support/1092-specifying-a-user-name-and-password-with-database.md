---
title: "Specifying a user name and password with DATABASE"
source: "fgl-topics/c_fgl_Connections_020.html"
breadcrumb: "SQL support > Database connections > Database user authentication > Specifying a user name and password with DATABASE"
type: "concept"
description: "The DATABASE instruction does not support the USER / USING clause as CONNECT TO does. If you don't use an automatic user authentication method of the database server, you must provide a user name and ..."
---

# Specifying a user name and password with DATABASE

The [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") instruction does not
support the `USER`/`USING` clause as [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") does. If you don't use an
automatic user authentication method of the database server, you must provide a user name and
password in some way.

The best way to identify database users is to replace every `DATABASE`
instruction by a `CONNECT TO` with `USER/USING` clause.
However, it is also possible to provide the user name and password
with the user authentication callback function, by defining
a global FGLPROFILE entry.

In a development environment, a default login and password can
be specified with the `dbi.database.dbname.username` and
`dbi.database.dbname.password` FGLPROFILE
entries. This solution must not be used in a production environment
because the password is not encrypted. For backward compatibility
reasons, when using the IBM®
Informix® driver, these
FGLPROFILE entries are ignored by the `DATABASE` instruction,
only the `CONNECT TO` instruction takes external
(or callback) login parameters into account.

Login parameters can also be provided in the connection string
used in the database name specification in `DATABASE` instruction.

## Related links

**Related concepts**  

[Indirect database specification method](1078-indirect-database-specification-method.md "Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.")

[User authentication callback function](1093-user-authentication-callback-function.md "User authentication callback function")
