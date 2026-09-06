---
title: "Database schema extractor"
source: "fgl-topics/c_fgl_MigI4GL_008.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > Installation and setup topics > Database schema extractor"
type: "concept"
---

# Database schema extractor

> Before compiling .4gl or .per files, you must extract the database schema with the fgldbsch tool.

The fgldbsch tool will produce an .sch file, and optionally .val and .att
files. The fgldbsch tool can extract database schemas from Informix®, as well as from other databases such as Oracle® and SQL Server, but you must be aware of data type
conversion rules.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
