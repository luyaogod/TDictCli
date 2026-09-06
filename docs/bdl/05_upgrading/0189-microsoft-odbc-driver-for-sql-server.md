---
title: "Microsoft ODBC Driver for SQL Server"
source: "fgl-topics/c_fgl_Migrate_to_310_sqlserver_2016.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Microsoft ODBC Driver for SQL Server"
type: "concept"
description: "Support for Microsoft ODBC Driver for SQL Server"
---

# Microsoft ODBC Driver for SQL Server

> Support for Microsoft® ODBC Driver for SQL Server

## Feature summary

Genero BDL 3.10 now supports Microsoft ODBC Driver v13 and v17 for SQL Server, respectively with the
`dbmsnc_13` and `dbmsnc_17` ODI drivers.

These ODI drivers are available on Windows® and Linux® platforms, to
connect to:

- Microsoft SQL Server 2016,
- Microsoft SQL Server 2017,
- Microsoft Azure SQL Database.

## Using Microsoft ODBC on Windows platforms

The new ODI drivers for Microsoft SQL Server can be used to connect from Windows platforms to
Microsoft SQL Server 2016, 2017 and Azure DB.

The Windows ODI drivers for Microsoft ODBC for SQL Server need the following DLLs:

- The `dbmsnc_13` ODI driver requires MSODBCSQL13.DLL.
- The `dbmsnc_17` ODI driver requires MSODBCSQL17.DLL.

> **Note:**
>
> Microsoft ODBC 11 for SQL Server (MSODBCSQL11.DLL) is not
> supported.

For backward compatibility, the `dbmsnc_11` driver using Microsoft SQL Native
Client 11 (SQLNCLI11.DLL) is still supported. However, it is recommended to
upgrade to the `dbmsnc_13` or `dbmsnc_17` drivers using Microsoft ODBC
for SQL Server.

> **Important:**
>
> With Genero BDL 3.10, the generic ODI driver named `dbmsnc` maps to the most
> recent SNC driver `dbmsnc_17`, requiring Microsoft ODBC 17 for SQL Server
> (MSODBCSQL17.DLL) in the ODBC data source definition. If you do not set up the
> ODBC data source configuration to use MSODBCSQL17.DLL, the SNC ODI driver will
> report an invalid ODBC driver error.
>
> With Genero 3.00, you still have to use `dbmsnc_11` to connect to SQL Server 2016.
> For a cross-version configuration of Genero, use explicit ODI driver names
> `dbmsnc_11`, `dbmsnc_13` or `dbmsnc_17`, instead of the
> generic driver name "`dbmsnc`". This is required because `dbmsnc` maps
> to `dbmsnc_17` in BDL 3.10 and `dbmsnc` maps to
> `dbmsnc_11` in BDL 3.00.

## Using Microsoft ODBC on Linux platforms

The new ODI drivers for Microsoft SQL Server can be used to connect from Linux platforms to
Microsoft SQL Server 2016, 2017 and Azure DB.

The Linux ODI drivers for Microsoft ODBC for SQL Server are linked to the following shared
libraries:

- The `dbmsnc_13` ODI driver requires libmsodbcsql-13.so.
- The `dbmsnc_17` ODI driver requires libmsodbcsql-17.so.

On Linux
platforms, the `dbmsnc_nn` drivers are directly linked to the
corresponding libmsodbcsql-nn.so ODBC driver library and to
the `libodbcinst.so.n` library. The unixODBC software must be
installed on the system. The SNC drivers will be able to connect to SQL Server, as long as the
dynamic linker can find the Microsoft ODBC driver library. The
libmsodbcsql-nn.so shared library is a symbolic link located
in /usr/lib64, which points to the real ODBC shared library.

## Azure SQL Database

The ODI drivers for Microsoft ODBC for SQL Server can be used to connect to Microsoft Azure SQL
Databases, from Windows or Linux platforms.

To establish a TCP connection to a server created through the Azure portal, you need to configure
the firewalls on the server side and on your computer and network.

For more details, check <http://azure.microsoft.com/en-us/>.

## Related links

**Related concepts**  

[Microsoft SQL Server](../10_sql-support/1068-microsoft-sql-server.md "Microsoft SQL Server")

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")
