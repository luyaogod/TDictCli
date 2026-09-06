---
title: "Install Microsoft SQL Server and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagmsv_003.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Installation (Runtime Configuration) > Install Microsoft™ SQL Server and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: Install Microsoft™ SQL Server, or setup a Microsoft Azure SQL Database in the cloud. Important: The ..."
---

# Install Microsoft SQL Server and create a database - database configuration/design tasks

If you are tasked with installing and configuring the database, here is a list of steps to be
taken:

1. Install Microsoft™ SQL Server, or setup a Microsoft Azure SQL Database in the cloud.

   > **Important:**
   >
   > The collation of the SQL Server instance defines the collation for the `model`
   > system database, which is used to create the `tempdb` system database at SQL Server
   > startup. The collation of the `tempdb` system database matters, if your programs
   > create temporary tables with `CREATE TABLE #tabname`. For more details, see [Temporary tables](1291-temporary-tables.md).
2. Create an SQL Server database entity with the SQL Server Management Studio.

   In the database properties:

   1. Choose a "Database collation", to define the character set for CHAR/VARCHAR columns:

      - Since SQL Server 2019, the "\_UTF8" collation option specifies that character strings are UTF-8
        encoded in CHAR/VARCHAR columns. See [CHAR and VARCHAR data types](1273-char-and-varchar-data-types.md) for more details.
      - Choose the "CS" collation option to get a case-sensitive collation, or the "CI" option to get a
        case-insensitive collation.
      - The "SC" collation option defines if NCHAR/NVARCHAR columns store UTF-16 characters, or UCS-2
        when SC is not used.
   2. The "Compatibility level" must be set to an SQL Server version supported by the Genero
      BDL.
   3. The "ANSI NULL Default" option must be set to true, if you want to have the same default NULL
      constraint as in Informix® (a column created without a
      NULL constraint will allow null values, users must specify NOT NULL to deny nulls).
   4. The "Quoted Identifiers Enabled" option must be set to false, to use database object names
      without quotes.
3. Create and declare a database user dedicated to your application: the application
   administrator.
4. If you plan to use SERIAL emulation based on triggers using a registration table,
   create the SERIALREG table and create the serial triggers for all tables using a SERIAL.

   See [SERIAL and BIGSERIAL data types](1277-serial-and-bigserial-data-types.md).
5. Create the application tables.

   Convert Informix data types to SQL Server data types. See [SQL types mapping: SQL Server](1271-sql-types-mapping-sql-server.md). In order to make application tables
   visible to all users, make sure that the tables are created with the 'dbo' owner.
