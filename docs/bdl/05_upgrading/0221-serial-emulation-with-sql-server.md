---
title: "SERIAL emulation with SQL Server"
source: "fgl-topics/c_fgl_Migrate_to_300_sqlserver_serial.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > SERIAL emulation with SQL Server"
type: "concept"
---

# SERIAL emulation with SQL Server

> The SERIAL and BIGSERIAL types can be emulated with triggers and sequences when using SQL Server 2012 and higher.

By default when using SQL Server, the SERIAL and BIGSERIAL types are emulated with IDENTITY
columns. This native sequence generator is the fastest and preferred solution. However, it
requires removing the serial column in all INSERT statements, which can lead to a large change
in your legacy code.

Until version 3.00, it was possible to workaround this limitation by using the "regtable" serial
emulation. But this solution required using a dedicated SERIALREG table that needed
to be updated for each INSERT statement. This resulted in poor performances, when
concurrent programs create rows in the same tables (locking issues in
SERIALREG).

Starting with Genero 3.00, it is now possible to use a serial emulation based on triggers and
sequences. Sequences were introduced in SQL Server version 2012, so you need at least a 2012
server in order to use this emulation:

```
dbi.database.mydb.ifxemul.datatype.serial.emulation = "trigseq"
```

## Related links

**Related concepts**  

[SERIAL and BIGSERIAL data types](../10_sql-support/1277-serial-and-bigserial-data-types.md "SERIAL and BIGSERIAL data types")

[FGLPROFILE entries for core language](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")

[Auto-incremented columns (serials)](../10_sql-support/1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
