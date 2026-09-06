---
title: "Microsoft SQL Server"
source: "fgl-topics/c_fgl_odiagmsv_001.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server"
type: "concept"
description: "Supported versions Important: When connecting to Microsoft™ SQL Server, the ODI driver will check the SQL Server compatibility level, and produce an SQL error, if the compatibility level does not ..."
---

# Microsoft SQL Server

## Supported versions

> **Important:**
>
> When connecting to Microsoft™ SQL
> Server, the ODI driver will check the SQL Server compatibility level, and produce an SQL error, if
> the compatibility level does not match one of the versions listed in this topic.

Genero BDL supports the following Microsoft SQL Server
versions / Compatibility Level:

- Microsoft SQL Server 2017 / Compatibility Level 140
- Microsoft SQL Server 2019 / Compatibility Level 150
- Microsoft SQL Server 2022 / Compatibility Level 160
- Microsoft SQL Server 2025 / Compatibility Level 170

Genero BDL supports the following Microsoft Azure SQL
Database versions:

- Microsoft Azure SQL Database v12 (with Compatibility Level 170, 160)

> **Note:**
>
> SQL Server defines the compatibility level or a database
> with:
>
> ```
> ALTER DATABASE db-name SET COMPATIBILITY_LEVEL = compat-level
> ```
>
> The
> compatiblity level is the actual version number that is checked by the ODI drivers at connection
> time. Even if you are connected to a recent SQL Server version, if the compatibility level is too
> old, the driver will refuse to connect.

## Child topics

- [Purpose of the Microsoft SQL Server SQL guide](1257-purpose-of-the-microsoft-sql-server-sql-guide.md)
- [Installation (Runtime Configuration)](1258-installation-runtime-configuration.md): Microsoft SQL Server related installation topics.
- [Database concepts](1264-database-concepts.md): Microsoft SQL Server related database concepts topics.
- [Data dictionary](1270-data-dictionary.md): Microsoft SQL Server related data dictionary topics.
- [Data manipulation](1287-data-manipulation.md): Microsoft SQL Server related data manipulation topics.
- [BDL programming](1299-bdl-programming.md): Microsoft SQL Server related programming topics.
