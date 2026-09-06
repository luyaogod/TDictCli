---
title: "Easysoft ODBC for SQL Server"
source: "fgl-topics/c_fgl_odiagmsv_esm_config.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database > Easysoft ODBC for SQL Server"
type: "concept"
description: "Configure the Easysoft ODBC for SQL Server data source that will be used in the ODI \" source \" connection parameter: Software requirements When using the ESM database driver ( dbmesm ), the EasySoft ..."
---

# Easysoft ODBC for SQL Server

Configure the Easysoft ODBC for SQL Server data source that will be used in the ODI "[`source`](1072-database-source-specification-source.md)" connection parameter:

## Software requirements

When using the ESM database driver (`dbmesm`), the EasySoft ODBC driver for SQL
Server ([www.easysoft.com](http://www.easysoft.com))
and the unixODBC software must be installed on the system.

Minimum required Easysoft ODBC for SQL Server version: 2.3.0.

The ESM database driver is linked directly to the libessqlsrv.so and to the
`libodbcinst.so.n` shared library. Verify the environment variable
(LD\_LIBRARY\_PATH or equivalent) specifies the search path for that database client shared
library.

Make sure the EasySoft environment variables are properly set. Check for example EASYSOFT\_ROOT
(the path to the installation directory).

## ODBC data source configuration

On Linux® / UNIX® platforms, create the odbc.ini file to
define the ODBC data source parameters to connect to the database server. The ODBCINI environment
variable must point to the odbc.ini file. A second ODBC configuration file named
`odbcinst.ini` defines the name and path to the installed ODBC drivers. However, the
odbcinst.ini file is not required for Genero BDL ODI drivers, since the ODI
drivers are directly linked to the vendor's ODBC driver shared library.

On Microsoft™ Windows®, use the ODBC Data Source Administrator tool to configure your data sources.

## Specific ODBC settings

Set the following data source parameters when using the EasySoft database
client:

```
AnsiNPW=Yes
Mars_Connection=No
QuotedId=No
```

## Client locale settings

Configure the client character set for EasySoft with the `Client_CSet` parameter
in odbc.ini. The client character set is an iconv name and must match the [locale](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.") of your Genero
application.

To support all UNICODE characters when using UTF-8 with
`NCHAR/NVARCHAR` columns, you need to define `Client_CSet=UTF-8` and
`Server_UCSet=UTF-16LE`.

When using CHAR/VARCHAR types in the database and when the database
collation is different from the client locale, you must also set the `Server_CSet`
parameter to an iconv name corresponding to the database collation. Some examples:

- If `Client_CSet=ISO-8859-15` and the db collation is `Latin1_*`
  (=CP1252), you must set `Server_CSet=WINDOWS-1252` (otherwise, the characters €, Š,
  š, Ž, ž, Œ, œ, Ÿ which are encoded differently
- If `Client_CSet=BIG5` and the db collation is
  `Chinese_Taiwan_Stroke_BIN`, you must set
  `Server_CSet=BIG5HKSCS`.

On both Linux and Windows, ODBC character string bindings is controlled by the
`widechar` option, which is automatically selected according to the application
locale and length semantics. For more details see [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).

## ODBC data source example

UNIX ODBCINI sample for EasySoft ODBC for
SQL Server driver:

```
[esm_msvtest1_ida_utf8_2017]
Driver=Easysoft ODBC-SQL Server
Description=Easysoft SQL Server ODBC driver
Server=ida
Port=1683
Database=msvtest1
Mars_Connection=No
Logging=No
LogFile=/tmp/odbc.log
#QuotedId=No
AnsiNPW=Yes
Language=
Version7=No
ClientLB=No
Failover_Partner=
VarMaxAsLong=No
DisguiseWide=No
DisguiseLong=No
Trusted_Connection=No
Trusted_Domain=
IPv6=No
Client_CSet=UTF-8
Server_UCSet=UTF-16LE
```

With the above ODBC data source definition, the ODI "`source`" parameter needs to
be defined as
follows:

```
dbi.database.dbname.source = "esm_msvtest1_ida_utf8_2017"
```
