---
title: "Specifying the output filename"
source: "fgl-topics/c_fgl_DatabaseSchema_010.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Specifying the output filename"
type: "concept"
description: "By default, the generated schema files get the name of the database source specified with the -db option. The name of the schema file can be forced with the -of filename option. Specify the output ..."
---

# Specifying the output filename

By default, the generated schema files get the name of the database source specified with the
`-db` option.

The name of the schema file can be forced with the `-of
filename` option.

Specify the output filename without the .sch extension.

```
fgldbsch -db test1 -of myschema
```

The filename specified with the `-of` option will also be used to generate the
files containing column validation rules and column attributes (extracted from IBM® Informix® syscolval and syscolatt
tables).

## Related links

**Related concepts**  

[Column Validation File (.val)](0796-column-validation-file-val.md "The .val database schema file holds functional and display attributes of database table columns.")

[Column Video Attributes File (.att)](0797-column-video-attributes-file-att.md "The .att database schema file contains the default video attributes of database table columns.")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
