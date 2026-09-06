---
title: "The LVARCHAR type in IBM Informix databases"
source: "fgl-topics/c_fgl_Migrate_to_240_lvarchar.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > The LVARCHAR type in IBM® Informix® databases"
type: "concept"
---

# The LVARCHAR type in IBM Informix databases

> Native LVARCHAR type of Informix is now mapped by default to a large VARCHAR in schema file.

Starting with version 2.40, the fgldbsch database schema extractor converts
now by default IBM®
Informix® LVARCHAR(N) types to VARCHAR2(N) with type code
201. Before 2.40, you had to pass `-cv` AAAB... option to avoid a conversion error
when generating the schema file.

The static SQL syntax has been enhanced, to support the LVARCHAR type name in DDL statements such
as CREATE TABLE. The non-Informix ODI
drivers have been adapted to convert LVARCHAR type names to VARCHAR.

Two-Pass reports can now use VARCHAR types with a size greater than 255 bytes (the temporary
table will be created with an LVARCHAR column). However, the index is created as well, and IBM
Informix IDS (version 11 when writing these lines) has a
size limitation for indexes. You may get an SQL error -517 if the VARCHAR variable used to group /
order rows in the report routine exceeds ~350 bytes (see IDS SQL error -517 for details).

## Related links

**Related concepts**  

[Database schema extractor options](../09_advanced-features/0798-database-schema-extractor-options.md "The fgldbsch tool extracts the schema description for an existing database.")

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
