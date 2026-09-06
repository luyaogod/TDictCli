---
title: "CREATE SEQUENCE"
source: "fgl-topics/c_fgl_static_sql_CREATE_SEQUENCE.html"
breadcrumb: "SQL support > Static SQL statements > CREATE SEQUENCE"
type: "concept"
---

# CREATE SEQUENCE

> Creates a new sequence object in the database.

## Syntax

```
CREATE SEQUENCE [ IF NOT EXISTS ] sequence-name
[ INCREMENT BY integer 
| START WITH integer
| NOMAXVALUE
| MAXVALUE integer
| NOMINVALUE
| MINVALUE integer
| CYCLE
| NOCYCLE
| CACHE integer
| NOCACHE
| ORDER
| NOORDER
]
```

## Related links

**Related concepts**  

[ALTER SEQUENCE](1126-alter-sequence.md "Modifies the definition of an existing sequence in the database.")

[DROP SEQUENCE](1127-drop-sequence.md "Drops a sequence object from the database.")
