---
title: "Scrollable cursors"
source: "fgl-topics/c_fgl_odiagmys_034.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > Scrollable cursors"
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

## Oracle® MySQL and MariaDB

MySQL and MariaDB do not support scrollable cursors.

## Solution

The MySQL and MariaDB database drivers emulate scrollable cursors by fetching rows in a temporary
file.

> **Important:**
>
> With MySQL and MariadDB, is it NOT possible to use [LOB columns](1333-text-and-byte-lob-types.md) in a scrollable cursor. If
> `TEXT/BYTE` columns are used with a scrollable cursor, the `OPEN`
> instruction will produce the SQL error [-611](../15_library-reference/4483-genero-bdl-errors.md) (as with Informix).

## Related links

**Related concepts**  

[Scrollable cursors](1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.")
