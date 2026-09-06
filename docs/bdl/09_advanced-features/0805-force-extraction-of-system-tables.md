---
title: "Force extraction of system tables"
source: "fgl-topics/c_fgl_DatabaseSchema_009.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Force extraction of system tables"
type: "concept"
description: "By default fgldbsch does not extract the definition of database system tables. Use the -st option to extract schema information of system tables. fgldbsch -db test1 -st"
---

# Force extraction of system tables

By default fgldbsch does not extract the definition of database system
tables.

Use the `-st` option to extract schema information of system tables.

```
fgldbsch -db test1 -st
```

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
