---
title: "FreeTDS driver supports SQL Server 2008, 2012, 2014"
source: "fgl-topics/c_fgl_Migrate_to_300_freetds.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > FreeTDS driver supports SQL Server 2008, 2012, 2014"
type: "concept"
---

# FreeTDS driver supports SQL Server 2008, 2012, 2014

> The FreeTDS driver can now be used for SQL Server versions > 2005.

Before Genero version 3.00, the FreeTDS driver could only be used to connect to SQL Server 2005.
Starting with Genero 3.00 the `dbmftm` driver can connect to SQL Server 2008, 2012
and 2014.

With SQL Server version >= 2008, date/time types used to store DATE and DATETIME values are
different to those used with SQL Server version 2005. See [DATE and DATETIME data types](../10_sql-support/1275-date-and-datetime-data-types.md) for
more details.

> **Important:**
>
> For SQL Server version 2008, 2012 and 2014, you must set
> `TDS_Version=7.3` in odbc.ini. Using TDS version 8.0 will cause
> problems (tested with FreeTDS 0.95.5 to 0.95.19)

## Related links

**Related tasks**  

[Prepare the runtime environment - connecting to the database](../10_sql-support/1260-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
