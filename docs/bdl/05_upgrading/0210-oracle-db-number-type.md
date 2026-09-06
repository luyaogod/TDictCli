---
title: "Oracle DB NUMBER type"
source: "fgl-topics/c_fgl_Migrate_to_300_oracle_number.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Oracle DB NUMBER type"
type: "concept"
description: "The NUMBER/FLOAT Oracle data type can now be extracted by fgldbsch to create .sch files."
---

# Oracle DB NUMBER type

> The NUMBER/FLOAT Oracle® data type can now be extracted by fgldbsch to create .sch files.

Before Genero 3.00, columns using the native Oracle NUMBER/NUMBER(p>32) type (with up to 38 significant digits), or the FLOAT(b) type (when
(b/3)>32), were not allowed by the fgldbsch schema extractor. This restriction
was applied to avoid the risk of overflow errors, if the Oracle NUMBER/FLOAT column contains values that do not fit into a BDL
DECIMAL(32,s) type.

Starting with Genero 3.00, fgldbsch can map NUMBER/FLOAT native Oracle types to BDL DECIMAL(32) or DECIMAL(32,s)
types, when specifying the flag `B` with the `-cv` option for these
native types:

- NUMBER (floating point number) is extracted as DECIMAL(32)
- NUMBER(p>32) (scale defaults to 0) is extracted as DECIMAL(32,0)
- NUMBER(p>32,s) or NUMBER(\*,s) is extracted as DECIMAL(32,s)
- FLOAT(b) is extracted as DECIMAL(b/3) or FLOAT

For more details about Oracle type
conversion rules and `-cv` type positions, run [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") with the `-cx ora`
option.

This new behavior has been introduced to simplify integration with existing Oracle databases, to extract .sch schema
from databases using column types that have no exact equivalent BDL type. When designing new database
tables, it is recommended that you only use DECIMAL(p,s), with p<=32 to achieve maximum portability.
When fetching numeric values with more than 32 significant digits into BDL decimals, values will be
rounded for DECIMAL(32), or raise an overflow error -1226 for DECIMAL(32,s).

## Related links

**Related concepts**  

[Numeric data types](../10_sql-support/1374-numeric-data-types.md "Numeric data types")

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
