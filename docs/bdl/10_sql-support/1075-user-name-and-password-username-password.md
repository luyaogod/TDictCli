---
title: "User name and password (username/password)"
source: "fgl-topics/c_fgl_Connections_010.html"
breadcrumb: "SQL support > Database connections > Connection parameters > User name and password (username/password)"
type: "concept"
description: "In database connection parameters, the username and password parameters define the default database user, when the program uses the DATABASE instruction or the CONNECT TO instruction without the ..."
---

# User name and password (username/password)

In database connection parameters, the `username` and `password`
parameters define the default database user, when the program uses the [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") instruction or the [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction without the
`USER/USING` clause.

The `username` and `password` [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") entries are not encrypted. These
parameters are provided to simplify migration and are not recommended in production. It is better to
use `CONNECT TO` with a `USER` / `USING` clause to
avoid any security hole, setup OS user authentication, or use the connection callback method.
Example of database servers supporting OS user authentication are: IBM® Informix®, Oracle®, and SQL Server.

> **Important:**
>
> **Do not write clear user passwords in your sources! It is recommended
> that `username` and `password` parameters are set from
> variables.**

For backward compatibility reasons, when using the IBM Informix driver,
the username / password specification is ignored by the `DATABASE` instruction,
only the `CONNECT TO` instruction takes external (or
callback) login parameters into account.

## Related links

**Related concepts**  

[Database user authentication](1090-database-user-authentication.md "Different database user authentication methods exist.")
