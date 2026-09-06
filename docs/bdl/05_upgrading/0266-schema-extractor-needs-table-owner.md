---
title: "Schema extractor needs table owner"
source: "fgl-topics/c_fgl_Migrate_to_230_fgldbsch_schema_owner.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > Schema extractor needs table owner"
type: "concept"
---

# Schema extractor needs table owner

> The fgldbsch schema extractor requires a -ow option to distinguish different database users/shemas.

Starting with version 2.30, the fgldbsch schema extractor will always use a
table owner / schema to select tables from databases where several schemas can hold tables with
the same name.

The table owner can be specified with the `-ow` option, and defaults to the user
name passed with the `-un` option, or to the current database user if no
`-up` option was given. The last case can occur when the database connection
information is taken from the FGLPROFILE configuration file, or when the OS user authentication
is used.

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
