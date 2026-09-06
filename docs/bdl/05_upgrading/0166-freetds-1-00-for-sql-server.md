---
title: "FreeTDS 1.00 for SQL Server"
source: "fgl-topics/c_fgl_Migrate_to_320_freetds_version.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > FreeTDS 1.00 for SQL Server"
type: "concept"
---

# FreeTDS 1.00 for SQL Server

> Genero 3.20 requires FreeTDS version 1.00+ to connect to SQL Server.

Starting with Genero 3.20, the `dbmftm` driver to connect to SQL Server requires
[FreeTDS](http://www.freetds.org) version 1.00.

> **Important:**
>
> The minimum required version of FreeTDS is **1.00.104**.

Download the latest stable release of FreeTDS 1.00 from <http://www.freetds.org/files/stable/>.

> **Note:**
>
> FreeTDS 1.00 fixes several bugs related to date/time data, allowing [better support for date/time types](0165-datetime-sql-type-mappings.md "For some databases, the type mapping for DATETIME HOUR TO MINUTE has changed.").

For more details about FreeTDS / SQL Server configuration, see [SQL Server client environment](../10_sql-support/1260-prepare-the-runtime-environment-connecting-to-the-database.md).
