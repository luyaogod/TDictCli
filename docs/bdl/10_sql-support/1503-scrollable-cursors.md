---
title: "Scrollable cursors"
source: "fgl-topics/c_fgl_odiagsqt_026.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > Scrollable cursors"
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

## SQLite

SQLite does not support scrollable cursors.

## Solution

The SQLite database driver emulates scrollable cursors by fetching rows in a temporary file.

> **Important:**
>
> With SQLite is it NOT possible to use [LOB columns](1486-text-and-byte-lob-types.md) in a scrollable cursor. If `TEXT/BYTE` columns are used with a
> scrollable cursor, the `OPEN` instruction will produce the SQL error [-611](../15_library-reference/4483-genero-bdl-errors.md) (as with Informix).

See [Scrollable cursors](1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.") for more details about scroll
cursor emulation.
