---
title: "Opening a database connection"
source: "fgl-topics/c_fgl_Connections_003.html"
breadcrumb: "SQL support > Database connections > Opening a database connection"
type: "concept"
---

# Opening a database connection

> A database connection identifies the SQL database server and the database entity the program connects to, in order to execute SQL statements.

To connect to a database server, the database driver needs to be loaded, and the SQL data
source must be provided. Additionally, user authentication with user name / password may also
be needed. All these parameters define connection information.

There are different ways to give connection information, and it is possible to mix the
different methods to specify connection parameters. However, if provided, the database user
name and password have to be specified together with the same method.

A database connection is performed in programs with the `DATABASE` or
`CONNECT TO` instruction:

```
CONNECT TO dbspec [USER username USING password]
```

or

```
DATABASE dbspec
```

Prefer the `CONNECT TO` instruction, as it allows to specify a user name and
password.

For portability reasons, it is not recommended that you use database vendor specific syntax
(such as '`dbname@dbserver`') in the `DATABASE` or `CONNECT
TO` instructions: Connections must be identified in programs by a single name, while
connection parameters are provided in external files.

Indirect database specification uses entries in the [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") configuration file: When a `DATABASE` or `CONNECT
TO` instruction is executed with the parameter *dbspec*, the runtime system first looks
into FGLPROFILE for entries starting with `dbi.database.dbspec`, and uses
these connection parameters if found. Otherwise, the runtime system will do direct database
specification, by using the dbspec string to connect to the server.
> **Important:**
>
> - When using FGLPROFILE entries for database connection parameters, keep in mind that entries must
>   be written in lowercase.
> - With the Informix driver, the `dbi.database.dbspec.username`
>   and `dbi.database.dbspec.password` FGLPROFILE entries are not
>   taken into account by the `DATABASE dbspec` instruction. The
>   `username` and `password` entries are used only with the CONNECT TO
>   "dbspec" instruction.

Use a string variable with the `DATABASE` or `CONNECT TO` statement,
in order to specify the database source at runtime. This solution gives you the best flexibility.

The string variable can be set from your own configuration file, from a program argument or from
an environment variable.

## Example

```
MAIN
  DEFINE db, us, pwd CHAR(50)
  LET db = fgl_getenv("MYDBSOURCE")
  LET us = arg_val(2)
  LET pwd = arg_val(3)
  CONNECT TO db USER us USING pwd
  ...
END MAIN
```

## Related links

**Related concepts**  

[Direct database specification method](1077-direct-database-specification-method.md "Genero BDL applies direct database source specification when no FGLPROFILE entry corresponds to the database name used in programs.")

[Indirect database specification method](1078-indirect-database-specification-method.md "Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.")

[Unique session mode connection instructions](1095-unique-session-mode-connection-instructions.md "Opening and closing a database for a unique session.")

[Multi-session mode connection instructions](1098-multi-session-mode-connection-instructions.md "Opening and closing a database for a unique session.")
