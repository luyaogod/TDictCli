---
title: "Database connections"
source: "fgl-topics/c_fgl_Connections_001.html"
breadcrumb: "SQL support > Database connections"
type: "concept"
---

# Database connections

> Explains how to manage database connections in a program.


## Child topics

- [Understanding database connections](1060-understanding-database-connections.md): This is an introduction to database connections.
- [Opening a database connection](1061-opening-a-database-connection.md): A database connection identifies the SQL database server and the database entity the program connects to, in order to execute SQL statements.
- [Database client environment](1062-database-client-environment.md): To connect to a database server, Genero BDL programs use vendor's database client software.
- [Connection parameters](1071-connection-parameters.md): This section describes the different parameters which need to be specified in order to connect to a database.
- [Connection parameters in database specification](1076-connection-parameters-in-database-specification.md): Connection parameters can be provided in the database specification string passed to the DATABASE and CONNECT TO instructions.
- [Direct database specification method](1077-direct-database-specification-method.md): Genero BDL applies direct database source specification when no FGLPROFILE entry corresponds to the database name used in programs.
- [Indirect database specification method](1078-indirect-database-specification-method.md): Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.
- [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md): Emulation of Informix® specific SQL features can be controlled with FGLPROFILE entries.
- [Database type specific parameters in FGLPROFILE](1080-database-type-specific-parameters-in-fglprofile.md): Specific connection parameters can be configured with FGLPROFILE entries.
- [SQL connection identifier](1089-sql-connection-identifier.md): Database client programs can be identified by name with some database server types.
- [Database user authentication](1090-database-user-authentication.md): Different database user authentication methods exist.
- [Unique session mode connection instructions](1095-unique-session-mode-connection-instructions.md): Opening and closing a database for a unique session.
- [Multi-session mode connection instructions](1098-multi-session-mode-connection-instructions.md): Opening and closing a database for a unique session.
- [Miscellaneous SQL statements](1103-miscellaneous-sql-statements.md): These are particular SQL statements supported in the static SQL syntax.
