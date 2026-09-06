---
title: "Install Dameng and create a database - database configuration/design tasks"
source: "fgl-topics/t_fgl_odiagdmg_003.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Installation (Runtime Configuration) > Install Dameng® and create a database - database configuration/design tasks"
type: "task"
description: "If you are tasked with installing and configuring the database, here is a list of steps to be taken: Install the Dameng® software on your server. Create a Dameng database entity: dbname To create the ..."
---

# Install Dameng and create a database - database configuration/design tasks

If you are tasked with installing and configuring the database, here is a list of steps
to be taken:

1. Install the Dameng® software on your server.
2. Create a Dameng database entity: dbname

   To create the database entity in Dameng, use the graphical tool that can be started with the
   $DM\_HOME/tool/dbca.sh shell script. Make sure to create the Dameng database with
   UTF-8 locale.

   When creating a Dameng database/instance, define the dminit options to get case-insensitive
   character match, UTF-8 encoding (default is byte length semantics):

   - CASE\_SENSITIVE=N
   - CHARSET=1 (UTF-8)
3. Check that the database is created, by using the disql SQL command
   interpreted:

   ```
   $ disql sysdba/@[::]:5237
   password: 

   Server[LOCALHOST:5236]:mode is normal, state is open
   login used time : 3.012(ms)
   disql V8
   SQL>
   ```
4. If needed, you can change the password policy:

   The password policy is defined in a Dameng server by the `PWD_POLICY` system
   parameter.

   The following SQL query shows the current password
   policy:

   ```
   SQL> SELECT * FROM V$PARAMETER WHERE NAME= 'PWD_POLICY';
   ```

   See Dameng documentation for more details.
5. Define a database user dedicated to your application: the application administrator. This user
   will manage the database schema of the application.

   Use the disql tool to create the DB
   user:

   ```
   $ disql sysdba/@[::]:5237
   password:
   ...
   SQL> CREATE USER dmguser IDENTIFIED BY "fourjs"
          PASSWORD_POLICY 0 LIMIT PASSWORD_LIFE_TIME UNLIMITED;
   ...
   ```
6. Give all requested database administrator privileges to the application administrator.

   Grant the privileges to create tables to the new created user as
   follows:

   ```
   SQL> GRANT DBA TO dmguser;
   ...
   SQL> GRANT RESOURCE TO dmguser;
   ...
   ```
7. If you plan to use temporary table emulation, create the TEMPTABS table space:

   For more details, see [Temporary tables](1241-temporary-tables.md).
8. Connect with the disql tool as the application administrator:

   Open a new database connection:

   ```
   $ disql dmguser/fourjs@[::]:5237

   Server[:::5237]:mode is normal, state is open
   login used time : 2.225(ms)
   disql V8
   ```
9. Create the application tables with `CREATE TABLE` statements.

   Convert Informix® data types to Dameng data types. See issue [Data Type
   Conversion Table](1226-sql-types-mapping-dameng.md)  for more details.
