---
title: "ALTER TABLE"
source: "fgl-topics/c_fgl_static_sql_ALTER_TABLE.html"
breadcrumb: "SQL support > Static SQL statements > ALTER TABLE"
type: "concept"
---

# ALTER TABLE

> Modifies the definition of an existing table in the database.

## Syntax

```
ALTER TABLE table-specification
(
[ DROP ( column-name [,...] )
| ADD ( column-name data-type
          [ DEFAULT default-value ] [ NOT NULL ]
          [ PRIMARY KEY [ CONSTRAINT constraint-name ]
          | UNIQUE [ CONSTRAINT constraint-name ]
          | CHECK ( sql-condition ) [ CONSTRAINT constraint-name ]
          | REFERENCES table-name
                [ ( column-name [,...] ) ]
                [ ON DELETE CASCADE ]
                [ CONSTRAINT constraint-name ]
          ]
          [ BEFORE column-name
          [,...]
        )
| MODIFY ( column-name data-type
          [ DEFAULT default-value ] [ NOT NULL ]
          [ PRIMARY KEY [ CONSTRAINT constraint-name ]
          | UNIQUE [ CONSTRAINT constraint-name ]
          | CHECK ( sql-condition ) [ CONSTRAINT constraint-name ]
          | REFERENCES table-name
                [ ( column-name [,...] ) ]
                [ ON DELETE CASCADE ]
                [ CONSTRAINT constraint-name ]
          ]
          [,...]
        )
| DROP CONSTRAINT constraint-name
| ADD CONSTRAINT
    { PRIMARY KEY ( column-name [,...] ) [ CONSTRAINT constraint-name ]
    | UNIQUE ( column-name [,...] ) [ CONSTRAINT constraint-name ]
    | CHECK ( sql-condition ) [ CONSTRAINT constraint-name ]
    | FOREIGN KEY ( column-name [,...] )
         REFERENCES table-name
             [ ( column-name [,...] ) ]
             [ ON DELETE CASCADE ]
             [ CONSTRAINT constraint-name ]
    }
| LOCK MODE ( { PAGE | ROW } )
| MODIFY NEXT SIZE integer
] [,...]
)
```

## Related links

**Related concepts**  

[CREATE TABLE](1128-create-table.md "Creates a new table object in the database.")

[DROP TABLE](1130-drop-table.md "Drops a table object from the database.")
