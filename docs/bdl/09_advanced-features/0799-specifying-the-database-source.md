---
title: "Specifying the database source"
source: "fgl-topics/c_fgl_DatabaseSchema_004.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Specifying the database source"
type: "concept"
description: "The -db dbname option must be used to define the database source to which to connect. The dbname and related database connection parameters can be present in the FGLPROFILE file. Otherwise, related ..."
---

# Specifying the database source

The `-db dbname` option must be used to define the database
source to which to connect.

The `dbname` and related database connection parameters can be present in the
[FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") file. Otherwise,
related options have to be provided with the fgldbsch command such as
`-dv` for the driver.

```
fgldbsch -db test1
```

## Related links

**Related concepts**  

[Database source specification (source)](../10_sql-support/1072-database-source-specification-source.md "Database source specification (source)")

[Database schema extractor options](0798-database-schema-extractor-options.md "The fgldbsch tool extracts the schema description for an existing database.")
