---
title: "IBM Informix synonym tables"
source: "fgl-topics/c_fgl_DatabaseSchema_014.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > IBM® Informix® synonym tables"
type: "concept"
description: "When using IBM® Informix®, fgldbsch extracts synonyms by default. Only PUBLIC synonyms are extracted, to avoid duplicates in the .sch file, when the same name is used by several synonyms for different ..."
---

# IBM Informix synonym tables

When using IBM® Informix®, fgldbsch extracts synonyms by default. Only PUBLIC synonyms are
extracted, to avoid duplicates in the .sch file, when the same name is used by
several synonyms for different table owners.

To extract PRIVATE synonyms, use the `-ow` option to specify the owner of the
tables and synonyms.

```
fgldbsch -db test1 -ow mike
```

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
