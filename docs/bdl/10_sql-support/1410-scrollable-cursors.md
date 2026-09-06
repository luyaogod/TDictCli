---
title: "Scrollable cursors"
source: "fgl-topics/c_fgl_odiagora_051.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > Scrollable cursors"
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

## ORACLE

Oracle OCI supports native scrollable cursors with the
`OCI_STMT_SCROLLABLE_READONLY` statement attribute.

## Solution

The Oracle database driver uses native scrollable cursors by setting the
`OCI_STMT_SCROLLABLE_READONLY` statement attribute.

With Oracle is it possible to use [LOB columns](1380-text-and-byte-lob-types.md) in a
scrollable cursor. However, you should consider to use only simple data types for scrollable
cursors, and fetch `TEXT/BYTE` data in a secondary `SELECT` statement
using the primary key of the current row.

## Related links

**Related concepts**  

[Scrollable cursors](1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.")
