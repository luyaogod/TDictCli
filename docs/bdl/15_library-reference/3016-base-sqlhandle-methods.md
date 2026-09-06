---
title: "base.SqlHandle methods"
source: "fgl-topics/c_fgl_ClassSqlHandle_methods.html"
breadcrumb: "Library reference > Built-in packages > The base package > The SqlHandle class > base.SqlHandle methods"
type: "concept"
description: "Table 1. Class methods Name Description base.SqlHandle.create () RETURNS base.SqlHandle Create a new base.SqlHandle object. Table 2. Object methods Name Description close () Closes the SQL handle ..."
---

# base.SqlHandle methods

| Name | Description |
| --- | --- |
| base.SqlHandle.create() RETURNS base.SqlHandle | Create a new base.SqlHandle object. |

| Name | Description |
| --- | --- |
| close() | Closes the SQL handle (cursor). |
| execute() | Executes a simple SQL statement (without result set). |
| fetch() | Fetches a new row from the SQL result set. |
| fetchAbsolute(position INTEGER) | Fetches a specified row in a scrollable SQL result set. |
| fetchFirst() | Fetches the first row in a scrollable SQL result set. |
| fetchLast() | Fetches the last row in a scrollable SQL result set. |
| fetchPrevious() | Fetches the previous row in a scrollable SQL result set. |
| fetchRelative( position INTEGER) | Fetches a row relative to the current row in a scrollable SQL result set. |
| flush() | Flushes the rows from the insert cursor buffer. |
| getResultCount() RETURNS INTEGER | Returns the number of result set columns produced by the SQL statement. |
| getResultName( index INTEGER ) RETURNS STRING | Returns the name of a column in the result set produced by the SQL statement. |
| getResultType( index INTEGER ) RETURNS STRING | Returns the Genero type name of a column in the result set produced by the SQL statement. |
| getResultValue( index INTEGER ) RETURNS fgl-type | Returns the value of a column in the result set produced by the SQL statement. |
| open() | Opens the SQL handle (SELECT or INSERT cursor). |
| openCursorWithHold() | Opens the SQL handle with holdable option. |
| openScrollCursor() | Opens the SQL handle (with scrollable option). |
| openScrollCursorWithHold() | Opens the SQL handle with scrollable and holdable option. |
| prepare( sql STRING ) | Prepares an SQL statement for the SQL handle. |
| put() | Put a new row in the insert cursor buffer. |
| setParameter( index INTEGER, value primitive-type ) | Sets the value of an SQL parameter for this SQL handle. |
| setParameterType( index INTEGER, type STRING ) | Defines the type of an SQL parameter for this SQL handle. |
