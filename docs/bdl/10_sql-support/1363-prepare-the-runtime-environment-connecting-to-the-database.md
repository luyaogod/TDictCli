---
title: "Prepare the runtime environment - connecting to the database"
source: "fgl-topics/t_fgl_odiagora_004.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Installation (Runtime Configuration) > Prepare the runtime environment - connecting to the database"
type: "task"
description: "In order to connect to Oracle®, you must have a database driver \" dbmora \" in $FGLDIR/dbdrivers . If you want to connect to a remote Oracle server from an application server, you must install the ..."
---

# Prepare the runtime environment - connecting to the database

1. In order to connect to Oracle®, you
   must have a database driver "`dbmora`" in
   $FGLDIR/dbdrivers.
2. If you want to connect to a remote Oracle server from an application server, you must install the Oracle Client Software on your application server and configure
   this.
3. Make sure that the ORACLE client environment variables are properly set.

   Check Oracle client configuration envrionment variables such as TNS\_ADMIN. See the Oracle documentation for more details.
4. Verify the environment variable defining the search path for database client shared libraries
   (libclntsh.so on UNIX™,
   OCI.DLL on Windows®)

   | Oracle version | Shared library environment setting |
| --- | --- |
| Oracle 19c and higher | *UNIX*: Define LD\_LIBRARY\_PATH (or its equivalent):Database Server: Add `$ORACLE_HOME/lib`Instant Client: Add install-directory*Windows*: Define PATH:Database Server: Add `%ORACLE_HOME%\bin`Instant Client: Add install-directory |
5. Check the database client locale settings (NLS\_LANG, NLS\_DATE\_FORMAT, etc).

   The database client locale must match the locale used by the runtime system (LC\_ALL,
   LANG).
6. If you are using the TNS protocol, verify if the Oracle listener is started on the server.
7. To test the client environment settings, you can try to connect to the Oracle server with the SQL\*Plus tool:

   ```
   $ sqlplus username/password@service
   ```
