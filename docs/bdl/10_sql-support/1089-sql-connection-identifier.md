---
title: "SQL connection identifier"
source: "fgl-topics/c_fgl_Connections_017.html"
breadcrumb: "SQL support > Database connections > SQL connection identifier"
type: "concept"
---

# SQL connection identifier

> Database client programs can be identified by name with some database server types.

## Purpose of SQL connection identifiers

When connecting to databases such as Informix, Oracle, PostgreSQL or SQL Server, it is possible
to define a name for the database client program.

This allows for the association of a name with a database client connection in the database
server, to be used in logging and trace utilities.

For example, if you don't define the application identifier for SQL Server client programs, all
connections will get the name "ODBC".

## SQL application identifier with Informix

IBM® Informix®
clients can define the `CLIENT_LABEL` environment variable with value that identifies
the client program.

```
export CLIENT_LABEL="myappid"
```

This identifier is then passed to the Informix server and can be found in the system tables
`sysmaster.sysenvses`:

```
SELECT envses_sid, envses_value
    FROM sysenvses
   WHERE envses_name = 'CLIENT_LABEL'
```

Read Informix documentation about the `CLIENT_LABEL` environment variable for more
details.

## SQL application identifier with SQL Server

The SQL Server client application identifier can be defined by using the
`APP=name` ODBC connection string parameter. The ODBC application
identifier will be available in the SQL Server session with `sp_who` and
`APP_NAME()` SQL functions. This name will also be visible in the
`"ApplicationName"` column of SQL Profiler trace logs.

To pass ODBC connection string parameters to SQL Server, use the
`?options` notation in the `"source"` connection
parameter.

With an FGLPROFILE entry:

```
dbi.database.mydb.source = "mydatasource?APP=myappid;"
```

At runtime with connection parameters in database specification:

```
CONNECT TO "mydb+source='mydatasource?APP=myappid;'" USER un USING up
```

ODBC connection string parameters must be terminated by a `;` semi-colon.

See also the SQL Server `sys.sp_set_session_context()` system stored procedure, to
set key/value pairs for the current SQL session.

## SQL application identifier with PostgreSQL

The PostgreSQL client application identifier can be defined with the
`application_name` PostgreSQL URL connection string parameter.

To pass URL connection string parameters to PostgreSQL, use the
`?options` notation in the `"source"` connection
parameter:

With an FGLPROFILE entry:

```
dbi.database.mydb.source = "mydatabase?application_name=myappid"
```

At runtime with connection parameters in database specification:

```
CONNECT TO "mydb+source='mydatabase?application_name=myappid'" USER un USING up
```

PostgreSQL URL connection string parameters must be separated by a `&`
ampersand.

## SQL application info with Oracle database

With Oracle database and the ODI driver for Oracle Instant Client, once connected to the server,
it is possible to use the `DBMS_SESSION` and `DBMS_APPLICATION_INFO`
PL/SQL packages, to specify application information with simple [stored procedure calls](1411-stored-procedure-calls.md).

You can for example use the following stored procedures:

- `DBMS_SESSION.SET_IDENTIFIER()` - to identify the application and/or the end
  user.
- `DBMS_APPLICATION_INFO.SET_MODULE()` - to identify the current module.
- `DBMS_APPLICATION_INFO.SET_ACTION()` - to identify the current action.
- `DBMS_APPLICATION_INFO.SET_CLIENT_INFO()` - to specify more details.

Read to the Oracle database documentation for more details.

## Related links

**Related concepts**  

[Database source specification (source)](1072-database-source-specification-source.md "Database source specification (source)")

[Connection parameters in database specification](1076-connection-parameters-in-database-specification.md "Connection parameters can be provided in the database specification string passed to the DATABASE and CONNECT TO instructions.")
