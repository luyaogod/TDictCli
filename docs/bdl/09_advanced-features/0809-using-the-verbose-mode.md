---
title: "Using the verbose mode"
source: "fgl-topics/c_fgl_DatabaseSchema_013.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Using the verbose mode"
type: "concept"
description: "By default, fgldbsch extracts the database schema silently without any output. Use the -v option to get verbose output from fgldbsch : fgldbsch -db test1 -v Do not base other tools or development ..."
---

# Using the verbose mode

By default, fgldbsch extracts the database schema silently without any
output.

Use the `-v` option to get verbose output from
fgldbsch:

```
fgldbsch -db test1 -v
```

Do not base other tools or development procedures on the output format of the
fgldbsch `-v` option: The output can change in later versions.

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
