---
title: "Connection parameters in database specification"
source: "fgl-topics/c_fgl_Connections_011.html"
breadcrumb: "SQL support > Database connections > Connection parameters in database specification"
type: "concept"
---

# Connection parameters in database specification

> Connection parameters can be provided in the database specification string passed to the DATABASE and CONNECT TO instructions.

## Using connection parameters at runtime

In the database name specification of [`CONNECT
TO`](1099-connect-to.md "Opens a new database session in multi-session mode.") or [`DATABASE`](1096-database.md "Opens a new database connection in unique-session mode.")
instructions, a `+` plus sign starts the list of connection specification
parameters.

The connection specification parameters override the `dbi.database`
connection parameters defined in [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files").

In this example, `driver`, `source` and `resource`
parameters are specified in the database specification string of the `CONNECT TO`
instruction:

```
MAIN
  DEFINE db, un, up STRING
  LET db = "stores+driver='dbmora',source='orcl',resource='myconfig'"
  LET un = ...
  LET up = ...
  CONNECT TO db USER un USING up
  ...
END MAIN
```

> **Important:**
>
> - Do not hard code connection parameters in programs to be installed on a production site.
>   Instead, build the connection string at runtime, or consider using the [indirect database specification method](1078-indirect-database-specification-method.md "Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.").
> - Do not specify the `username` and `password` parameters in
>   connection specification parameters. Instead, provide the SQL user credentials with the
>   `USER/USING` clause of `CONNECT TO`.
> - Consider backslash interpretation in connection strings, as described below in [Connection parameter parsing](1076-connection-parameters-in-database-specification.md).

## Syntax for connection parameters

Each parameter is defined with a name followed by an equal sign and a value enclosed in single
quotes. Connection parameters must be separated by a comma:

```
dbname+parameter='value'[,...]
```

In this syntax, parameter can be one of the following:

| Parameter | Description |
| --- | --- |
| `resource` | Specifies which '`dbi.database`' entries have to be read from the FGLPROFILE configuration file.When this property is set, the database interface reads `dbi.database.name.*` entries, where *name* is the value specified for the `resource` parameter. |
| `driver` | Defines the database driver library to be loaded (filename without extension). |
| `source` | Specifies the data source of the database. |
| `username` | Defines the name of the database user.**Important:**Consider using `CONNECT TO` with `USER/USING` clause instead! |
| `password` | Defines the password of the database user.**Important:**Do not write clear user passwords in your sources! This parameter should be set from a variable value. |

## Passing a plus sign in a connection parameter

With some databases, the [`source`](1072-database-source-specification-source.md)
connection parameter can take different forms, that may contain a `+` sign. When
specified directly as default `source` in the database connection specification, this
`+` sign will be interpreted as the starting character for connection specification
parameters, and produce error [-6373](../15_library-reference/4483-genero-bdl-errors.md): Invalid database connection string.

To workaround the `+` sign interpretation, put the `source`
parameter explicitly in the connection specification.

For example, with SQL Server, you can specify ODBC connection string parameters with the
`?options`
syntax:

```
CONNECT TO "mydsn?APP=myappid;" USER un USING up
```

In the above example, the connection specification passed to the `CONNECT TO`
instruction defines implicitly the `"source"` connection parameter. This string
contains the ODBC data source name (`mydsn`), and, after the `?`
question mark, the ODBC connection string parameter `APP=myappid;` defining the [SQL client application identifier](1089-sql-connection-identifier.md "Database client programs can be identified by name with some database server types.").

When using a plus sign in the ODBC connection string parameters, it will produce the error -6373.
This can happend for example when using authentication credentials, as with the
`CRED` parameter in the next example:

```
CONNECT TO "mydsn?APP=myappid;CRED=ZXB+2A" USER un USING up
```

To solve this issue, use an explicit `"source"` parameter in the connection
specification of `CONNECT
TO`:

```
CONNECT TO "mydbc+source='mydsn?APP=myappid;CRED=ZXB+2A'" USER un USING up
```

The plus sign in the `"CRED"` ODBC parameter value will be considered as part of
the `"source"` parameter value.

## Connection parameter parsing

The parameters and values after the `+` sign are parsed, and backslashes are
interpreted as in a [string literal](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.").

For example, `\n` becomes a new line character in the final value used in the
database driver.

To pass a backslash to the driver (for example when specifying a Windows® path), it must be doubled in the
string value.

Furthermore, a backslash in a string literal of the source code must be quadrupled.

Consider implementing a utility function to double backslashes in a
string:

```
FUNCTION escape_backslashes(str)
    DEFINE str STRING
    DEFINE buf base.StringBuffer
    LET buf = base.StringBuffer.create()
    CALL buf.append(str)
    CALL buf.replace("\\","\\\\",0)
    RETURN buf.toString()
END FUNCTION
```

## Related links

**Related concepts**  

[Indirect database specification method](1078-indirect-database-specification-method.md "Genero BDL allows to define database connection parameters in FGLPROFILE, that can be referenced by a single identifier in programs.")

[Database driver specification (driver)](1073-database-driver-specification-driver.md "Database driver specification (driver)")

[Database source specification (source)](1072-database-source-specification-source.md "Database source specification (source)")

[User name and password (username/password)](1075-user-name-and-password-username-password.md "User name and password (username/password)")
