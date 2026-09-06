---
title: "Schema extraction tool changes"
source: "fgl-topics/c_fgl_Migrate_to_200_fgldbsch.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Schema extraction tool changes"
type: "concept"
---

# Schema extraction tool changes

> The fgldbsch schema extractor is recommended, and has been enhanced.

## Unique tool

Version prior to 2.00 two schema extractors were provided: fglschema and [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database."). The first one can only extract schemas from
Informix® databases, while the second
one can extract schemas from all supported databases.

Starting with version 2.00, the
fgldbsch tool has been extended to support the old fglschema
options, and fglschema has been replaced by a simple script calling
fgldbsch. When you call fglschema, you actually call
fgldbsch. We recommend that you use fgldbsch with its specific
command line options.

## System tables

In 2.0x, fgldbsch does not extract system tables by default. You must specify
the `-st` option to get the system tables description in the schema
files.

## Remote synonyms

The original fglschema tool searched for *remote synonyms* with Informix databases. The fgldbsch tool of
version 2.00 does not search for remote synonyms.

## Public and private synonyms

Since version 1.32.1b (build 620.313), fgldbsch does not extract *private
synonyms* anymore. Only *public synonyms* are extracted. The .sch
schema files do not contain table owners, and if two *private synonyms* have the same names,
there is no way to distinguish them in the schema files. Therefore, to avoid any mistakes,
*private synonyms* are not extracted anymore.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")

[Database schema extractor options](../09_advanced-features/0798-database-schema-extractor-options.md "The fgldbsch tool extracts the schema description for an existing database.")
