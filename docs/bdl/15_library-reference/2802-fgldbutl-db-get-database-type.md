---
title: "fgldbutl.db_get_database_type()"
source: "fgl-topics/c_fgl_utility_functions_DB_GET_DATABASE_TYPE.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_get_database_type()"
type: "concept"
---

# fgldbutl.db_get_database_type()

> Returns the database type for the current connection.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION db_get_database_type()
  RETURNS STRING
```

## Usage

After connecting to the database, you can get the type of the database server with this
function.

> **Important:**
>
> This function is deprecated, use the [`fgl_db_driver_type()`](2737-fgl-db-driver-type.md "Returns the 3-letter identifier/code of the current database driver.") function instead. Note that
> `db_get_database_type()` returns a code of the type of database engine (in
> uppercase), while `fgl_db_driver_type()` returns the ODI driver code (in
> lowercase).

| DB Code | Description |
| --- | --- |
| `IFX` | IBM® Informix® |
| `MDB` | MariaDB |
| `MYS` | Oracle® MySQL |
| `MSV` | Microsoft™ SQL Server |
| `ORA` | Oracle Database |
| `PGS` | PostgreSQL |
| `SQT` | SQLite |
