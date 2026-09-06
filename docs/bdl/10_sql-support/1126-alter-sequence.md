---
title: "ALTER SEQUENCE"
source: "fgl-topics/c_fgl_static_sql_ALTER_SEQUENCE.html"
breadcrumb: "SQL support > Static SQL statements > ALTER SEQUENCE"
type: "concept"
---

# ALTER SEQUENCE

> Modifies the definition of an existing sequence in the database.

## Syntax

```
ALTER SEQUENCE sequence-name
[ INCREMENT BY integer
| RESTART WITH integer
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

[CREATE SEQUENCE](1125-create-sequence.md "Creates a new sequence object in the database.")

[DROP SEQUENCE](1127-drop-sequence.md "Drops a sequence object from the database.")