8. Disable installation by Oracle client of SIGINT signal handler.

   By default, the OCI client library installs a signal handlers for the SIGINT signal. Normally
   this should not be done by a database client library: It's in the hands of the main code
   (fglrun / `DEFER INTERRUPT`) to install SIGINT signal handlers and
   call appropriate DB client APIs to cancel an SQL query (OCIBreak()).

   To workaround this problem, the ODI driver for Oracle resets the fglrun signal
   handler for SIGINT, after each connectiong to Oracle.

   Additionally, you can disable OCI client library SIGINT signal handler installation, by defining
   the `DISABLE_INTERRUPT` sqlnet.ora file:

   ```
   DISABLE_INTERRUPT=ON
   ```

   For more details, search the Oracle bug database for [Bug 12877221](https://support.oracle.com/epmos/faces/BugDisplay?id=12877221).
9. Disable installation by Oracle client of various signal handlers for the OCI diagnostic
   framework.

   By default, the OCI client library installs signal handlers for the OCI diagnostic framework.
   These signal handlers are registered when connecting to the database, and can conflict with the
   signal handlers installed by the fglrun runtime system.

   Unless you are explicitly asked to use OCI diagnostic framework, there is no need to let OCI use
   these signals.

   To disable OCI client library signal handlers installations, define the following parameters in
   the sqlnet.ora file:

   ```
   DIAG_ADR_ENABLED=FALSE
   DIAG_DDE_ENABLED=FALSE
   DIAG_SIGHANDLER_ENABLED=FALSE
   ```

   For more details, search Oracle OCI
   documentation about "Fault Diagnosability in OCI".
10. Set up the FGLPROFILE entries for [database connections](1059-database-connections.md "Explains how to manage database connections in a program.").
    1. Set up FGLPROFILE for the serial emulation method.

       The following entry defines the serial emulation method. You can use the sequence-based
       trigger or the serialreg-based trigger
       method:

       ```
       dbi.database.dbname.ifxemul.datatype.serial.emulation = "(native|regtable)"
       ```

       The
       value '`native`' selects the sequence-based method, and the value
       '`regtable`' selects the serialreg-based method. This entry has no effect if
       `dbi.database.dbname.ifxemul.datatype.serial` is set to
       'false'.

       The default is serial emulation enabled with native method (sequence-based). See
       issue [SERIAL and BIGSERIAL data types](1377-serial-and-bigserial-data-types.md) for more details.
    2. The "`source`" parameter defines the TNS name of the Oracle database.

       ```
       dbi.database.dbname.source = "stock"
       ```
    3. Define the database schema selection if needed.

       The following entry defines the database schema to be used by the application. The database
       interface automatically executes an "`ALTER SESSION SET CURRENT_SCHEMA
       owner`" instruction to switch to a specific
       schema:

       ```
       dbi.database.dbname.ora.schema = "name"
       ```

       Here dbname identifies the database name used in the BDL program
       (`DATABASE dbname`) and name is the schema name
       to be used in the `ALTER SESSION` instruction. If this entry is not defined, no
       `ALTER SESSION` instruction is executed and the current schema defaults to the user's
       name.
    4. Define pre-fetch parameters.

       Oracle offers high performance by
       pre-fetching rows in memory. The pre-fetching parameters can be tuned with the following
       entries:

       ```
       dbi.database.dbname.ora.prefetch.rows = integer
       dbi.database.dbname.ora.prefetch.memory = integer # in bytes
       ```

       These values will be applied to all application cursors.

       The interface pre-fetches rows up to the `prefetch.rows` limit unless the
       `prefetch.memory` limit is reached, in which case the interface returns as many rows
       as will fit in a buffer of size `prefetch.memory`. By default, pre-fetching is on and
       defaults to 10 rows; the memory parameter is set to zero, so the memory size is not included in
       computing the number of rows to prefetch.
    5. If needed, define a specific command to generate session identifiers with this FGLPROFILE
       setting:

       ```
       dbi.database.dbname.ora.sid.command = "SELECT ..."
       ```

       This unique session identifier will be used to create table names for temporary table
       emulation.

       By default, the database driver will use "`SELECT USERENV('SESSIONID') FROM
       DUAL`".
    6. If needed, define a specific command to generate session identifiers with this FGLPROFILE
       setting:

       ```
       dbi.database.dbname.ora.sid.command = "SELECT ..."
       ```

       This unique session identifier will be used to create table names for temporary table
       emulation.

       By default, the database driver will use "`SELECT USERENV('SESSIONID') FROM
       DUAL`".
    7. The default temporary table emulation uses regular permanent tables.

       If this does not fit your needs, you can use global temporary tables with this FGLPROFILE
       setting:

       ```
       dbi.database.dbname.ifxemul.temptables.emulation = "global"
       ```
    8. By default, the tablespace for the default temporary table emulation is
       `TEMPTABS`. For global temporary table emulation, there is no tablespace used by
       default.

       If required, define the tablespace to be used for temporary table emulations (note that this
       parameter applies to all temporary table emulation
       methods):

       ```
       dbi.database.dbname.ora.temptables.tablespace = "mytemptabs"
       ```
    9. By default, no schema is used for the default temporary table emulation. For global temporary
       table emulation, the schema `TEMPTABS` is used by default.

       If required, define the schema to be used for temporary table emulations (note that these
       parameters applies to all temporary table emulation methods). Define the "source" parameter to use
       the current user login, or a `SELECT` command to produce the schema
       name:

       ```
       # Get the schema from the current user name specified in connection command:
       dbi.database.dbname.ora.temptables.schema.source = "login"
       # or, get the schema from a SELECT statement:
       dbi.database.dbname.ora.temptables.schema.source = "command"
       dbi.database.dbname.ora.temptables.schema.command =
           "SELECT SYS_CONTEXT('USERENV','SESSION_USER') FROM DUAL"
       ```
    10. Starting with Oracle 23c and the `dbmora_23` ODI driver, you can used native
        Oracle SQL `BOOLEAN` types to store FGL `BOOLEAN` values.

        Prior to Oracle 23c, when using the `dbmora_18` ODI driver, FGL
        `BOOLEAN` types could only be stored in `CHAR(1)` columns. To keep
        using `CHAR(1)` SQL types for FGL `BOOLEAN` values, define the
        following FGLPROFILE
        entry:

        ```
        dbi.database.dbname.ora.boolean.aschar = true
        ```

        For more details, see [BOOLEAN data type](1372-boolean-data-type.md).
