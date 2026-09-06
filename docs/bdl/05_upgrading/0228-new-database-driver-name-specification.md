---
title: "New database driver name specification"
source: "fgl-topics/c_fgl_Migrate_to_251_odi_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > New database driver name specification"
type: "concept"
---

# New database driver name specification

> Allows database driver specification without target database version information.

Starting with version 2.51, the database drivers follow a new filename convention, which allows you
to specify a generic name based on the target database type without any database version information.

> **Important:**
>
> Most database driver names have changed. You need to re-configure the ["`driver`"](../10_sql-support/1073-database-driver-specification-driver.md) entry in your FGLPROFILE settings
> (or database connection string parameters), to match the new driver names. If you are using the default
> driver (`dbmdefault`), there is no configuration change needed. To simplify upgrading,
> the runtime system identifies old driver names and converts them to new names. However, it is recommended
> that you consider using the generic driver name corresponding to the type of database your applications
> connect to. The error [-6366](../15_library-reference/4483-genero-bdl-errors.md)
> occurs if the runtime system is not able to load the specified database driver, or cannot identify an
> old driver name.

Before version 2.51, it was required to specify the exact database type and version, to match both
the database client and the server version. For example, when using Oracle® 11.2 (server and client):

```
dbi.database.stores.driver = "dbmoraB2x"
```

Starting with 2.51, you can now, for example, specify a generic driver name without version, which
can connect to any database server version supported by the DB vendor client/server protocol. The
generic name defines a database driver for the latest database client version that is available on the
platform:

```
dbi.database.stores.driver = "dbmora"
```

Each generic database driver name has also a human-readable alias, such as "informix" or "oracle".

```
dbi.database.stores.driver = "oracle"
```

To simplify driver specification, install the latest database client software that corresponds to
the generic driver name, especially if it does not require a database server upgrade.

For some database client software, additional database drivers are still provided for older database
client versions (if available on the platform). In such cases, the driver filename gets a version identifier.

For example:

- `dbmora_11` (Oracle 11g client)
- `dbmmys_5_1` (Oracle MySQL client 5.1.x)
- `dbmsnc_10` (SQL Server Native Client 10 (SQLNCLI10.DLL))
- `dbmsnc_9` (SQL Server Native Client 9 (SQLNCLI.DLL))

Such database drivers with version info are provided to follow db client library dependency rules,
as defined by the database vendors. For example, on a Linux®
platform, Oracle MySQL version 5.1.x provides
the db client library named `libmysqlclient.so.16`. In this filename, "16" is the version
number that defines the shared library compatibility. The database driver that was compiled and linked
in a compatible db client environment is `dbmmys_5_1`. This database driver is linked to
`libmysqlclient.so.16`. Starting with Oracle MySQL version 5.5.x, the db client library version number has been
incremented to 18 (linked to `libmysqlclient.so.18`). The driver to be used with that
library version is `dbmmys_5_5`, which was compiled and linked with a 5.5.x
environment.

## Related links

**Related concepts**  

[Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md "Database driver specification (driver)")

[FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")
