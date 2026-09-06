---
title: "CLOSE (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_CLOSE.html"
breadcrumb: "SQL support > Result set processing > CLOSE (result set cursor)"
type: "concept"
---

# CLOSE (result set cursor)

> Closes a database cursor and frees resources allocated on the database server for the result set.

## Syntax

```
CLOSE cid
```

1. cid is the identifier of the database cursor.

## Usage

The `CLOSE` instruction
releases the resources allocated for the result set on the database
server.

After using the `CLOSE` instruction, you must reopen the cursor with [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") before retrieving values with [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.").

It is recommended that you close the cursor when the result set is no longer used, this saves
resources on the database client and database server side.

## Example

```
MAIN
   DATABASE stores 
   DECLARE c1 CURSOR FOR SELECT * FROM customer 
   OPEN c1
   CLOSE c1
   OPEN c1
   CLOSE c1
END MAIN
```

## Related links

**Related concepts**  

[FREE (result set cursor)](1154-free-result-set-cursor.md "Releases SQL cursor resources allocated by the DECLARE instruction.")
