---
title: "Cursors WITH HOLD"
source: "fgl-topics/c_fgl_odiagsqt_037.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > Cursors WITH HOLD"
type: "concept"
description: "Informix® Informix closes opened cursors automatically when a transaction ends, unless the WITH HOLD option is used in the DECLARE instruction: DECLARE c1 CURSOR WITH HOLD FOR SELECT ... OPEN c1 BEGIN ..."
---

# Cursors WITH HOLD

## Informix®

Informix closes opened cursors automatically when a
transaction ends, unless the `WITH HOLD` option is used in the
`DECLARE`
instruction:

```
DECLARE c1 CURSOR WITH HOLD FOR SELECT ...
OPEN c1
BEGIN WORK
FETCH c1 ...
COMMIT WORK
FETCH c1 ...
CLOSE c1
```

## SQLite

SQLite does not close cursors when a transaction ends.

SQLite does not support the `FOR UPDATE` close in `SELECT` syntax:
Therefore, there cannot be a combination of `WITH HOLD` + `SELECT … FOR
UPDATE`.

## Solution

BDL cursors declared `WITH HOLD` remain open even
after terminating a transaction with a `COMMIT WORK` or `ROLLBACK
WORK`.

For consistency with other database brands, database cursors that are
not declared `WITH HOLD` are automatically closed, when a `COMMIT
WORK` or `ROLLBACK WORK` is performed.

> **Important:**
>
> Opening a `WITH HOLD` cursor
> declared with a `SELECT FOR UPDATE` results in an SQL error; in the same conditions,
> this does not normally appear with Informix. Review the
> program logic in order to find another way to set locks.

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
