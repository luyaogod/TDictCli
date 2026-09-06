---
title: "FREE (insert cursor)"
source: "fgl-topics/c_fgl_InsertCursors_FREE.html"
breadcrumb: "SQL support > SQL insert cursors > FREE (insert cursor)"
type: "concept"
---

# FREE (insert cursor)

> Releases resources allocated for an insert cursor.

## Syntax

```
FREE cid
```

1. cid is the identifier of the insert cursor.

## Usage

After executing the `FREE` statement,
all resources allocated to the insert cursor are released.

It is recommended that the cursor be explicitly closed with the [`CLOSE`](1169-close-insert-cursor.md "Flushes and closes an insert cursor.") instruction, before it is
freed.

If you release cursor resources with this instruction, you cannot use the cursor unless you
declare the cursor again.

## Related links

**Related concepts**  

[DECLARE (insert cursor)](1165-declare-insert-cursor.md "The DECLARE with an INSERT instruction defines an insert cursor.")
