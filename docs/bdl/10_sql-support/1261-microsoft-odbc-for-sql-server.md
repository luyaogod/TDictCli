---
title: "Microsoft ODBC for SQL Server"
source: "fgl-topics/c_fgl_odiagmsv_snc_config.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database > Microsoft ODBC for SQL Server"
type: "concept"
description: "Configure the Microsoft ODBC for SQL Server data source that will be used in the ODI \" source \" connection parameter: Software requirements When using the SNC database driver ( dbmsnc ), the ..."
---

# Microsoft ODBC for SQL Server

Configure the Microsoft ODBC for SQL Server data source that will be used in the ODI "[`source`](1072-database-source-specification-source.md)" connection parameter:

## Software requirements

When using the SNC database driver (`dbmsnc`), the "Microsoft ODBC for SQL Server"
software installed must be installed on the computer running Genero applications, see [msdn.microsoft.com](http://msdn.microsoft.com):

- On Microsoft® Windows® platforms:

  - With the [Microsoft ODBC 18 driver for SQL Server](https://docs.microsoft.com/en-us/sql/connect/odbc/download-odbc-driver-for-sql-server?view=sql-server-ver15)
    (MSODBCSQL18.DLL), use the `dbmsnc_18` ODI driver.

    Minimum
    Microsoft ODBC for SQL Server version: 18.5.2.
- On Linux® platforms:

  - With [Microsoft ODBC 18 driver for SQL Server](https://docs.microsoft.com/en-us/sql/connect/odbc/download-odbc-driver-for-sql-server?view=sql-server-ver15), use the
    `dbmsnc_18` ODI driver.

    Minimum Microsoft ODBC for SQL Server version:
    18.5.2.

On Windows platforms, the SNC ODI driver is linked with
ODBC32.DLL; There is no need to set the PATH environment variable to a specific database client
library path: The ODBC32.DLL driver manager will find the MS ODBC driver from the data source
settings.

On Linux
platforms, the `dbmsnc_nn` drivers are directly linked to the
corresponding libmsodbcsql-nn.so ODBC driver library and to
the `libodbcinst.so.n` library. The unixODBC software must be
installed on the system. The SNC drivers will be able to connect to SQL Server, as long as the
dynamic linker can find the Microsoft ODBC driver library. The
libmsodbcsql-nn.so shared library is a symbolic link located
in /usr/lib64, which points to the real ODBC shared library.

## ODBC data source configuration

On Linux / UNIX® platforms, create the odbc.ini file to
define the ODBC data source parameters to connect to the database server. The ODBCINI environment
variable must point to the odbc.ini file. A second ODBC configuration file named
`odbcinst.ini` defines the name and path to the installed ODBC drivers. However, the
odbcinst.ini file is not required for Genero BDL ODI drivers, since the ODI
drivers are directly linked to the vendor's ODBC driver shared library.

On Microsoft Windows, use the ODBC Data Source Administrator tool to configure your data sources.

## Specific ODBC settings

By default, Microsoft ODBC 18 for SQL Server enables connection encryption, that can lead to the following
ODBC error when TLS/SSL certificates are not properly
configured:

```
[Microsoft][ODBC Driver 18 for SQL Server]SSL Provider: [error:1416F086:SSL routines:
tls_process_server_certificate:certificate verify failed:self signed certificate]
```

To
disable encryption on Linux platforms, set the following ODBC option in the data source definition
file:

```
Encrypt = No
```

On Microsoft Windows, set "Connection Encryption" to "Optional" in the 4th
panel of the ODBC data source configuration application.

## Client locale settings

On Windows, the MS ODBC database
client locale is defined by the Windows regional settings for non-unicode applications, and must match the BDL application
locale. The BDL application locale is usually also defined by the regional settings (the ACP), or
for UTF-8 it can be set with `set LANG=.utf8`.

On Windows, the ODBC data source
option "Perform translation for character data" option must be disabled, when using UTF-8
application locale, NLS\_LENGTH\_SEMANTICS=BYTE, and the database has \_UTF8 collation for CHAR/VARCHAR
columns.

On Linux, with the MS ODBC driver, no ODBC configuration is required: Character set conversions
and ODBC bindings are automatically deduced from to C application locale (LANG/LC\_ALL).

On both Linux and Windows, ODBC character string bindings is controlled by the
`widechar` option, which is automatically selected according to the application
locale and length semantics. For more details see [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).

## ODBC data source example

UNIX ODBCINI sample for MS ODBC driver for
SQL Server:

```
[snc_msvtest1_dirac_utf8]
Driver            = /usr/lib64/libmsodbcsql-18.so
Description       = SQL Server ODBC 18
#Server           = [protocol:]server[,port]
Server            = tcp:dirac,1433
Database          = msvtest1
#-- Always Encrypted (Column Encryption)
#  ColumnEncryption = Enabled
#-- Transport encryption with SSL/TLS
#  Encrypt = Yes/No
#  TrustServerCertificate = xxx
#  Trusted_Connection=yes
```

With the above ODBC data source definition, the ODI "`source`" parameter needs to
be defined as
follows:

```
dbi.database.dbname.source = "snc_msvtest1_dirac_utf8"
```
