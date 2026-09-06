---
title: "Database schema extractor"
source: "fgl-topics/c_fgl_Mig0000_006.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > Installation and setup topics > Database schema extractor"
type: "concept"
---

# Database schema extractor

> Before compiling .4gl or .per files with Four Js Business Development Suite (BDS) or with Genero Business Development Language (BDS), you need to extract the database schema as a .sch file. However, the extraction tools differ.

BDS provides the fglschema tool, while Genero BDL provides the
fgldbsch tool. The fglschema tool could only extract schemas
from Informix® databases; the fgldbsch
can extract database schemas from any SQL server supported by Genero BDL.

The fglschema tool is still supported in Genero BDL for backward
compatibility, but fglschema actually calls fgldbsch.

> **Note:**
>
> Genero BDL allows you to centralize new widget types and attributes in the .val file.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
