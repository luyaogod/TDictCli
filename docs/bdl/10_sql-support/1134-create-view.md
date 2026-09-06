---
title: "CREATE VIEW"
source: "fgl-topics/c_fgl_static_sql_CREATE_VIEW.html"
breadcrumb: "SQL support > Static SQL statements > CREATE VIEW"
type: "concept"
---

# CREATE VIEW

> Creates a new view object in the database.

## Syntax

```
CREATE VIEW [ IF NOT EXISTS ] view-name
[ ( column-alias-name [,...] ) ]
    AS sub-query
[ WITH CHECK OPTION ]
```

where *sub-query* is a limited syntax of the `SELECT`
statement.

## Related links

**Related concepts**  

[DROP VIEW](1135-drop-view.md "Drops a view object from the database.")
