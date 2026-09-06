---
title: "Using SQL Server 2008 date/time types"
source: "fgl-topics/c_fgl_Migrate_to_210_sql_server_datetime.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.10 upgrade guide > Using SQL Server 2008 date/time types"
type: "concept"
---

# Using SQL Server 2008 date/time types

> SQL Server 2008 introduces new SQL type to store date/time information.

Starting with Genero 2.10, SQL Server 2008 is now supported with the `dbmsncA0`
driver.

SQL Server 2008 introduces new data types to store date/time information, namely TIME(n), DATE,
DATETIME2(n) and DATETIMEOFFSET(n). These data types offer better precision for data storage.

When connected to SQL Server 2008 and higher, Genero database drivers for SQL Server use the
TIME(n), DATE and DATETIME2(n) data types, to store Genero BDL DATE and DATETIME information.

> **Note:**
>
> SQL Server `DATETIMEOFFSET(n)` data type is not supported by
> Genero: This type includes timezone information, that would be lost when fetching the value into a
> `DATETIME` variable. If timezone information needs to be stored, use a dedicated
> column for this data, or create a view on the base table, that splits the
> `DATETIMEOFFSET` column into a `DATETIME2`
> (`cast(colname as datetime2)`) and timezome data
> (`datepart(tz, colname)`).

Upgrading to SQL Server 2008 requires planning the SQL Server database migration, in order to use
the new date/time types provided in SQL Server 2008.

Up to Genero version 2.50, ODI drivers are built for a given database client and database server
version. For example, `dbmftm90` (using FreeTDS) is designed for SQL Server 2005, and
will use old date/time types. Pay attention to the fact that drivers designed for an old SQL Server
version (2005), can also work with a more recent SQL Server version (2008), but will act as if they
were connected to SQL Server 2005 (using old date/time types).

Starting with Genero 3.00 (2.51 for Genero Mobile), [ODI
drivers](../10_sql-support/1073-database-driver-specification-driver.md) can detect the database server version at runtime, and adapt the SQL conversions to
the target server: Assuming that the FreeTDS version supports the required TDS protocol version to
connect to the target server version, the `dbmftm_0` ODI driver can connect to SQL
Server 2005 or 2008, and adapt the date/time type usage to the targeted SQL Server version.

However, since Genero 3.10, SQL Server 2005 is no longer supported and only new SQL Server 2008
date/time types can be used.

For more details about BDL to SQL Server type mappings, see [DATE and DATETIME data types](../10_sql-support/1275-date-and-datetime-data-types.md) and [SQL types mapping: SQL Server](../10_sql-support/1271-sql-types-mapping-sql-server.md).
