---
title: "TEXT/BYTE support with FTM/ESM database drivers"
source: "fgl-topics/c_fgl_Migrate_to_250_text_byte_msv.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > TEXT/BYTE support with FTM/ESM database drivers"
type: "concept"
---

# TEXT/BYTE support with FTM/ESM database drivers

> FTM and ESM database drivers TEXT/BYTE type mapping has changed.

Since version 2.50, the TEXT and BYTE data types are now converted respectively to VARCHAR(MAX)
and VARBINARY(MAX) data types, the recommended LOB types introduced in SQL Server 2005. Before
version 2.50, the TEXT and BYTE data types were converted to TEXT and IMAGE data types,
respectively, in SQL Server.

It is still possible to use SQL Server TEXT and IMAGE types, but if you create or alter tables
in an FGL program, the VARCHAR(MAX) and VARBINARY(MAX) types will be used instead.

## Related links

**Related concepts**  

[TEXT and BYTE (LOB) types](../10_sql-support/1282-text-and-byte-lob-types.md "TEXT and BYTE (LOB) types")
