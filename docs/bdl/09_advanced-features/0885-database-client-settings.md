---
title: "Database client settings"
source: "fgl-topics/c_fgl_localization_024.html"
breadcrumb: "Advanced features > Localization > Application locale > Database client settings"
type: "concept"
---

# Database client settings

> This section describes the settings defining the locale for the database client.

Each database software has its own client character set configuration.

> **Important:**
>
> In order to properly send/receive character string data to/from
> the database server, the database client locale and application locale settings must
> match.

| Database Client | Locale settings | Notes |
| --- | --- | --- |
| Oracle® database server | The database client locale must be set with NLS\_LANG environment variable. | By default, the client locale is set from the database server locale.The locale can also be defined after connection, with the ALTER SESSION instruction. However, this is not supported by Genero BDL. |
| IBM® Informix® | The database client locale is defined by the CLIENT\_LOCALE environment variable, and DB\_LOCALE must be set to the database locale. | If CLIENT\_LOCALE is not defined, other settings are used when defined (DBDATE / DBTIME / GL\_DATE / GL\_DATETIME, as well as standard LC\_\* variables). |
| Microsoft® SQL Server with SNC driver (Microsoft ODBC). | On Windows® platforms, the database client locale is defined by the language settings for non-Unicode applications. The current ANSI code page (ACP) is used by the SQL Server client and the Genero runtime system.On Linux® platforms, the database client locale is always UTF-16. The ODI driver uses the MS ODBC Wide Char API, and makes the required character set conversions between the application locale and UTF-16. | See Microsoft ODBC documentation for more details regarding client character set configuration. |
| Microsoft SQL Server with FTM driver (FreeTDS). | The database client character set is defined by the `client charset` parameter in freetds.conf, or with the `ClientCharset` parameter in the DSN of the odbc.ini file. | See FreeTDS documentation for more details regarding client character set configuration. |
| Microsoft SQL Server with ESM driver (Easysoft). | The database client character set is defined by the `Client_CSet` parameter in the DSN of the odbc.ini file.Depending on the application locale and SQL Server CHAR/VARCHAR or NCHAR/NVARCHAR usage, you might also need to define the `Server_CSet` and/or the `Server_UCSet` parameters. | To support all UNICODE characters when using UTF-8 with `NCHAR/NVARCHAR` columns, you need to define `Client_CSet=UTF-8` and `Server_UCSet=UTF-16LE`.When using CHAR/VARCHAR types in the database and when the database collation is different from the client locale, you must also set the `Server_CSet` parameter to an iconv name corresponding to the database collation. Some examples:If `Client_CSet=ISO-8859-15` and the db collation is `Latin1_*` (=CP1252), you must set `Server_CSet=WINDOWS-1252` (otherwise, the characters €, Š, š, Ž, ž, Œ, œ, Ÿ which are encoded differentlyIf `Client_CSet=BIG5` and the db collation is `Chinese_Taiwan_Stroke_BIN`, you must set `Server_CSet=BIG5HKSCS`. |
| PostgreSQL | The database client locale must be set with the PGCLIENTENCODING environment variable, or with the `client_encoding` configuration parameter in postgresql.conf.The default is to use the database encoding. | After the database connection, the locale can be set with the SET CLIENT\_ENCODING instruction. This is not recommended with Genero BDL.Check the pg\_conversion system table for available character set conversions. |
| Oracle® MySQL and MariaDB | The database client locale is defined by the `default-character-set` option in the MySQL configuration file. | MySQL/MariaDB support the `SET NAMES` SQL instruction, to change the character set after connecting to the database. However, this is not supported with Genero: The driver needs to know the character set at connection initialization. Use the `default-character-set` configuration option. |
| SQLite | No database client locale configuration is required with SQLite: The database driver makes the appropriate charset conversions when needed. | SQLite databases use UTF-8 encoding. If the locale used by the runtime system (LANG/LC\_ALL) is not UTF-8, Genero will do the appropriate character set conversions. |
| Dameng® | The database client character set is defined by the `PAGE_MODE` parameter in the dm\_svc.conf configuration file. | See Dameng documentation for more details regarding client character set configuration. |

## Related links

**Related concepts**  

[Database client environment](../10_sql-support/1062-database-client-environment.md "To connect to a database server, Genero BDL programs use vendor's database client software.")

[SQL character type for Unicode/UTF-8](../10_sql-support/1014-sql-character-type-for-unicode-utf-8.md "This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.")
