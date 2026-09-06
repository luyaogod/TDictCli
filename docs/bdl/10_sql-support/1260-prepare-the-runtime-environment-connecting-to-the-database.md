---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagmsv_004.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "Genero BDL provides several ODI database drivers, based on different ODBC clients: To each ODI driver corresponds an ODBC database client software that must be installed, and configured with a ..."
---

# Prepare the runtime environment - connecting to the database

1. Genero BDL provides several ODI database drivers, based on different ODBC clients:

   To each ODI driver corresponds an ODBC database client software that must be installed, and
   configured with a dedicated ODBC data source, using vendor-specific configuration options:

   1. For the `dbmsnc` ODI driver, see [Microsoft ODBC for SQL Server](1261-microsoft-odbc-for-sql-server.md)
   2. For the `dbmesm` ODI driver, see [Easysoft ODBC for SQL Server](1262-easysoft-odbc-for-sql-server.md)
   3. For the `dbmftm` ODI driver, see [FreeTDS ODBC for SQL Server](1263-freetds-odbc-for-sql-server.md)
2. Set up the FGLPROFILE entries for [database
   connections](1059-database-connections.md "Explains how to manage database connections in a program.").
   1. Define the SQL Server database driver correspondig to the ODBC database client used:

      For Microsoft ODBC for SQL
      Server:

      ```
      dbi.database.dbname.driver = "dbmsnc"
      ```

      For Easysoft ODBC for SQL
      Server:

      ```
      dbi.database.dbname.driver = "dbmesm"
      ```

      For FreeTDS
      ODBC:

      ```
      dbi.database.dbname.driver = "dbmftm"
      ```
   2. The "`source`" parameter defines the name of the ODBC data source.

      ```
      dbi.database.dbname.source = "test1"
      ```
   3. Only if needed, force the `widechar` FGLPROFILE parameter.

      ```
      dbi.database.dbname.snc.widechar = false
      ```

      In most cases, it is sufficient to let ODI driver select the defaut snc.widechar setting
      according to the application locale and database collation. For more details, read [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md).
   4. If needed, enable the `ifxemul.nationalchars` FGLPROFILE parameter to get
      automatic conversion of CHAR/VARCHAR SQL types to NCHAR/NVARCHAR types, in DDL statements executed
      by programs.

      ```
      dbi.database.dbname.ifxemul.nationalchar = true
      ```

      This
      option is required when creating SQL tables in programs with CREATE TABLE, to match existing
      database tables using NCHAR/NVARCHAR columns. Read [Automatic conversion of CHAR/VARCHAR type names](1273-char-and-varchar-data-types.md) for more details.
   5. If required, define the serial emulation method to "trigseq", when the INSERT statements use
      all columns of the table, including the serial column. For more details, see [SERIAL and BIGSERIAL data types](1277-serial-and-bigserial-data-types.md).

      ```
      dbi.database.dbname.ifxemul.datatype.serial.emulation = "trigseq"
      ```
   6. If needed, define the login timeout with the following FGLPROFILE entry:

      ```
      dbi.database.stores.driver-code.logintime = 5
      ```
   7. If needed, define the number of rows to be fetched at once on the application side, for each
      single `FETCH` instruction:

      ```
      dbi.database.stores.driver-code.prefetch.rows = 50
      ```

      The default is 10 rows. This is usually sufficient for regular interactive applications. Increase
      this parameter only in case of batch programs processing large result sets. The bigger this
      parameter is, the more memory is used by each program.
   8. If needed, add ODBC connection string parameters with the
      `datasource?options` notation, in the
      [`source` parameter](1072-database-source-specification-source.md) of the connection.
      You can for example define the [SQL client application
      identifier for SQL Server](1089-sql-connection-identifier.md "Database client programs can be identified by name with some database server types.").

      ```
      dbi.database.dbname.source = "test1?APP=myappid;"
      ```

      The `source` parameter can also be defined at runtime in the [database specification](1076-connection-parameters-in-database-specification.md "Connection parameters can be provided in the database specification string passed to the DATABASE and CONNECT TO instructions.") of `CONNECT TO`
      instruction.

## Child topics

- [Microsoft ODBC for SQL Server](1261-microsoft-odbc-for-sql-server.md)
- [Easysoft ODBC for SQL Server](1262-easysoft-odbc-for-sql-server.md)
- [FreeTDS ODBC for SQL Server](1263-freetds-odbc-for-sql-server.md)
