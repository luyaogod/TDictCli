---
title: "CREATE TABLE"
source: "fgl-topics/c_fgl_static_sql_CREATE_TABLE.html"
breadcrumb: "SQL support > Static SQL statements > CREATE TABLE"
type: "concept"
---

# CREATE TABLE

> Creates a new table object in the database.

## Syntax

```
CREATE [TEMP] TABLE [ IF NOT EXISTS ] table-specification
(
{ column-name data-type
       [ DEFAULT default-value ] [ NOT NULL ]
       [ PRIMARY KEY [ CONSTRAINT constraint-name ]
       | UNIQUE [ CONSTRAINT constraint-name ]
       | CHECK ( sql-condition ) [ CONSTRAINT constraint-name ]
       | REFERENCES table-name
             [ ( column-name [,...] ) ]
             [ ON DELETE CASCADE ]
             [ CONSTRAINT constraint-name ]
       ]
} [,...]
[ ,
  PRIMARY KEY ( column-name [,...] ) [ CONSTRAINT constraint-name ]
| UNIQUE ( column-name [,...] ) [ CONSTRAINT constraint-name ]
| CHECK ( sql-condition ) [ CONSTRAINT constraint-name ]
| FOREIGN KEY ( column-name [,...] )
     REFERENCES table-name
             [ ( column-name [,...] ) ]
             [ ON DELETE CASCADE ]
             [ CONSTRAINT constraint-name ]
] [,...]
)
[ WITH NO LOG ]
[ IN tablespace-name ]
[ EXTENT SIZE integer ]
[ NEXT SIZE integer ]
[ LOCK MODE { PAGE | ROW } ]
```

## Related links

**Related concepts**  

[ALTER TABLE](1129-alter-table.md "Modifies the definition of an existing table in the database.")

[DROP TABLE](1130-drop-table.md "Drops a table object from the database.")
