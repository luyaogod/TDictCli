---
title: "Scrollable cursors"
source: "fgl-topics/c_fgl_odiagpgs_042.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Scrollable cursors"
type: "concept"
description: "Informix® Informix SQL and Genero BDL support scrollable cursors when you specify the SCROLL clause in the DECLARE cursor instruction: DECLARE c1 SCROLL CURSOR FOR SELECT ... Important: Informix does ..."
---

# Scrollable cursors

## Informix®

Informix SQL and Genero BDL support [scrollable
cursors](1150-declare-result-set-cursor.md) when you specify the `SCROLL` clause in the `DECLARE`
cursor instruction:

```
DECLARE c1 SCROLL CURSOR FOR SELECT ...
```

> **Important:**
>
> Informix does not allow to fetch `TEXT/BYTE` columns with
> scrollable cursors. If you declare a scroll cursor with a `SELECT` containing
> `TEXT/BYTE` columns, Informix will produce the SQL error [-611](../15_library-reference/4483-genero-bdl-errors.md) when executing the
> `OPEN` instruction.

## PostgreSQL

PostgreSQL supports native scrollable cursors: The client application must use the
`SCROLL` option when declaring a server cursor.

## Solution

The PostgreSQL database driver uses native scrollable cursors by declaring server cursors with
the `SCROLL` option.

## Related links

**Related concepts**  

[Scrollable cursors](1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.")
