---
title: "Install Oracle and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagora_003.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Installation (Runtime Configuration) > Install Oracle and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: Install the ORACLE database software on your computer. Create an ORACLE instance with multitenant ..."
---

# Install Oracle and create a database - database configuration/design tasks

If you are tasked with installing and configuring the database, here is a list of
steps to be taken:

1. Install the ORACLE database software on your computer.
2. Create an ORACLE instance with multitenant option, to get a Container DataBase (CDB) and create
   Pluggable DataBases (PDBs).

   Specify the database character set when creating the database instance. If you plan to use
   UTF-8, use also [character length
   semantics](1373-char-and-varchar-data-types.md) (See NLS\_LENGTH\_SEMANTICS session parameter).
3. Create a database PDB dedicated to your application.
   1. Connect as system user with:

      ```
      $ sqlplus / AS SYSDBA
      ```
   2. Create a pluggable database and its PDB administrator user. Below is a basic PDB creation
      example using Oracle Managed Files; Consider planing the PDB creation with your Oracle database
      administrator:

      ```
      CREATE PLUGGABLE DATABASE mypdb
        ADMIN USER pdbadmin IDENTIFIED BY password ROLES = (DBA)
        DEFAULT TABLESPACE mypdb_01
           DATAFILE 'path_01' SIZE 250M AUTOEXTEND ON ;
      ```
   3. For now the PDB is only mounted, it must be opened for regular usage:

      ```
      ALTER PLUGGABLE DATABASE mypdb OPEN;
      ```
   4. PDBs must be identified as separate database services (i.e. different from the CDB service). By
      default Oracle creates a database service with the same name as the PDB. To access the PDB through
      TNS, create the mypdb record in TNSNAMES.ORA file in
      addition to the default database service (`ORC*`):

      ```
      tnsname = 
       (DESCRIPTION = 
          (ADDRESS = (PROTOCOL = TCP)(HOST = localhost)(PORT = 1521))
          (CONNECT_DATA = 
             (SERVER = DEDICATED)
             (SERVICE_NAME = mypdb)
          )
       )
      ```
   5. By default when Oracle starts, the PDBs are mounted but are not open for regular usage. Create
      a database trigger to open all PDBs automatically:

      ```
      CREATE OR REPLACE TRIGGER open_pdbs
        AFTER STARTUP ON DATABASE
      BEGIN
          EXECUTE IMMEDIATE 'ALTER PLUGGABLE DATABASE ALL OPEN';
      END open_pdbs;
      /
      ```
   6. Or, save the state of a specific PDB when it is open:

      ```
      ALTER PLUGGABLE DATABASE test1 OPEN;
      ALTER PLUGGABLE DATABASE test1 SAVE STATE;
      ```
   7. Re-connect as PDB administrator and create a user dedicated to application tables
      administration, and grant permissions to that user:

      ```
      CONNECT pdbadmin/password@mypdb
      CREATE USER appadmin IDENTIFIED BY password;
      GRANT DBA, UNLIMITED TABLESPACE TO appadmin;
      ```
4. If programs create temporary tables, you must create a dedicated tablespace and schema
   depending on the type of temporary table emulation used.

   For more details about temporay table emulations, see [Temporary tables](1390-temporary-tables.md).
5. Create the application tables by connecting to the database context as the application
   administrator:

   ```
   $ sqlplus appadmin/password@tnsname
   ```

   Convert Informix® data types to Oracle data types.
   See issue [data type Conversion Tables](1371-sql-types-mapping-oracle-database.md)
   for more details.
6. If you plan to use SERIAL emulation, you must choose a serial emulation method.

   Select the best emulation technique that matches your needs. You need to prepare the database
   depending on the emulation type. For more details, see [SERIAL and BIGSERIAL data types](1377-serial-and-bigserial-data-types.md).
