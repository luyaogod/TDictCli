---
title: "FREE (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_FREE.html"
breadcrumb: "SQL support > Result set processing > FREE (result set cursor)"
type: "concept"
---

# FREE (result set cursor)

> Releases SQL cursor resources allocated by the DECLARE instruction.

## Syntax

```
FREE cid
```

1. cid is the identifier of the database cursor.

## Usage

The `FREE` instruction releases the resources allocated for a cursor created by a
[`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") instruction.

If no [`CLOSE`](1153-close-result-set-cursor.md "Closes a database cursor and frees resources allocated on the database server for the result set.") instruction
was executed, the cursor is automatically closed when doing a `FREE`.

When cursor resources are released with `FREE`, the cursor must be declared
again before usage.

Free the cursor when the result set is no longer used by the program; this saves resources
on the database client and database server side.

## Example

```
MAIN
   DEFINE i, j INTEGER
   DATABASE stores
   FOR i=1 TO 10
       DECLARE c1 CURSOR FOR SELECT * FROM customer
       FOR j=1 TO 10
           OPEN c1
           FETCH c1
           CLOSE c1
       END FOR
       FREE c1
   END FOR
END MAIN
```

## Related links

**Related concepts**  

[DECLARE (result set cursor)](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.")
