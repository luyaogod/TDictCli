---
title: "FreeTDS ODBC for SQL Server"
source: "fgl-topics/c_fgl_odiagmsv_ftm_config.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database > FreeTDS ODBC for SQL Server"
type: "concept"
description: "Configure the FreeTDS ODBC data source that will be used in the ODI \" source \" connection parameter: Software requirements When using the FTM database driver, the FreeTDS driver must be installed, see ..."
---

# FreeTDS ODBC for SQL Server

Configure the FreeTDS ODBC data source that will be used in the ODI "[`source`](1072-database-source-specification-source.md)" connection parameter:

## Software requirements

When using the FTM database driver, the FreeTDS driver must be installed, see [www.freetds.org](http://www.freetds.org).

Minimum required FreeTDS version is 1.5.168

With the FTM driver, there is no need to install a driver manager like unixODBC: The FTM database
driver is linked directly with the libtdsodbc.so shared library. Verify the
environment variable (LD\_LIBRARY\_PATH or equivalent) specifies the search path for that database
client shared library.

Make sure the FreeTDS environment variables are properly set. Check for example FREETDS (the path
to the configuration file).

## ODBC data source configuration

On Linux® / UNIX® platforms, create the odbc.ini file to
define the ODBC data source parameters to connect to the database server. The ODBCINI environment
variable must point to the odbc.ini file. A second ODBC configuration file named
`odbcinst.ini` defines the name and path to the installed ODBC drivers. However, the
odbcinst.ini file is not required for Genero BDL ODI drivers, since the ODI
drivers are directly linked to the vendor's ODBC driver shared library.

On Microsoft™ Windows®, use the ODBC Data Source Administrator tool to configure your data sources.

## Specific ODBC settings

Set the TDS protocol version depending on the SQL Server version, by setting the `tds
version` parameter in freetds.conf or `TDS_Version` in
odbc.ini. For example, for SQL Server version 2012 and 2014, use
`TDS_Version=7.3`. For more details, see the [FreeTDS documentation](http://www.freetds.org/userguide/choosingtdsprotocol.htm).

## Client locale settings

Define the client character set for FreeTDS (`client charset` parameter in
freetds.conf or `ClientCharset` parameter in
odbc.ini). You may need to link FreeTDS with the libiconv library to support
character set conversions.

On both Linux and Windows, ODBC character string bindings is controlled by the
`widechar` option, which is automatically selected according to the application
locale and length semantics. For more details see [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).

## FreeTDS ODBC data source example

UNIX ODBCINI sample for FreeTDS
driver:

```
[ftm_msvtest1_ida_utf8_2017]
Description     = SQL Server 2017
Server          = ida
Database        = msvtest1
Port            = 1433
TDS_Version     = 7.3
ClientCharset   = UTF-8
#dump_file = /tmp/freetds.log
#dump_file_append = yes
```

With the above ODBC data source definition, the ODI "`source`" parameter needs to
be defined as
follows:

```
dbi.database.dbname.source = "ftm_msvtest1_ida_utf8_2017"
```
