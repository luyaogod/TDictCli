---
title: "FETCH (result set cursor)"
source: "fgl-topics/c_fgl_result_sets_FETCH.html"
breadcrumb: "SQL support > Result set processing > FETCH (result set cursor)"
type: "concept"
---

# FETCH (result set cursor)

> Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.

## Syntax

```
FETCH [ fetch-option ] cid
   [ INTO fvar [,...] ]
```

where *fetch-option* is one
of:

```
{
  NEXT
| { PREVIOUS | PRIOR }
| CURRENT
| FIRST
| LAST
| ABSOLUTE position
| RELATIVE offset
}
```

1. cid is the identifier of the database cursor.
2. fvar is a variable used as fetch buffer.
3. Fetch options different from `NEXT` can only be used with scrollable
   cursors.
4. position is an positive integer expression.
5. offset is a positive or negative integer expression.

## Usage

The `FETCH` instruction retrieves a row from a result set of an opened cursor.

Before fetching rows, the cursor SQL statement must be executed with the [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") instruction.

Fetching rows can have specific behavior when the cursor was declared `FOR UPDATE`
to perform a [positioned update or delete](1156-positioned-updates-deletes.md "Describes row modification based on a FOR UPDATE cursor.").

The `INTO` clause can be used to provide the fetch buffers that receive the result
set column values.

A sequential cursor can fetch only the next row in sequence from the result set.

`FETCH [NEXT]` (the default) retrieves the next row in the result set. If the row
pointer was on the last row before executing the instruction, the SQL code is set to 100
(`NOTFOUND`), and the row pointer remains on the last row. (if you issue a
`FETCH PREVIOUS` at this time, you get the next-to-last row).

`FETCH PREVIOUS` retrieves the previous row in the result set. If the row pointer
was on the first row before executing the instruction, the SQL code is set to [`NOTFOUND/100`](../08_language-basics/0575-notfound.md "NOTFOUND is a predefined constant used to check if an SQL statement returns rows.")), and the row pointer
remains on the first row. (if you issue a `FETCH NEXT` at this time, you get the
second row).

`FETCH CURRENT` retrieves the current row in the result set.

`FETCH FIRST` retrieves the first row in the result set.

`FETCH LAST` retrieves the last row in the result set.

`FETCH ABSOLUTE position` retrieves the row at
position in the result set. If the position is not correct,
the SQL code is set to 100 (`NOTFOUND`). Absolute row positions are numbered from
1.

`FETCH RELATIVE offset` clause moves offset
rows in the result set and returns the row at the current position. The offset can be a negative
value. If the offset is not correct, the SQL code is set to 100
(`NOTFOUND`). If offset is zero, the current row is fetched.

## Example

```
MAIN
   DEFINE cust_rec RECORD
            cnum INTEGER,
            cname CHAR(20)
       END RECORD
   DATABASE stores 
   DECLARE c1 SCROLL CURSOR FOR SELECT customer_num, cust_name FROM customer 
   OPEN c1
   FETCH c1 INTO cust_rec.*
   FETCH LAST c1 INTO cust_rec.* 
   FETCH PREVIOUS c1 INTO  cust_rec.* 
   FETCH FIRST c1 INTO cust_rec.*
   FETCH LAST c1 -- INTO clause is optional
   FETCH FIRST c1 -- INTO clause is optional
END MAIN
```

## Related links

**Related concepts**  

[FOREACH (result set cursor)](1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor.")

[SQL execution diagnostics](0987-sql-execution-diagnostics.md "If an SQL statement execution fails, error description can be found in the sqlca.sqlcode, SQLSTATE, status and SQLERRMESSAGE predefined registers.")
