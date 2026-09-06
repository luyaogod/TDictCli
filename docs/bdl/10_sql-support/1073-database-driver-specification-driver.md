---
title: "Database driver specification (driver)"
source: "fgl-topics/c_fgl_Connections_009.html"
breadcrumb: "SQL support > Database connections > Connection parameters > Database driver specification (driver)"
type: "concept"
description: "In database connection parameters, the driver parameter identifies the type of database driver to be used. The driver must correspond to the database client software. Important: Pay attention to the ..."
---

# Database driver specification (driver)

In database connection parameters, the `driver` parameter identifies the type of
database driver to be used.

The driver must correspond to the database client software.
> **Important:**
>
> Pay attention
> to the binary architecture of the database client software: Genero runtime system and database
> client binaries must match. For example, a 32 bit Oracle® client can not be used with a Genero 64 bit runtime system.

We distinguish two types of database driver names:

- Generic driver names ("`dbmora`", "`dbmsnc`"), and aliases
  ("`oracle`", "`sqlserver`")
- Version-stamped driver names ("`dbmora_23`", "`dbmsnc_18`",
  "`dbmmys_8_4`")

A driver name "`dbmxxx`" identifies a generic driver name for
the database server identified by the code `xxx`.

For example, in FGLPROFILE, to define the database driver for the Oracle OCI client (code "`ora`"), use the name
"`dbmora`":

```
dbi.database.stores.driver = "dbmora"
```

For convenience, it is also possible to specify a long name (alias) such as
"`oracle`" or "`sqlserver`", as defined in the database driver table
below.

The generic driver names (like `dbmora`) require the latest database client
software available on the platform. This can change from one Genero BDL release to another, when
supporting a new database client version. Use the version-stamped driver name (like
`dbmora_18`), to point to a specific database client version, for example when the
most recent database client software is not available on the platform.

Check for library dependency on your system, to identify the database client library required by
the driver with the generic name. The driver definition table below lists the driver names for each
supported database client types and versions. For example, on Linux® platform, use the ldd
command:

```
$ ldd $FGLDIR/dbdrivers/dbmmys.so
      ...
      libmysqlclient.so.24 => ...
      ...
```

Drivers with generic names are compatible with the latest database client version available on
the platform. Depending on the platform, the same generic driver name can refer to different
database client software versions.

To limit the number of drivers, if the database client software is ABI compatible across several
versions, the ODI drivers are build with the oldest database client version that is compatible with
the latest available database client versions. For example, the `dbmpgs_9` driver is
linked to libpq.so.5, supported by various PostgreSQL client versions.

> **Note:**
>
> A given driver (combined with the corresponding database client software library) can connect
> to a database server of an older version, if the database vendor client/server protocol supports the
> combination. For example, you can use a PostgreSQL client version 12 to connect to a PostgreSQL 11
> server. The ODI driver will then adapt SQL translations and emulations to the target database server
> version.

A default driver can be specified with the `dbi.default.driver` FGLPROFILE entry.
This driver will be used for all database connections that do not specify the driver explicitly:

```
dbi.default.driver = "dbmora"
```

If this entry is not defined, and if no driver parameter is specified for the data source, the
driver name defaults to `dbmdefault`. This default driver is a copy of the database
driver that was chosen during installation.

| Versionned name / Generic name / Alias | Code | Database client software version | UNIX™ shared objects | Microsoft™ Windows® DLLs | macOS® dynamic libraries |
| --- | --- | --- | --- | --- | --- |
| `dbmdmg_8` | `dmg` | Dameng 8 DPI client | libdmdpi.so | N/A | N/A |
| `dbmesm_1` | `esm` | Easysoft ODBC 2.3 + for SQL Server | libessqlsrv.so / libodbcinst.so.n | N/A | N/A |
| `dbmifx_9` | `ifx` | IBM® Informix® CSDK 4.50 | libifsql.so, libifasf.so, libifgen.so, libifos.so, libifgls.so, libifglx.so | isqlt09a.dll, igl4n304.dll, iglxn304.dll, igo4n304.dll | libifsql.dylib, libifasf.dylib, libifgen.dylib, libifos.dylib, libifgls.dylib, libifglx.dylib |
| `dbmifx_15` | `ifx` | IBM Informix CSDK 15 and higher | libifsql15a.so, libifasf15a.so, libifgen15a.so, libifos15a.so, libifgls.so, libifglx.so | isqlt15a.dll, igl4n304.dll, igo4n304.dll | libifsql15a.dylib, libifasf15a.dylib, libifgen15a.dylib, libifos15a.dylib, libifgls.dylib, libifglx.dylib |
| `dbmmdb_10_2` | `mdb` | MariaDB Client 10.2.x, 11.x and 12.x | libmariadb.so.3 | libmariadb.dll | libmariadb.3.dylib |
| `dbmmys_8_4` | `mys` | Oracle MySQL Client 8.4.x LTS (see note) | libmysqlclient.so.24 | libmysql.dll | libmysqlclient.24.dylib |
| `dbmmys_9_7` | `mys` | Oracle MySQL Client 9.7.x LTS | libmysqlclient.so.24 | libmysql.dll | libmysqlclient.24.dylib |
| `dbmora_18` | `ora` | OCI Client 18, 19, 21 | libclntsh.so.18.1 | oci.dll | libclntsh.dylib.18.1 |
| `dbmora_23` | `ora` | OCI Client 23 | libclntsh.so.23.1 | oci.dll | libclntsh.dylib.23.1 |
| `dbmpgs_9` | `pgs` | PostgreSQL Client 14 to 18 | libpq.so.5 | libpq.dll | libpq.5.dylib |
| `dbmsnc_18` | `snc` | Microsoft ODBC 18 for SQL Server | libmsodbcsql-18.so / libodbcinst.so.n | odbc32.dll / MSODBCSQL18.DLL | libmsodbcsql.18.dylib |
| `dbmftm_0` | `ftm` | FreeTDS ODBC version 1.5+ | libtdsodbc.so.0 | N/A | N/A |
| `dbmsqt_3` | `sqt` | SQLite 3.x | libsqlite3.so.0 | N/A (statically linked) | libsqlite3.dylib |

> **Note:**
>
> The MySQL 8.2 and 8.3 versions are Innovation Releases (IR), while MySQL versions 8.0 and 8.4 are
> Long-Term Support releases (LTS). The `dbmmys_8_2` driver provided in previous Genero
> versions has been desupported in favor to MySQL 8.4 with the new `dbmmys_8_4` ODI
> driver, requiring libmysqlclient.so.24. If you need to access a MySQL 8.2 or
> 8.3 server, install MySQL Connector client with libmysqlclient.so.24 and use
> the `dbmmys_8_4` ODI driver.

## Related links

**Related concepts**  

[Default database driver](1074-default-database-driver.md "Default database driver")

[fgl\_db\_driver\_type()](../15_library-reference/2737-fgl-db-driver-type.md "Returns the 3-letter identifier/code of the current database driver.")
