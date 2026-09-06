---
title: "Extracting the database schema with fgldbsch"
source: "fgl-topics/c_fgl_Migrate_to_220_fgldbsch.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Extracting the database schema with fgldbsch"
type: "concept"
---

# Extracting the database schema with fgldbsch

> The fgldbsch database schema extraction tool has been updated to map native database types to newly-added types.

Version 2.20 implements new data types such as BIGINT and BOOLEAN. The fgldbsch database schema
extraction tool has been reviewed to map native database types to these new types when possible.
Pay attention to these changes, when extracting a schema from your database.

For example, before version 2.20, fgldbsch converted an Oracle
NUMBER(20,0) to a DECIMAL(20,0) by default. Now, since 2.20 provides
the BIGINT native FGL type, it can be used to store a NUMBER(20,0)
from Oracle.

You can get the previous behavior by using a conversion directive with the `-cv`
option of fgldbsch.

To see the new conversion rules, run the fgldbsch tool with the `-ct` option.

## Related links

**Related concepts**  

[BIGINT](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.")

[BOOLEAN](../08_language-basics/0556-boolean.md "The BOOLEAN data type stores a logical value, TRUE or FALSE.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
