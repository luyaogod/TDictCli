---
title: "Extracting definition of a single table"
source: "fgl-topics/c_fgl_DatabaseSchema_011.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Extracting definition of a single table"
type: "concept"
description: "In some cases, you may just want to extract schema file of new created tables. Use the -tn tabname option, to extract schema information of a specific table. fgldbsch -db test1 -tn customers"
---

# Extracting definition of a single table

In some cases, you may just want to extract schema file of new created tables.

Use the `-tn tabname` option, to extract schema information of
a specific table.

```
fgldbsch -db test1 -tn customers
```

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
