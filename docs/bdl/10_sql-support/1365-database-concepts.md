---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagora_010.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Database concepts > Database concepts"
type: "concept"
description: "Informix® servers can handle multiple database entities. ORACLE supports the multitenant option, to create several Pluggable DataBases (PDBs) in a Container DataBase (CDB). See ORACLE documentation ..."
---

# Database concepts

Informix® servers can handle multiple database
entities.

ORACLE supports the multitenant option, to create several Pluggable DataBases (PDBs) in a
Container DataBase (CDB). See ORACLE documentation for more details.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.

ORACLE can manage multiple schemas, but by default other users must give the owner name
as prefix to the table name:

```
SELECT * FROM stores.customer
```

## Solution 1: Using Pluggable DataBases

Oracle supports the multitenant database concept, where you can create several
pluggable databases (PDBs) in a root container (CDB). Consider using this feature, if you need to
create several copies of the same database entity, that can be accessed/seen as individual data
sources.

## Solution 2: Using a dedicated schema

In an Oracle database, each user can manage his own database schema. You can dedicate
a database user to administer each occurrence of the application database.

Any user can select the current database schema with the following SQL
command:

```
ALTER SESSION SET CURRENT_SCHEMA = "schema"
```

Using this instruction, any user can access the tables without giving the owner
prefix as long as the table owner has granted the privileges to access the
tables.

You can make the database interface select the current schema automatically with the
following fglprofile
entry:

```
dbi.database.dbname.schema = "schema"
```

When using multiple database schemas, it is recommended that you create them in
separated tablespaces to enable independent backups and keep logical sets of tables
together. The simplest way is to define a default tablespace when creating the
schema owner:

```
 CREATE USER user IDENTIFIED BY password
    DEFAULT TABLESPACE deftablespace
     TEMPORARY TABLESPACE tmptablespace
```
