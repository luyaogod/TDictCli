---
title: "CREATE INDEX"
source: "fgl-topics/c_fgl_static_sql_CREATE_INDEX.html"
breadcrumb: "SQL support > Static SQL statements > CREATE INDEX"
type: "concept"
---

# CREATE INDEX

> Creates a new index object in the database.

## Syntax

```
CREATE [ UNIQUE | CLUSTER | UNIQUE CLUSTER ] INDEX
       [ IF NOT EXISTS ] index-name
 ON table-specification
    ( column-name [ ASCENCDING | DESCENDING ]  [,...] )
```

## Related links

**Related concepts**  

[ALTER INDEX](1132-alter-index.md "Modifies the definition of an existing index in the database.")

[DROP INDEX](1133-drop-index.md "Drops an index object from the database.")
