---
title: "Controlling the character case"
source: "fgl-topics/c_fgl_DatabaseSchema_012.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Controlling the character case"
type: "concept"
description: "By default, table and column names are converted to lower case letters to enforce compatibility with IBM® Informix®. Force lower case, upper case or case-sensitive table and column names by using the ..."
---

# Controlling the character case

By default, table and column names are converted to lower case letters to enforce compatibility
with IBM® Informix®.

Force lower case, upper case or case-sensitive table and column names by using the
`-cl`, `-cu` or `-cc`
options.

```
fgldbsch -db test1 -cc
```

As a general rule, it is strongly recommended to keep table and
column names in lowercase, in all areas (including the objects
created in the database entity).

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
