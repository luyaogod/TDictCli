---
title: "The base.SQLHandle built-in class"
source: "fgl-topics/c_fgl_sql_programming_sql_handle.html"
breadcrumb: "SQL support > Dynamic SQL management > The base.SQLHandle built-in class"
type: "concept"
---

# The base.SQLHandle built-in class

> Handle SQL queries with a 3GL API.

Genero BDL provides a 3GL API to execute SQL queries and introspect result set column information
with the [`base.SqlHandle`](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets.") built-in
class.

The class implements typical SQL statement execution methods existing in well-known APIs, such
as:

- `prepare("sql-text")`
- `setParameter()`
- `execute()`
- `open()`, `openScrollCursor()`
- `fetch()`, `fetchFirst()`, `fetchLast()`, ...

The class also implements introspection methods for the result set columns:

- `getResultCount()`
- `getResultType(index)`
- `getResultName(index)`
- `getResultValue(index)`

This class is provided to allow generic code implementation for specific needs. Consider using
traditional static and dynamic SQL instruction for regular code implementing your business rules;
the 3GL code based on the SqlHandle class is not as readable as static or dynamic SQL.

## Related links

**Related concepts**  

[The SqlHandle class](../15_library-reference/3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets.")
