---
title: "Understanding database result sets"
source: "fgl-topics/c_fgl_result_sets_002.html"
breadcrumb: "SQL support > Result set processing > Understanding database result sets"
type: "concept"
---

# Understanding database result sets

> This is an introduction to database result sets.

A database result set is a group of rows produced by an SQL statement such as
`SELECT`. The result set is maintained by the database server. In a program, you
handle a result set with a database cursor.

First you must declare the database cursor with the [`DECLARE`](1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") instruction. This instruction sends the SQL statement to the
database server for parsing, validation and to generate the execution plan.

![Database result set diagram](../_images/RSTFig01.jpg)

*Database result set*

The result set is produced after execution of the SQL statement, when the database cursor is
associated with the result set by the [`OPEN`](1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") instruction. At this point, no data rows are transmitted to the
program. You must use the [`FETCH`](1152-fetch-result-set-cursor.md "Moves a cursor to a new row in the corresponding result set and retrieves the row values into fetch buffers.")
instruction to retrieve data rows from the database server.

![FETCH instruction diagram](../_images/RSTFig02.jpg)

*FETCH instruction*

When finished with the result set processing, you must [`CLOSE`](1153-close-result-set-cursor.md "Closes a database cursor and frees resources allocated on the database server for the result set.") the cursor to release the resources allocated for the result set on
the database server. The cursor can be reopened if needed. If the SQL statement is no longer needed,
you can free the resources allocated to statement execution with the [`FREE`](1154-free-result-set-cursor.md "Releases SQL cursor resources allocated by the DECLARE instruction.") instruction.

![FREE instruction diagram](../_images/RSTFig03.jpg)

*FREE instruction*

The scope of reference of a database cursor is local to a module,
so a cursor that was declared in one source file cannot be referenced
in a statement in another file.

The language supports sequential cursors and scrollable
cursors. Sequential cursors, which are unidirectional,
are used to retrieve rows for a `REPORT`,
for example. Scrollable cursors allow you to move backwards
or to an absolute or relative position in the result set. Specify
whether a cursor is scrollable with the `SCROLL` option
of the `DECLARE` instruction.

For better code readability, use a [`FOREACH /
END FOREACH`](1155-foreach-result-set-cursor.md "Processes a series of data rows returned from a database cursor.") loop, to perform the equivalent of an `OPEN` +
`FETCH` (in `WHILE` loop) + `CLOSE`.

## Related links

**Related concepts**  

[Reports](../12_reports/2461-reports.md "Reports")
