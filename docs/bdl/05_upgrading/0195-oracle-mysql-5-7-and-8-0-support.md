---
title: "Oracle MySQL 5.7 and 8.0 support"
source: "fgl-topics/c_fgl_Migrate_to_310_mysql_5_7.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Oracle® MySQL 5.7 and 8.0 support"
type: "concept"
description: "Support for Oracle MySQL version 5.7 and 8.0"
---

# Oracle MySQL 5.7 and 8.0 support

> Support for Oracle® MySQL version 5.7 and 8.0

## ODI driver for Oracle MySQL 8.0

Genero BDL 3.10.08 now supports Oracle
MySQL 8.0 with the `dbmmys / dbmmys_8_0` driver.

> **Important:**
>
> MySQL 8.0 uses the version number `.21` for the
> libmysqlclient.so library. The `dbmmys` generic driver name maps
> now to the `dbmmys_8_0` driver, which is linked to
> libmysqlclient.so.21. In prior Genero versions, the generic driver name
> `dbmmys` was an alias for `dbmmys_5_5` or `dbmmys_5_7`.
> When using MySQL versions older than 8.0, you need to use the exact driver name with the version
> number. See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) for more details.

## ODI driver for Oracle MySQL 5.7

Genero BDL 3.10 supports Oracle MySQL
5.7 with the `dbmmys_5_7` driver.

> **Note:**
>
> MySQL 5.7 extended support ended in Oct 2023, Genero BDL (starting with version 5.00)
> consequently desupports the 5.7 version in favor to MySQL 8.0 or higher.

> **Important:**
>
> MySQL 5.7 uses the version number .20 for the
> libmysqlclient.so library. The `dbmmys_5_7` driver is linked to
> libmysqlclient.so.20. With MySQL 5.7, you need to use the
> `dbmmys_5_7` driver. See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) for more
> details.

## Oracle MySQL JSON data type

MySQL 5.7 introduces the JSON data type to store and handle JSON documents.

This new type can be used in Genero BDL programs, by using TEXT data
type:

```
MAIN
    DEFINE p_js TEXT
    LOCATE p_js IN MEMORY
    CONNECT TO "test1+driver='dbmmys'" USER "mysuser" USING "fourjs"
    EXECUTE IMMEDIATE "create table t1 ( pk int, js json )"
    LET p_js = '{"id": "9999", "name": "Tom Baker"}'
    INSERT INTO t1 VALUES ( 1, p_js )
    LET p_js = NULL
    SELECT js INTO p_js FROM t1 WHERE pk = 1
    DISPLAY p_js
END MAIN
```

When producing a database schema file from a MySQL 5.7+ database, the fgldbsch
tool will convert JSON columns to TEXT.

## Related links

**Related concepts**  

[Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md "Oracle MySQL / MariaDB")

[What is JSON?](../09_advanced-features/0955-what-is-json.md "JSON (JavaScript Object Notation) is a well known lightweight data-interchange format for JavaScript.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
