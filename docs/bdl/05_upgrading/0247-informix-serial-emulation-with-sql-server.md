---
title: "Informix SERIAL emulation with SQL Server"
source: "fgl-topics/c_fgl_Migrate_to_240_sqlserver_serial.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Informix® SERIAL emulation with SQL Server"
type: "concept"
---

# Informix SERIAL emulation with SQL Server

> SERIAL type emulation has been enhanced for SQL Server.

## Using SCOPE\_IDENTITY() to get the last sequence

Before version 2.40, the SQL Server drivers (SNC, MSV, ESM, FTM) used the @@IDENTITY expression
to retrieve the last generated identity column, if the native serial emulation is configured. But
@@IDENTITY is not recommended, because it can return an identity value generated for another table
in a trigger of the main table.

Starting with 2.40, the SQL Server drivers use the `SCOPE_IDENTITY()`
function, which returns the last number generated in the current scope (ignoring identity numbers
generated in triggers).

## Regtable serial emulation trigger code change

When using the "regtable" serial emulation, the code of the triggers has changed in version 2.40,
using now the `SET NOCOUNT ON` instruction. Existing serial triggers created by prior
versions must be reviewed, to have the same trigger body in all tables, otherwise an SQL error is
raised when executing `INSERT` statements.

## Related links

**Related concepts**  

[SERIAL and BIGSERIAL data types](../10_sql-support/1277-serial-and-bigserial-data-types.md "SERIAL and BIGSERIAL data types")
