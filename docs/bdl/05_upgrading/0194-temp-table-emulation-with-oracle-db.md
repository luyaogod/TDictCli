---
title: "Temp table emulation with Oracle DB"
source: "fgl-topics/c_fgl_Migrate_to_310_oracle_temptabs.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Temp table emulation with Oracle DB"
type: "concept"
description: "Controlling the Oracle schema and the tablespace for tables created by Informix temporary table emulation."
---

# Temp table emulation with Oracle DB

> Controlling the Oracle® schema and the tablespace for tables created by Informix® temporary table emulation.

In order to support Informix CREATE TEMP
TABLE and SELECT ... INTO TEMP statements, the Oracle ODI driver converts the SQL text to Oracle syntax. Two emulation methods are supported by Genero: "default" and
"global" emulation. First method creates regular tables, while second method uses Oracle global temporary tables. The emulation
method used is defined by the following FGLPROFILE entry:

```
dbi.database.mydb.ifxemul.temptables.emulation = { "default" | "global" }
```

If your Genero application requires real Oracle database users for each end user, you need to specify the common schema where
application tables are defined, to access them without a schema prefix in SQL statements. You
typically do this with the following FGLPROFILE entry:

```
dbi.database.mydb.ora.schema = "app_owner"
```

However, if your application programs create Informix-style temporary tables, each Oracle DB user needs CREATE ANY TABLE and DROP
ANY TABLE privileges to create the tables in the common schema. This is not suitable in an
organization using a strong security policy.

The Oracle tables created by the
"default" emulation are by default created in the TEMPTABS tablespace, while global temporary tables
are created in the TEMPTABS schema. Before Genero 3.10, it was not possible to control this.

Starting with Genero 3.10, it is possible to specify the schema and the tablespace where Oracle tables are created for temporary table
emulation:

```
dbi.database.dbname.ora.temptables.schema.source = { "login" | "command" }
dbi.database.dbname.ora.temptables.schema.command = "select-statement"
dbi.database.dbname.ora.temptables.tablespace = "tablespace-name"
```

These FGLPROFILE parameters apply to both default and global temporary table emulation methods.
When using "default" emulation, you typically configure these entries to add the current user name
as schema for the CREATE TABLE Oracle
statements and avoid granting CREATE/DROP ANY TABLE privileges to all users. With the "global"
emulation, you can specify a common schema (different from the default TEMPTABS schema), to share
the global temporary table among several database users.

## Related links

**Related concepts**  

[Temporary tables](../10_sql-support/1390-temporary-tables.md "Temporary tables")

[Oracle DB specific FGLPROFILE parameters](../10_sql-support/1081-oracle-db-specific-fglprofile-parameters.md "Oracle DB specific FGLPROFILE parameters")
