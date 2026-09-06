---
title: "Understanding database connections"
source: "fgl-topics/c_fgl_Connections_002.html"
breadcrumb: "SQL support > Database connections > Understanding database connections"
type: "concept"
---

# Understanding database connections

> This is an introduction to database connections.

A database connection is a session of work, opened by the program to communicate
with a specific database server, in order to execute SQL statements as a specific user.

Before working with database connections, make sure you have properly installed and configured
all software, using the correct database client software/environment, and BDL database driver.
It is very important to understand database client settings, regarding user authentication as
well as database client character set configuration.

Note that on some platforms like on mobile devices, Genero BDL includes the SQLite lightweight
database library, which is the default. Therefore, when executing programs on these platforms,
there is no need to install a database client software and configure the database driver for the
runtime system.

A database connection is initiated with the [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.") instruction, or with the [`CONNECT TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") instruction: The `CONNECT TO` instruction allows
to specify database user credentials with the `USER/USING` clauses.

Multiple database connections can be established in a Genero BDL program.

![Schema example of a program using three database connections](../_images/DBCFig01.jpg)

*Schema example of a program using three database connections*

The database user can be identified explicitly for each connection. Usually, the user is
identified by a login and a password, or by using the authentication mechanism of the operating
system (or even from a tier security system).

Database connection instructions can not be prepared and executed as dynamic SQL
statements.

There are two kind of connection modes: [unique-session](1095-unique-session-mode-connection-instructions.md "Opening and closing a database for a unique session.") and [multi-session](1098-multi-session-mode-connection-instructions.md "Opening and closing a database for a unique session.") mode. When using the `DATABASE` and `CLOSE
DATABASE` instructions, the program is in unique-session mode. When using the
`CONNECT TO`, `SET CONNECTION` and `DISCONNECT`
instructions, the program is in multi-session mode.

> **Important:**
>
> It is not possible to mix unique-session and multi-session modes.

Once connected to a database server, the program uses the current session to execute SQL
statements in that context.

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

[Database driver specification (driver)](1073-database-driver-specification-driver.md "Database driver specification (driver)")
