---
title: "Dameng specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_dmg.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > Dameng® specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .dmg.schema Name of the database schema to be selected after connection is established. dbi.database.stores.dmg.schema = \"store2\" Set this parameter to a specific schema in order ..."
---

# Dameng specific FGLPROFILE parameters

## `dbi.database.dsname.dmg.schema`

Name of the database schema to be selected after connection is established.

```
dbi.database.stores.dmg.schema = "store2"
```

Set this parameter to a specific schema in order to share the same table with all users.

## `dbi.database.dsname.dmg.temptables.tablespace`

Defines the tablespace to be used for temporary table emulation.

```
dbi.database.stores.dmg.temptables.tablespace = "mytemptabs"
```

By default, the driver uses the TEMPTABS tablespace. The tablespace specified in the
`dmg.temptables.tablespace` entry must be a permanent tablespace.

For more details, refer to [Temporary tables](1241-temporary-tables.md).

## `dbi.database.dsname.dmg.sid.command`

SQL command (SELECT) to generate a unique session id (used for temp table names).

```
dbi.database.stores.dmg.sid.command = "SELECT my_unique_id() FROM dual"
```

By default, the driver uses "`SELECT
SYS_CONTEXT('USERENV','SESSIONID')||'_'||TO_CHAR(SYSTIMESTAMP,'FF') FROM dual`".

This parameter gives you the freedom to provide your own way to generate a unique id.

The SELECT statement must return a single row with one single column.

For more details, refer to [Temporary tables](1241-temporary-tables.md).
